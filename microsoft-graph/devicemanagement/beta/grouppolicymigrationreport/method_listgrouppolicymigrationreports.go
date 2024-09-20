package grouppolicymigrationreport

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

type ListGroupPolicyMigrationReportsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicyMigrationReport
}

type ListGroupPolicyMigrationReportsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicyMigrationReport
}

type ListGroupPolicyMigrationReportsOperationOptions struct {
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

func DefaultListGroupPolicyMigrationReportsOperationOptions() ListGroupPolicyMigrationReportsOperationOptions {
	return ListGroupPolicyMigrationReportsOperationOptions{}
}

func (o ListGroupPolicyMigrationReportsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListGroupPolicyMigrationReportsOperationOptions) ToOData() *odata.Query {
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

func (o ListGroupPolicyMigrationReportsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListGroupPolicyMigrationReportsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupPolicyMigrationReportsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupPolicyMigrationReports - Get groupPolicyMigrationReports from deviceManagement. A list of Group Policy
// migration reports.
func (c GroupPolicyMigrationReportClient) ListGroupPolicyMigrationReports(ctx context.Context, options ListGroupPolicyMigrationReportsOperationOptions) (result ListGroupPolicyMigrationReportsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListGroupPolicyMigrationReportsCustomPager{},
		Path:          "/deviceManagement/groupPolicyMigrationReports",
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
		Values *[]beta.GroupPolicyMigrationReport `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupPolicyMigrationReportsComplete retrieves all the results into a single object
func (c GroupPolicyMigrationReportClient) ListGroupPolicyMigrationReportsComplete(ctx context.Context, options ListGroupPolicyMigrationReportsOperationOptions) (ListGroupPolicyMigrationReportsCompleteResult, error) {
	return c.ListGroupPolicyMigrationReportsCompleteMatchingPredicate(ctx, options, GroupPolicyMigrationReportOperationPredicate{})
}

// ListGroupPolicyMigrationReportsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyMigrationReportClient) ListGroupPolicyMigrationReportsCompleteMatchingPredicate(ctx context.Context, options ListGroupPolicyMigrationReportsOperationOptions, predicate GroupPolicyMigrationReportOperationPredicate) (result ListGroupPolicyMigrationReportsCompleteResult, err error) {
	items := make([]beta.GroupPolicyMigrationReport, 0)

	resp, err := c.ListGroupPolicyMigrationReports(ctx, options)
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

	result = ListGroupPolicyMigrationReportsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
