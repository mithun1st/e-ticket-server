package appenviroment

var _enviromentE EnviromentType

func Set(e EnviromentType) {
	_enviromentE = e
}

func GetCuerrentStatus() EnviromentType {
	return _enviromentE
}
