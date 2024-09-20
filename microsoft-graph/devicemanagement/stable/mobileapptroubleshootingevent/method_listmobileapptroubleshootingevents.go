package mobileapptroubleshootingevent

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

type ListMobileAppTroubleshootingEventsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.MobileAppTroubleshootingEvent
}

type ListMobileAppTroubleshootingEventsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.MobileAppTroubleshootingEvent
}

type ListMobileAppTroubleshootingEventsOperationOptions struct {
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

func DefaultListMobileAppTroubleshootingEventsOperationOptions() ListMobileAppTroubleshootingEventsOperationOptions {
	return ListMobileAppTroubleshootingEventsOperationOptions{}
}

func (o ListMobileAppTroubleshootingEventsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMobileAppTroubleshootingEventsOperationOptions) ToOData() *odata.Query {
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

func (o ListMobileAppTroubleshootingEventsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListMobileAppTroubleshootingEventsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMobileAppTroubleshootingEventsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMobileAppTroubleshootingEvents - List mobileAppTroubleshootingEvents. List properties and relationships of the
// mobileAppTroubleshootingEvent objects.
func (c MobileAppTroubleshootingEventClient) ListMobileAppTroubleshootingEvents(ctx context.Context, options ListMobileAppTroubleshootingEventsOperationOptions) (result ListMobileAppTroubleshootingEventsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMobileAppTroubleshootingEventsCustomPager{},
		Path:          "/deviceManagement/mobileAppTroubleshootingEvents",
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
		Values *[]stable.MobileAppTroubleshootingEvent `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMobileAppTroubleshootingEventsComplete retrieves all the results into a single object
func (c MobileAppTroubleshootingEventClient) ListMobileAppTroubleshootingEventsComplete(ctx context.Context, options ListMobileAppTroubleshootingEventsOperationOptions) (ListMobileAppTroubleshootingEventsCompleteResult, error) {
	return c.ListMobileAppTroubleshootingEventsCompleteMatchingPredicate(ctx, options, MobileAppTroubleshootingEventOperationPredicate{})
}

// ListMobileAppTroubleshootingEventsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MobileAppTroubleshootingEventClient) ListMobileAppTroubleshootingEventsCompleteMatchingPredicate(ctx context.Context, options ListMobileAppTroubleshootingEventsOperationOptions, predicate MobileAppTroubleshootingEventOperationPredicate) (result ListMobileAppTroubleshootingEventsCompleteResult, err error) {
	items := make([]stable.MobileAppTroubleshootingEvent, 0)

	resp, err := c.ListMobileAppTroubleshootingEvents(ctx, options)
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

	result = ListMobileAppTroubleshootingEventsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
