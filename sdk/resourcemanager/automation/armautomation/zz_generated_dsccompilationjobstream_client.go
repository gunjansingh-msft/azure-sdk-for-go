//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DscCompilationJobStreamClient contains the methods for the DscCompilationJobStream group.
// Don't use this type directly, use NewDscCompilationJobStreamClient() instead.
type DscCompilationJobStreamClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDscCompilationJobStreamClient creates a new instance of DscCompilationJobStreamClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDscCompilationJobStreamClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DscCompilationJobStreamClient, error) {
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
	client := &DscCompilationJobStreamClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// ListByJob - Retrieve all the job streams for the compilation Job.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// jobID - The job id.
// options - DscCompilationJobStreamClientListByJobOptions contains the optional parameters for the DscCompilationJobStreamClient.ListByJob
// method.
func (client *DscCompilationJobStreamClient) ListByJob(ctx context.Context, resourceGroupName string, automationAccountName string, jobID string, options *DscCompilationJobStreamClientListByJobOptions) (DscCompilationJobStreamClientListByJobResponse, error) {
	req, err := client.listByJobCreateRequest(ctx, resourceGroupName, automationAccountName, jobID, options)
	if err != nil {
		return DscCompilationJobStreamClientListByJobResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscCompilationJobStreamClientListByJobResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DscCompilationJobStreamClientListByJobResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByJobHandleResponse(resp)
}

// listByJobCreateRequest creates the ListByJob request.
func (client *DscCompilationJobStreamClient) listByJobCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, jobID string, options *DscCompilationJobStreamClientListByJobOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/compilationjobs/{jobId}/streams"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	urlPath = strings.ReplaceAll(urlPath, "{jobId}", url.PathEscape(jobID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByJobHandleResponse handles the ListByJob response.
func (client *DscCompilationJobStreamClient) listByJobHandleResponse(resp *http.Response) (DscCompilationJobStreamClientListByJobResponse, error) {
	result := DscCompilationJobStreamClientListByJobResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobStreamListResult); err != nil {
		return DscCompilationJobStreamClientListByJobResponse{}, err
	}
	return result, nil
}
