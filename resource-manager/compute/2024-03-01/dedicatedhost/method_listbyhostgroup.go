package dedicatedhost

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

type ListByHostGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DedicatedHost
}

type ListByHostGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DedicatedHost
}

type ListByHostGroupCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByHostGroupCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByHostGroup ...
func (c DedicatedHostClient) ListByHostGroup(ctx context.Context, id commonids.DedicatedHostGroupId) (result ListByHostGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByHostGroupCustomPager{},
		Path:       fmt.Sprintf("%s/hosts", id.ID()),
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
		Values *[]DedicatedHost `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByHostGroupComplete retrieves all the results into a single object
func (c DedicatedHostClient) ListByHostGroupComplete(ctx context.Context, id commonids.DedicatedHostGroupId) (ListByHostGroupCompleteResult, error) {
	return c.ListByHostGroupCompleteMatchingPredicate(ctx, id, DedicatedHostOperationPredicate{})
}

// ListByHostGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DedicatedHostClient) ListByHostGroupCompleteMatchingPredicate(ctx context.Context, id commonids.DedicatedHostGroupId, predicate DedicatedHostOperationPredicate) (result ListByHostGroupCompleteResult, err error) {
	items := make([]DedicatedHost, 0)

	resp, err := c.ListByHostGroup(ctx, id)
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

	result = ListByHostGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
