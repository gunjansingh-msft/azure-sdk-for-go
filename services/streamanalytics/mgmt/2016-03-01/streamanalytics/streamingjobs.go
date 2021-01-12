package streamanalytics

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

// StreamingJobsClient is the stream Analytics Client
type StreamingJobsClient struct {
	BaseClient
}

// NewStreamingJobsClient creates an instance of the StreamingJobsClient client.
func NewStreamingJobsClient(subscriptionID string) StreamingJobsClient {
	return NewStreamingJobsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewStreamingJobsClientWithBaseURI creates an instance of the StreamingJobsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewStreamingJobsClientWithBaseURI(baseURI string, subscriptionID string) StreamingJobsClient {
	return StreamingJobsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrReplace creates a streaming job or replaces an already existing streaming job.
// Parameters:
// streamingJob - the definition of the streaming job that will be used to create a new streaming job or
// replace the existing one.
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// jobName - the name of the streaming job.
// ifMatch - the ETag of the streaming job. Omit this value to always overwrite the current record set. Specify
// the last-seen ETag value to prevent accidentally overwriting concurrent changes.
// ifNoneMatch - set to '*' to allow a new streaming job to be created, but to prevent updating an existing
// record set. Other values will result in a 412 Pre-condition Failed response.
func (client StreamingJobsClient) CreateOrReplace(ctx context.Context, streamingJob StreamingJob, resourceGroupName string, jobName string, ifMatch string, ifNoneMatch string) (result StreamingJobsCreateOrReplaceFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingJobsClient.CreateOrReplace")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrReplacePreparer(ctx, streamingJob, resourceGroupName, jobName, ifMatch, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "CreateOrReplace", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrReplaceSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "CreateOrReplace", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrReplacePreparer prepares the CreateOrReplace request.
func (client StreamingJobsClient) CreateOrReplacePreparer(ctx context.Context, streamingJob StreamingJob, resourceGroupName string, jobName string, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}", pathParameters),
		autorest.WithJSON(streamingJob),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrReplaceSender sends the CreateOrReplace request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingJobsClient) CreateOrReplaceSender(req *http.Request) (future StreamingJobsCreateOrReplaceFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client StreamingJobsClient) (sj StreamingJob, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsCreateOrReplaceFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("streamanalytics.StreamingJobsCreateOrReplaceFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		if sj.Response.Response, err = future.GetResult(sender); err == nil && sj.Response.Response.StatusCode != http.StatusNoContent {
			sj, err = client.CreateOrReplaceResponder(sj.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsCreateOrReplaceFuture", "Result", sj.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrReplaceResponder handles the response to the CreateOrReplace request. The method always
// closes the http.Response Body.
func (client StreamingJobsClient) CreateOrReplaceResponder(resp *http.Response) (result StreamingJob, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a streaming job.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// jobName - the name of the streaming job.
func (client StreamingJobsClient) Delete(ctx context.Context, resourceGroupName string, jobName string) (result StreamingJobsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingJobsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, jobName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client StreamingJobsClient) DeletePreparer(ctx context.Context, resourceGroupName string, jobName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingJobsClient) DeleteSender(req *http.Request) (future StreamingJobsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client StreamingJobsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("streamanalytics.StreamingJobsDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client StreamingJobsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets details about the specified streaming job.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// jobName - the name of the streaming job.
// expand - the $expand OData query parameter. This is a comma-separated list of additional streaming job
// properties to include in the response, beyond the default set returned when this parameter is absent. The
// default set is all streaming job properties other than 'inputs', 'transformation', 'outputs', and
// 'functions'.
func (client StreamingJobsClient) Get(ctx context.Context, resourceGroupName string, jobName string, expand string) (result StreamingJob, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingJobsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, jobName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client StreamingJobsClient) GetPreparer(ctx context.Context, resourceGroupName string, jobName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingJobsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client StreamingJobsClient) GetResponder(resp *http.Response) (result StreamingJob, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all of the streaming jobs in the given subscription.
// Parameters:
// expand - the $expand OData query parameter. This is a comma-separated list of additional streaming job
// properties to include in the response, beyond the default set returned when this parameter is absent. The
// default set is all streaming job properties other than 'inputs', 'transformation', 'outputs', and
// 'functions'.
func (client StreamingJobsClient) List(ctx context.Context, expand string) (result StreamingJobListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingJobsClient.List")
		defer func() {
			sc := -1
			if result.sjlr.Response.Response != nil {
				sc = result.sjlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.sjlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "List", resp, "Failure sending request")
		return
	}

	result.sjlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.sjlr.hasNextLink() && result.sjlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client StreamingJobsClient) ListPreparer(ctx context.Context, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.StreamAnalytics/streamingjobs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingJobsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client StreamingJobsClient) ListResponder(resp *http.Response) (result StreamingJobListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client StreamingJobsClient) listNextResults(ctx context.Context, lastResults StreamingJobListResult) (result StreamingJobListResult, err error) {
	req, err := lastResults.streamingJobListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client StreamingJobsClient) ListComplete(ctx context.Context, expand string) (result StreamingJobListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingJobsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, expand)
	return
}

// ListByResourceGroup lists all of the streaming jobs in the specified resource group.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// expand - the $expand OData query parameter. This is a comma-separated list of additional streaming job
// properties to include in the response, beyond the default set returned when this parameter is absent. The
// default set is all streaming job properties other than 'inputs', 'transformation', 'outputs', and
// 'functions'.
func (client StreamingJobsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, expand string) (result StreamingJobListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingJobsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.sjlr.Response.Response != nil {
				sc = result.sjlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.sjlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.sjlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.sjlr.hasNextLink() && result.sjlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client StreamingJobsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingJobsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client StreamingJobsClient) ListByResourceGroupResponder(resp *http.Response) (result StreamingJobListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client StreamingJobsClient) listByResourceGroupNextResults(ctx context.Context, lastResults StreamingJobListResult) (result StreamingJobListResult, err error) {
	req, err := lastResults.streamingJobListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client StreamingJobsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, expand string) (result StreamingJobListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingJobsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, expand)
	return
}

// Start starts a streaming job. Once a job is started it will start processing input events and produce output.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// jobName - the name of the streaming job.
// startJobParameters - parameters applicable to a start streaming job operation.
func (client StreamingJobsClient) Start(ctx context.Context, resourceGroupName string, jobName string, startJobParameters *StartStreamingJobParameters) (result StreamingJobsStartFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingJobsClient.Start")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StartPreparer(ctx, resourceGroupName, jobName, startJobParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "Start", nil, "Failure preparing request")
		return
	}

	result, err = client.StartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "Start", nil, "Failure sending request")
		return
	}

	return
}

// StartPreparer prepares the Start request.
func (client StreamingJobsClient) StartPreparer(ctx context.Context, resourceGroupName string, jobName string, startJobParameters *StartStreamingJobParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if startJobParameters != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(startJobParameters))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingJobsClient) StartSender(req *http.Request) (future StreamingJobsStartFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client StreamingJobsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsStartFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("streamanalytics.StreamingJobsStartFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client StreamingJobsClient) StartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Stop stops a running streaming job. This will cause a running streaming job to stop processing input events and
// producing output.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// jobName - the name of the streaming job.
func (client StreamingJobsClient) Stop(ctx context.Context, resourceGroupName string, jobName string) (result StreamingJobsStopFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingJobsClient.Stop")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StopPreparer(ctx, resourceGroupName, jobName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "Stop", nil, "Failure preparing request")
		return
	}

	result, err = client.StopSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "Stop", nil, "Failure sending request")
		return
	}

	return
}

// StopPreparer prepares the Stop request.
func (client StreamingJobsClient) StopPreparer(ctx context.Context, resourceGroupName string, jobName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingJobsClient) StopSender(req *http.Request) (future StreamingJobsStopFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client StreamingJobsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsStopFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("streamanalytics.StreamingJobsStopFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client StreamingJobsClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update updates an existing streaming job. This can be used to partially update (ie. update one or two properties) a
// streaming job without affecting the rest the job definition.
// Parameters:
// streamingJob - a streaming job object. The properties specified here will overwrite the corresponding
// properties in the existing streaming job (ie. Those properties will be updated). Any properties that are set
// to null here will mean that the corresponding property in the existing input will remain the same and not
// change as a result of this PATCH operation.
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// jobName - the name of the streaming job.
// ifMatch - the ETag of the streaming job. Omit this value to always overwrite the current record set. Specify
// the last-seen ETag value to prevent accidentally overwriting concurrent changes.
func (client StreamingJobsClient) Update(ctx context.Context, streamingJob StreamingJob, resourceGroupName string, jobName string, ifMatch string) (result StreamingJob, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingJobsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, streamingJob, resourceGroupName, jobName, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.StreamingJobsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client StreamingJobsClient) UpdatePreparer(ctx context.Context, streamingJob StreamingJob, resourceGroupName string, jobName string, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}", pathParameters),
		autorest.WithJSON(streamingJob),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingJobsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client StreamingJobsClient) UpdateResponder(resp *http.Response) (result StreamingJob, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
