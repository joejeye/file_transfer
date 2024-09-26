package global_config

const (
	PeerDiscoveryPacketSizeLimitBytes = 4096
	PDServerListenPort                = "32461" // PD stands for Peer Discovery
	PDClientListenPort                = "32462"
	PDTimeOutMilliSec                 = 200
)

type ConfigLock struct {
	PDServerListenPort                string
	PDClientListenPort                string
	PeerDiscoveryPacketSizeLimitBytes int
	PDTimeOutMilliSec                 int
}

func GetLockedConfig() ConfigLock {
	return ConfigLock{
		PDServerListenPort:                PDServerListenPort,
		PDClientListenPort:                PDClientListenPort,
		PeerDiscoveryPacketSizeLimitBytes: PeerDiscoveryPacketSizeLimitBytes,
		PDTimeOutMilliSec:                 PDTimeOutMilliSec,
	}
}
