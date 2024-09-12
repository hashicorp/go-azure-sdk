package calendarcalendarpermission

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

type ListCalendarPermissionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CalendarPermission
}

type ListCalendarPermissionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CalendarPermission
}

type ListCalendarPermissionsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListCalendarPermissionsOperationOptions() ListCalendarPermissionsOperationOptions {
	return ListCalendarPermissionsOperationOptions{}
}

func (o ListCalendarPermissionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCalendarPermissionsOperationOptions) ToOData() *odata.Query {
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

func (o ListCalendarPermissionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCalendarPermissionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarPermissionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarPermissions - Get calendarPermissions from me. The permissions of the users with whom the calendar is
// shared.
func (c CalendarCalendarPermissionClient) ListCalendarPermissions(ctx context.Context, id beta.MeCalendarId, options ListCalendarPermissionsOperationOptions) (result ListCalendarPermissionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCalendarPermissionsCustomPager{},
		Path:          fmt.Sprintf("%s/calendarPermissions", id.ID()),
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
		Values *[]beta.CalendarPermission `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCalendarPermissionsComplete retrieves all the results into a single object
func (c CalendarCalendarPermissionClient) ListCalendarPermissionsComplete(ctx context.Context, id beta.MeCalendarId, options ListCalendarPermissionsOperationOptions) (ListCalendarPermissionsCompleteResult, error) {
	return c.ListCalendarPermissionsCompleteMatchingPredicate(ctx, id, options, CalendarPermissionOperationPredicate{})
}

// ListCalendarPermissionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarCalendarPermissionClient) ListCalendarPermissionsCompleteMatchingPredicate(ctx context.Context, id beta.MeCalendarId, options ListCalendarPermissionsOperationOptions, predicate CalendarPermissionOperationPredicate) (result ListCalendarPermissionsCompleteResult, err error) {
	items := make([]beta.CalendarPermission, 0)

	resp, err := c.ListCalendarPermissions(ctx, id, options)
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

	result = ListCalendarPermissionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
