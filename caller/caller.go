package caller

type Call struct {
	ChainID  string `json:"chainId"`
	HTTPCall `json:"httpCall"`
}
