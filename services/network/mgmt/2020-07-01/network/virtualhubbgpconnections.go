package network

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// VirtualHubBgpConnectionsClient is the network Client
type VirtualHubBgpConnectionsClient struct {
	BaseClient
}

// NewVirtualHubBgpConnectionsClient creates an instance of the VirtualHubBgpConnectionsClient client.
func NewVirtualHubBgpConnectionsClient(subscriptionID string) VirtualHubBgpConnectionsClient {
	return NewVirtualHubBgpConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualHubBgpConnectionsClientWithBaseURI creates an instance of the VirtualHubBgpConnectionsClient client using
// a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewVirtualHubBgpConnectionsClientWithBaseURI(baseURI string, subscriptionID string) VirtualHubBgpConnectionsClient {
	return VirtualHubBgpConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List retrieves the details of all VirtualHubBgpConnections.
// Parameters:
// resourceGroupName - the resource group name of the VirtualHub.
// virtualHubName - the name of the VirtualHub.
func (client VirtualHubBgpConnectionsClient) List(ctx context.Context, resourceGroupName string, virtualHubName string) (result ListVirtualHubBgpConnectionResultsPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualHubBgpConnectionsClient.List")
		defer func() {
			sc := -1
			if result.lvhbcr.Response.Response != nil {
				sc = result.lvhbcr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, virtualHubName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubBgpConnectionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lvhbcr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualHubBgpConnectionsClient", "List", resp, "Failure sending request")
		return
	}

	result.lvhbcr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubBgpConnectionsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.lvhbcr.hasNextLink() && result.lvhbcr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualHubBgpConnectionsClient) ListPreparer(ctx context.Context, resourceGroupName string, virtualHubName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/bgpConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualHubBgpConnectionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualHubBgpConnectionsClient) ListResponder(resp *http.Response) (result ListVirtualHubBgpConnectionResults, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client VirtualHubBgpConnectionsClient) listNextResults(ctx context.Context, lastResults ListVirtualHubBgpConnectionResults) (result ListVirtualHubBgpConnectionResults, err error) {
	req, err := lastResults.listVirtualHubBgpConnectionResultsPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualHubBgpConnectionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualHubBgpConnectionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubBgpConnectionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualHubBgpConnectionsClient) ListComplete(ctx context.Context, resourceGroupName string, virtualHubName string) (result ListVirtualHubBgpConnectionResultsIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualHubBgpConnectionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, virtualHubName)
	return
}

// ListAdvertisedRoutes retrieves a list of routes the virtual hub bgp connection is advertising to the specified peer.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the virtual hub.
// connectionName - the name of the virtual hub bgp connection.
func (client VirtualHubBgpConnectionsClient) ListAdvertisedRoutes(ctx context.Context, resourceGroupName string, hubName string, connectionName string) (result VirtualHubBgpConnectionsListAdvertisedRoutesFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualHubBgpConnectionsClient.ListAdvertisedRoutes")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListAdvertisedRoutesPreparer(ctx, resourceGroupName, hubName, connectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubBgpConnectionsClient", "ListAdvertisedRoutes", nil, "Failure preparing request")
		return
	}

	result, err = client.ListAdvertisedRoutesSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubBgpConnectionsClient", "ListAdvertisedRoutes", nil, "Failure sending request")
		return
	}

	return
}

// ListAdvertisedRoutesPreparer prepares the ListAdvertisedRoutes request.
func (client VirtualHubBgpConnectionsClient) ListAdvertisedRoutesPreparer(ctx context.Context, resourceGroupName string, hubName string, connectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":    autorest.Encode("path", connectionName),
		"hubName":           autorest.Encode("path", hubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{hubName}/bgpConnections/{connectionName}/advertisedRoutes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAdvertisedRoutesSender sends the ListAdvertisedRoutes request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualHubBgpConnectionsClient) ListAdvertisedRoutesSender(req *http.Request) (future VirtualHubBgpConnectionsListAdvertisedRoutesFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client VirtualHubBgpConnectionsClient) (prl PeerRouteList, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VirtualHubBgpConnectionsListAdvertisedRoutesFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("network.VirtualHubBgpConnectionsListAdvertisedRoutesFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		if prl.Response.Response, err = future.GetResult(sender); err == nil && prl.Response.Response.StatusCode != http.StatusNoContent {
			prl, err = client.ListAdvertisedRoutesResponder(prl.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "network.VirtualHubBgpConnectionsListAdvertisedRoutesFuture", "Result", prl.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// ListAdvertisedRoutesResponder handles the response to the ListAdvertisedRoutes request. The method always
// closes the http.Response Body.
func (client VirtualHubBgpConnectionsClient) ListAdvertisedRoutesResponder(resp *http.Response) (result PeerRouteList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListLearnedRoutes retrieves a list of routes the virtual hub bgp connection has learned.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the virtual hub.
// connectionName - the name of the virtual hub bgp connection.
func (client VirtualHubBgpConnectionsClient) ListLearnedRoutes(ctx context.Context, resourceGroupName string, hubName string, connectionName string) (result VirtualHubBgpConnectionsListLearnedRoutesFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualHubBgpConnectionsClient.ListLearnedRoutes")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListLearnedRoutesPreparer(ctx, resourceGroupName, hubName, connectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubBgpConnectionsClient", "ListLearnedRoutes", nil, "Failure preparing request")
		return
	}

	result, err = client.ListLearnedRoutesSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubBgpConnectionsClient", "ListLearnedRoutes", nil, "Failure sending request")
		return
	}

	return
}

// ListLearnedRoutesPreparer prepares the ListLearnedRoutes request.
func (client VirtualHubBgpConnectionsClient) ListLearnedRoutesPreparer(ctx context.Context, resourceGroupName string, hubName string, connectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":    autorest.Encode("path", connectionName),
		"hubName":           autorest.Encode("path", hubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{hubName}/bgpConnections/{connectionName}/learnedRoutes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListLearnedRoutesSender sends the ListLearnedRoutes request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualHubBgpConnectionsClient) ListLearnedRoutesSender(req *http.Request) (future VirtualHubBgpConnectionsListLearnedRoutesFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client VirtualHubBgpConnectionsClient) (prl PeerRouteList, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VirtualHubBgpConnectionsListLearnedRoutesFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("network.VirtualHubBgpConnectionsListLearnedRoutesFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		if prl.Response.Response, err = future.GetResult(sender); err == nil && prl.Response.Response.StatusCode != http.StatusNoContent {
			prl, err = client.ListLearnedRoutesResponder(prl.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "network.VirtualHubBgpConnectionsListLearnedRoutesFuture", "Result", prl.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// ListLearnedRoutesResponder handles the response to the ListLearnedRoutes request. The method always
// closes the http.Response Body.
func (client VirtualHubBgpConnectionsClient) ListLearnedRoutesResponder(resp *http.Response) (result PeerRouteList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
