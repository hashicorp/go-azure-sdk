package mhsmlistprivateendpointconnections

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MHSMPrivateEndpointConnectionsListByResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]MHSMPrivateEndpointConnection
}

type MHSMPrivateEndpointConnectionsListByResourceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []MHSMPrivateEndpointConnection
}

// MHSMPrivateEndpointConnectionsListByResource ...
func (c MHSMListPrivateEndpointConnectionsClient) MHSMPrivateEndpointConnectionsListByResource(ctx context.Context, id ManagedHSMId) (result MHSMPrivateEndpointConnectionsListByResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/privateEndpointConnections", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]MHSMPrivateEndpointConnection `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// MHSMPrivateEndpointConnectionsListByResourceComplete retrieves all the results into a single object
func (c MHSMListPrivateEndpointConnectionsClient) MHSMPrivateEndpointConnectionsListByResourceComplete(ctx context.Context, id ManagedHSMId) (MHSMPrivateEndpointConnectionsListByResourceCompleteResult, error) {
	return c.MHSMPrivateEndpointConnectionsListByResourceCompleteMatchingPredicate(ctx, id, MHSMPrivateEndpointConnectionOperationPredicate{})
}

// MHSMPrivateEndpointConnectionsListByResourceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MHSMListPrivateEndpointConnectionsClient) MHSMPrivateEndpointConnectionsListByResourceCompleteMatchingPredicate(ctx context.Context, id ManagedHSMId, predicate MHSMPrivateEndpointConnectionOperationPredicate) (result MHSMPrivateEndpointConnectionsListByResourceCompleteResult, err error) {
	items := make([]MHSMPrivateEndpointConnection, 0)

	resp, err := c.MHSMPrivateEndpointConnectionsListByResource(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = MHSMPrivateEndpointConnectionsListByResourceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
