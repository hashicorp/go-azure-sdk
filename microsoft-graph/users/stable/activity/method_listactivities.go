package activity

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListActivitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserActivity
}

type ListActivitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserActivity
}

type ListActivitiesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListActivitiesOperationOptions() ListActivitiesOperationOptions {
	return ListActivitiesOperationOptions{}
}

func (o ListActivitiesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListActivitiesOperationOptions) ToOData() *odata.Query {
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

func (o ListActivitiesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListActivitiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListActivitiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListActivities - Get activities from users. The user's activities across devices. Read-only. Nullable.
func (c ActivityClient) ListActivities(ctx context.Context, id stable.UserId, options ListActivitiesOperationOptions) (result ListActivitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListActivitiesCustomPager{},
		Path:          fmt.Sprintf("%s/activities", id.ID()),
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
		Values *[]stable.UserActivity `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListActivitiesComplete retrieves all the results into a single object
func (c ActivityClient) ListActivitiesComplete(ctx context.Context, id stable.UserId, options ListActivitiesOperationOptions) (ListActivitiesCompleteResult, error) {
	return c.ListActivitiesCompleteMatchingPredicate(ctx, id, options, UserActivityOperationPredicate{})
}

// ListActivitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ActivityClient) ListActivitiesCompleteMatchingPredicate(ctx context.Context, id stable.UserId, options ListActivitiesOperationOptions, predicate UserActivityOperationPredicate) (result ListActivitiesCompleteResult, err error) {
	items := make([]stable.UserActivity, 0)

	resp, err := c.ListActivities(ctx, id, options)
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

	result = ListActivitiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
