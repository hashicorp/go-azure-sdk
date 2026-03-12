package playwrightworkspacequotas

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByPlaywrightWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PlaywrightWorkspaceQuota
}

type ListByPlaywrightWorkspaceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PlaywrightWorkspaceQuota
}

type ListByPlaywrightWorkspaceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByPlaywrightWorkspaceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByPlaywrightWorkspace ...
func (c PlaywrightWorkspaceQuotasClient) ListByPlaywrightWorkspace(ctx context.Context, id PlaywrightWorkspaceId) (result ListByPlaywrightWorkspaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByPlaywrightWorkspaceCustomPager{},
		Path:       fmt.Sprintf("%s/quotas", id.ID()),
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
		Values *[]PlaywrightWorkspaceQuota `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByPlaywrightWorkspaceComplete retrieves all the results into a single object
func (c PlaywrightWorkspaceQuotasClient) ListByPlaywrightWorkspaceComplete(ctx context.Context, id PlaywrightWorkspaceId) (ListByPlaywrightWorkspaceCompleteResult, error) {
	return c.ListByPlaywrightWorkspaceCompleteMatchingPredicate(ctx, id, PlaywrightWorkspaceQuotaOperationPredicate{})
}

// ListByPlaywrightWorkspaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PlaywrightWorkspaceQuotasClient) ListByPlaywrightWorkspaceCompleteMatchingPredicate(ctx context.Context, id PlaywrightWorkspaceId, predicate PlaywrightWorkspaceQuotaOperationPredicate) (result ListByPlaywrightWorkspaceCompleteResult, err error) {
	items := make([]PlaywrightWorkspaceQuota, 0)

	resp, err := c.ListByPlaywrightWorkspace(ctx, id)
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

	result = ListByPlaywrightWorkspaceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
