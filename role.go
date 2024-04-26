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

// GetOrder - Returns a specifc order
func (c *Client) GetRole(roleId string, authToken *string) (*RoleByPermissions, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/organization/roles/%s/permissions", c.HostURL, orderID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	permissions := Permissions{}
	err = json.Unmarshal(body, &permissions)
	if err != nil {
		return nil, err
	}

	return &permissions, nil
}

// CreateOrder - Create new order
func (c *Client) CreateRole(name *string, description *string, permissions Permissions, authToken *string) error {
	// Create the JSON request body for creating the role
	roleBody := Role{
		Name:        *name, // Dereference the pointer to get the value
		Description: *description,
	}
	rb, err := json.Marshal(roleBody)
	if err != nil {
		return err
	}

	// Create the HTTP request to create the role
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/organization/createRole/%s", c.HostURL, c.OrgID), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	// Perform the request to create the role
	body, err := c.doRequest(req, authToken)
	if err != nil {
		return err
	}

	// Unmarshal the response body into a RoleItem to get the created role ID
	createdRole := RoleItem{}
	err = json.Unmarshal(body, &createdRole)
	if err != nil {
		return err
	}

	// Create the JSON request body for setting permissions
	var permissionsList []map[string]string
	for _, scope := range permissions.Scope {
		permissionsList = append(permissionsList, map[string]string{"permission_name": scope})
	}
	rbPermissions, err := json.Marshal(permissionsList)
	if err != nil {
		return err
	}

	// Create the HTTP request to set permissions for the role
	reqPermissions, err := http.NewRequest("POST", fmt.Sprintf("%s/organization/role/%s/permissions", c.HostURL, createdRole.ID), strings.NewReader(string(rbPermissions)))
	if err != nil {
		return err
	}

	// Perform the request to set permissions for the role
	_, err = c.doRequest(reqPermissions, authToken)
	if err != nil {
		return err
	}

	return nil
}



// DeleteOrder - Deletes an order
func (c *Client) DeleteOrder(roleID string, authToken *string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/organization/role/%s", c.HostURL, orderID), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return err
	}

	if string(body) != "Deleted order" {
		return errors.New(string(body))
	}

	return nil
}
