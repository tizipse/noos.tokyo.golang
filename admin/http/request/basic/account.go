package basic

type ToAccountOfPermissions struct {
	Module string `json:"module" query:"module" form:"module" valid:"required" label:"模块"`
}
