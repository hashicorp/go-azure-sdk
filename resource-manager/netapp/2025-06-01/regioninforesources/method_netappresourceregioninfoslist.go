package regioninforesources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetAppResourceRegionInfosListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RegionInfoResource
}

type NetAppResourceRegionInfosListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []RegionInfoResource
}

type NetAppResourceRegionInfosListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *NetAppResourceRegionInfosListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// NetAppResourceRegionInfosList ...
func (c RegionInfoResourcesClient) NetAppResourceRegionInfosList(ctx context.Context, id LocationId) (result NetAppResourceRegionInfosListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &NetAppResourceRegionInfosListCustomPager{},
		Path:       fmt.Sprintf("%s/regionInfos", id.ID()),
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
		Values *[]RegionInfoResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// NetAppResourceRegionInfosListComplete retrieves all the results into a single object
func (c RegionInfoResourcesClient) NetAppResourceRegionInfosListComplete(ctx context.Context, id LocationId) (NetAppResourceRegionInfosListCompleteResult, error) {
	return c.NetAppResourceRegionInfosListCompleteMatchingPredicate(ctx, id, RegionInfoResourceOperationPredicate{})
}

// NetAppResourceRegionInfosListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RegionInfoResourcesClient) NetAppResourceRegionInfosListCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate RegionInfoResourceOperationPredicate) (result NetAppResourceRegionInfosListCompleteResult, err error) {
	items := make([]RegionInfoResource, 0)

	resp, err := c.NetAppResourceRegionInfosList(ctx, id)
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

	result = NetAppResourceRegionInfosListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
