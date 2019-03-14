package customobject

import (
	lrvalidate "bitbucket.org/nombiezinja/lr-go-sdk/internal/validate"

	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
)

// PostCustomObjectCreateByUID is used to write information in JSON format to the custom object for the specified account.
// Post parameters: custom data to be created in the object.
// Required query parameter: objectname - string
// Required template parameter: uid - string
// Please ensure this feature is enabled for your LoginRadius account
func (lr Loginradius) PostCustomObjectCreateByUID(uid string, queries interface{}, body interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"objectname": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req, err := lr.Client.NewPostReq("/identity/v2/manage/account/"+uid+"/customobject", body, validatedQueries)
	if err != nil {
		return nil, err
	}

	lr.Client.AddApiCredentialsToReqHeader(req)
	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// PostCustomObjectCreateByToken is used to write information in JSON format to the custom object for the specified account.
// Post parameters: custom data to be created in the object.
// Required query parameter: objectname - string
// Required template parameter: token - string
// Please ensure this feature is enabled for your LoginRadius account
func (lr Loginradius) PostCustomObjectCreateByToken(queries interface{}, body interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"objectname": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req, err := lr.Client.NewPostReqWithToken("/identity/v2/auth/customobject", body, validatedQueries)
	if err != nil {
		return nil, err
	}

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

//GetCustomObjectByObjectRecordIDAndUID is used to retrieve the Custom Object data for the specified account.
// Required query parameter: objectname - string
// Required template parameter: uid - string, objectrecordid
// Please ensure this feature is enabled for your LoginRadius account
func (lr Loginradius) GetCustomObjectByObjectRecordIDAndUID(uid, objectRecordID string, queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"objectname": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req := lr.Client.NewGetReq("/identity/v2/manage/account/"+uid+"/customobject/"+objectRecordID, validatedQueries)
	lr.Client.AddApiCredentialsToReqHeader(req)
	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetCustomObjectByObjectRecordIDAndToken is used to retrieve the Custom Object data for the specified account.
// Required query parameter: objectname - string; apikey - string
// Required template parameter: objectrecordid
// Please ensure this feature is enabled for your LoginRadius account
func (lr Loginradius) GetCustomObjectByObjectRecordIDAndToken(objectRecordID string, queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"objectname": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/customobject/"+objectRecordID, validatedQueries)
	if err != nil {
		return nil, err
	}
	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetCustomObjectByToken is used to retrieve the specified Custom Object data for the specified account.
// Required parameters: objectname - string; apikey - string
func (lr Loginradius) GetCustomObjectByToken(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"objectname": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/customobject", validatedQueries)
	if err != nil {
		return nil, err
	}

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetCustomObjectByUID is used to retrieve all the custom objects by UID from cloud storage.
// Required parameters: objectname - string
// Required template parameter: uid
func (lr Loginradius) GetCustomObjectByUID(uid string, queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"objectname": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req := lr.Client.NewGetReq("/identity/v2/manage/account/"+uid+"/customobject/", validatedQueries)
	lr.Client.AddApiCredentialsToReqHeader(req)
	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// // GetCustomObjectByUID is used to retrieve all the custom objects by UID from cloud storage.
// func GetCustomObjectByUID(objectName, uid string) (CustomObjectMulti, error) {
// 	data := new(CustomObjectMulti)
// 	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN") + "/identity/v2/manage/account/"+uid+"/customobject/", "")
// 	if reqErr != nil {
// 		return *data, reqErr
// 	}

// 	q := req.URL.Query()
// 	q.Add("objectname", objectName)
// 	req.URL.RawQuery = q.Encode()
// 	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
// 	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
// 	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

// 	err := RunRequest(req, data)
// 	return *data, err
// }

// // PutCustomObjectUpdateByUID is used to update the specified custom object data of a specified account.
// // If the value of updatetype is 'replace' then it will fully replace custom object with new custom object and
// // if the value of updatetype is partialreplace then it will perform an upsert type operation.
// // Post parameters are the fields that need to be changed.
// func PutCustomObjectUpdateByUID(objectName, updateType, uid, objectRecordID string, body interface{}) (CustomObject, error) {
// 	data := new(CustomObject)
// 	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN") + "/identity/v2/manage/account/"+uid+"/customobject/"+
// 		objectRecordID, body)
// 	if reqErr != nil {
// 		return *data, reqErr
// 	}

// 	q := req.URL.Query()
// 	q.Add("objectname", objectName)
// 	q.Add("updatetype", updateType)
// 	req.URL.RawQuery = q.Encode()
// 	req.Header.Add("content-Type", "application/json")
// 	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
// 	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

// 	err := RunRequest(req, data)
// 	return *data, err
// }

// // PutCustomObjectUpdateByToken is used to update the specified custom object data of a specified account.
// // If the value of updatetype is 'replace' then it will fully replace custom object with new custom object and
// // if the value of updatetype is partialreplace then it will perform an upsert type operation.
// // Post parameters are the fields that need to be changed.
// func PutCustomObjectUpdateByToken(objectName, updateType, authorization,
// 	objectRecordID string, body interface{}) (CustomObject, error) {
// 	data := new(CustomObject)
// 	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN") + "/identity/v2/auth/customobject/"+objectRecordID, body)
// 	if reqErr != nil {
// 		return *data, reqErr
// 	}

// 	q := req.URL.Query()
// 	q.Add("apikey", os.Getenv("APIKEY"))
// 	q.Add("objectname", objectName)
// 	q.Add("updatetype", updateType)
// 	req.URL.RawQuery = q.Encode()
// 	req.Header.Add("content-Type", "application/json")
// 	req.Header.Add("Authorization", "Bearer "+authorization)

// 	err := RunRequest(req, data)
// 	return *data, err
// }

// DeleteCustomObjectByObjectRecordIDAndUID is used to remove the
// specified Custom Object data using ObjectRecordId of specified account.
// Required template parameters: uid, objectrecordid
// Required query parameters: objectname
func (lr Loginradius) DeleteCustomObjectByObjectRecordIDAndUID(uid, objectRecordId string, queries interface{}) (*httprutils.Response, error) {

	allowedQueries := map[string]bool{"objectname": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	req := lr.Client.NewDeleteReq("/identity/v2/manage/account/" + uid + "/customobject/" + objectRecordId)
	req.QueryParams = validatedQueries
	lr.Client.AddApiCredentialsToReqHeader(req)
	req.Headers["content-Type"] = "application/json"
	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// DeleteCustomObjectByObjectRecordIDAndUID is used to remove the
// specified Custom Object data using ObjectRecordId of specified account.
// Required template parameters: objectrecordid
// Required query parameters: objectname
func (lr Loginradius) DeleteCustomObjectByObjectRecordIDAndToken(objectRecordId string, queries interface{}) (*httprutils.Response, error) {

	allowedQueries := map[string]bool{"objectname": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	req, err := lr.Client.NewDeleteReqWithToken("/identity/v2/auth/customobject/"+objectRecordId, "", validatedQueries)
	if err != nil {
		return nil, err
	}
	req.QueryParams = validatedQueries
	req.Headers["content-Type"] = "application/json"
	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// DeleteCustomObjectByObjectRecordIDAndToken is used to remove the
// // specified Custom Object data using ObjectRecordId of specified account.
// func DeleteCustomObjectByObjectRecordIDAndToken(objectName, authorization, objectRecordID string) (CustomObject, error) {
// 	data := new(CustomObject)
// 	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN")+"/identity/v2/auth/customobject/"+objectRecordID, "")
// 	if reqErr != nil {
// 		return *data, reqErr
// 	}

// 	q := req.URL.Query()
// 	q.Add("apikey", os.Getenv("APIKEY"))
// 	q.Add("objectname", objectName)
// 	req.URL.RawQuery = q.Encode()
// 	req.Header.Add("content-Type", "application/json")
// 	req.Header.Add("Authorization", "Bearer "+authorization)

// 	err := RunRequest(req, data)
// 	return *data, err
// }
