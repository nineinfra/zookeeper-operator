package controller

import (
	zookeeperv1 "github.com/nineinfra/zookeeper-operator/api/v1"
	"strings"
)

func ClusterResourceName(cluster *zookeeperv1.ZookeeperCluster, suffixs ...string) string {
	return cluster.Name + DefaultNameSuffix + strings.Join(suffixs, "-")
}

func ClusterResourceLabels(cluster *zookeeperv1.ZookeeperCluster) map[string]string {
	return map[string]string{
		"cluster": cluster.Name,
		"app":     DefaultClusterSign,
	}
}

func GetStorageClassName(cluster *zookeeperv1.ZookeeperCluster) string {
	if cluster.Spec.Resource.StorageClass != "" {
		return cluster.Spec.Resource.StorageClass
	}
	return DefaultStorageClass
}
