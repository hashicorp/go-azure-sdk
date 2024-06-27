package billingroleassignment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByCustomerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BillingRoleAssignment
}

type ListByCustomerCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingRoleAssignment
}

type ListByCustomerOperationOptions struct {
	Filter *string
	Skip   *int64
	Top    *int64
}

func DefaultListByCustomerOperationOptions() ListByCustomerOperationOptions {
	return ListByCustomerOperationOptions{}
}

func (o ListByCustomerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByCustomerOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByCustomerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Skip != nil {
		out.Append("skip", fmt.Sprintf("%v", *o.Skip))
	}
	if o.Top != nil {
		out.Append("top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// ListByCustomer ...
func (c BillingRoleAssignmentClient) ListByCustomer(ctx context.Context, id BillingProfileCustomerId, options ListByCustomerOperationOptions) (result ListByCustomerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/billingRoleAssignments", id.ID()),
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
		Values *[]BillingRoleAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByCustomerComplete retrieves all the results into a single object
func (c BillingRoleAssignmentClient) ListByCustomerComplete(ctx context.Context, id BillingProfileCustomerId, options ListByCustomerOperationOptions) (ListByCustomerCompleteResult, error) {
	return c.ListByCustomerCompleteMatchingPredicate(ctx, id, options, BillingRoleAssignmentOperationPredicate{})
}

// ListByCustomerCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BillingRoleAssignmentClient) ListByCustomerCompleteMatchingPredicate(ctx context.Context, id BillingProfileCustomerId, options ListByCustomerOperationOptions, predicate BillingRoleAssignmentOperationPredicate) (result ListByCustomerCompleteResult, err error) {
	items := make([]BillingRoleAssignment, 0)

	resp, err := c.ListByCustomer(ctx, id, options)
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

	result = ListByCustomerCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
