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
    Permissions []string `json:"Permissions"`
}



type Role struct {
	
	Name string   `json:"name"`
	Description string   `json:"description"`
	Permissions []string  `json:"permissions"`
}
type RoleItem struct {
	ID          string  `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string `json:"description,omitempty"` // Use a pointer for nullable fields
}



// Ingredient -
