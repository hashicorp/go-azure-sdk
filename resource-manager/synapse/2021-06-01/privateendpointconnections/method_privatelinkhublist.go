package privateendpointconnections

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkHubListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateEndpointConnectionForPrivateLinkHub
}

type PrivateLinkHubListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PrivateEndpointConnectionForPrivateLinkHub
}

// PrivateLinkHubList ...
func (c PrivateEndpointConnectionsClient) PrivateLinkHubList(ctx context.Context, id PrivateLinkHubId) (result PrivateLinkHubListOperationResponse, err error) {
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
		Values *[]PrivateEndpointConnectionForPrivateLinkHub `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// PrivateLinkHubListComplete retrieves all the results into a single object
func (c PrivateEndpointConnectionsClient) PrivateLinkHubListComplete(ctx context.Context, id PrivateLinkHubId) (PrivateLinkHubListCompleteResult, error) {
	return c.PrivateLinkHubListCompleteMatchingPredicate(ctx, id, PrivateEndpointConnectionForPrivateLinkHubOperationPredicate{})
}

// PrivateLinkHubListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateEndpointConnectionsClient) PrivateLinkHubListCompleteMatchingPredicate(ctx context.Context, id PrivateLinkHubId, predicate PrivateEndpointConnectionForPrivateLinkHubOperationPredicate) (result PrivateLinkHubListCompleteResult, err error) {
	items := make([]PrivateEndpointConnectionForPrivateLinkHub, 0)

	resp, err := c.PrivateLinkHubList(ctx, id)
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

	result = PrivateLinkHubListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
