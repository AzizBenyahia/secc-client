package secc



// // Coffee -
// type User struct {
// 	Email         string `json:"email"`
// 	OrganizationID int    `json:"organizationId"`
// 	Role          Role   `json:"role"`
// }
type Permissions struct {
	Scope []string `json:"scope"`
	// jwt.StandardClaims
}
type RoleByPermissions struct {
	
	Permissions Permissions   `json:"Permissions"`
}

type Role struct {
	
	Name string   `json:"description"`
	Description string   `json:"description"`

}

type RoleItem struct {
	ID    int         `json:"id,omitempty"`
	Name   string `json:"Name"`
	Description      string   `json:"description"`
}


// Ingredient -
