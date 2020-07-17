package contact

func Init() {
	InitConfig()
	InitLog()
	InitFlag()
	InitGin()
}

func Close() {
	LogClose()
}
