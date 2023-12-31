package compute

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
)

// DedicatedHostsClient is the compute Client
type DedicatedHostsClient struct {
	BaseClient
}

// NewDedicatedHostsClient creates an instance of the DedicatedHostsClient client.
func NewDedicatedHostsClient(subscriptionID string) DedicatedHostsClient {
	return NewDedicatedHostsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDedicatedHostsClientWithBaseURI creates an instance of the DedicatedHostsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDedicatedHostsClientWithBaseURI(baseURI string, subscriptionID string) DedicatedHostsClient {
	return DedicatedHostsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a dedicated host .
// Parameters:
// resourceGroupName - the name of the resource group.
// hostGroupName - the name of the dedicated host group.
// hostName - the name of the dedicated host .
// parameters - parameters supplied to the Create Dedicated Host.
func (client DedicatedHostsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHost) (result DedicatedHostsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHostsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.DedicatedHostProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.DedicatedHostProperties.PlatformFaultDomain", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.DedicatedHostProperties.PlatformFaultDomain", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}},
				}},
				{Target: "parameters.Sku", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("compute.DedicatedHostsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, hostGroupName, hostName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DedicatedHostsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHost) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostGroupName":     autorest.Encode("path", hostGroupName),
		"hostName":          autorest.Encode("path", hostName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2023-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedHostsClient) CreateOrUpdateSender(req *http.Request) (future DedicatedHostsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DedicatedHostsClient) CreateOrUpdateResponder(resp *http.Response) (result DedicatedHost, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a dedicated host.
// Parameters:
// resourceGroupName - the name of the resource group.
// hostGroupName - the name of the dedicated host group.
// hostName - the name of the dedicated host.
func (client DedicatedHostsClient) Delete(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string) (result DedicatedHostsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHostsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, hostGroupName, hostName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DedicatedHostsClient) DeletePreparer(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostGroupName":     autorest.Encode("path", hostGroupName),
		"hostName":          autorest.Encode("path", hostName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2023-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedHostsClient) DeleteSender(req *http.Request) (future DedicatedHostsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DedicatedHostsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves information about a dedicated host.
// Parameters:
// resourceGroupName - the name of the resource group.
// hostGroupName - the name of the dedicated host group.
// hostName - the name of the dedicated host.
// expand - the expand expression to apply on the operation. 'InstanceView' will retrieve the list of instance
// views of the dedicated host. 'UserData' is not supported for dedicated host.
func (client DedicatedHostsClient) Get(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, expand InstanceViewTypes) (result DedicatedHost, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHostsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, hostGroupName, hostName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DedicatedHostsClient) GetPreparer(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, expand InstanceViewTypes) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostGroupName":     autorest.Encode("path", hostGroupName),
		"hostName":          autorest.Encode("path", hostName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2023-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(expand)) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedHostsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DedicatedHostsClient) GetResponder(resp *http.Response) (result DedicatedHost, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAvailableSizes lists all available dedicated host sizes to which the specified dedicated host can be resized.
// NOTE: The dedicated host sizes provided can be used to only scale up the existing dedicated host.
// Parameters:
// resourceGroupName - the name of the resource group.
// hostGroupName - the name of the dedicated host group.
// hostName - the name of the dedicated host.
func (client DedicatedHostsClient) ListAvailableSizes(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string) (result DedicatedHostSizeListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHostsClient.ListAvailableSizes")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: hostGroupName,
			Constraints: []validation.Constraint{{Target: "hostGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: hostName,
			Constraints: []validation.Constraint{{Target: "hostName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("compute.DedicatedHostsClient", "ListAvailableSizes", err.Error())
	}

	req, err := client.ListAvailableSizesPreparer(ctx, resourceGroupName, hostGroupName, hostName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "ListAvailableSizes", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAvailableSizesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "ListAvailableSizes", resp, "Failure sending request")
		return
	}

	result, err = client.ListAvailableSizesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "ListAvailableSizes", resp, "Failure responding to request")
		return
	}

	return
}

// ListAvailableSizesPreparer prepares the ListAvailableSizes request.
func (client DedicatedHostsClient) ListAvailableSizesPreparer(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostGroupName":     autorest.Encode("path", hostGroupName),
		"hostName":          autorest.Encode("path", hostName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2023-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}/hostSizes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAvailableSizesSender sends the ListAvailableSizes request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedHostsClient) ListAvailableSizesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAvailableSizesResponder handles the response to the ListAvailableSizes request. The method always
// closes the http.Response Body.
func (client DedicatedHostsClient) ListAvailableSizesResponder(resp *http.Response) (result DedicatedHostSizeListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByHostGroup lists all of the dedicated hosts in the specified dedicated host group. Use the nextLink property in
// the response to get the next page of dedicated hosts.
// Parameters:
// resourceGroupName - the name of the resource group.
// hostGroupName - the name of the dedicated host group.
func (client DedicatedHostsClient) ListByHostGroup(ctx context.Context, resourceGroupName string, hostGroupName string) (result DedicatedHostListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHostsClient.ListByHostGroup")
		defer func() {
			sc := -1
			if result.dhlr.Response.Response != nil {
				sc = result.dhlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByHostGroupNextResults
	req, err := client.ListByHostGroupPreparer(ctx, resourceGroupName, hostGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "ListByHostGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByHostGroupSender(req)
	if err != nil {
		result.dhlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "ListByHostGroup", resp, "Failure sending request")
		return
	}

	result.dhlr, err = client.ListByHostGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "ListByHostGroup", resp, "Failure responding to request")
		return
	}
	if result.dhlr.hasNextLink() && result.dhlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByHostGroupPreparer prepares the ListByHostGroup request.
func (client DedicatedHostsClient) ListByHostGroupPreparer(ctx context.Context, resourceGroupName string, hostGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostGroupName":     autorest.Encode("path", hostGroupName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2023-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByHostGroupSender sends the ListByHostGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedHostsClient) ListByHostGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByHostGroupResponder handles the response to the ListByHostGroup request. The method always
// closes the http.Response Body.
func (client DedicatedHostsClient) ListByHostGroupResponder(resp *http.Response) (result DedicatedHostListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByHostGroupNextResults retrieves the next set of results, if any.
func (client DedicatedHostsClient) listByHostGroupNextResults(ctx context.Context, lastResults DedicatedHostListResult) (result DedicatedHostListResult, err error) {
	req, err := lastResults.dedicatedHostListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "listByHostGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByHostGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "listByHostGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByHostGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "listByHostGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByHostGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client DedicatedHostsClient) ListByHostGroupComplete(ctx context.Context, resourceGroupName string, hostGroupName string) (result DedicatedHostListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHostsClient.ListByHostGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByHostGroup(ctx, resourceGroupName, hostGroupName)
	return
}

// Restart restart the dedicated host. The operation will complete successfully once the dedicated host has restarted
// and is running. To determine the health of VMs deployed on the dedicated host after the restart check the Resource
// Health Center in the Azure Portal. Please refer to
// https://docs.microsoft.com/azure/service-health/resource-health-overview for more details.
// Parameters:
// resourceGroupName - the name of the resource group.
// hostGroupName - the name of the dedicated host group.
// hostName - the name of the dedicated host.
func (client DedicatedHostsClient) Restart(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string) (result DedicatedHostsRestartFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHostsClient.Restart")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RestartPreparer(ctx, resourceGroupName, hostGroupName, hostName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "Restart", nil, "Failure preparing request")
		return
	}

	result, err = client.RestartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "Restart", result.Response(), "Failure sending request")
		return
	}

	return
}

// RestartPreparer prepares the Restart request.
func (client DedicatedHostsClient) RestartPreparer(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostGroupName":     autorest.Encode("path", hostGroupName),
		"hostName":          autorest.Encode("path", hostName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2023-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}/restart", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RestartSender sends the Restart request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedHostsClient) RestartSender(req *http.Request) (future DedicatedHostsRestartFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// RestartResponder handles the response to the Restart request. The method always
// closes the http.Response Body.
func (client DedicatedHostsClient) RestartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update update a dedicated host .
// Parameters:
// resourceGroupName - the name of the resource group.
// hostGroupName - the name of the dedicated host group.
// hostName - the name of the dedicated host .
// parameters - parameters supplied to the Update Dedicated Host operation.
func (client DedicatedHostsClient) Update(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHostUpdate) (result DedicatedHostsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHostsClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, hostGroupName, hostName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client DedicatedHostsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHostUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostGroupName":     autorest.Encode("path", hostGroupName),
		"hostName":          autorest.Encode("path", hostName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2023-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedHostsClient) UpdateSender(req *http.Request) (future DedicatedHostsUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client DedicatedHostsClient) UpdateResponder(resp *http.Response) (result DedicatedHost, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
