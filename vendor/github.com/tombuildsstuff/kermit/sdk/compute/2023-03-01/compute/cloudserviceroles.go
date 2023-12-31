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
	"github.com/Azure/go-autorest/tracing"
)

// CloudServiceRolesClient is the compute Client
type CloudServiceRolesClient struct {
	BaseClient
}

// NewCloudServiceRolesClient creates an instance of the CloudServiceRolesClient client.
func NewCloudServiceRolesClient(subscriptionID string) CloudServiceRolesClient {
	return NewCloudServiceRolesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCloudServiceRolesClientWithBaseURI creates an instance of the CloudServiceRolesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewCloudServiceRolesClientWithBaseURI(baseURI string, subscriptionID string) CloudServiceRolesClient {
	return CloudServiceRolesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a role from a cloud service.
// Parameters:
// roleName - name of the role.
// resourceGroupName - name of the resource group.
// cloudServiceName - name of the cloud service.
func (client CloudServiceRolesClient) Get(ctx context.Context, roleName string, resourceGroupName string, cloudServiceName string) (result CloudServiceRole, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudServiceRolesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, roleName, resourceGroupName, cloudServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRolesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRolesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRolesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client CloudServiceRolesClient) GetPreparer(ctx context.Context, roleName string, resourceGroupName string, cloudServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cloudServiceName":  autorest.Encode("path", cloudServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"roleName":          autorest.Encode("path", roleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-09-04"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roles/{roleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CloudServiceRolesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CloudServiceRolesClient) GetResponder(resp *http.Response) (result CloudServiceRole, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of all roles in a cloud service. Use nextLink property in the response to get the next page of
// roles. Do this till nextLink is null to fetch all the roles.
// Parameters:
// resourceGroupName - name of the resource group.
// cloudServiceName - name of the cloud service.
func (client CloudServiceRolesClient) List(ctx context.Context, resourceGroupName string, cloudServiceName string) (result CloudServiceRoleListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudServiceRolesClient.List")
		defer func() {
			sc := -1
			if result.csrlr.Response.Response != nil {
				sc = result.csrlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, cloudServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRolesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.csrlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRolesClient", "List", resp, "Failure sending request")
		return
	}

	result.csrlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRolesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.csrlr.hasNextLink() && result.csrlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client CloudServiceRolesClient) ListPreparer(ctx context.Context, resourceGroupName string, cloudServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cloudServiceName":  autorest.Encode("path", cloudServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-09-04"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roles", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client CloudServiceRolesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client CloudServiceRolesClient) ListResponder(resp *http.Response) (result CloudServiceRoleListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client CloudServiceRolesClient) listNextResults(ctx context.Context, lastResults CloudServiceRoleListResult) (result CloudServiceRoleListResult, err error) {
	req, err := lastResults.cloudServiceRoleListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.CloudServiceRolesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.CloudServiceRolesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRolesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client CloudServiceRolesClient) ListComplete(ctx context.Context, resourceGroupName string, cloudServiceName string) (result CloudServiceRoleListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudServiceRolesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, cloudServiceName)
	return
}
