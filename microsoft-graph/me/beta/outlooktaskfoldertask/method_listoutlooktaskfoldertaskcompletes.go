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

type ListOutlookTaskFolderTaskCompletesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.OutlookTask
}

type ListOutlookTaskFolderTaskCompletesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.OutlookTask
}

type ListOutlookTaskFolderTaskCompletesOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultListOutlookTaskFolderTaskCompletesOperationOptions() ListOutlookTaskFolderTaskCompletesOperationOptions {
	return ListOutlookTaskFolderTaskCompletesOperationOptions{}
}

func (o ListOutlookTaskFolderTaskCompletesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOutlookTaskFolderTaskCompletesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListOutlookTaskFolderTaskCompletesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOutlookTaskFolderTaskCompletesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOutlookTaskFolderTaskCompletesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOutlookTaskFolderTaskCompletes - Invoke action complete. Complete an Outlook task which sets the
// completedDateTime property to the current date, and the status property to completed. If you are completing a task in
// a recurring series, in the response, the task collection will contain the completed task in the series, and the next
// task in the series. The completedDateTime property represents the date when the task is finished. The time portion of
// completedDateTime is set to midnight UTC by default. By default, this operation (and the POST, GET, and PATCH task
// operations) returns date-related properties in UTC. You can use the Prefer: outlook.timezone header to have all the
// date-related properties in the response represented in a time zone different than UTC.
func (c OutlookTaskFolderTaskClient) ListOutlookTaskFolderTaskCompletes(ctx context.Context, id beta.MeOutlookTaskFolderIdTaskId, options ListOutlookTaskFolderTaskCompletesOperationOptions) (result ListOutlookTaskFolderTaskCompletesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListOutlookTaskFolderTaskCompletesCustomPager{},
		Path:          fmt.Sprintf("%s/complete", id.ID()),
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

// ListOutlookTaskFolderTaskCompletesComplete retrieves all the results into a single object
func (c OutlookTaskFolderTaskClient) ListOutlookTaskFolderTaskCompletesComplete(ctx context.Context, id beta.MeOutlookTaskFolderIdTaskId, options ListOutlookTaskFolderTaskCompletesOperationOptions) (ListOutlookTaskFolderTaskCompletesCompleteResult, error) {
	return c.ListOutlookTaskFolderTaskCompletesCompleteMatchingPredicate(ctx, id, options, OutlookTaskOperationPredicate{})
}

// ListOutlookTaskFolderTaskCompletesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OutlookTaskFolderTaskClient) ListOutlookTaskFolderTaskCompletesCompleteMatchingPredicate(ctx context.Context, id beta.MeOutlookTaskFolderIdTaskId, options ListOutlookTaskFolderTaskCompletesOperationOptions, predicate OutlookTaskOperationPredicate) (result ListOutlookTaskFolderTaskCompletesCompleteResult, err error) {
	items := make([]beta.OutlookTask, 0)

	resp, err := c.ListOutlookTaskFolderTaskCompletes(ctx, id, options)
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

	result = ListOutlookTaskFolderTaskCompletesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
