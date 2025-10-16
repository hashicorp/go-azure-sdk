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

type ListByMongoClusterOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateEndpointConnectionResource
}

type ListByMongoClusterCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PrivateEndpointConnectionResource
}

type ListByMongoClusterCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByMongoClusterCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByMongoCluster ...
func (c PrivateEndpointConnectionsClient) ListByMongoCluster(ctx context.Context, id MongoClusterId) (result ListByMongoClusterOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByMongoClusterCustomPager{},
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
		Values *[]PrivateEndpointConnectionResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByMongoClusterComplete retrieves all the results into a single object
func (c PrivateEndpointConnectionsClient) ListByMongoClusterComplete(ctx context.Context, id MongoClusterId) (ListByMongoClusterCompleteResult, error) {
	return c.ListByMongoClusterCompleteMatchingPredicate(ctx, id, PrivateEndpointConnectionResourceOperationPredicate{})
}

// ListByMongoClusterCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateEndpointConnectionsClient) ListByMongoClusterCompleteMatchingPredicate(ctx context.Context, id MongoClusterId, predicate PrivateEndpointConnectionResourceOperationPredicate) (result ListByMongoClusterCompleteResult, err error) {
	items := make([]PrivateEndpointConnectionResource, 0)

	resp, err := c.ListByMongoCluster(ctx, id)
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

	result = ListByMongoClusterCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
