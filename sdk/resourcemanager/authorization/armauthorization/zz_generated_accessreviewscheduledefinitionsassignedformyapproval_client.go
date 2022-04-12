//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// AccessReviewScheduleDefinitionsAssignedForMyApprovalClient contains the methods for the AccessReviewScheduleDefinitionsAssignedForMyApproval group.
// Don't use this type directly, use NewAccessReviewScheduleDefinitionsAssignedForMyApprovalClient() instead.
type AccessReviewScheduleDefinitionsAssignedForMyApprovalClient struct {
	host string
	pl   runtime.Pipeline
}

// NewAccessReviewScheduleDefinitionsAssignedForMyApprovalClient creates a new instance of AccessReviewScheduleDefinitionsAssignedForMyApprovalClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAccessReviewScheduleDefinitionsAssignedForMyApprovalClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*AccessReviewScheduleDefinitionsAssignedForMyApprovalClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AccessReviewScheduleDefinitionsAssignedForMyApprovalClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// List - Get access review instances assigned for my approval.
// If the operation fails it returns an *azcore.ResponseError type.
// options - AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListOptions contains the optional parameters for the
// AccessReviewScheduleDefinitionsAssignedForMyApprovalClient.List method.
func (client *AccessReviewScheduleDefinitionsAssignedForMyApprovalClient) List(options *AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListOptions) *runtime.Pager[AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse]{
		More: func(page AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse) (AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AccessReviewScheduleDefinitionsAssignedForMyApprovalClient) listCreateRequest(ctx context.Context, options *AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/accessReviewScheduleDefinitions"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AccessReviewScheduleDefinitionsAssignedForMyApprovalClient) listHandleResponse(resp *http.Response) (AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse, error) {
	result := AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewScheduleDefinitionListResult); err != nil {
		return AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse{}, err
	}
	return result, nil
}
