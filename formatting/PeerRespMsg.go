package formatting

// PeerRespMsg is the message format for the response to a peer discovery message.
type PeerRespMsg struct {
	Name              string `json:"Name"`
	FileReceptionPort string `json:"FileReceptionPort"`
}
