package secc

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetAllOrders - Returns all user's order
func (c *Client) GetAllRoles(authToken *string,orgId *string) (*[]RoleItem, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/organization/roles/%s", c.HostURL,orgId), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	RoleItem := []RoleItem{}
	err = json.Unmarshal(body, &RoleItem)
	if err != nil {
		return nil, err
	}

	return &RoleItem, nil
}
type RolePermissions struct {
    Permissions []string `json:"Permissions"`
}
func (c *Client) GetRole(roleId string, authToken *string) (*RolePermissions, error) {
    req, err := http.NewRequest("GET", fmt.Sprintf("%s/organization/role/%s/permissions", c.HostURL, roleId), nil)
    if err != nil {
        return nil, err
    }

    body, err := c.doRequest(req, authToken)
    if err != nil {
        return nil, err
    }
    var rolePermissions RolePermissions
    err = json.Unmarshal(body, &rolePermissions.Permissions)
    if err != nil {
        return nil, err
    }

    return &rolePermissions, nil
}
	// CreateOrder - Create new order
	func (c *Client) CreateRole(role *Role, authToken *string)   {
		// Initialize the roleData map
		roleData := make(map[string]interface{})
		roleData["name"] = role.Name
		roleData["description"] = role.Description
	
		// Marshal roleData to JSON
		rb, err := json.Marshal(roleData)
		if err != nil {
			return 
		}
	
		// Create the HTTP request to create the role
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/organization/createRole/%s", c.HostURL, c.OrgID), strings.NewReader(string(rb)))
		if err != nil {
			return 
		}
	
		// Perform the request to create the role
		body, _ := c.doRequest(req, authToken)

		
		// Unmarshal the response body into a RoleItem to get the created role ID
		createdRole := RoleItem{}
		err = json.Unmarshal(body, &createdRole)

		if err != nil {
			return 
		}
	
		// Create the JSON request body for setting permissions
		var permissionsList []map[string]string
		for _, scope := range role.Permissions {
			permissionsList = append(permissionsList, map[string]string{"permission_name": scope})
		}

		fmt.Println(permissionsList)
		rbPermissions, err := json.Marshal(permissionsList)
		if err != nil {
			return 
		}
	
		// Create the HTTP request to set permissions for the role
		reqPermissions, err := http.NewRequest("POST", fmt.Sprintf("%s/organization/role/%s/permissions", c.HostURL, createdRole.ID), strings.NewReader(string(rbPermissions)))
		if err != nil {
			return 
		}
	
		// Perform the request to set permissions for the role
		_, err = c.doRequest(reqPermissions, authToken)
		if err != nil {
			return 
		}
	    return 

		}
		

// DeleteOrder - Deletes an order
func (c *Client) DeleteRole(roleID string, authToken *string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/organization/role/%s", c.HostURL, roleID), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return err
	}

	if string(body) != "Deleted Role" {
		return errors.New(string(body))
	}

	return nil
}
