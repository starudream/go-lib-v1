package config

import (
	"strings"

	"github.com/spf13/viper"

	"github.com/starudream/go-lib/codec/json"
	"github.com/starudream/go-lib/constant"
	"github.com/starudream/go-lib/internal/ilog"
)

var _v = func() *viper.Viper {
	v := New()
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
	}

	ilog.X.Debug().Msgf("settings: %s", json.MustMarshalString(v.AllSettings()))

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
	GetUint                 = _v.GetUint
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
