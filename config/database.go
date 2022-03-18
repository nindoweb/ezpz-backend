package config

func DBConfig() map[string]string {
	return map[string]string{
		"host":     "localhost",
		"port":     "27017",
		"username": "",
		"password": "",
		"name":     "ezpz",
	}
}
