package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyVpnPbrRouteEntryWeight invokes the vpc.ModifyVpnPbrRouteEntryWeight API synchronously
func (client *Client) ModifyVpnPbrRouteEntryWeight(request *ModifyVpnPbrRouteEntryWeightRequest) (response *ModifyVpnPbrRouteEntryWeightResponse, err error) {
	response = CreateModifyVpnPbrRouteEntryWeightResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVpnPbrRouteEntryWeightWithChan invokes the vpc.ModifyVpnPbrRouteEntryWeight API asynchronously
func (client *Client) ModifyVpnPbrRouteEntryWeightWithChan(request *ModifyVpnPbrRouteEntryWeightRequest) (<-chan *ModifyVpnPbrRouteEntryWeightResponse, <-chan error) {
	responseChan := make(chan *ModifyVpnPbrRouteEntryWeightResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVpnPbrRouteEntryWeight(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyVpnPbrRouteEntryWeightWithCallback invokes the vpc.ModifyVpnPbrRouteEntryWeight API asynchronously
func (client *Client) ModifyVpnPbrRouteEntryWeightWithCallback(request *ModifyVpnPbrRouteEntryWeightRequest, callback func(response *ModifyVpnPbrRouteEntryWeightResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVpnPbrRouteEntryWeightResponse
		var err error
		defer close(result)
		response, err = client.ModifyVpnPbrRouteEntryWeight(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyVpnPbrRouteEntryWeightRequest is the request struct for api ModifyVpnPbrRouteEntryWeight
type ModifyVpnPbrRouteEntryWeightRequest struct {
	*requests.RpcRequest
	RouteSource          string           `position:"Query" name:"RouteSource"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	NewWeight            requests.Integer `position:"Query" name:"NewWeight"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Weight               requests.Integer `position:"Query" name:"Weight"`
	VpnGatewayId         string           `position:"Query" name:"VpnGatewayId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Priority             requests.Integer `position:"Query" name:"Priority"`
	RouteDest            string           `position:"Query" name:"RouteDest"`
	NextHop              string           `position:"Query" name:"NextHop"`
	OverlayMode          string           `position:"Query" name:"OverlayMode"`
}

// ModifyVpnPbrRouteEntryWeightResponse is the response struct for api ModifyVpnPbrRouteEntryWeight
type ModifyVpnPbrRouteEntryWeightResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyVpnPbrRouteEntryWeightRequest creates a request to invoke ModifyVpnPbrRouteEntryWeight API
func CreateModifyVpnPbrRouteEntryWeightRequest() (request *ModifyVpnPbrRouteEntryWeightRequest) {
	request = &ModifyVpnPbrRouteEntryWeightRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyVpnPbrRouteEntryWeight", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyVpnPbrRouteEntryWeightResponse creates a response to parse from ModifyVpnPbrRouteEntryWeight response
func CreateModifyVpnPbrRouteEntryWeightResponse() (response *ModifyVpnPbrRouteEntryWeightResponse) {
	response = &ModifyVpnPbrRouteEntryWeightResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
