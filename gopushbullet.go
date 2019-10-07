package gopushbullet

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	pushbullet "github.com/xconstruct/go-pushbullet"
)

// PushBulletConfig - configuration for pushbullet
type PushBulletConfig struct {
	Pushbulletchanneltag string `yaml:"pushbulletchanneltag"`
	PushbulletAPIKey     string `yaml:"pushbulletAPIKey"`
	PushbulletEnable     bool   `yaml:"pushbulletEnable"`
}

// PBNotify struct for doing pushbullet notification
type PBNotify struct {
	config *PushBulletConfig
}

func readConfig(path string) (*PushBulletConfig, error) {
	v := viper.New()
	v.AddConfigPath(filepath.Dir(path))
	filename := filepath.Base(path)
	v.SetConfigName(strings.TrimSuffix(filename, filepath.Ext(filename)))
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var cf = &PushBulletConfig{}
	err = v.UnmarshalKey("pbconfig", cf)
	if err != nil {
		return nil, err
	}
	return cf, nil
}

// New : Create notify
func New(path string) (*PBNotify, error) {
	config, err := readConfig(path)
	if err != nil {
		return nil, err
	}
	return &PBNotify{
		config: config,
	}, nil
}

// Notify : push
func (notify *PBNotify) Notify(title string, msg interface{}) error {
	config := notify.config
	if !config.PushbulletEnable {
		return nil
	}
	pb := pushbullet.New(config.PushbulletAPIKey)
	sub, err := pb.Subscription(config.Pushbulletchanneltag)
	if err != nil {
		return err
	}
	pbmsg := fmt.Sprint(msg)
	return sub.PushNote(title, pbmsg)
}
