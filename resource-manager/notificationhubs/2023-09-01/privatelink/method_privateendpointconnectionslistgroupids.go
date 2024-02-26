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

type PrivateEndpointConnectionsListGroupIdsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateLinkResource
}

type PrivateEndpointConnectionsListGroupIdsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PrivateLinkResource
}

// PrivateEndpointConnectionsListGroupIds ...
func (c PrivateLinkClient) PrivateEndpointConnectionsListGroupIds(ctx context.Context, id NamespaceId) (result PrivateEndpointConnectionsListGroupIdsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/privateLinkResources", id.ID()),
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
		Values *[]PrivateLinkResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// PrivateEndpointConnectionsListGroupIdsComplete retrieves all the results into a single object
func (c PrivateLinkClient) PrivateEndpointConnectionsListGroupIdsComplete(ctx context.Context, id NamespaceId) (PrivateEndpointConnectionsListGroupIdsCompleteResult, error) {
	return c.PrivateEndpointConnectionsListGroupIdsCompleteMatchingPredicate(ctx, id, PrivateLinkResourceOperationPredicate{})
}

// PrivateEndpointConnectionsListGroupIdsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateLinkClient) PrivateEndpointConnectionsListGroupIdsCompleteMatchingPredicate(ctx context.Context, id NamespaceId, predicate PrivateLinkResourceOperationPredicate) (result PrivateEndpointConnectionsListGroupIdsCompleteResult, err error) {
	items := make([]PrivateLinkResource, 0)

	resp, err := c.PrivateEndpointConnectionsListGroupIds(ctx, id)
	if err != nil {
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

	result = PrivateEndpointConnectionsListGroupIdsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
