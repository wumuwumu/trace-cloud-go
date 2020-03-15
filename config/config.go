package config


type Config struct {
	MongoDB struct{
		URI	string `json:"uri"`
	} `json:"mongodb"`
}

var C Config
