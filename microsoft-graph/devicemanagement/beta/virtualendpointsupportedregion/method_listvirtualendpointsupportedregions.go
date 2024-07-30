package virtualendpointsupportedregion

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

type ListVirtualEndpointSupportedRegionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPCSupportedRegion
}

type ListVirtualEndpointSupportedRegionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPCSupportedRegion
}

type ListVirtualEndpointSupportedRegionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEndpointSupportedRegionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEndpointSupportedRegions ...
func (c VirtualEndpointSupportedRegionClient) ListVirtualEndpointSupportedRegions(ctx context.Context) (result ListVirtualEndpointSupportedRegionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListVirtualEndpointSupportedRegionsCustomPager{},
		Path:       "/deviceManagement/virtualEndpoint/supportedRegions",
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
		Values *[]beta.CloudPCSupportedRegion `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualEndpointSupportedRegionsComplete retrieves all the results into a single object
func (c VirtualEndpointSupportedRegionClient) ListVirtualEndpointSupportedRegionsComplete(ctx context.Context) (ListVirtualEndpointSupportedRegionsCompleteResult, error) {
	return c.ListVirtualEndpointSupportedRegionsCompleteMatchingPredicate(ctx, CloudPCSupportedRegionOperationPredicate{})
}

// ListVirtualEndpointSupportedRegionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointSupportedRegionClient) ListVirtualEndpointSupportedRegionsCompleteMatchingPredicate(ctx context.Context, predicate CloudPCSupportedRegionOperationPredicate) (result ListVirtualEndpointSupportedRegionsCompleteResult, err error) {
	items := make([]beta.CloudPCSupportedRegion, 0)

	resp, err := c.ListVirtualEndpointSupportedRegions(ctx)
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

	result = ListVirtualEndpointSupportedRegionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
