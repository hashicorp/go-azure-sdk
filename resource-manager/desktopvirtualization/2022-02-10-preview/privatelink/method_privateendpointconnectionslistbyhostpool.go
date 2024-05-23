package privatelink

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionsListByHostPoolOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateEndpointConnectionWithSystemData
}

type PrivateEndpointConnectionsListByHostPoolCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PrivateEndpointConnectionWithSystemData
}

// PrivateEndpointConnectionsListByHostPool ...
func (c PrivateLinkClient) PrivateEndpointConnectionsListByHostPool(ctx context.Context, id HostPoolId) (result PrivateEndpointConnectionsListByHostPoolOperationResponse, err error) {
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
		Values *[]PrivateEndpointConnectionWithSystemData `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// PrivateEndpointConnectionsListByHostPoolComplete retrieves all the results into a single object
func (c PrivateLinkClient) PrivateEndpointConnectionsListByHostPoolComplete(ctx context.Context, id HostPoolId) (PrivateEndpointConnectionsListByHostPoolCompleteResult, error) {
	return c.PrivateEndpointConnectionsListByHostPoolCompleteMatchingPredicate(ctx, id, PrivateEndpointConnectionWithSystemDataOperationPredicate{})
}

// PrivateEndpointConnectionsListByHostPoolCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateLinkClient) PrivateEndpointConnectionsListByHostPoolCompleteMatchingPredicate(ctx context.Context, id HostPoolId, predicate PrivateEndpointConnectionWithSystemDataOperationPredicate) (result PrivateEndpointConnectionsListByHostPoolCompleteResult, err error) {
	items := make([]PrivateEndpointConnectionWithSystemData, 0)

	resp, err := c.PrivateEndpointConnectionsListByHostPool(ctx, id)
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

	result = PrivateEndpointConnectionsListByHostPoolCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
