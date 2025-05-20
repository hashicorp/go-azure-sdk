package cloudhsmclusterprivateendpointconnectionsget

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionsListByCloudHsmClusterOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateEndpointConnection
}

type PrivateEndpointConnectionsListByCloudHsmClusterCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PrivateEndpointConnection
}

type PrivateEndpointConnectionsListByCloudHsmClusterCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *PrivateEndpointConnectionsListByCloudHsmClusterCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// PrivateEndpointConnectionsListByCloudHsmCluster ...
func (c CloudHSMClusterPrivateEndpointConnectionsGetClient) PrivateEndpointConnectionsListByCloudHsmCluster(ctx context.Context, id CloudHsmClusterId) (result PrivateEndpointConnectionsListByCloudHsmClusterOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &PrivateEndpointConnectionsListByCloudHsmClusterCustomPager{},
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

// PrivateEndpointConnectionsListByCloudHsmClusterComplete retrieves all the results into a single object
func (c CloudHSMClusterPrivateEndpointConnectionsGetClient) PrivateEndpointConnectionsListByCloudHsmClusterComplete(ctx context.Context, id CloudHsmClusterId) (PrivateEndpointConnectionsListByCloudHsmClusterCompleteResult, error) {
	return c.PrivateEndpointConnectionsListByCloudHsmClusterCompleteMatchingPredicate(ctx, id, PrivateEndpointConnectionOperationPredicate{})
}

// PrivateEndpointConnectionsListByCloudHsmClusterCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CloudHSMClusterPrivateEndpointConnectionsGetClient) PrivateEndpointConnectionsListByCloudHsmClusterCompleteMatchingPredicate(ctx context.Context, id CloudHsmClusterId, predicate PrivateEndpointConnectionOperationPredicate) (result PrivateEndpointConnectionsListByCloudHsmClusterCompleteResult, err error) {
	items := make([]PrivateEndpointConnection, 0)

	resp, err := c.PrivateEndpointConnectionsListByCloudHsmCluster(ctx, id)
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

	result = PrivateEndpointConnectionsListByCloudHsmClusterCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
