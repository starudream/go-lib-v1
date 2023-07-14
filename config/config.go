package config

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/starudream/go-lib/codec/json"
	"github.com/starudream/go-lib/constant"

	"github.com/starudream/go-lib/internal/ilog"
	"github.com/starudream/go-lib/internal/viper"
)

var _v = func() *viper.Viper {
	v := viper.NewWithOptions(viper.WithLogger(&logger{}))
	v.SetEnvPrefix(strings.ToUpper(constant.PREFIX))
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	v.AutomaticEnv()

	var err error

	cp := v.GetString("config.path")
	if cp != "" {
		err = vReadFromFile(v, cp)
		if err != nil {
			ilog.X.Fatal().Msgf("read config file error file=%s\n%s", v.ConfigFileUsed(), err.Error())
		}
	} else {
		ns := func() (ns []string) {
			wd, _ := os.Getwd()
			hd, _ := os.UserHomeDir()
			ep, en := func() (ep, en string) {
				e, _ := os.Executable()
				if e != "" {
					ep = filepath.Dir(e)
					en = strings.TrimSuffix(filepath.Base(e), filepath.Ext(e))
				}
				return
			}()
			nm := map[string]struct{}{}
			if wd != "" {
				nm[filepath.Join(wd, "config")] = struct{}{}
				if en != "" {
					nm[filepath.Join(wd, en)] = struct{}{}
				}
			}
			if ep != "" {
				nm[filepath.Join(ep, "config")] = struct{}{}
				if en != "" {
					nm[filepath.Join(ep, en)] = struct{}{}
				}
			}
			if hd != "" {
				if en != "" {
					nm[filepath.Join(hd, ".config", en)] = struct{}{}
					nm[filepath.Join(hd, ".config", en, "config")] = struct{}{}
				}
				if constant.NAME != "" {
					nm[filepath.Join(hd, ".config", constant.NAME)] = struct{}{}
					nm[filepath.Join(hd, ".config", constant.NAME, "config")] = struct{}{}
				}
			}
			for n := range nm {
				ns = append(ns, n)
			}
			sort.Strings(ns)
			return
		}()
		ilog.X.Info().Msgf("search config file in %s", json.MustMarshalIndent(ns))
		for i := 0; i < len(ns); i++ {
			err = vReadFromFile(v, ns[i])
			if err == nil {
				break
			}
		}
	}

	if err == nil {
		ilog.X.Info().Msgf("read config file success file=%s", v.ConfigFileUsed())

		ss := v.AllSettings()
		if len(ss) != 0 {
			ilog.X.Debug().Msgf("settings: %s", json.MustMarshalString(ss))
		}
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

func vReadFromFile(v *viper.Viper, name string) (err error) {
	for i := 0; i < len(viper.SupportedExts); i++ {
		v.SetConfigFile(name + "." + viper.SupportedExts[i])
		err = v.ReadInConfig()
		if err == nil {
			return
		}
	}
	return
}
