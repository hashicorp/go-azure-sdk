package privateendpointconnection

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPrivateEndpointConnectionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateEndpointConnection
}

type GetPrivateEndpointConnectionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PrivateEndpointConnection
}

type GetPrivateEndpointConnectionsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *GetPrivateEndpointConnectionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetPrivateEndpointConnections ...
func (c PrivateEndpointConnectionClient) GetPrivateEndpointConnections(ctx context.Context, id MasterSiteId) (result GetPrivateEndpointConnectionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &GetPrivateEndpointConnectionsCustomPager{},
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

// GetPrivateEndpointConnectionsComplete retrieves all the results into a single object
func (c PrivateEndpointConnectionClient) GetPrivateEndpointConnectionsComplete(ctx context.Context, id MasterSiteId) (GetPrivateEndpointConnectionsCompleteResult, error) {
	return c.GetPrivateEndpointConnectionsCompleteMatchingPredicate(ctx, id, PrivateEndpointConnectionOperationPredicate{})
}

// GetPrivateEndpointConnectionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateEndpointConnectionClient) GetPrivateEndpointConnectionsCompleteMatchingPredicate(ctx context.Context, id MasterSiteId, predicate PrivateEndpointConnectionOperationPredicate) (result GetPrivateEndpointConnectionsCompleteResult, err error) {
	items := make([]PrivateEndpointConnection, 0)

	resp, err := c.GetPrivateEndpointConnections(ctx, id)
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

	result = GetPrivateEndpointConnectionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
