package config

func DBConfig() map[string]string {
	return map[string]string{
		"host":     "127.0.0.1",
		"port":     "27017",
		"username": "",
		"password": "",
		"name":     "",
	}
}
