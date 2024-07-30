package macossoftwareupdateaccountsummarycategorysummary

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

type ListMacOSSoftwareUpdateAccountSummaryCategorySummariesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MacOSSoftwareUpdateCategorySummary
}

type ListMacOSSoftwareUpdateAccountSummaryCategorySummariesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MacOSSoftwareUpdateCategorySummary
}

type ListMacOSSoftwareUpdateAccountSummaryCategorySummariesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMacOSSoftwareUpdateAccountSummaryCategorySummariesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMacOSSoftwareUpdateAccountSummaryCategorySummaries ...
func (c MacOSSoftwareUpdateAccountSummaryCategorySummaryClient) ListMacOSSoftwareUpdateAccountSummaryCategorySummaries(ctx context.Context, id DeviceManagementMacOSSoftwareUpdateAccountSummaryId) (result ListMacOSSoftwareUpdateAccountSummaryCategorySummariesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMacOSSoftwareUpdateAccountSummaryCategorySummariesCustomPager{},
		Path:       fmt.Sprintf("%s/categorySummaries", id.ID()),
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
		Values *[]beta.MacOSSoftwareUpdateCategorySummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMacOSSoftwareUpdateAccountSummaryCategorySummariesComplete retrieves all the results into a single object
func (c MacOSSoftwareUpdateAccountSummaryCategorySummaryClient) ListMacOSSoftwareUpdateAccountSummaryCategorySummariesComplete(ctx context.Context, id DeviceManagementMacOSSoftwareUpdateAccountSummaryId) (ListMacOSSoftwareUpdateAccountSummaryCategorySummariesCompleteResult, error) {
	return c.ListMacOSSoftwareUpdateAccountSummaryCategorySummariesCompleteMatchingPredicate(ctx, id, MacOSSoftwareUpdateCategorySummaryOperationPredicate{})
}

// ListMacOSSoftwareUpdateAccountSummaryCategorySummariesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MacOSSoftwareUpdateAccountSummaryCategorySummaryClient) ListMacOSSoftwareUpdateAccountSummaryCategorySummariesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementMacOSSoftwareUpdateAccountSummaryId, predicate MacOSSoftwareUpdateCategorySummaryOperationPredicate) (result ListMacOSSoftwareUpdateAccountSummaryCategorySummariesCompleteResult, err error) {
	items := make([]beta.MacOSSoftwareUpdateCategorySummary, 0)

	resp, err := c.ListMacOSSoftwareUpdateAccountSummaryCategorySummaries(ctx, id)
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

	result = ListMacOSSoftwareUpdateAccountSummaryCategorySummariesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
