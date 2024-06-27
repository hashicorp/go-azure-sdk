package billingsubscription

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByEnrollmentAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BillingSubscription
}

type ListByEnrollmentAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingSubscription
}

type ListByEnrollmentAccountOperationOptions struct {
	Count   *bool
	Filter  *string
	OrderBy *string
	Search  *string
	Skip    *int64
	Top     *int64
}

func DefaultListByEnrollmentAccountOperationOptions() ListByEnrollmentAccountOperationOptions {
	return ListByEnrollmentAccountOperationOptions{}
}

func (o ListByEnrollmentAccountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByEnrollmentAccountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByEnrollmentAccountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Count != nil {
		out.Append("count", fmt.Sprintf("%v", *o.Count))
	}
	if o.Filter != nil {
		out.Append("filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.OrderBy != nil {
		out.Append("orderBy", fmt.Sprintf("%v", *o.OrderBy))
	}
	if o.Search != nil {
		out.Append("search", fmt.Sprintf("%v", *o.Search))
	}
	if o.Skip != nil {
		out.Append("skip", fmt.Sprintf("%v", *o.Skip))
	}
	if o.Top != nil {
		out.Append("top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// ListByEnrollmentAccount ...
func (c BillingSubscriptionClient) ListByEnrollmentAccount(ctx context.Context, id EnrollmentAccountId, options ListByEnrollmentAccountOperationOptions) (result ListByEnrollmentAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/billingSubscriptions", id.ID()),
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
		Values *[]BillingSubscription `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByEnrollmentAccountComplete retrieves all the results into a single object
func (c BillingSubscriptionClient) ListByEnrollmentAccountComplete(ctx context.Context, id EnrollmentAccountId, options ListByEnrollmentAccountOperationOptions) (ListByEnrollmentAccountCompleteResult, error) {
	return c.ListByEnrollmentAccountCompleteMatchingPredicate(ctx, id, options, BillingSubscriptionOperationPredicate{})
}

// ListByEnrollmentAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BillingSubscriptionClient) ListByEnrollmentAccountCompleteMatchingPredicate(ctx context.Context, id EnrollmentAccountId, options ListByEnrollmentAccountOperationOptions, predicate BillingSubscriptionOperationPredicate) (result ListByEnrollmentAccountCompleteResult, err error) {
	items := make([]BillingSubscription, 0)

	resp, err := c.ListByEnrollmentAccount(ctx, id, options)
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

	result = ListByEnrollmentAccountCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
