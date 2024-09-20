package cloudpcconnectivityissue

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

type ListCloudPCConnectivityIssuesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPCConnectivityIssue
}

type ListCloudPCConnectivityIssuesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPCConnectivityIssue
}

type ListCloudPCConnectivityIssuesOperationOptions struct {
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

func DefaultListCloudPCConnectivityIssuesOperationOptions() ListCloudPCConnectivityIssuesOperationOptions {
	return ListCloudPCConnectivityIssuesOperationOptions{}
}

func (o ListCloudPCConnectivityIssuesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCloudPCConnectivityIssuesOperationOptions) ToOData() *odata.Query {
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

func (o ListCloudPCConnectivityIssuesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCloudPCConnectivityIssuesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCloudPCConnectivityIssuesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCloudPCConnectivityIssues - Get cloudPCConnectivityIssues from deviceManagement. The list of CloudPC Connectivity
// Issue.
func (c CloudPCConnectivityIssueClient) ListCloudPCConnectivityIssues(ctx context.Context, options ListCloudPCConnectivityIssuesOperationOptions) (result ListCloudPCConnectivityIssuesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCloudPCConnectivityIssuesCustomPager{},
		Path:          "/deviceManagement/cloudPCConnectivityIssues",
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
		Values *[]beta.CloudPCConnectivityIssue `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCloudPCConnectivityIssuesComplete retrieves all the results into a single object
func (c CloudPCConnectivityIssueClient) ListCloudPCConnectivityIssuesComplete(ctx context.Context, options ListCloudPCConnectivityIssuesOperationOptions) (ListCloudPCConnectivityIssuesCompleteResult, error) {
	return c.ListCloudPCConnectivityIssuesCompleteMatchingPredicate(ctx, options, CloudPCConnectivityIssueOperationPredicate{})
}

// ListCloudPCConnectivityIssuesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CloudPCConnectivityIssueClient) ListCloudPCConnectivityIssuesCompleteMatchingPredicate(ctx context.Context, options ListCloudPCConnectivityIssuesOperationOptions, predicate CloudPCConnectivityIssueOperationPredicate) (result ListCloudPCConnectivityIssuesCompleteResult, err error) {
	items := make([]beta.CloudPCConnectivityIssue, 0)

	resp, err := c.ListCloudPCConnectivityIssues(ctx, options)
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

	result = ListCloudPCConnectivityIssuesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
