/* 
 * JumpCloud APIs
 *
 * V1 and V2 versions of JumpCloud's API. The previous version of JumpCloud's API. This set of endpoints allows JumpCloud customers to manage commands, systems, & system users.
 *
 * OpenAPI spec version: 1.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package v1

import (
	"net/url"
	"strings"
	"encoding/json"
)

type SearchApi struct {
	Configuration *Configuration
}

func NewSearchApi() *SearchApi {
	configuration := NewConfiguration()
	return &SearchApi{
		Configuration: configuration,
	}
}

func NewSearchApiWithBasePath(basePath string) *SearchApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &SearchApi{
		Configuration: configuration,
	}
}

/**
 * List System Users
 * Return System Users in multi-record format allowing for the passing of the &#39;filter&#39; parameter. The WILL NOT allow you to add a new system.
 *
 * @param body 
 * @param contentType 
 * @param accept 
 * @param limit The number of records to return at once.
 * @param skip The offset into the records to return.
 * @param fields The comma separated fileds included in the returned records. If omitted the default list of fields will be returned. 
 * @return *InlineResponse2004
 */
func (a SearchApi) POSTSearchSystemusers(body Search, contentType string, accept string, limit int32, skip int32, fields string) (*InlineResponse2004, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/search/systemusers"

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
	localVarQueryParams.Add("fields", a.Configuration.APIClient.ParameterToString(fields, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json; charset=utf-8",
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
	var successPayload = new(InlineResponse2004)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "POSTSearchSystemusers", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}
