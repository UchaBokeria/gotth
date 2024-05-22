package category

type CategoryDto struct {
	ID   	string 	`param:"id" form:"id"`
	Public 	string 	`form:"Public"`
	Name 	string 	`form:"Name"`
	Icon 	[]byte 	`form:"Icon"`
}