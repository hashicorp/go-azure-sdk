package usageright

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

type ListUsageRightsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UsageRight
}

type ListUsageRightsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UsageRight
}

type ListUsageRightsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUsageRightsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUsageRights ...
func (c UsageRightClient) ListUsageRights(ctx context.Context) (result ListUsageRightsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUsageRightsCustomPager{},
		Path:       "/me/usageRights",
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
		Values *[]beta.UsageRight `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUsageRightsComplete retrieves all the results into a single object
func (c UsageRightClient) ListUsageRightsComplete(ctx context.Context) (ListUsageRightsCompleteResult, error) {
	return c.ListUsageRightsCompleteMatchingPredicate(ctx, UsageRightOperationPredicate{})
}

// ListUsageRightsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UsageRightClient) ListUsageRightsCompleteMatchingPredicate(ctx context.Context, predicate UsageRightOperationPredicate) (result ListUsageRightsCompleteResult, err error) {
	items := make([]beta.UsageRight, 0)

	resp, err := c.ListUsageRights(ctx)
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

	result = ListUsageRightsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
