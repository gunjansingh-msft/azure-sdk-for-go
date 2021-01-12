package servicemap

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PortsClient is the service Map API Reference
type PortsClient struct {
	BaseClient
}

// NewPortsClient creates an instance of the PortsClient client.
func NewPortsClient(subscriptionID string) PortsClient {
	return NewPortsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPortsClientWithBaseURI creates an instance of the PortsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPortsClientWithBaseURI(baseURI string, subscriptionID string) PortsClient {
	return PortsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get returns the specified port. The port must be live during the specified time interval. If the port is not live
// during the interval, status 404 (Not Found) is returned.
// Parameters:
// resourceGroupName - resource group name within the specified subscriptionId.
// workspaceName - OMS workspace containing the resources of interest.
// machineName - machine resource name.
// portName - port resource name.
// startTime - UTC date and time specifying the start time of an interval. When not specified the service uses
// DateTime.UtcNow - 10m
// endTime - UTC date and time specifying the end time of an interval. When not specified the service uses
// DateTime.UtcNow
func (client PortsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, machineName string, portName string, startTime *date.Time, endTime *date.Time) (result Port, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PortsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_-]+`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_][a-zA-Z0-9_-]+[a-zA-Z0-9_]`, Chain: nil}}},
		{TargetValue: machineName,
			Constraints: []validation.Constraint{{Target: "machineName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "machineName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: portName,
			Constraints: []validation.Constraint{{Target: "portName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "portName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicemap.PortsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, machineName, portName, startTime, endTime)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PortsClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, machineName string, portName string, startTime *date.Time, endTime *date.Time) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"machineName":       autorest.Encode("path", machineName),
		"portName":          autorest.Encode("path", portName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if startTime != nil {
		queryParameters["startTime"] = autorest.Encode("query", *startTime)
	}
	if endTime != nil {
		queryParameters["endTime"] = autorest.Encode("query", *endTime)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/ports/{portName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PortsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PortsClient) GetResponder(resp *http.Response) (result Port, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetLiveness obtains the liveness status of the port during the specified time interval.
// Parameters:
// resourceGroupName - resource group name within the specified subscriptionId.
// workspaceName - OMS workspace containing the resources of interest.
// machineName - machine resource name.
// portName - port resource name.
// startTime - UTC date and time specifying the start time of an interval. When not specified the service uses
// DateTime.UtcNow - 10m
// endTime - UTC date and time specifying the end time of an interval. When not specified the service uses
// DateTime.UtcNow
func (client PortsClient) GetLiveness(ctx context.Context, resourceGroupName string, workspaceName string, machineName string, portName string, startTime *date.Time, endTime *date.Time) (result Liveness, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PortsClient.GetLiveness")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_-]+`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_][a-zA-Z0-9_-]+[a-zA-Z0-9_]`, Chain: nil}}},
		{TargetValue: machineName,
			Constraints: []validation.Constraint{{Target: "machineName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "machineName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: portName,
			Constraints: []validation.Constraint{{Target: "portName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "portName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicemap.PortsClient", "GetLiveness", err.Error())
	}

	req, err := client.GetLivenessPreparer(ctx, resourceGroupName, workspaceName, machineName, portName, startTime, endTime)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "GetLiveness", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetLivenessSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "GetLiveness", resp, "Failure sending request")
		return
	}

	result, err = client.GetLivenessResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "GetLiveness", resp, "Failure responding to request")
		return
	}

	return
}

// GetLivenessPreparer prepares the GetLiveness request.
func (client PortsClient) GetLivenessPreparer(ctx context.Context, resourceGroupName string, workspaceName string, machineName string, portName string, startTime *date.Time, endTime *date.Time) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"machineName":       autorest.Encode("path", machineName),
		"portName":          autorest.Encode("path", portName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if startTime != nil {
		queryParameters["startTime"] = autorest.Encode("query", *startTime)
	}
	if endTime != nil {
		queryParameters["endTime"] = autorest.Encode("query", *endTime)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/ports/{portName}/liveness", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetLivenessSender sends the GetLiveness request. The method will close the
// http.Response Body if it receives an error.
func (client PortsClient) GetLivenessSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetLivenessResponder handles the response to the GetLiveness request. The method always
// closes the http.Response Body.
func (client PortsClient) GetLivenessResponder(resp *http.Response) (result Liveness, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAcceptingProcesses returns a collection of processes accepting on the specified port
// Parameters:
// resourceGroupName - resource group name within the specified subscriptionId.
// workspaceName - OMS workspace containing the resources of interest.
// machineName - machine resource name.
// portName - port resource name.
// startTime - UTC date and time specifying the start time of an interval. When not specified the service uses
// DateTime.UtcNow - 10m
// endTime - UTC date and time specifying the end time of an interval. When not specified the service uses
// DateTime.UtcNow
func (client PortsClient) ListAcceptingProcesses(ctx context.Context, resourceGroupName string, workspaceName string, machineName string, portName string, startTime *date.Time, endTime *date.Time) (result ProcessCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PortsClient.ListAcceptingProcesses")
		defer func() {
			sc := -1
			if result.pc.Response.Response != nil {
				sc = result.pc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_-]+`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_][a-zA-Z0-9_-]+[a-zA-Z0-9_]`, Chain: nil}}},
		{TargetValue: machineName,
			Constraints: []validation.Constraint{{Target: "machineName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "machineName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: portName,
			Constraints: []validation.Constraint{{Target: "portName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "portName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicemap.PortsClient", "ListAcceptingProcesses", err.Error())
	}

	result.fn = client.listAcceptingProcessesNextResults
	req, err := client.ListAcceptingProcessesPreparer(ctx, resourceGroupName, workspaceName, machineName, portName, startTime, endTime)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "ListAcceptingProcesses", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAcceptingProcessesSender(req)
	if err != nil {
		result.pc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "ListAcceptingProcesses", resp, "Failure sending request")
		return
	}

	result.pc, err = client.ListAcceptingProcessesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "ListAcceptingProcesses", resp, "Failure responding to request")
		return
	}
	if result.pc.hasNextLink() && result.pc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListAcceptingProcessesPreparer prepares the ListAcceptingProcesses request.
func (client PortsClient) ListAcceptingProcessesPreparer(ctx context.Context, resourceGroupName string, workspaceName string, machineName string, portName string, startTime *date.Time, endTime *date.Time) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"machineName":       autorest.Encode("path", machineName),
		"portName":          autorest.Encode("path", portName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if startTime != nil {
		queryParameters["startTime"] = autorest.Encode("query", *startTime)
	}
	if endTime != nil {
		queryParameters["endTime"] = autorest.Encode("query", *endTime)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/ports/{portName}/acceptingProcesses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAcceptingProcessesSender sends the ListAcceptingProcesses request. The method will close the
// http.Response Body if it receives an error.
func (client PortsClient) ListAcceptingProcessesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAcceptingProcessesResponder handles the response to the ListAcceptingProcesses request. The method always
// closes the http.Response Body.
func (client PortsClient) ListAcceptingProcessesResponder(resp *http.Response) (result ProcessCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAcceptingProcessesNextResults retrieves the next set of results, if any.
func (client PortsClient) listAcceptingProcessesNextResults(ctx context.Context, lastResults ProcessCollection) (result ProcessCollection, err error) {
	req, err := lastResults.processCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicemap.PortsClient", "listAcceptingProcessesNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAcceptingProcessesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicemap.PortsClient", "listAcceptingProcessesNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAcceptingProcessesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "listAcceptingProcessesNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAcceptingProcessesComplete enumerates all values, automatically crossing page boundaries as required.
func (client PortsClient) ListAcceptingProcessesComplete(ctx context.Context, resourceGroupName string, workspaceName string, machineName string, portName string, startTime *date.Time, endTime *date.Time) (result ProcessCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PortsClient.ListAcceptingProcesses")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAcceptingProcesses(ctx, resourceGroupName, workspaceName, machineName, portName, startTime, endTime)
	return
}

// ListConnections returns a collection of connections established via the specified port.
// Parameters:
// resourceGroupName - resource group name within the specified subscriptionId.
// workspaceName - OMS workspace containing the resources of interest.
// machineName - machine resource name.
// portName - port resource name.
// startTime - UTC date and time specifying the start time of an interval. When not specified the service uses
// DateTime.UtcNow - 10m
// endTime - UTC date and time specifying the end time of an interval. When not specified the service uses
// DateTime.UtcNow
func (client PortsClient) ListConnections(ctx context.Context, resourceGroupName string, workspaceName string, machineName string, portName string, startTime *date.Time, endTime *date.Time) (result ConnectionCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PortsClient.ListConnections")
		defer func() {
			sc := -1
			if result.cc.Response.Response != nil {
				sc = result.cc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_-]+`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_][a-zA-Z0-9_-]+[a-zA-Z0-9_]`, Chain: nil}}},
		{TargetValue: machineName,
			Constraints: []validation.Constraint{{Target: "machineName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "machineName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: portName,
			Constraints: []validation.Constraint{{Target: "portName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "portName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicemap.PortsClient", "ListConnections", err.Error())
	}

	result.fn = client.listConnectionsNextResults
	req, err := client.ListConnectionsPreparer(ctx, resourceGroupName, workspaceName, machineName, portName, startTime, endTime)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "ListConnections", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListConnectionsSender(req)
	if err != nil {
		result.cc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "ListConnections", resp, "Failure sending request")
		return
	}

	result.cc, err = client.ListConnectionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "ListConnections", resp, "Failure responding to request")
		return
	}
	if result.cc.hasNextLink() && result.cc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListConnectionsPreparer prepares the ListConnections request.
func (client PortsClient) ListConnectionsPreparer(ctx context.Context, resourceGroupName string, workspaceName string, machineName string, portName string, startTime *date.Time, endTime *date.Time) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"machineName":       autorest.Encode("path", machineName),
		"portName":          autorest.Encode("path", portName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if startTime != nil {
		queryParameters["startTime"] = autorest.Encode("query", *startTime)
	}
	if endTime != nil {
		queryParameters["endTime"] = autorest.Encode("query", *endTime)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/ports/{portName}/connections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListConnectionsSender sends the ListConnections request. The method will close the
// http.Response Body if it receives an error.
func (client PortsClient) ListConnectionsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListConnectionsResponder handles the response to the ListConnections request. The method always
// closes the http.Response Body.
func (client PortsClient) ListConnectionsResponder(resp *http.Response) (result ConnectionCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listConnectionsNextResults retrieves the next set of results, if any.
func (client PortsClient) listConnectionsNextResults(ctx context.Context, lastResults ConnectionCollection) (result ConnectionCollection, err error) {
	req, err := lastResults.connectionCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicemap.PortsClient", "listConnectionsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListConnectionsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicemap.PortsClient", "listConnectionsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListConnectionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "listConnectionsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListConnectionsComplete enumerates all values, automatically crossing page boundaries as required.
func (client PortsClient) ListConnectionsComplete(ctx context.Context, resourceGroupName string, workspaceName string, machineName string, portName string, startTime *date.Time, endTime *date.Time) (result ConnectionCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PortsClient.ListConnections")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListConnections(ctx, resourceGroupName, workspaceName, machineName, portName, startTime, endTime)
	return
}
