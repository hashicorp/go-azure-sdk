package calendargroupcalendar

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

type GetCalendarGroupCalendarSchedulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ScheduleInformation
}

type GetCalendarGroupCalendarSchedulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ScheduleInformation
}

type GetCalendarGroupCalendarSchedulesOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultGetCalendarGroupCalendarSchedulesOperationOptions() GetCalendarGroupCalendarSchedulesOperationOptions {
	return GetCalendarGroupCalendarSchedulesOperationOptions{}
}

func (o GetCalendarGroupCalendarSchedulesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetCalendarGroupCalendarSchedulesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o GetCalendarGroupCalendarSchedulesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type GetCalendarGroupCalendarSchedulesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetCalendarGroupCalendarSchedulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetCalendarGroupCalendarSchedules - Invoke action getSchedule. Get the free/busy availability information for a
// collection of users, distributions lists, or resources (rooms or equipment) for a specified time period.
func (c CalendarGroupCalendarClient) GetCalendarGroupCalendarSchedules(ctx context.Context, id beta.MeCalendarGroupIdCalendarId, input GetCalendarGroupCalendarSchedulesRequest, options GetCalendarGroupCalendarSchedulesOperationOptions) (result GetCalendarGroupCalendarSchedulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &GetCalendarGroupCalendarSchedulesCustomPager{},
		Path:          fmt.Sprintf("%s/getSchedule", id.ID()),
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
		Values *[]beta.ScheduleInformation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetCalendarGroupCalendarSchedulesComplete retrieves all the results into a single object
func (c CalendarGroupCalendarClient) GetCalendarGroupCalendarSchedulesComplete(ctx context.Context, id beta.MeCalendarGroupIdCalendarId, input GetCalendarGroupCalendarSchedulesRequest, options GetCalendarGroupCalendarSchedulesOperationOptions) (GetCalendarGroupCalendarSchedulesCompleteResult, error) {
	return c.GetCalendarGroupCalendarSchedulesCompleteMatchingPredicate(ctx, id, input, options, ScheduleInformationOperationPredicate{})
}

// GetCalendarGroupCalendarSchedulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarGroupCalendarClient) GetCalendarGroupCalendarSchedulesCompleteMatchingPredicate(ctx context.Context, id beta.MeCalendarGroupIdCalendarId, input GetCalendarGroupCalendarSchedulesRequest, options GetCalendarGroupCalendarSchedulesOperationOptions, predicate ScheduleInformationOperationPredicate) (result GetCalendarGroupCalendarSchedulesCompleteResult, err error) {
	items := make([]beta.ScheduleInformation, 0)

	resp, err := c.GetCalendarGroupCalendarSchedules(ctx, id, input, options)
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

	result = GetCalendarGroupCalendarSchedulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
