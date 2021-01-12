package peering

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

// LegacyPeeringsClient is the peering Client
type LegacyPeeringsClient struct {
	BaseClient
}

// NewLegacyPeeringsClient creates an instance of the LegacyPeeringsClient client.
func NewLegacyPeeringsClient(subscriptionID string) LegacyPeeringsClient {
	return NewLegacyPeeringsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLegacyPeeringsClientWithBaseURI creates an instance of the LegacyPeeringsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewLegacyPeeringsClientWithBaseURI(baseURI string, subscriptionID string) LegacyPeeringsClient {
	return LegacyPeeringsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists all of the legacy peerings under the given subscription matching the specified kind and location.
// Parameters:
// peeringLocation - the location of the peering.
// kind - the kind of the peering.
func (client LegacyPeeringsClient) List(ctx context.Context, peeringLocation string, kind string) (result ListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LegacyPeeringsClient.List")
		defer func() {
			sc := -1
			if result.lr.Response.Response != nil {
				sc = result.lr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, peeringLocation, kind)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.LegacyPeeringsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "peering.LegacyPeeringsClient", "List", resp, "Failure sending request")
		return
	}

	result.lr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.LegacyPeeringsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.lr.hasNextLink() && result.lr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client LegacyPeeringsClient) ListPreparer(ctx context.Context, peeringLocation string, kind string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version":     APIVersion,
		"kind":            autorest.Encode("query", kind),
		"peeringLocation": autorest.Encode("query", peeringLocation),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Peering/legacyPeerings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LegacyPeeringsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LegacyPeeringsClient) ListResponder(resp *http.Response) (result ListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client LegacyPeeringsClient) listNextResults(ctx context.Context, lastResults ListResult) (result ListResult, err error) {
	req, err := lastResults.listResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "peering.LegacyPeeringsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "peering.LegacyPeeringsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.LegacyPeeringsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client LegacyPeeringsClient) ListComplete(ctx context.Context, peeringLocation string, kind string) (result ListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LegacyPeeringsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, peeringLocation, kind)
	return
}
