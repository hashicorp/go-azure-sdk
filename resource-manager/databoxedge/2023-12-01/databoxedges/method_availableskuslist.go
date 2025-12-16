package databoxedges

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvailableSkusListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DataBoxEdgeSku
}

type AvailableSkusListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DataBoxEdgeSku
}

type AvailableSkusListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *AvailableSkusListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AvailableSkusList ...
func (c DataboxedgesClient) AvailableSkusList(ctx context.Context, id commonids.SubscriptionId) (result AvailableSkusListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &AvailableSkusListCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.DataBoxEdge/availableSkus", id.ID()),
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
		Values *[]DataBoxEdgeSku `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AvailableSkusListComplete retrieves all the results into a single object
func (c DataboxedgesClient) AvailableSkusListComplete(ctx context.Context, id commonids.SubscriptionId) (AvailableSkusListCompleteResult, error) {
	return c.AvailableSkusListCompleteMatchingPredicate(ctx, id, DataBoxEdgeSkuOperationPredicate{})
}

// AvailableSkusListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DataboxedgesClient) AvailableSkusListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate DataBoxEdgeSkuOperationPredicate) (result AvailableSkusListCompleteResult, err error) {
	items := make([]DataBoxEdgeSku, 0)

	resp, err := c.AvailableSkusList(ctx, id)
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

	result = AvailableSkusListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
