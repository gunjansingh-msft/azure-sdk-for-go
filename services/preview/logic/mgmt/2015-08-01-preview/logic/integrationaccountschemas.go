package logic

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

// IntegrationAccountSchemasClient is the REST API for Azure Logic Apps.
type IntegrationAccountSchemasClient struct {
	BaseClient
}

// NewIntegrationAccountSchemasClient creates an instance of the IntegrationAccountSchemasClient client.
func NewIntegrationAccountSchemasClient(subscriptionID string) IntegrationAccountSchemasClient {
	return NewIntegrationAccountSchemasClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIntegrationAccountSchemasClientWithBaseURI creates an instance of the IntegrationAccountSchemasClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewIntegrationAccountSchemasClientWithBaseURI(baseURI string, subscriptionID string) IntegrationAccountSchemasClient {
	return IntegrationAccountSchemasClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an integration account schema.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
// schemaName - the integration account schema name.
// schema - the integration account schema.
func (client IntegrationAccountSchemasClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string, schema IntegrationAccountSchema) (result IntegrationAccountSchema, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountSchemasClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, integrationAccountName, schemaName, schema)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountSchemasClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountSchemasClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountSchemasClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client IntegrationAccountSchemasClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string, schema IntegrationAccountSchema) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"schemaName":             autorest.Encode("path", schemaName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}", pathParameters),
		autorest.WithJSON(schema),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountSchemasClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client IntegrationAccountSchemasClient) CreateOrUpdateResponder(resp *http.Response) (result IntegrationAccountSchema, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an integration account schema.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
// schemaName - the integration account schema name.
func (client IntegrationAccountSchemasClient) Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountSchemasClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, integrationAccountName, schemaName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountSchemasClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountSchemasClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountSchemasClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client IntegrationAccountSchemasClient) DeletePreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"schemaName":             autorest.Encode("path", schemaName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountSchemasClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client IntegrationAccountSchemasClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets an integration account schema.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
// schemaName - the integration account schema name.
func (client IntegrationAccountSchemasClient) Get(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string) (result IntegrationAccountSchema, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountSchemasClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, integrationAccountName, schemaName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountSchemasClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountSchemasClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountSchemasClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client IntegrationAccountSchemasClient) GetPreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"schemaName":             autorest.Encode("path", schemaName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountSchemasClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IntegrationAccountSchemasClient) GetResponder(resp *http.Response) (result IntegrationAccountSchema, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of integration account schemas.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
// top - the number of items to be included in the result.
// filter - the filter to apply on the operation.
func (client IntegrationAccountSchemasClient) List(ctx context.Context, resourceGroupName string, integrationAccountName string, top *int32, filter string) (result IntegrationAccountSchemaListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountSchemasClient.List")
		defer func() {
			sc := -1
			if result.iaslr.Response.Response != nil {
				sc = result.iaslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, integrationAccountName, top, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountSchemasClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.iaslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountSchemasClient", "List", resp, "Failure sending request")
		return
	}

	result.iaslr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountSchemasClient", "List", resp, "Failure responding to request")
		return
	}
	if result.iaslr.hasNextLink() && result.iaslr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client IntegrationAccountSchemasClient) ListPreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, top *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountSchemasClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client IntegrationAccountSchemasClient) ListResponder(resp *http.Response) (result IntegrationAccountSchemaListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client IntegrationAccountSchemasClient) listNextResults(ctx context.Context, lastResults IntegrationAccountSchemaListResult) (result IntegrationAccountSchemaListResult, err error) {
	req, err := lastResults.integrationAccountSchemaListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.IntegrationAccountSchemasClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic.IntegrationAccountSchemasClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountSchemasClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client IntegrationAccountSchemasClient) ListComplete(ctx context.Context, resourceGroupName string, integrationAccountName string, top *int32, filter string) (result IntegrationAccountSchemaListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountSchemasClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, integrationAccountName, top, filter)
	return
}
