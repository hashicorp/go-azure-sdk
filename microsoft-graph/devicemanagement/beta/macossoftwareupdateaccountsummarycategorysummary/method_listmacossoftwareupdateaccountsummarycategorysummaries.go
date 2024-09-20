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

type ListMacOSSoftwareUpdateAccountSummaryCategorySummariesOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListMacOSSoftwareUpdateAccountSummaryCategorySummariesOperationOptions() ListMacOSSoftwareUpdateAccountSummaryCategorySummariesOperationOptions {
	return ListMacOSSoftwareUpdateAccountSummaryCategorySummariesOperationOptions{}
}

func (o ListMacOSSoftwareUpdateAccountSummaryCategorySummariesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMacOSSoftwareUpdateAccountSummaryCategorySummariesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListMacOSSoftwareUpdateAccountSummaryCategorySummariesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListMacOSSoftwareUpdateAccountSummaryCategorySummaries - Get categorySummaries from deviceManagement. Summary of the
// updates by category.
func (c MacOSSoftwareUpdateAccountSummaryCategorySummaryClient) ListMacOSSoftwareUpdateAccountSummaryCategorySummaries(ctx context.Context, id beta.DeviceManagementMacOSSoftwareUpdateAccountSummaryId, options ListMacOSSoftwareUpdateAccountSummaryCategorySummariesOperationOptions) (result ListMacOSSoftwareUpdateAccountSummaryCategorySummariesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMacOSSoftwareUpdateAccountSummaryCategorySummariesCustomPager{},
		Path:          fmt.Sprintf("%s/categorySummaries", id.ID()),
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
func (c MacOSSoftwareUpdateAccountSummaryCategorySummaryClient) ListMacOSSoftwareUpdateAccountSummaryCategorySummariesComplete(ctx context.Context, id beta.DeviceManagementMacOSSoftwareUpdateAccountSummaryId, options ListMacOSSoftwareUpdateAccountSummaryCategorySummariesOperationOptions) (ListMacOSSoftwareUpdateAccountSummaryCategorySummariesCompleteResult, error) {
	return c.ListMacOSSoftwareUpdateAccountSummaryCategorySummariesCompleteMatchingPredicate(ctx, id, options, MacOSSoftwareUpdateCategorySummaryOperationPredicate{})
}

// ListMacOSSoftwareUpdateAccountSummaryCategorySummariesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MacOSSoftwareUpdateAccountSummaryCategorySummaryClient) ListMacOSSoftwareUpdateAccountSummaryCategorySummariesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementMacOSSoftwareUpdateAccountSummaryId, options ListMacOSSoftwareUpdateAccountSummaryCategorySummariesOperationOptions, predicate MacOSSoftwareUpdateCategorySummaryOperationPredicate) (result ListMacOSSoftwareUpdateAccountSummaryCategorySummariesCompleteResult, err error) {
	items := make([]beta.MacOSSoftwareUpdateCategorySummary, 0)

	resp, err := c.ListMacOSSoftwareUpdateAccountSummaryCategorySummaries(ctx, id, options)
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
