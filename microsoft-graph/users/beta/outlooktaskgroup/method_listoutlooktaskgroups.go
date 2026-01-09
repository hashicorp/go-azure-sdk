package outlooktaskgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOutlookTaskGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.OutlookTaskGroup
}

type ListOutlookTaskGroupsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.OutlookTaskGroup
}

type ListOutlookTaskGroupsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListOutlookTaskGroupsOperationOptions() ListOutlookTaskGroupsOperationOptions {
	return ListOutlookTaskGroupsOperationOptions{}
}

func (o ListOutlookTaskGroupsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOutlookTaskGroupsOperationOptions) ToOData() *odata.Query {
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

func (o ListOutlookTaskGroupsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOutlookTaskGroupsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOutlookTaskGroupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOutlookTaskGroups - Get taskGroups from users. The user's Outlook task groups. Read-only. Nullable.
func (c OutlookTaskGroupClient) ListOutlookTaskGroups(ctx context.Context, id beta.UserId, options ListOutlookTaskGroupsOperationOptions) (result ListOutlookTaskGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOutlookTaskGroupsCustomPager{},
		Path:          fmt.Sprintf("%s/outlook/taskGroups", id.ID()),
		RetryFunc:     options.RetryFunc,
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
		Values *[]beta.OutlookTaskGroup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOutlookTaskGroupsComplete retrieves all the results into a single object
func (c OutlookTaskGroupClient) ListOutlookTaskGroupsComplete(ctx context.Context, id beta.UserId, options ListOutlookTaskGroupsOperationOptions) (ListOutlookTaskGroupsCompleteResult, error) {
	return c.ListOutlookTaskGroupsCompleteMatchingPredicate(ctx, id, options, OutlookTaskGroupOperationPredicate{})
}

// ListOutlookTaskGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OutlookTaskGroupClient) ListOutlookTaskGroupsCompleteMatchingPredicate(ctx context.Context, id beta.UserId, options ListOutlookTaskGroupsOperationOptions, predicate OutlookTaskGroupOperationPredicate) (result ListOutlookTaskGroupsCompleteResult, err error) {
	items := make([]beta.OutlookTaskGroup, 0)

	resp, err := c.ListOutlookTaskGroups(ctx, id, options)
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

	result = ListOutlookTaskGroupsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
