package model

type Api struct {
	TenantId string

	Path       string
	EnableCors bool

	Method          string
	BackendType     int
	FunctionName    string
	FunctionVersion string
	BackendUrl      string

	AuthType            int
	AuthFunctionName    string
	AuthFunctionVersion string
}
