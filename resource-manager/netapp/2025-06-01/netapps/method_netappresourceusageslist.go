package netapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetAppResourceUsagesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]UsageResult
}

type NetAppResourceUsagesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []UsageResult
}

type NetAppResourceUsagesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *NetAppResourceUsagesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// NetAppResourceUsagesList ...
func (c NetappsClient) NetAppResourceUsagesList(ctx context.Context, id LocationId) (result NetAppResourceUsagesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &NetAppResourceUsagesListCustomPager{},
		Path:       fmt.Sprintf("%s/usages", id.ID()),
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
		Values *[]UsageResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// NetAppResourceUsagesListComplete retrieves all the results into a single object
func (c NetappsClient) NetAppResourceUsagesListComplete(ctx context.Context, id LocationId) (NetAppResourceUsagesListCompleteResult, error) {
	return c.NetAppResourceUsagesListCompleteMatchingPredicate(ctx, id, UsageResultOperationPredicate{})
}

// NetAppResourceUsagesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetappsClient) NetAppResourceUsagesListCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate UsageResultOperationPredicate) (result NetAppResourceUsagesListCompleteResult, err error) {
	items := make([]UsageResult, 0)

	resp, err := c.NetAppResourceUsagesList(ctx, id)
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

	result = NetAppResourceUsagesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
