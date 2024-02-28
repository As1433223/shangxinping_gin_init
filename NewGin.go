package NewGin

import "os"

func NewGin() {
	var err error
	//todo: router
	_, err = os.Stat("./router")
	if err != nil {
		os.MkdirAll("./router", os.ModeDevice)
		os.Create("./router/router.go")
	}
	//todo: controllers
	_, err = os.Stat("./controllers")
	if err != nil {
		os.MkdirAll("./controllers", os.ModeDevice)
		os.Create("./controllers/index.go")
	}
	//todo: models
	_, err = os.Stat("./models")
	if err != nil {
		os.MkdirAll("./models", os.ModeDevice)
		os.Create("./models/init.go")
		os.Create("./models/mysql.go")
		os.Create("./models/es.go")
	}
	//todo: views
	_, err = os.Stat("./views")
	if err != nil {
		os.MkdirAll("./views", os.ModeDevice)
		os.Create("./views/index.html")
	}
	//todo: global
	_, err = os.Stat("./global")
	if err != nil {
		os.MkdirAll("./global", os.ModeDevice)
		os.Create("./global/global.go")
		os.Create("./global/global_conf.go")
		os.Create("./global/global_init.go")
	}
	_, err = os.Stat("./common")
	if err != nil {
		os.MkdirAll("./common", os.ModeDevice)
	}
}
