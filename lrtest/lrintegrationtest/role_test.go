package lrintegrationtest

import (
	"reflect"
	"testing"

	"bitbucket.org/nombiezinja/lr-go-sdk/api/role"
	lrbody "bitbucket.org/nombiezinja/lr-go-sdk/lrbody"
	lrjson "bitbucket.org/nombiezinja/lr-go-sdk/lrjson"
)

func TestPostRolesCreate(t *testing.T) {
	_, _, _, tearDown := setupRole(t)
	defer tearDown(t)
}

func TestDeleteAccountRole(t *testing.T) {
	_, _, _, tearDown := setupRole(t)
	defer tearDown(t)
}

func TestGetContextRolesPermissions(t *testing.T) {
	_, _, uid, _, _, lrclient, tearDown := setupLogin(t)
	defer tearDown(t)
	_, err := role.Loginradius(role.Loginradius{lrclient}).GetContextRolesPermissions(uid)
	if err != nil {
		t.Errorf("Error calling GetContextRolesPermissions: %v", err)
	}
}

func TestGetRolesList(t *testing.T) {
	_, rolename, lrclient, tearDown := setupRole(t)
	defer tearDown(t)
	res, err := role.Loginradius(role.Loginradius{lrclient}).GetRolesList()
	if err != nil {
		t.Errorf("Error calling GetRolesList: %v", err)
	}
	roles, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil {
		t.Errorf("Error returned from GetRolesList: %v", err)
	}
	exists := false
	for _, r := range roles["data"].([]interface{}) {
		if r.(map[string]interface{})["Name"].(string) == rolename {
			exists = true
		}
	}

	if !exists {
		t.Errorf("Error returning created role %v from GetRolesList: %v", rolename, roles)
	}
}

func TestGetRolesByUID(t *testing.T) {
	_, _, uid, _, lrclient, tearDownAccount := setupAccount(t)
	defer tearDownAccount(t)

	_, rolename, lrclient, tearDownRole := setupRole(t)
	defer tearDownRole(t)

	_, err := role.Loginradius(role.Loginradius{lrclient}).PutRolesAssignToUser(
		uid,
		lrbody.RoleList{[]string{rolename}},
	)

	if err != nil {
		t.Errorf("Error calling PutRolesAssignToUser for GetRolesByUID%v", err)
	}

	res, err := role.Loginradius(role.Loginradius{lrclient}).GetRolesByUID(uid)
	if err != nil {
		t.Errorf("Error calling GetRolesByUID%v", err)
	}

	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !reflect.DeepEqual(data["Roles"].([]interface{})[0].(string), rolename) {
		t.Errorf("Error returned from GetRolesByUID %v, %v", err, data)
	}
}

func TestPutAccountAddPermissionsToRole(t *testing.T) {
	_, _, _, _, lrclient, tearDownAccount := setupAccount(t)
	defer tearDownAccount(t)

	_, rolename, lrclient, tearDownRole := setupRole(t)
	defer tearDownRole(t)

	permissionName := "example_permission_name"

	res, err := role.Loginradius(role.Loginradius{lrclient}).PutAccountAddPermissionsToRole(
		rolename,
		lrbody.PermissionList{[]string{permissionName}},
	)
	if err != nil {
		t.Errorf("Error calling PutAccountAddPermissionsToRole: %v", err)
	}

	data, err := lrjson.DynamicUnmarshal(res.Body)
	included := false
	for k, _ := range data["Permissions"].(map[string]interface{}) {
		if k == permissionName {
			included = true
		}
	}
	if err != nil || !included {
		t.Errorf("Error returned from PutAccountAddPermissionsToRole %v, %v", err, data)
	}
}

func TestPutRolesAssignToUser(t *testing.T) {
	_, _, uid, _, lrclient, tearDownAccount := setupAccount(t)
	defer tearDownAccount(t)

	_, rolename, lrclient, tearDownRole := setupRole(t)
	defer tearDownRole(t)

	res, err := role.Loginradius(role.Loginradius{lrclient}).PutRolesAssignToUser(
		uid,
		// lrbody.RoleList{[]string{rolename}},
		[]byte(`{"roles": ["example_role_name"]}`),
	)

	if err != nil {
		t.Errorf("Error calling PutRolesAssignToUser %v", err)
	}

	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !reflect.DeepEqual(data["Roles"].([]interface{})[0].(string), rolename) {
		t.Errorf("Error returned from PutRolesAssignToUser %v, %v, %v", err, data["Roles"], []string{rolename})
	}
}

// func TestPutRolesAssignToUserInvalid(t *testing.T) {
// 	fmt.Println("Starting test TestPutRolesAssignToUserInvalid")
// 	_, _, testuid, _, teardownAccount := setupAccount(t)
// 	defer teardownAccount(t)
// 	invalid := InvalidBody{"invalid"}
// 	_, err := PutRolesAssignToUser(testuid, invalid)
// 	if err == nil {
// 		t.Errorf("Should be error")
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Test complete")
// }

// func TestPutRolesUpsertContext(t *testing.T) {
// 	fmt.Println("Starting test TestPutRolesUpsertContext")
// 	_, _, testuid, _, teardownAccount := setupAccount(t)
// 	defer teardownAccount(t)
// 	roleName, teardownRole := setupRole(t)
// 	defer teardownRole(t)
// 	roleContext := RoleContext{"contextTest", []string{roleName}, []string{"permission1"}, ""}
// 	roleContextContainer := RoleContextContainer{[]RoleContext{roleContext}}
// 	_, err := PutRolesUpsertContext(testuid, roleContextContainer)
// 	if err != nil {
// 		t.Errorf("Error setting role context for user")
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Test complete")
// }

// func TestPutRolesUpsertContextInvalid(t *testing.T) {
// 	fmt.Println("Starting test TestPutRolesUpsertContextInvalid")
// 	_, _, testuid, _, teardownAccount := setupAccount(t)
// 	defer teardownAccount(t)
// 	invalid := InvalidBody{"invalid"}
// 	_, err := PutRolesUpsertContext(testuid, invalid)
// 	if err == nil {
// 		t.Errorf("Should be error")
// 	}
// 	fmt.Println("Test complete")
// }

// func TestDeleteRolesAssignedToUser(t *testing.T) {
// 	fmt.Println("Starting test TestDeleteRolesAssignedToUser")
// 	_, _, testuid, _, teardownAccount := setupAccount(t)
// 	defer teardownAccount(t)
// 	roleName, teardownRole := setupRole(t)
// 	defer teardownRole(t)
// 	roles := RoleList{[]string{roleName}}
// 	_, err := PutRolesAssignToUser(testuid, roles)
// 	if err != nil {
// 		t.Errorf("Error setting role for user")
// 		fmt.Println(err)
// 	}
// 	_, err2 := DeleteRolesAssignedToUser(testuid, roles)
// 	if err2 != nil {
// 		t.Errorf("Error deleting role for user")
// 		fmt.Println(err2)
// 	}
// 	fmt.Println("Test complete")
// }

// func TestDeleteRolesAssignedToUserInvalid(t *testing.T) {
// 	fmt.Println("Starting test TestDeleteRolesAssignedToUserInvalid")
// 	_, _, testuid, _, teardownAccount := setupAccount(t)
// 	defer teardownAccount(t)
// 	invalid := InvalidBody{"invalid"}
// 	_, err := PutRolesAssignToUser(testuid, invalid)
// 	if err == nil {
// 		t.Errorf("Error should be error")
// 	}
// 	fmt.Println("Test complete")
// }

func TestDeleteRolesAccountRemovePermissions(t *testing.T) {
	_, _, _, _, lrclient, tearDownAccount := setupAccount(t)
	defer tearDownAccount(t)

	permissionName, rolename, lrclient, tearDownRole := setupRole(t)
	defer tearDownRole(t)

	res, err := role.Loginradius(role.Loginradius{lrclient}).DeleteRolesAccountRemovePermissions(
		rolename,
		lrbody.PermissionList{[]string{permissionName}},
	)

	if err != nil {
		t.Errorf("Error calling DeleteRolesAccountRemovePermissions: %v", err)
	}

	data, err := lrjson.DynamicUnmarshal(res.Body)
	included := false
	for k, _ := range data["Permissions"].(map[string]interface{}) {
		if k == permissionName {
			included = true
		}
	}
	if err != nil || included {
		t.Errorf("Error returned from DeleteRolesAccountRemovePermission %v, %v", err, data)
	}
}

// func TestDeleteRolesAccountRemovePermissionsInvalid(t *testing.T) {
// 	fmt.Println("Starting test TestDeleteRolesAccountRemovePermissions")
// 	roleName, teardownTestCase := setupRole(t)
// 	defer teardownTestCase(t)
// 	invalid := InvalidBody{"invalid"}
// 	_, err := PutAccountAddPermissionsToRole(roleName, invalid)
// 	if err == nil {
// 		t.Errorf("Should be error")
// 	}
// 	fmt.Println("Test complete")
// }

// func TestDeleteContextFromRole(t *testing.T) {
// 	fmt.Println("Starting test TestDeleteContextFromRole")
// 	_, _, testuid, _, teardownAccount := setupAccount(t)
// 	defer teardownAccount(t)
// 	roleName, teardownRole := setupRole(t)
// 	defer teardownRole(t)
// 	roleContext := RoleContext{"contextTest", []string{roleName}, []string{"permission1"}, ""}
// 	roleContextContainer := RoleContextContainer{[]RoleContext{roleContext}}
// 	_, err := PutRolesUpsertContext(testuid, roleContextContainer)
// 	if err != nil {
// 		t.Errorf("Error adding contexts and roles to user")
// 		fmt.Println(err)
// 	}
// 	_, err2 := DeleteContextFromRole(testuid, "contextTest")
// 	if err2 != nil {
// 		t.Errorf("Error deleting role context")
// 		fmt.Println(err2)
// 	}
// 	fmt.Println("Test complete")
// }

// func TestDeleteContextFromRoleInvalid(t *testing.T) {
// 	fmt.Println("Starting test TestDeleteContextFromRoleInvalid")
// 	_, _, testuid, _, teardownAccount := setupAccount(t)
// 	defer teardownAccount(t)
// 	invalid := InvalidBody{"invalid"}
// 	_, err := PutRolesUpsertContext(testuid, invalid)
// 	if err == nil {
// 		t.Errorf("Should be error")
// 	}
// 	fmt.Println("Test complete")
// }

// func TestDeleteRoleFromContext(t *testing.T) {
// 	fmt.Println("Starting test TestDeleteRoleFromContext")
// 	_, _, testuid, _, teardownAccount := setupAccount(t)
// 	defer teardownAccount(t)
// 	roleName, teardownRole := setupRole(t)
// 	defer teardownRole(t)
// 	roleContext := RoleContext{"contextTest", []string{roleName}, []string{"permission1"}, ""}
// 	roleContextContainer := RoleContextContainer{[]RoleContext{roleContext}}
// 	roles := RoleList{[]string{roleName}}
// 	_, err := PutRolesUpsertContext(testuid, roleContextContainer)
// 	if err != nil {
// 		t.Errorf("Error adding contexts and roles to user")
// 		fmt.Println(err)
// 	}
// 	_, err2 := DeleteRoleFromContext(testuid, "contextTest", roles)
// 	if err2 != nil {
// 		t.Errorf("Error deleting role context")
// 		fmt.Println(err2)
// 	}
// 	fmt.Println("Test complete")
// }

// func TestDeleteRoleFromContextInvalid(t *testing.T) {
// 	fmt.Println("Starting test TestDeleteRoleFromContextInvalid")
// 	_, _, testuid, _, teardownAccount := setupAccount(t)
// 	defer teardownAccount(t)
// 	invalid := InvalidBody{"invalid"}
// 	_, err := PutRolesUpsertContext(testuid, invalid)
// 	if err == nil {
// 		t.Errorf("Should be error")
// 	}
// 	fmt.Println("Test complete")
// }

// func TestDeleteAdditionalPermissionFromContext(t *testing.T) {
// 	fmt.Println("Starting test TestDeleteAdditionalPermissionFromContext")
// 	_, _, testuid, _, teardownAccount := setupAccount(t)
// 	defer teardownAccount(t)
// 	roleName, teardownRole := setupRole(t)
// 	defer teardownRole(t)
// 	roleContext := RoleContext{"contextTest", []string{roleName}, []string{"permission1"}, ""}
// 	roleContextContainer := RoleContextContainer{[]RoleContext{roleContext}}
// 	permissions := DeletePermissionList{[]string{"permission1"}}
// 	_, err := PutRolesUpsertContext(testuid, roleContextContainer)
// 	if err != nil {
// 		t.Errorf("Error adding contexts and roles to user")
// 		fmt.Println(err)
// 	}
// 	_, err2 := DeleteAdditionalPermissionFromContext(testuid, "contextTest", permissions)
// 	if err2 != nil {
// 		t.Errorf("Error deleting role context")
// 		fmt.Println(err2)
// 	}
// 	fmt.Println("Test complete")
// }

// func TestDeleteAdditionalPermissionFromContextInvalid(t *testing.T) {
// 	fmt.Println("Starting test TestDeleteAdditionalPermissionFromContextInvalid")
// 	_, _, testuid, _, teardownAccount := setupAccount(t)
// 	defer teardownAccount(t)
// 	invalid := InvalidBody{"invalid"}
// 	_, err := PutRolesUpsertContext(testuid, invalid)
// 	if err == nil {
// 		t.Errorf("Should be error")
// 	}
// 	fmt.Println("Test complete")
// }
