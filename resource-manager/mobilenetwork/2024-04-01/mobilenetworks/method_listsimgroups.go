package mobilenetworks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSimGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SimGroup
}

type ListSimGroupsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SimGroup
}

type ListSimGroupsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListSimGroupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSimGroups ...
func (c MobileNetworksClient) ListSimGroups(ctx context.Context, id MobileNetworkId) (result ListSimGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &ListSimGroupsCustomPager{},
		Path:       fmt.Sprintf("%s/listSimGroups", id.ID()),
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
		Values *[]SimGroup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSimGroupsComplete retrieves all the results into a single object
func (c MobileNetworksClient) ListSimGroupsComplete(ctx context.Context, id MobileNetworkId) (ListSimGroupsCompleteResult, error) {
	return c.ListSimGroupsCompleteMatchingPredicate(ctx, id, SimGroupOperationPredicate{})
}

// ListSimGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MobileNetworksClient) ListSimGroupsCompleteMatchingPredicate(ctx context.Context, id MobileNetworkId, predicate SimGroupOperationPredicate) (result ListSimGroupsCompleteResult, err error) {
	items := make([]SimGroup, 0)

	resp, err := c.ListSimGroups(ctx, id)
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

	result = ListSimGroupsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
