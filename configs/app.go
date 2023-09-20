package configs

type AppEnv struct {
	Port         string
	UrlStatic    string
	UrlTemplates string
}

var app = AppEnv{}

func init() {
	app = AppEnv{
		Port:         ":8888",
		UrlStatic:    "/web",
		UrlTemplates: "web/*",
	}
}

func GetAppEnv() AppEnv {
	return app
}
