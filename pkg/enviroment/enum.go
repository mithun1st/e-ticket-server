package appenviroment

type EnviromentType int

const (
	Development EnviromentType = iota + 1
	Staging
	Production
)

func (e EnviromentType) ToTitle() string {

	switch e {
	case Development:
		return "Development"

	case Staging:
		return "Staging"

	case Production:
		return "Production"

	default:
		return ""
	}
}
