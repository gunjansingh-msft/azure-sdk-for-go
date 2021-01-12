package consumption

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// MarketplacesByEnrollmentAccountsClient is the consumption management client provides access to consumption resources
// for Azure Enterprise Subscriptions.
type MarketplacesByEnrollmentAccountsClient struct {
	BaseClient
}

// NewMarketplacesByEnrollmentAccountsClient creates an instance of the MarketplacesByEnrollmentAccountsClient client.
func NewMarketplacesByEnrollmentAccountsClient(subscriptionID string) MarketplacesByEnrollmentAccountsClient {
	return NewMarketplacesByEnrollmentAccountsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMarketplacesByEnrollmentAccountsClientWithBaseURI creates an instance of the
// MarketplacesByEnrollmentAccountsClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewMarketplacesByEnrollmentAccountsClientWithBaseURI(baseURI string, subscriptionID string) MarketplacesByEnrollmentAccountsClient {
	return MarketplacesByEnrollmentAccountsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists the marketplaces for a scope by enrollmentAccountId and current billing period. Marketplaces are
// available via this API only for May 1, 2014 or later.
// Parameters:
// enrollmentAccountID - enrollmentAccount ID
// filter - may be used to filter marketplaces by properties/usageEnd (Utc time), properties/usageStart (Utc
// time), properties/resourceGroup, properties/instanceName or properties/instanceId. The filter supports 'eq',
// 'lt', 'gt', 'le', 'ge', and 'and'. It does not currently support 'ne', 'or', or 'not'.
// top - may be used to limit the number of results to the most recent N marketplaces.
// skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls.
func (client MarketplacesByEnrollmentAccountsClient) List(ctx context.Context, enrollmentAccountID string, filter string, top *int32, skiptoken string) (result MarketplacesListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MarketplacesByEnrollmentAccountsClient.List")
		defer func() {
			sc := -1
			if result.mlr.Response.Response != nil {
				sc = result.mlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("consumption.MarketplacesByEnrollmentAccountsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, enrollmentAccountID, filter, top, skiptoken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesByEnrollmentAccountsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.mlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesByEnrollmentAccountsClient", "List", resp, "Failure sending request")
		return
	}

	result.mlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesByEnrollmentAccountsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.mlr.hasNextLink() && result.mlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client MarketplacesByEnrollmentAccountsClient) ListPreparer(ctx context.Context, enrollmentAccountID string, filter string, top *int32, skiptoken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"enrollmentAccountId": autorest.Encode("path", enrollmentAccountID),
	}

	const APIVersion = "2018-03-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/enrollmentAccounts/{enrollmentAccountId}/providers/Microsoft.Consumption/marketplaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client MarketplacesByEnrollmentAccountsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MarketplacesByEnrollmentAccountsClient) ListResponder(resp *http.Response) (result MarketplacesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client MarketplacesByEnrollmentAccountsClient) listNextResults(ctx context.Context, lastResults MarketplacesListResult) (result MarketplacesListResult, err error) {
	req, err := lastResults.marketplacesListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "consumption.MarketplacesByEnrollmentAccountsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "consumption.MarketplacesByEnrollmentAccountsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesByEnrollmentAccountsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client MarketplacesByEnrollmentAccountsClient) ListComplete(ctx context.Context, enrollmentAccountID string, filter string, top *int32, skiptoken string) (result MarketplacesListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MarketplacesByEnrollmentAccountsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, enrollmentAccountID, filter, top, skiptoken)
	return
}

// ListByBillingPeriod lists the marketplaces for a scope by billing period and enrollmentAccountId. Marketplaces are
// available via this API only for May 1, 2014 or later.
// Parameters:
// enrollmentAccountID - enrollmentAccount ID
// billingPeriodName - billing Period Name.
// filter - may be used to filter marketplaces by properties/usageEnd (Utc time), properties/usageStart (Utc
// time), properties/resourceGroup, properties/instanceName or properties/instanceId. The filter supports 'eq',
// 'lt', 'gt', 'le', 'ge', and 'and'. It does not currently support 'ne', 'or', or 'not'.
// top - may be used to limit the number of results to the most recent N marketplaces.
// skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls.
func (client MarketplacesByEnrollmentAccountsClient) ListByBillingPeriod(ctx context.Context, enrollmentAccountID string, billingPeriodName string, filter string, top *int32, skiptoken string) (result MarketplacesListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MarketplacesByEnrollmentAccountsClient.ListByBillingPeriod")
		defer func() {
			sc := -1
			if result.mlr.Response.Response != nil {
				sc = result.mlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("consumption.MarketplacesByEnrollmentAccountsClient", "ListByBillingPeriod", err.Error())
	}

	result.fn = client.listByBillingPeriodNextResults
	req, err := client.ListByBillingPeriodPreparer(ctx, enrollmentAccountID, billingPeriodName, filter, top, skiptoken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesByEnrollmentAccountsClient", "ListByBillingPeriod", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingPeriodSender(req)
	if err != nil {
		result.mlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesByEnrollmentAccountsClient", "ListByBillingPeriod", resp, "Failure sending request")
		return
	}

	result.mlr, err = client.ListByBillingPeriodResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesByEnrollmentAccountsClient", "ListByBillingPeriod", resp, "Failure responding to request")
		return
	}
	if result.mlr.hasNextLink() && result.mlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByBillingPeriodPreparer prepares the ListByBillingPeriod request.
func (client MarketplacesByEnrollmentAccountsClient) ListByBillingPeriodPreparer(ctx context.Context, enrollmentAccountID string, billingPeriodName string, filter string, top *int32, skiptoken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingPeriodName":   autorest.Encode("path", billingPeriodName),
		"enrollmentAccountId": autorest.Encode("path", enrollmentAccountID),
	}

	const APIVersion = "2018-03-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/enrollmentAccounts/{enrollmentAccountId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}/providers/Microsoft.Consumption/marketplaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingPeriodSender sends the ListByBillingPeriod request. The method will close the
// http.Response Body if it receives an error.
func (client MarketplacesByEnrollmentAccountsClient) ListByBillingPeriodSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingPeriodResponder handles the response to the ListByBillingPeriod request. The method always
// closes the http.Response Body.
func (client MarketplacesByEnrollmentAccountsClient) ListByBillingPeriodResponder(resp *http.Response) (result MarketplacesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingPeriodNextResults retrieves the next set of results, if any.
func (client MarketplacesByEnrollmentAccountsClient) listByBillingPeriodNextResults(ctx context.Context, lastResults MarketplacesListResult) (result MarketplacesListResult, err error) {
	req, err := lastResults.marketplacesListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "consumption.MarketplacesByEnrollmentAccountsClient", "listByBillingPeriodNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingPeriodSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "consumption.MarketplacesByEnrollmentAccountsClient", "listByBillingPeriodNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingPeriodResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesByEnrollmentAccountsClient", "listByBillingPeriodNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingPeriodComplete enumerates all values, automatically crossing page boundaries as required.
func (client MarketplacesByEnrollmentAccountsClient) ListByBillingPeriodComplete(ctx context.Context, enrollmentAccountID string, billingPeriodName string, filter string, top *int32, skiptoken string) (result MarketplacesListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MarketplacesByEnrollmentAccountsClient.ListByBillingPeriod")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingPeriod(ctx, enrollmentAccountID, billingPeriodName, filter, top, skiptoken)
	return
}
