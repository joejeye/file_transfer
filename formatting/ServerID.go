package formatting

type ServerID struct {
	PeerRespMsg
	ServerIp string `json:"ServerIp"`
}
