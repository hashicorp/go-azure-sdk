package outlooktaskfoldertask

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

type ListOutlookTaskFolderTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.OutlookTask
}

type ListOutlookTaskFolderTasksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.OutlookTask
}

type ListOutlookTaskFolderTasksOperationOptions struct {
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

func DefaultListOutlookTaskFolderTasksOperationOptions() ListOutlookTaskFolderTasksOperationOptions {
	return ListOutlookTaskFolderTasksOperationOptions{}
}

func (o ListOutlookTaskFolderTasksOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOutlookTaskFolderTasksOperationOptions) ToOData() *odata.Query {
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

func (o ListOutlookTaskFolderTasksOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOutlookTaskFolderTasksCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOutlookTaskFolderTasksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOutlookTaskFolderTasks - List tasks (deprecated). Get all the Outlook tasks in the specified folder. By default,
// this operation (and the POST, PATCH, and complete task operations) returns date-related properties in UTC. You can
// use a Prefer: outlook.timezone request header to have all the date-related properties in the response represented in
// a time zone different than UTC. See an example for getting a single task. You can apply the header similarly to get
// multiple tasks. If there is more than one task group, and you want to get all the tasks in a specific task group,
// first get all the task folders in that task group, and then get the tasks in each of these task folders.
func (c OutlookTaskFolderTaskClient) ListOutlookTaskFolderTasks(ctx context.Context, id beta.MeOutlookTaskFolderId, options ListOutlookTaskFolderTasksOperationOptions) (result ListOutlookTaskFolderTasksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOutlookTaskFolderTasksCustomPager{},
		Path:          fmt.Sprintf("%s/tasks", id.ID()),
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
		Values *[]beta.OutlookTask `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOutlookTaskFolderTasksComplete retrieves all the results into a single object
func (c OutlookTaskFolderTaskClient) ListOutlookTaskFolderTasksComplete(ctx context.Context, id beta.MeOutlookTaskFolderId, options ListOutlookTaskFolderTasksOperationOptions) (ListOutlookTaskFolderTasksCompleteResult, error) {
	return c.ListOutlookTaskFolderTasksCompleteMatchingPredicate(ctx, id, options, OutlookTaskOperationPredicate{})
}

// ListOutlookTaskFolderTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OutlookTaskFolderTaskClient) ListOutlookTaskFolderTasksCompleteMatchingPredicate(ctx context.Context, id beta.MeOutlookTaskFolderId, options ListOutlookTaskFolderTasksOperationOptions, predicate OutlookTaskOperationPredicate) (result ListOutlookTaskFolderTasksCompleteResult, err error) {
	items := make([]beta.OutlookTask, 0)

	resp, err := c.ListOutlookTaskFolderTasks(ctx, id, options)
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

	result = ListOutlookTaskFolderTasksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
