package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/starudream/go-lib/codec/json"
	"github.com/starudream/go-lib/constant"

	"github.com/starudream/go-lib/internal/ilog"
	"github.com/starudream/go-lib/internal/viper"
)

var _v = func() *viper.Viper {
	v := viper.NewWithOptions(viper.WithLogger(&logger{}))
	v.SetConfigName("config")
	v.SetEnvPrefix(strings.ToUpper(constant.PREFIX))
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	v.AutomaticEnv()

	cp := v.GetString("config.path")
	if cp != "" {
		v.SetConfigFile(cp)
		err := v.ReadInConfig()
		if err != nil {
			ilog.X.Fatal().Msgf("read config file error file=%s\n%s", v.ConfigFileUsed(), err.Error())
		} else {
			ilog.X.Info().Msgf("read config file success file=%s", v.ConfigFileUsed())
		}
	} else {
		if dir, err := os.Getwd(); err == nil {
			v.AddConfigPath(dir)
		}
		if file, err := os.Executable(); err == nil {
			v.AddConfigPath(filepath.Dir(file))
		}
		if dir, err := os.UserHomeDir(); err == nil {
			v.AddConfigPath(dir)
		}
		if dir, err := os.UserConfigDir(); err == nil && constant.NAME != "" {
			v.AddConfigPath(filepath.Join(dir, constant.NAME))
		}
		err := v.ReadInConfig()
		if err == nil {
			ilog.X.Info().Msgf("read config file success file=%s", v.ConfigFileUsed())
		}
	}

	ss := v.AllSettings()
	if len(ss) != 0 {
		ilog.X.Debug().Msgf("settings: %s", json.MustMarshalString(ss))
	}

	return v
}()

var (
	AllKeys     = _v.AllKeys
	AllSettings = _v.AllSettings

	IsSet    = _v.IsSet
	InConfig = _v.InConfig

	Get                     = _v.Get
	GetString               = _v.GetString
	GetBool                 = _v.GetBool
	GetInt                  = _v.GetInt
	GetInt32                = _v.GetInt32
	GetInt64                = _v.GetInt64
	GetUint                 = _v.GetUint
	GetUint16               = _v.GetUint16
	GetUint32               = _v.GetUint32
	GetUint64               = _v.GetUint64
	GetFloat64              = _v.GetFloat64
	GetTime                 = _v.GetTime
	GetDuration             = _v.GetDuration
	GetIntSlice             = _v.GetIntSlice
	GetStringSlice          = _v.GetStringSlice
	GetStringMap            = _v.GetStringMap
	GetStringMapString      = _v.GetStringMapString
	GetStringMapStringSlice = _v.GetStringMapStringSlice
	GetSizeInBytes          = _v.GetSizeInBytes

	Set        = _v.Set
	SetDefault = _v.SetDefault

	BindPFlag      = _v.BindPFlag
	BindPFlags     = _v.BindPFlags
	BindFlagValue  = _v.BindFlagValue
	BindFlagValues = _v.BindFlagValues
)

func UnmarshalKeyTo[T any](key string) (t T, err error) {
	err = _v.UnmarshalKey(key, &t)
	return
}
