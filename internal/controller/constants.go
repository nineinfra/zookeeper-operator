package controller

import "strconv"

const (
	// DefaultNameSuffix is the default name suffix of the resources of the zookeeper
	DefaultNameSuffix = "-zookeeper"

	// DefaultClusterSign is the default cluster sign of the zookeeper
	DefaultClusterSign = "zookeeper"

	// DefaultStorageClass is the default storage class of the zookeeper
	DefaultStorageClass = "nineinfra-default"

	DefaultReplicas = 3

	DefaultClusterDomainName = "clusterDomain"
	DefaultClusterDomain     = "cluster.local"

	DefaultDataVolumeName = "data"

	DefaultLogVolumeName = "log"

	DefaultConfigNameSuffix      = "-config"
	DefaultHeadlessSvcNameSuffix = "-headless"

	DefaultZookeeperHome     = "/opt/zookeeper"
	DefaultZooConfigFileName = "zoo.cfg"
	DefaultLogConfigFileName = "log4j.properties"
	//DefaultStartScriptFileName = "zkStart.sh"

	DefaultDataPath = "/opt/zookeeper/data"
	DefaultLogPath  = "/opt/zookeeper/logs"

	DefaultClientPortName = "client"
	DefaultClientPort     = 2181

	DefaultQuorumPortName = "quorum"
	DefaultQuorumPort     = 2888

	DefaultElectionPortName = "election"
	DefaultElectionPort     = 3888

	DefaultMetricsPortName = "metrics"
	DefaultMetricsPort     = 7000

	DefaultAdminPortName = "admin"
	DefaultAdminPortPort = 8080

	// DefaultTerminationGracePeriod is the default time given before the
	// container is stopped. This gives clients time to disconnect from a
	// specific node gracefully.
	DefaultTerminationGracePeriod = 30

	// DefaultZookeeperVolumeSize is the default volume size for the
	// Zookeeper cache volume
	DefaultZookeeperVolumeSize    = "20Gi"
	DefaultZookeeperLogVolumeSize = "5Gi"

	// DefaultReadinessProbeInitialDelaySeconds is the default initial delay (in seconds)
	// for the readiness probe
	DefaultReadinessProbeInitialDelaySeconds = 40

	// DefaultReadinessProbePeriodSeconds is the default probe period (in seconds)
	// for the readiness probe
	DefaultReadinessProbePeriodSeconds = 10

	// DefaultReadinessProbeFailureThreshold is the default probe failure threshold
	// for the readiness probe
	DefaultReadinessProbeFailureThreshold = 10

	// DefaultReadinessProbeSuccessThreshold is the default probe success threshold
	// for the readiness probe
	DefaultReadinessProbeSuccessThreshold = 1

	// DefaultReadinessProbeTimeoutSeconds is the default probe timeout (in seconds)
	// for the readiness probe
	DefaultReadinessProbeTimeoutSeconds = 10

	// DefaultLivenessProbeInitialDelaySeconds is the default initial delay (in seconds)
	// for the liveness probe
	DefaultLivenessProbeInitialDelaySeconds = 40

	// DefaultLivenessProbePeriodSeconds is the default probe period (in seconds)
	// for the liveness probe
	DefaultLivenessProbePeriodSeconds = 10

	// DefaultLivenessProbeFailureThreshold is the default probe failure threshold
	// for the liveness probe
	DefaultLivenessProbeFailureThreshold = 10

	// DefaultLivenessProbeSuccessThreshold is the default probe success threshold
	// for the readiness probe
	DefaultLivenessProbeSuccessThreshold = 1

	// DefaultLivenessProbeTimeoutSeconds is the default probe timeout (in seconds)
	// for the liveness probe
	DefaultLivenessProbeTimeoutSeconds = 10
)

var DefaultZooConfKeyValue = map[string]string{
	"clientPort":                    strconv.Itoa(DefaultClientPort),
	"admin.serverPort":              strconv.Itoa(DefaultAdminPortPort),
	"dataDir":                       DefaultDataPath,
	"dataLogDir":                    DefaultLogPath,
	"4lw.commands.whitelist":        "cons, envi, conf, crst, srvr, stat, mntr, ruok",
	"admin.enableServer":            "true",
	"reconfigEnabled":               "false",
	"skipACL":                       "yes",
	"metricsProvider.className":     "org.apache.zookeeper.metrics.prometheus.PrometheusMetricsProvider",
	"metricsProvider.httpPort":      strconv.Itoa(DefaultMetricsPort),
	"metricsProvider.exportJvmInfo": "true",
	"initLimit":                     strconv.Itoa(10),
	"syncLimit":                     strconv.Itoa(2),
	"tickTime":                      strconv.Itoa(2000),
	"globalOutstandingLimit":        strconv.Itoa(1000),
	"preAllocSize":                  strconv.Itoa(65536),
	"snapCount":                     strconv.Itoa(10000),
	"commitLogCount":                strconv.Itoa(500),
	"snapSizeLimitInKb":             strconv.Itoa(4194304),
	"maxCnxns":                      strconv.Itoa(0),
	"maxClientCnxns":                strconv.Itoa(60),
	"minSessionTimeout":             strconv.Itoa(4000),
	"maxSessionTimeout":             strconv.Itoa(40000),
	"autopurge.snapRetainCount":     strconv.Itoa(3),
	"autopurge.purgeInterval":       strconv.Itoa(1),
	"quorumListenOnAllIPs":          strconv.FormatBool(true),
}
