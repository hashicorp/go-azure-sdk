package calendargroupcalendarpermission

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

type ListCalendarGroupCalendarPermissionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.CalendarPermission
}

type ListCalendarGroupCalendarPermissionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.CalendarPermission
}

type ListCalendarGroupCalendarPermissionsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListCalendarGroupCalendarPermissionsOperationOptions() ListCalendarGroupCalendarPermissionsOperationOptions {
	return ListCalendarGroupCalendarPermissionsOperationOptions{}
}

func (o ListCalendarGroupCalendarPermissionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCalendarGroupCalendarPermissionsOperationOptions) ToOData() *odata.Query {
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

func (o ListCalendarGroupCalendarPermissionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCalendarGroupCalendarPermissionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarGroupCalendarPermissionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarGroupCalendarPermissions - Get calendarPermissions from me. The permissions of the users with whom the
// calendar is shared.
func (c CalendarGroupCalendarPermissionClient) ListCalendarGroupCalendarPermissions(ctx context.Context, id stable.MeCalendarGroupIdCalendarId, options ListCalendarGroupCalendarPermissionsOperationOptions) (result ListCalendarGroupCalendarPermissionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCalendarGroupCalendarPermissionsCustomPager{},
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
		Values *[]stable.CalendarPermission `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCalendarGroupCalendarPermissionsComplete retrieves all the results into a single object
func (c CalendarGroupCalendarPermissionClient) ListCalendarGroupCalendarPermissionsComplete(ctx context.Context, id stable.MeCalendarGroupIdCalendarId, options ListCalendarGroupCalendarPermissionsOperationOptions) (ListCalendarGroupCalendarPermissionsCompleteResult, error) {
	return c.ListCalendarGroupCalendarPermissionsCompleteMatchingPredicate(ctx, id, options, CalendarPermissionOperationPredicate{})
}

// ListCalendarGroupCalendarPermissionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarGroupCalendarPermissionClient) ListCalendarGroupCalendarPermissionsCompleteMatchingPredicate(ctx context.Context, id stable.MeCalendarGroupIdCalendarId, options ListCalendarGroupCalendarPermissionsOperationOptions, predicate CalendarPermissionOperationPredicate) (result ListCalendarGroupCalendarPermissionsCompleteResult, err error) {
	items := make([]stable.CalendarPermission, 0)

	resp, err := c.ListCalendarGroupCalendarPermissions(ctx, id, options)
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

	result = ListCalendarGroupCalendarPermissionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
