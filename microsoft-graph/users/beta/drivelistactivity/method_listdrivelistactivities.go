package drivelistactivity

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

type ListDriveListActivitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ItemActivityOLD
}

type ListDriveListActivitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ItemActivityOLD
}

type ListDriveListActivitiesOperationOptions struct {
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

func DefaultListDriveListActivitiesOperationOptions() ListDriveListActivitiesOperationOptions {
	return ListDriveListActivitiesOperationOptions{}
}

func (o ListDriveListActivitiesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveListActivitiesOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveListActivitiesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveListActivitiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveListActivitiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveListActivities - Get activities from users. The recent activities that took place within this list.
func (c DriveListActivityClient) ListDriveListActivities(ctx context.Context, id beta.UserIdDriveId, options ListDriveListActivitiesOperationOptions) (result ListDriveListActivitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDriveListActivitiesCustomPager{},
		Path:          fmt.Sprintf("%s/list/activities", id.ID()),
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
		Values *[]beta.ItemActivityOLD `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDriveListActivitiesComplete retrieves all the results into a single object
func (c DriveListActivityClient) ListDriveListActivitiesComplete(ctx context.Context, id beta.UserIdDriveId, options ListDriveListActivitiesOperationOptions) (ListDriveListActivitiesCompleteResult, error) {
	return c.ListDriveListActivitiesCompleteMatchingPredicate(ctx, id, options, ItemActivityOLDOperationPredicate{})
}

// ListDriveListActivitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveListActivityClient) ListDriveListActivitiesCompleteMatchingPredicate(ctx context.Context, id beta.UserIdDriveId, options ListDriveListActivitiesOperationOptions, predicate ItemActivityOLDOperationPredicate) (result ListDriveListActivitiesCompleteResult, err error) {
	items := make([]beta.ItemActivityOLD, 0)

	resp, err := c.ListDriveListActivities(ctx, id, options)
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

	result = ListDriveListActivitiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
