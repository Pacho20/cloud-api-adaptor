package ibmcloud

import (
	"context"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/vpc-go-sdk/common"
)

type ClusterV2 struct {
	Service *core.BaseService
}

const DefaultServiceURL = "https://containers.cloud.ibm.com/global"

type ClusterOptions struct {
	Authenticator core.Authenticator
}

func NewClusterV2Service(options *ClusterOptions) (service *ClusterV2, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	err = core.ValidateStruct(options, "options")
	if err != nil {
		err = core.SDKErrorf(err, "", "invalid-global-options", common.GetComponentInfo())
		return
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		err = core.SDKErrorf(err, "", "new-base-error", common.GetComponentInfo())
		return
	}

	service = &ClusterV2{
		Service: baseService,
	}

	return
}

func (clusterApi *ClusterV2) GetSecurityGroups(clusterID string) (result []securityGroup, response *core.DetailedResponse, err error) {
	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(context.Background())
	builder.EnableGzipCompression = clusterApi.Service.GetEnableGzipCompression()

	// Construct the request URL
	_, err = builder.ResolveRequestURL(
		clusterApi.Service.Options.URL,
		"/network/v2/security-group/getSecurityGroups",
		nil,
	)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	builder.AddQuery("cluster", clusterID)
	builder.AddQuery("type", "cluster")

	// Add headers
	sdkHeaders := common.GetSdkHeaders("kubernetes_service_api", "V1", "GetSecurityGroups")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	// Build the request
	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse []securityGroup
	response, err = clusterApi.Service.Request(request, &rawResponse)
	if err != nil {
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		result = rawResponse
		response.Result = result
	}

	return
}

type securityGroup struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	Name         string `json:"name"`
	UserProvided bool   `json:"userProvided"`
	Shared       bool   `json:"shared"`
	WorkerPoolID string `json:"workerPoolID"`
}
