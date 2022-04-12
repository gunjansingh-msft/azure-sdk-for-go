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

// SoftwareUpdateConfigurationsClient contains the methods for the SoftwareUpdateConfigurations group.
// Don't use this type directly, use NewSoftwareUpdateConfigurationsClient() instead.
type SoftwareUpdateConfigurationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSoftwareUpdateConfigurationsClient creates a new instance of SoftwareUpdateConfigurationsClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSoftwareUpdateConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SoftwareUpdateConfigurationsClient, error) {
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
	client := &SoftwareUpdateConfigurationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Create a new software update configuration with the name given in the URI.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// softwareUpdateConfigurationName - The name of the software update configuration to be created.
// parameters - Request body.
// options - SoftwareUpdateConfigurationsClientCreateOptions contains the optional parameters for the SoftwareUpdateConfigurationsClient.Create
// method.
func (client *SoftwareUpdateConfigurationsClient) Create(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationName string, parameters SoftwareUpdateConfiguration, options *SoftwareUpdateConfigurationsClientCreateOptions) (SoftwareUpdateConfigurationsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, automationAccountName, softwareUpdateConfigurationName, parameters, options)
	if err != nil {
		return SoftwareUpdateConfigurationsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SoftwareUpdateConfigurationsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SoftwareUpdateConfigurationsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *SoftwareUpdateConfigurationsClient) createCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationName string, parameters SoftwareUpdateConfiguration, options *SoftwareUpdateConfigurationsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/softwareUpdateConfigurations/{softwareUpdateConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if softwareUpdateConfigurationName == "" {
		return nil, errors.New("parameter softwareUpdateConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{softwareUpdateConfigurationName}", url.PathEscape(softwareUpdateConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header.Set("clientRequestId", *options.ClientRequestID)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleResponse handles the Create response.
func (client *SoftwareUpdateConfigurationsClient) createHandleResponse(resp *http.Response) (SoftwareUpdateConfigurationsClientCreateResponse, error) {
	result := SoftwareUpdateConfigurationsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SoftwareUpdateConfiguration); err != nil {
		return SoftwareUpdateConfigurationsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - delete a specific software update configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// softwareUpdateConfigurationName - The name of the software update configuration to be created.
// options - SoftwareUpdateConfigurationsClientDeleteOptions contains the optional parameters for the SoftwareUpdateConfigurationsClient.Delete
// method.
func (client *SoftwareUpdateConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationName string, options *SoftwareUpdateConfigurationsClientDeleteOptions) (SoftwareUpdateConfigurationsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, softwareUpdateConfigurationName, options)
	if err != nil {
		return SoftwareUpdateConfigurationsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SoftwareUpdateConfigurationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return SoftwareUpdateConfigurationsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return SoftwareUpdateConfigurationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SoftwareUpdateConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationName string, options *SoftwareUpdateConfigurationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/softwareUpdateConfigurations/{softwareUpdateConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if softwareUpdateConfigurationName == "" {
		return nil, errors.New("parameter softwareUpdateConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{softwareUpdateConfigurationName}", url.PathEscape(softwareUpdateConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header.Set("clientRequestId", *options.ClientRequestID)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// GetByName - Get a single software update configuration by name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// softwareUpdateConfigurationName - The name of the software update configuration to be created.
// options - SoftwareUpdateConfigurationsClientGetByNameOptions contains the optional parameters for the SoftwareUpdateConfigurationsClient.GetByName
// method.
func (client *SoftwareUpdateConfigurationsClient) GetByName(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationName string, options *SoftwareUpdateConfigurationsClientGetByNameOptions) (SoftwareUpdateConfigurationsClientGetByNameResponse, error) {
	req, err := client.getByNameCreateRequest(ctx, resourceGroupName, automationAccountName, softwareUpdateConfigurationName, options)
	if err != nil {
		return SoftwareUpdateConfigurationsClientGetByNameResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SoftwareUpdateConfigurationsClientGetByNameResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SoftwareUpdateConfigurationsClientGetByNameResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByNameHandleResponse(resp)
}

// getByNameCreateRequest creates the GetByName request.
func (client *SoftwareUpdateConfigurationsClient) getByNameCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationName string, options *SoftwareUpdateConfigurationsClientGetByNameOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/softwareUpdateConfigurations/{softwareUpdateConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if softwareUpdateConfigurationName == "" {
		return nil, errors.New("parameter softwareUpdateConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{softwareUpdateConfigurationName}", url.PathEscape(softwareUpdateConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header.Set("clientRequestId", *options.ClientRequestID)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByNameHandleResponse handles the GetByName response.
func (client *SoftwareUpdateConfigurationsClient) getByNameHandleResponse(resp *http.Response) (SoftwareUpdateConfigurationsClientGetByNameResponse, error) {
	result := SoftwareUpdateConfigurationsClientGetByNameResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SoftwareUpdateConfiguration); err != nil {
		return SoftwareUpdateConfigurationsClientGetByNameResponse{}, err
	}
	return result, nil
}

// List - Get all software update configurations for the account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// options - SoftwareUpdateConfigurationsClientListOptions contains the optional parameters for the SoftwareUpdateConfigurationsClient.List
// method.
func (client *SoftwareUpdateConfigurationsClient) List(ctx context.Context, resourceGroupName string, automationAccountName string, options *SoftwareUpdateConfigurationsClientListOptions) (SoftwareUpdateConfigurationsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, automationAccountName, options)
	if err != nil {
		return SoftwareUpdateConfigurationsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SoftwareUpdateConfigurationsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SoftwareUpdateConfigurationsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *SoftwareUpdateConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *SoftwareUpdateConfigurationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/softwareUpdateConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header.Set("clientRequestId", *options.ClientRequestID)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SoftwareUpdateConfigurationsClient) listHandleResponse(resp *http.Response) (SoftwareUpdateConfigurationsClientListResponse, error) {
	result := SoftwareUpdateConfigurationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SoftwareUpdateConfigurationListResult); err != nil {
		return SoftwareUpdateConfigurationsClientListResponse{}, err
	}
	return result, nil
}
