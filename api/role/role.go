package role

import "bitbucket.org/nombiezinja/lr-go-sdk/httprutils"

// PostRolesCreate creates a role with permissions.
// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/roles-create
// Required query parameters: apikey, apisecret
// Required post parameter: roles - array
// Pass data in struct lrbody.Roles as body to help ensure parameters satisfy API requirements
func (lr Loginradius) PostRolesCreate(body interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewPostReq("/identity/v2/manage/role", body)
	if err != nil {
		return nil, err
	}
	lr.Client.AddApiCredentialsToReqHeader(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// DeleteAccountRole is used to delete the role.
// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/delete-role
// Required template parameter: role - string representing the rolename of the role to be deleted
func (lr Loginradius) DeleteAccountRole(role string) (*httprutils.Response, error) {
	req := lr.Client.NewDeleteReq("/identity/v2/manage/role/" + role)
	lr.Client.AddApiCredentialsToReqHeader(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetContextRolesPermissions gets the contexts that have been configured and the associated roles and permissions.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/get-context
// Required template parameter: uid - string representing uid of the user
func (lr Loginradius) GetContextRolesPermissions(uid string) (*httprutils.Response, error) {
	req := lr.Client.NewGetReq("/identity/v2/manage/account/" + uid + "/rolecontext")
	lr.Client.AddApiCredentialsToReqHeader(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetRolesList retrieves the complete list of created roles with permissions of your app.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/roles-list
func (lr Loginradius) GetRolesList() (*httprutils.Response, error) {
	req := lr.Client.NewGetReq("/identity/v2/manage/role")
	lr.Client.AddApiCredentialsToReqHeader(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetRolesByUID is used to retrieve all the assigned roles of a particular User.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/get-roles-by-uid
// Required template parameter: uid - string representing user's uid
func (lr Loginradius) GetRolesByUID(uid string) (*httprutils.Response, error) {
	req := lr.Client.NewGetReq("/identity/v2/manage/account/" + uid + "/role")
	lr.Client.AddApiCredentialsToReqHeader(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
	// 	data := new(RoleArray)
	// 	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN") + "/identity/v2/manage/account/"+uid+"/role", "")
	// 	if reqErr != nil {
	// 		return *data, reqErr
	// 	}

	// 	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	// 	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
	// 	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	// 	err := RunRequest(req, data)
	// 	return *data, err
}

// // PutAccountAddPermissionsToRole is used to add permissions to the role.
// // Post Parameters are permissions: string
// func PutAccountAddPermissionsToRole(role string, body interface{}) (Role, error) {
// 	data := new(Role)
// 	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN") + "/identity/v2/manage/role/"+role+"/permission", body)
// 	if reqErr != nil {
// 		return *data, reqErr
// 	}

// 	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
// 	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
// 	req.Header.Add("content-Type", "application/json")

// 	err := RunRequest(req, data)
// 	return *data, err
// }

// PutRolesAssignToUser is used to assign created roles to the user.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/assign-roles-by-uid
// Required template parameter: uid - string representing uid of the user
// Required post parameter: roles - array of string(s) representing role name(s)
func (lr Loginradius) PutRolesAssignToUser(uid string, body interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewPutReq("/identity/v2/manage/account/"+uid+"/role", body)
	if err != nil {
		return nil, err
	}
	lr.Client.AddApiCredentialsToReqHeader(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
	// 	data := new(RoleArray)
	// 	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN") + "/identity/v2/manage/account/"+uid+"/role", body)
	// 	if reqErr != nil {
	// 		return *data, reqErr
	// 	}

	// 	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	// 	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
	// 	req.Header.Add("content-Type", "application/json")

	// 	err := RunRequest(req, data)
	// 	return *data, err
}

// // PutRolesUpsertContext creates a Context with a set of Roles.
// // Post Parameters are rolecontext: string, context: string, roles: string, additionalpermissions: string and an
// // optional expiration: time.Time
// func PutRolesUpsertContext(uid string, body interface{}) (ContextRole, error) {
// 	data := new(ContextRole)
// 	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN") + "/identity/v2/manage/account/"+uid+"/rolecontext", body)
// 	if reqErr != nil {
// 		return *data, reqErr
// 	}

// 	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
// 	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
// 	req.Header.Add("content-Type", "application/json")

// 	err := RunRequest(req, data)
// 	return *data, err
// }

// // DeleteAccountRole is used to delete the role.
// func DeleteAccountRole(role string) (Role, error) {
// 	data := new(Role)
// 	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN") + "/identity/v2/manage/role/"+role, "")
// 	if reqErr != nil {
// 		return *data, reqErr
// 	}

// 	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
// 	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
// 	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

// 	err := RunRequest(req, data)
// 	return *data, err
// }

// // DeleteRolesAssignedToUser is used to unassign roles to the user.
// // Post Parameter is an array of roles
// func DeleteRolesAssignedToUser(uid string, body interface{}) (Role, error) {
// 	data := new(Role)
// 	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN") + "/identity/v2/manage/account/"+uid+"/role", body)
// 	if reqErr != nil {
// 		return *data, reqErr
// 	}

// 	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
// 	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
// 	req.Header.Add("content-Type", "application/json")

// 	err := RunRequest(req, data)
// 	return *data, err
// }

// // DeleteRolesAccountRemovePermissions is used to remove permissions to the role.
// // Post Parameter is the permissions from which you want to remove the role
// func DeleteRolesAccountRemovePermissions(roleName string, body interface{}) (Role, error) {
// 	data := new(Role)
// 	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN") + "/identity/v2/manage/role/"+roleName+"/permission", body)
// 	if reqErr != nil {
// 		return *data, reqErr
// 	}

// 	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
// 	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
// 	req.Header.Add("content-Type", "application/json")

// 	err := RunRequest(req, data)
// 	return *data, err
// }

// // DeleteContextFromRole deletes the specified Role Context
// func DeleteContextFromRole(uid, rolecontextname string) (Role, error) {
// 	data := new(Role)
// 	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN") + "/identity/v2/manage/account/"+uid+"/rolecontext/"+rolecontextname, "")
// 	if reqErr != nil {
// 		return *data, reqErr
// 	}

// 	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
// 	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
// 	req.Header.Add("content-Type", "application/json")

// 	err := RunRequest(req, data)
// 	return *data, err
// }

// // DeleteRoleFromContext deletes the specified Role from a Context.
// // Post Parameters is an array of roles
// func DeleteRoleFromContext(uid, rolecontextname string, body interface{}) (Role, error) {
// 	data := new(Role)
// 	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN") + "/identity/v2/manage/account/"+uid+
// 		"/rolecontext/"+rolecontextname+"/role", body)
// 	if reqErr != nil {
// 		return *data, reqErr
// 	}

// 	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
// 	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
// 	req.Header.Add("content-Type", "application/json")

// 	err := RunRequest(req, data)
// 	return *data, err
// }

// // DeleteAdditionalPermissionFromContext deletes Additional Permissions from Context.
// // Post Parameter is the array of strings which represent additional permissions
// func DeleteAdditionalPermissionFromContext(uid, rolecontextname string, body interface{}) (Role, error) {
// 	data := new(Role)
// 	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN") + "/identity/v2/manage/account/"+uid+
// 		"/rolecontext/"+rolecontextname+"/additionalpermission", body)
// 	if reqErr != nil {
// 		return *data, reqErr
// 	}

// 	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
// 	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
// 	req.Header.Add("content-Type", "application/json")

// 	err := RunRequest(req, data)
// 	return *data, err
// }
