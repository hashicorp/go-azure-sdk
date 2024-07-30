package virtualendpointonpremisesconnection

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListVirtualEndpointOnPremisesConnectionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPCOnPremisesConnection
}

type ListVirtualEndpointOnPremisesConnectionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPCOnPremisesConnection
}

type ListVirtualEndpointOnPremisesConnectionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEndpointOnPremisesConnectionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEndpointOnPremisesConnections ...
func (c VirtualEndpointOnPremisesConnectionClient) ListVirtualEndpointOnPremisesConnections(ctx context.Context) (result ListVirtualEndpointOnPremisesConnectionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListVirtualEndpointOnPremisesConnectionsCustomPager{},
		Path:       "/deviceManagement/virtualEndpoint/onPremisesConnections",
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
		Values *[]beta.CloudPCOnPremisesConnection `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualEndpointOnPremisesConnectionsComplete retrieves all the results into a single object
func (c VirtualEndpointOnPremisesConnectionClient) ListVirtualEndpointOnPremisesConnectionsComplete(ctx context.Context) (ListVirtualEndpointOnPremisesConnectionsCompleteResult, error) {
	return c.ListVirtualEndpointOnPremisesConnectionsCompleteMatchingPredicate(ctx, CloudPCOnPremisesConnectionOperationPredicate{})
}

// ListVirtualEndpointOnPremisesConnectionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointOnPremisesConnectionClient) ListVirtualEndpointOnPremisesConnectionsCompleteMatchingPredicate(ctx context.Context, predicate CloudPCOnPremisesConnectionOperationPredicate) (result ListVirtualEndpointOnPremisesConnectionsCompleteResult, err error) {
	items := make([]beta.CloudPCOnPremisesConnection, 0)

	resp, err := c.ListVirtualEndpointOnPremisesConnections(ctx)
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

	result = ListVirtualEndpointOnPremisesConnectionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
