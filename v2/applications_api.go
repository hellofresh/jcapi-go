/* 
 * JumpCloud APIs
 *
 * V1 & V2 versions of JumpCloud's API. The next version of JumpCloud's API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings. The most recent version of JumpCloud's API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings.
 *
 * OpenAPI spec version: 2.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package v2

import (
	"net/url"
	"strings"
	"encoding/json"
	"fmt"
)

type ApplicationsApi struct {
	Configuration *Configuration
}

func NewApplicationsApi() *ApplicationsApi {
	configuration := NewConfiguration()
	return &ApplicationsApi{
		Configuration: configuration,
	}
}

func NewApplicationsApiWithBasePath(basePath string) *ApplicationsApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &ApplicationsApi{
		Configuration: configuration,
	}
}

/**
 * List the associations of an Application
 * This endpoint will return the _direct_ associations of an Application. A direct association can be a non-homogenous relationship between 2 different objects. for example Applications and User Groups.   #### Sample Request &#x60;&#x60;&#x60; https://console.jumpcloud.com/api/v2/applications/{application_id}/associations?targets&#x3D;user_group &#x60;&#x60;&#x60;
 *
 * @param applicationId ObjectID of the Application.
 * @param targets 
 * @param contentType 
 * @param accept 
 * @param limit The number of records to return at once.
 * @param skip The offset into the records to return.
 * @return []GraphConnection
 */
func (a ApplicationsApi) GraphApplicationAssociationsList(applicationId string, targets []string, contentType string, accept string, limit int32, skip int32) ([]GraphConnection, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/applications/{application_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"application_id"+"}", fmt.Sprintf("%v", applicationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(x-api-key)' required
	// set key with prefix in header
	localVarHeaderParams["x-api-key"] = a.Configuration.GetAPIKeyWithPrefix("x-api-key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var targetsCollectionFormat = "csv"
	localVarQueryParams.Add("targets", a.Configuration.APIClient.ParameterToString(targets, targetsCollectionFormat))

	localVarQueryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
	localVarQueryParams.Add("skip", a.Configuration.APIClient.ParameterToString(skip, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// header params "Content-Type"
	localVarHeaderParams["Content-Type"] = a.Configuration.APIClient.ParameterToString(contentType, "")
	// header params "Accept"
	localVarHeaderParams["Accept"] = a.Configuration.APIClient.ParameterToString(accept, "")
	var successPayload = new([]GraphConnection)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GraphApplicationAssociationsList", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return *successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return *successPayload, localVarAPIResponse, err
}

/**
 * Manage the associations of an Application
 * This endpoint allows you to manage the _direct_ associations of an Application. A direct association can be a non-homogenous relationship between 2 different objects. for example Application and User Groups.  #### Sample Request &#x60;&#x60;&#x60; https://console.jumpcloud.com/api/v2/applications/{application_id}/associations &#x60;&#x60;&#x60;
 *
 * @param applicationId ObjectID of the Application.
 * @param contentType 
 * @param accept 
 * @param body 
 * @return void
 */
func (a ApplicationsApi) GraphApplicationAssociationsPost(applicationId string, contentType string, accept string, body GraphManagementReq) (*APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/applications/{application_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"application_id"+"}", fmt.Sprintf("%v", applicationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(x-api-key)' required
	// set key with prefix in header
	localVarHeaderParams["x-api-key"] = a.Configuration.GetAPIKeyWithPrefix("x-api-key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// header params "Content-Type"
	localVarHeaderParams["Content-Type"] = a.Configuration.APIClient.ParameterToString(contentType, "")
	// header params "Accept"
	localVarHeaderParams["Accept"] = a.Configuration.APIClient.ParameterToString(accept, "")
	// body params
	localVarPostBody = &body
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GraphApplicationAssociationsPost", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return localVarAPIResponse, err
	}
	return localVarAPIResponse, err
}

/**
 * List the Users associated with an Application
 * This endpoint will return Users associated with an Application. Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of attributes specifically set for this group.  The &#x60;paths&#x60; array enumerates each path from this Application to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this Application.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; https://console.jumpcloud.com/api/v2/applications/{application_id}/users &#x60;&#x60;&#x60;
 *
 * @param applicationId ObjectID of the Application.
 * @param contentType 
 * @param accept 
 * @param limit The number of records to return at once.
 * @param skip The offset into the records to return.
 * @return []GraphObjectWithPaths
 */
func (a ApplicationsApi) GraphApplicationTraverseUser(applicationId string, contentType string, accept string, limit int32, skip int32) ([]GraphObjectWithPaths, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/applications/{application_id}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"application_id"+"}", fmt.Sprintf("%v", applicationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(x-api-key)' required
	// set key with prefix in header
	localVarHeaderParams["x-api-key"] = a.Configuration.GetAPIKeyWithPrefix("x-api-key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
	localVarQueryParams.Add("skip", a.Configuration.APIClient.ParameterToString(skip, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// header params "Content-Type"
	localVarHeaderParams["Content-Type"] = a.Configuration.APIClient.ParameterToString(contentType, "")
	// header params "Accept"
	localVarHeaderParams["Accept"] = a.Configuration.APIClient.ParameterToString(accept, "")
	var successPayload = new([]GraphObjectWithPaths)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GraphApplicationTraverseUser", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return *successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return *successPayload, localVarAPIResponse, err
}

/**
 * List the User Groups associated with an Application
 * This endpoint will return User Groups associated with an Application. Each element will contain the group&#39;s type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of attributes specifically set for this group.  The &#x60;paths&#x60; array enumerates  each path from this Application to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this Application.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; https://console.jumpcloud.com/api/v2/applications/{application_id}/usergroups &#x60;&#x60;&#x60;
 *
 * @param applicationId ObjectID of the Application.
 * @param contentType 
 * @param accept 
 * @param limit The number of records to return at once.
 * @param skip The offset into the records to return.
 * @return []GraphObjectWithPaths
 */
func (a ApplicationsApi) GraphApplicationTraverseUserGroup(applicationId string, contentType string, accept string, limit int32, skip int32) ([]GraphObjectWithPaths, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/applications/{application_id}/usergroups"
	localVarPath = strings.Replace(localVarPath, "{"+"application_id"+"}", fmt.Sprintf("%v", applicationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(x-api-key)' required
	// set key with prefix in header
	localVarHeaderParams["x-api-key"] = a.Configuration.GetAPIKeyWithPrefix("x-api-key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
	localVarQueryParams.Add("skip", a.Configuration.APIClient.ParameterToString(skip, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// header params "Content-Type"
	localVarHeaderParams["Content-Type"] = a.Configuration.APIClient.ParameterToString(contentType, "")
	// header params "Accept"
	localVarHeaderParams["Accept"] = a.Configuration.APIClient.ParameterToString(accept, "")
	var successPayload = new([]GraphObjectWithPaths)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GraphApplicationTraverseUserGroup", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return *successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return *successPayload, localVarAPIResponse, err
}
