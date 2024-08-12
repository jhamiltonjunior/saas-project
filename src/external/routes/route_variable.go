package routes

func NewRouteVariableDB() string {
	return "root:0000@tcp(localhost:3306)/my_saas_app?charset=utf8mb4&parseTime=True&loc=Local"
}
