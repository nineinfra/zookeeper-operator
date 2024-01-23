package controller

import (
	"fmt"
	zookeeperv1 "github.com/nineinfra/zookeeper-operator/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
)

func volumeRequest(q resource.Quantity) corev1.ResourceList {
	m := make(corev1.ResourceList, 1)
	m[corev1.ResourceStorage] = q
	return m
}

func capacityPerVolume(capacity string) (*resource.Quantity, error) {
	totalQuantity, err := resource.ParseQuantity(capacity)
	if err != nil {
		return nil, err
	}
	return resource.NewQuantity(totalQuantity.Value(), totalQuantity.Format), nil
}

func getReplicas(cluster *zookeeperv1.ZookeeperCluster) int32 {
	if cluster.Spec.Resource.Replicas != 0 && cluster.Spec.Resource.Replicas%2 != 0 {
		return cluster.Spec.Resource.Replicas
	}
	return DefaultReplicas
}

func getZooConfigValue(cluster *zookeeperv1.ZookeeperCluster, key string, value string) string {
	if cluster.Spec.Conf != nil {
		if value, ok := cluster.Spec.Conf[key]; ok {
			return value
		}
	}
	return value
}

func constructZooConfig(cluster *zookeeperv1.ZookeeperCluster) string {
	zooConf := make(map[string]string)
	for k, v := range DefaultZooConfKeyValue {
		zooConf[k] = getZooConfigValue(cluster, k, v)
	}
	replicas := getReplicas(cluster)
	for i := 0; i < int(replicas); i++ {
		zooConf[fmt.Sprintf("server.%d", i+1)] = fmt.Sprintf("%s-%d.%s.%s.svc.%s:%d:%d",
			ClusterResourceName(cluster),
			i,
			ClusterResourceName(cluster),
			cluster.Namespace,
			GetClusterDomain(cluster),
			DefaultQuorumPort,
			DefaultElectionPort)
	}
	if cluster.Spec.Conf != nil {
		for k, v := range cluster.Spec.Conf {
			if _, ok := DefaultZooConfKeyValue[k]; !ok {
				zooConf[k] = v
			}
		}
	}
	return map2String(zooConf)
}

func constructLogConfig() string {
	tmpConf := map[string]string{
		"zookeeper.root.logger":                           "CONSOLE",
		"zookeeper.console.threshold":                     "INFO",
		"log4j.rootLogger":                                "CONSOLE",
		"log4j.appender.CONSOLE":                          "org.apache.log4j.ConsoleAppender",
		"log4j.appender.CONSOLE.Threshold":                "INFO",
		"log4j.appender.CONSOLE.layout":                   "org.apache.log4j.PatternLayout",
		"log4j.appender.CONSOLE.layout.ConversionPattern": "%d{ISO8601} [myid:%X{myid}] - %-5p [%t:%C{1}@%L] - %m%n",
	}
	return map2String(tmpConf)
}

func getImageConfig(cluster *zookeeperv1.ZookeeperCluster) zookeeperv1.ImageConfig {
	ic := zookeeperv1.ImageConfig{
		Repository:  cluster.Spec.Image.Repository,
		PullSecrets: cluster.Spec.Image.PullSecrets,
	}
	ic.Tag = cluster.Spec.Image.Tag
	if ic.Tag == "" {
		ic.Tag = cluster.Spec.Version
	}
	ic.PullPolicy = cluster.Spec.Image.PullPolicy
	if ic.PullPolicy == "" {
		ic.PullPolicy = string(corev1.PullIfNotPresent)
	}
	return ic
}

func (r *ZookeeperClusterReconciler) constructHeadlessService(cluster *zookeeperv1.ZookeeperCluster) (*corev1.Service, error) {
	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ClusterResourceName(cluster, DefaultHeadlessSvcNameSuffix),
			Namespace: cluster.Namespace,
			Labels:    ClusterResourceLabels(cluster),
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name: DefaultClientPortName,
					Port: DefaultClientPort,
				},
				{
					Name: DefaultQuorumPortName,
					Port: DefaultQuorumPort,
				},
				{
					Name: DefaultElectionPortName,
					Port: DefaultElectionPort,
				},
				{
					Name: DefaultMetricsPortName,
					Port: DefaultMetricsPort,
				},
				{
					Name: DefaultAdminPortName,
					Port: DefaultAdminPortPort,
				},
			},
			Selector:  ClusterResourceLabels(cluster),
			ClusterIP: corev1.ClusterIPNone,
		},
	}
	if err := ctrl.SetControllerReference(cluster, svc, r.Scheme); err != nil {
		return svc, err
	}
	return svc, nil
}

func (r *ZookeeperClusterReconciler) constructService(cluster *zookeeperv1.ZookeeperCluster) (*corev1.Service, error) {
	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ClusterResourceName(cluster),
			Namespace: cluster.Namespace,
			Labels:    ClusterResourceLabels(cluster),
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name: DefaultClientPortName,
					Port: DefaultClientPort,
				},
				{
					Name: DefaultAdminPortName,
					Port: DefaultAdminPortPort,
				},
			},
			Selector: ClusterResourceLabels(cluster),
			Type:     corev1.ServiceTypeClusterIP,
		},
	}
	if err := ctrl.SetControllerReference(cluster, svc, r.Scheme); err != nil {
		return svc, err
	}
	return svc, nil
}

func (r *ZookeeperClusterReconciler) constructConfigMap(cluster *zookeeperv1.ZookeeperCluster) (*corev1.ConfigMap, error) {
	cm := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ClusterResourceName(cluster, DefaultConfigNameSuffix),
			Namespace: cluster.Namespace,
			Labels:    ClusterResourceLabels(cluster),
		},
		Data: map[string]string{
			DefaultZooConfigFileName: constructZooConfig(cluster),
			DefaultLogConfigFileName: constructLogConfig(),
		},
	}
	if err := ctrl.SetControllerReference(cluster, cm, r.Scheme); err != nil {
		return cm, err
	}
	return cm, nil
}

func (r *ZookeeperClusterReconciler) getStorageRequests(cluster *zookeeperv1.ZookeeperCluster) (*resource.Quantity, error) {
	if cluster.Spec.Resource.ResourceRequirements.Requests != nil {
		if value, ok := cluster.Spec.Resource.ResourceRequirements.Requests["storage"]; ok {
			return &value, nil
		}
	}
	return capacityPerVolume(DefaultZookeeperVolumeSize)
}

func (r *ZookeeperClusterReconciler) defaultZookeeperPorts() []corev1.ContainerPort {
	return []corev1.ContainerPort{
		{
			Name:          DefaultClientPortName,
			ContainerPort: DefaultClientPort,
		},
		{
			Name:          DefaultQuorumPortName,
			ContainerPort: DefaultQuorumPort,
		},
		{
			Name:          DefaultElectionPortName,
			ContainerPort: DefaultElectionPort,
		},
		{
			Name:          DefaultMetricsPortName,
			ContainerPort: DefaultMetricsPort,
		},
		{
			Name:          DefaultAdminPortName,
			ContainerPort: DefaultAdminPortPort,
		},
	}
}

func (r *ZookeeperClusterReconciler) constructZookeeperPodSpec(cluster *zookeeperv1.ZookeeperCluster) corev1.PodSpec {
	tgp := int64(DefaultTerminationGracePeriod)
	ic := getImageConfig(cluster)
	var tmpPullSecrets []corev1.LocalObjectReference
	if ic.PullSecrets != "" {
		tmpPullSecrets = make([]corev1.LocalObjectReference, 0)
		tmpPullSecrets = append(tmpPullSecrets, corev1.LocalObjectReference{Name: ic.PullSecrets})
	}
	return corev1.PodSpec{
		Containers: []corev1.Container{
			{
				Name:            cluster.Name,
				Image:           ic.Repository + ":" + ic.Tag,
				ImagePullPolicy: corev1.PullPolicy(ic.PullPolicy),
				Ports:           r.defaultZookeeperPorts(),
				Env:             DefaultDownwardAPI(),
				ReadinessProbe: &corev1.Probe{
					ProbeHandler: corev1.ProbeHandler{
						Exec: &corev1.ExecAction{
							Command: []string{
								"/bin/bash",
								"-c",
								"echo ruok|nc 127.0.0.1 2181",
							},
						},
					},
					InitialDelaySeconds: DefaultReadinessProbeInitialDelaySeconds,
					PeriodSeconds:       DefaultReadinessProbePeriodSeconds,
					TimeoutSeconds:      DefaultReadinessProbeTimeoutSeconds,
					FailureThreshold:    DefaultReadinessProbeFailureThreshold,
					SuccessThreshold:    DefaultReadinessProbeSuccessThreshold,
				},
				LivenessProbe: &corev1.Probe{
					ProbeHandler: corev1.ProbeHandler{
						Exec: &corev1.ExecAction{
							Command: []string{
								"/bin/bash",
								"-c",
								fmt.Sprintf("%s/bin/zkServer.sh status", DefaultZookeeperHome)},
						},
					},
					InitialDelaySeconds: DefaultLivenessProbeInitialDelaySeconds,
					PeriodSeconds:       DefaultLivenessProbePeriodSeconds,
					TimeoutSeconds:      DefaultLivenessProbeTimeoutSeconds,
					FailureThreshold:    DefaultLivenessProbeFailureThreshold,
					SuccessThreshold:    DefaultLivenessProbeSuccessThreshold,
				},
				VolumeMounts: []corev1.VolumeMount{
					{
						Name:      ClusterResourceName(cluster, DefaultConfigNameSuffix),
						MountPath: fmt.Sprintf("%s/conf/%s", DefaultZookeeperHome, DefaultZooConfigFileName),
						SubPath:   DefaultZooConfigFileName,
					},
					{
						Name:      ClusterResourceName(cluster, DefaultConfigNameSuffix),
						MountPath: fmt.Sprintf("%s/conf/%s", DefaultZookeeperHome, DefaultLogConfigFileName),
						SubPath:   DefaultLogConfigFileName,
					},
					{
						Name:      DefaultDataVolumeName,
						MountPath: DefaultDataPath,
					},
					{
						Name:      DefaultLogVolumeName,
						MountPath: DefaultLogPath,
					},
				},
			},
		},
		ImagePullSecrets:              tmpPullSecrets,
		RestartPolicy:                 corev1.RestartPolicyAlways,
		TerminationGracePeriodSeconds: &tgp,
		Volumes: []corev1.Volume{
			{
				Name: ClusterResourceName(cluster, DefaultConfigNameSuffix),
				VolumeSource: corev1.VolumeSource{
					ConfigMap: &corev1.ConfigMapVolumeSource{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: ClusterResourceName(cluster, DefaultConfigNameSuffix),
						},
						Items: []corev1.KeyToPath{
							{
								Key:  DefaultZooConfigFileName,
								Path: DefaultZooConfigFileName,
							},
							{
								Key:  DefaultLogConfigFileName,
								Path: DefaultLogConfigFileName,
							},
						},
					},
				},
			},
			{
				Name: DefaultDataVolumeName,
				VolumeSource: corev1.VolumeSource{
					PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
						ClaimName: DefaultDataVolumeName,
						ReadOnly:  false,
					},
				},
			},
			{
				Name: DefaultLogVolumeName,
				VolumeSource: corev1.VolumeSource{
					PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
						ClaimName: DefaultLogVolumeName,
						ReadOnly:  false,
					},
				},
			},
		},
	}
}
func (r *ZookeeperClusterReconciler) constructZookeeperWorkload(cluster *zookeeperv1.ZookeeperCluster) (*appsv1.StatefulSet, error) {
	q, err := r.getStorageRequests(cluster)
	if err != nil {
		return nil, err
	}
	logq, err := capacityPerVolume(DefaultZookeeperLogVolumeSize)
	if err != nil {
		return nil, err
	}
	sc := GetStorageClassName(cluster)
	stsDesired := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ClusterResourceName(cluster),
			Namespace: cluster.Namespace,
			Labels:    ClusterResourceLabels(cluster),
		},
		Spec: appsv1.StatefulSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: ClusterResourceLabels(cluster),
			},
			ServiceName: ClusterResourceName(cluster),
			Replicas:    int32Ptr(getReplicas(cluster)),
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ClusterResourceLabels(cluster),
				},
				Spec: r.constructZookeeperPodSpec(cluster),
			},
			VolumeClaimTemplates: []corev1.PersistentVolumeClaim{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      DefaultDataVolumeName,
						Namespace: cluster.Namespace,
						Labels:    ClusterResourceLabels(cluster),
					},
					Spec: corev1.PersistentVolumeClaimSpec{
						StorageClassName: &sc,
						AccessModes: []corev1.PersistentVolumeAccessMode{
							corev1.ReadWriteOnce,
						},
						Resources: corev1.ResourceRequirements{
							Requests: volumeRequest(*q),
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      DefaultLogVolumeName,
						Namespace: cluster.Namespace,
						Labels:    ClusterResourceLabels(cluster),
					},
					Spec: corev1.PersistentVolumeClaimSpec{
						StorageClassName: &sc,
						AccessModes: []corev1.PersistentVolumeAccessMode{
							corev1.ReadWriteOnce,
						},
						Resources: corev1.ResourceRequirements{
							Requests: volumeRequest(*logq),
						},
					},
				},
			},
		},
	}

	if err := ctrl.SetControllerReference(cluster, stsDesired, r.Scheme); err != nil {
		return stsDesired, err
	}
	return stsDesired, nil
}
