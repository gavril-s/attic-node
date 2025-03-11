package config

type JSONStorage struct {
	Path     string `json:"path"`
	Capacity string `json:"capacity"`
}

type JSONConfig struct {
	MasterURL         string        `json:"master_url"`
	Host              string        `json:"host"`
	Port              uint16        `json:"port"`
	IsPersistent      bool          `json:"is_persistent"`
	AcceptedChunkSize string        `json:"accepted_chunk_size"`
	Storages          []JSONStorage `json:"storages"`
}
