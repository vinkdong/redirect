package types

import "github.com/VinkDong/gox/vtime"

type StaticMapper struct {
	Data     map[string]string `yaml:"data"`
	Template map[string]string `yaml:"template"`
	Time     VTimeConfig       `yaml:"time"`
}

type VTimeConfig struct {
	Keys  []string
	From vtime.Time
	To   vtime.Time
	Skip  []string
}

func (s *StaticMapper) Get(key string) string {
	data := s.Data
	if data, ok := data[key]; ok {
		return data
	}
	return ""
}

func (s *StaticMapper) Parser(key string) string {
	data := s.Data
	if data, ok := data[key]; ok {
		return data
	}
	return key
}

type Config struct {
	Static      StaticMapper `yaml:"static"`
	Destination map[string]string
	EnableSSL   bool
	Key         string
	Cert        string
}

type Context struct {
	Config *Config
}
