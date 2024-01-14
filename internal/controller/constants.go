package controller

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

	DefaultDataNameSuffix = "-data"

	DefaultLogNameSuffix = "-log"

	DefaultConfigNameSuffix      = "-config"
	DefaultHeadlessSvcNameSuffix = "-headless"

	DefaultZooConfigFileName   = "zoo.cfg"
	DefaultLogConfigFileName   = "log4j.properties"
	DefaultStartScriptFileName = "zkStart.sh"

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

	DefaultInitLimit = 10
	DefaultSyncLimit = 2
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
	DefaultReadinessProbeInitialDelaySeconds = 30

	// DefaultReadinessProbePeriodSeconds is the default probe period (in seconds)
	// for the readiness probe
	DefaultReadinessProbePeriodSeconds = 10

	// DefaultReadinessProbeFailureThreshold is the default probe failure threshold
	// for the readiness probe
	DefaultReadinessProbeFailureThreshold = 3

	// DefaultReadinessProbeSuccessThreshold is the default probe success threshold
	// for the readiness probe
	DefaultReadinessProbeSuccessThreshold = 1

	// DefaultReadinessProbeTimeoutSeconds is the default probe timeout (in seconds)
	// for the readiness probe
	DefaultReadinessProbeTimeoutSeconds = 10

	// DefaultLivenessProbeInitialDelaySeconds is the default initial delay (in seconds)
	// for the liveness probe
	DefaultLivenessProbeInitialDelaySeconds = 30

	// DefaultLivenessProbePeriodSeconds is the default probe period (in seconds)
	// for the liveness probe
	DefaultLivenessProbePeriodSeconds = 10

	// DefaultLivenessProbeFailureThreshold is the default probe failure threshold
	// for the liveness probe
	DefaultLivenessProbeFailureThreshold = 3

	// DefaultLivenessProbeSuccessThreshold is the default probe success threshold
	// for the readiness probe
	DefaultLivenessProbeSuccessThreshold = 1

	// DefaultLivenessProbeTimeoutSeconds is the default probe timeout (in seconds)
	// for the liveness probe
	DefaultLivenessProbeTimeoutSeconds = 10
)
