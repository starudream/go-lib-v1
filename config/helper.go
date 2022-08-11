package config

func GetStringNotEmpty(keys ...string) string {
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		if s := GetString(key); s != "" {
			return s
		}
	}
	return ""
}
