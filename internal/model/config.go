package model

type Config struct {
	PlaceApiKey  string  `yaml:"placeApiKey"`
	ImgUrlPrefix string  `yaml:"imgUrlPrefix"`
	BangsueLat   float64 `yaml:"bangsueLat"`
	BangsueLng   float64 `yaml:"bangsueLng"`
}
