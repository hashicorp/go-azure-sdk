package monthlyprintusagesummariesbyuser

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

type ListMonthlyPrintUsageSummariesByUsersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PrintUsageByUser
}

type ListMonthlyPrintUsageSummariesByUsersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PrintUsageByUser
}

type ListMonthlyPrintUsageSummariesByUsersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMonthlyPrintUsageSummariesByUsersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMonthlyPrintUsageSummariesByUsers ...
func (c MonthlyPrintUsageSummariesByUserClient) ListMonthlyPrintUsageSummariesByUsers(ctx context.Context) (result ListMonthlyPrintUsageSummariesByUsersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMonthlyPrintUsageSummariesByUsersCustomPager{},
		Path:       "/reports/monthlyPrintUsageSummariesByUser",
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
		Values *[]beta.PrintUsageByUser `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMonthlyPrintUsageSummariesByUsersComplete retrieves all the results into a single object
func (c MonthlyPrintUsageSummariesByUserClient) ListMonthlyPrintUsageSummariesByUsersComplete(ctx context.Context) (ListMonthlyPrintUsageSummariesByUsersCompleteResult, error) {
	return c.ListMonthlyPrintUsageSummariesByUsersCompleteMatchingPredicate(ctx, PrintUsageByUserOperationPredicate{})
}

// ListMonthlyPrintUsageSummariesByUsersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MonthlyPrintUsageSummariesByUserClient) ListMonthlyPrintUsageSummariesByUsersCompleteMatchingPredicate(ctx context.Context, predicate PrintUsageByUserOperationPredicate) (result ListMonthlyPrintUsageSummariesByUsersCompleteResult, err error) {
	items := make([]beta.PrintUsageByUser, 0)

	resp, err := c.ListMonthlyPrintUsageSummariesByUsers(ctx)
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

	result = ListMonthlyPrintUsageSummariesByUsersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
