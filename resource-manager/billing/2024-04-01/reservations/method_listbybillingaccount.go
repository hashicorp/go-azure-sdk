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

type ListByBillingAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Reservation
}

type ListByBillingAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Reservation
}

type ListByBillingAccountOperationOptions struct {
	Filter         *string
	OrderBy        *string
	RefreshSummary *string
	SelectedState  *string
	Skiptoken      *int64
	Take           *int64
}

func DefaultListByBillingAccountOperationOptions() ListByBillingAccountOperationOptions {
	return ListByBillingAccountOperationOptions{}
}

func (o ListByBillingAccountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByBillingAccountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListByBillingAccountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.OrderBy != nil {
		out.Append("orderBy", fmt.Sprintf("%v", *o.OrderBy))
	}
	if o.RefreshSummary != nil {
		out.Append("refreshSummary", fmt.Sprintf("%v", *o.RefreshSummary))
	}
	if o.SelectedState != nil {
		out.Append("selectedState", fmt.Sprintf("%v", *o.SelectedState))
	}
	if o.Skiptoken != nil {
		out.Append("skiptoken", fmt.Sprintf("%v", *o.Skiptoken))
	}
	if o.Take != nil {
		out.Append("take", fmt.Sprintf("%v", *o.Take))
	}
	return &out
}

type ListByBillingAccountCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByBillingAccountCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByBillingAccount ...
func (c ReservationsClient) ListByBillingAccount(ctx context.Context, id BillingAccountId, options ListByBillingAccountOperationOptions) (result ListByBillingAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListByBillingAccountCustomPager{},
		Path:          fmt.Sprintf("%s/reservations", id.ID()),
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

// ListByBillingAccountComplete retrieves all the results into a single object
func (c ReservationsClient) ListByBillingAccountComplete(ctx context.Context, id BillingAccountId, options ListByBillingAccountOperationOptions) (ListByBillingAccountCompleteResult, error) {
	return c.ListByBillingAccountCompleteMatchingPredicate(ctx, id, options, ReservationOperationPredicate{})
}

// ListByBillingAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ReservationsClient) ListByBillingAccountCompleteMatchingPredicate(ctx context.Context, id BillingAccountId, options ListByBillingAccountOperationOptions, predicate ReservationOperationPredicate) (result ListByBillingAccountCompleteResult, err error) {
	items := make([]Reservation, 0)

	resp, err := c.ListByBillingAccount(ctx, id, options)
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

	result = ListByBillingAccountCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
