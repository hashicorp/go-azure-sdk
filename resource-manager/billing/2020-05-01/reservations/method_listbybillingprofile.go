package reservations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByBillingProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Reservation
}

type ListByBillingProfileCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Reservation
}

type ListByBillingProfileOperationOptions struct {
	Filter         *string
	Orderby        *string
	RefreshSummary *string
	SelectedState  *string
}

func DefaultListByBillingProfileOperationOptions() ListByBillingProfileOperationOptions {
	return ListByBillingProfileOperationOptions{}
}

func (o ListByBillingProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByBillingProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByBillingProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Orderby != nil {
		out.Append("$orderby", fmt.Sprintf("%v", *o.Orderby))
	}
	if o.RefreshSummary != nil {
		out.Append("refreshSummary", fmt.Sprintf("%v", *o.RefreshSummary))
	}
	if o.SelectedState != nil {
		out.Append("selectedState", fmt.Sprintf("%v", *o.SelectedState))
	}
	return &out
}

type ListByBillingProfileCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByBillingProfileCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByBillingProfile ...
func (c ReservationsClient) ListByBillingProfile(ctx context.Context, id BillingProfileId, options ListByBillingProfileOperationOptions) (result ListByBillingProfileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Pager:         &ListByBillingProfileCustomPager{},
		Path:          fmt.Sprintf("%s/reservations", id.ID()),
		OptionsObject: options,
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
		Values *[]Reservation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByBillingProfileComplete retrieves all the results into a single object
func (c ReservationsClient) ListByBillingProfileComplete(ctx context.Context, id BillingProfileId, options ListByBillingProfileOperationOptions) (ListByBillingProfileCompleteResult, error) {
	return c.ListByBillingProfileCompleteMatchingPredicate(ctx, id, options, ReservationOperationPredicate{})
}

// ListByBillingProfileCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ReservationsClient) ListByBillingProfileCompleteMatchingPredicate(ctx context.Context, id BillingProfileId, options ListByBillingProfileOperationOptions, predicate ReservationOperationPredicate) (result ListByBillingProfileCompleteResult, err error) {
	items := make([]Reservation, 0)

	resp, err := c.ListByBillingProfile(ctx, id, options)
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

	result = ListByBillingProfileCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
