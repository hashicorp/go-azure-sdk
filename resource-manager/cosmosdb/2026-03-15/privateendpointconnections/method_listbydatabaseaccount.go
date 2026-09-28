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

type ListByDatabaseAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateEndpointConnection
}

type ListByDatabaseAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PrivateEndpointConnection
}

type ListByDatabaseAccountCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByDatabaseAccountCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByDatabaseAccount ...
func (c PrivateEndpointConnectionsClient) ListByDatabaseAccount(ctx context.Context, id DatabaseAccountId) (result ListByDatabaseAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByDatabaseAccountCustomPager{},
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
		Values *[]PrivateEndpointConnection `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByDatabaseAccountComplete retrieves all the results into a single object
func (c PrivateEndpointConnectionsClient) ListByDatabaseAccountComplete(ctx context.Context, id DatabaseAccountId) (ListByDatabaseAccountCompleteResult, error) {
	return c.ListByDatabaseAccountCompleteMatchingPredicate(ctx, id, PrivateEndpointConnectionOperationPredicate{})
}

// ListByDatabaseAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateEndpointConnectionsClient) ListByDatabaseAccountCompleteMatchingPredicate(ctx context.Context, id DatabaseAccountId, predicate PrivateEndpointConnectionOperationPredicate) (result ListByDatabaseAccountCompleteResult, err error) {
	items := make([]PrivateEndpointConnection, 0)

	resp, err := c.ListByDatabaseAccount(ctx, id)
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

	result = ListByDatabaseAccountCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
