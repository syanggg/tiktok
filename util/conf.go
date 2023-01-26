package util

import (
	"github.com/BurntSushi/toml"
)

type Mysql struct {
	Host string
	Port int
	Database string
	Username string
	Password string
	Charset string
	ParseTime bool `toml:"parse_time"`
	Loc string
}

type Jwt struct {
	Key string
}

type Tiktok struct {
	FeedLimit int `toml:"feed_limit"`
}

type Redis struct {
	Host string
	Port int
	Database int
}

type Server struct {
	Host string
	Port int
}

type Path struct {
	FfmpegPath string `toml:"ffmpeg_path"`
	StaticPath string `toml:"static_path"`
}

type Config struct{
	DB *Mysql `toml:"mysql"`
	Jwt *Jwt `toml:"jwt"`
	Tiktok *Tiktok `toml:"tiktok"`
	RDB *Redis `toml:"redis"`
	Server *Server `toml:"server"`
	Path *Path `toml:"path"`
}

var Info Config

func init()  {
	_,err := toml.DecodeFile("D:/project/tiktok/config/config.toml",&Info)
	if err != nil{
		panic(err)
	}
}

