package billingaccount

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BillingAccount
}

type ListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingAccount
}

type ListOperationOptions struct {
	Expand                           *string
	Filter                           *string
	IncludeAll                       *bool
	IncludeAllWithoutBillingProfiles *bool
	IncludeDeleted                   *bool
	IncludePendingAgreement          *bool
	IncludeResellee                  *bool
	LegalOwnerOID                    *string
	LegalOwnerTID                    *string
	Search                           *string
	Skip                             *int64
	Top                              *int64
}

func DefaultListOperationOptions() ListOperationOptions {
	return ListOperationOptions{}
}

func (o ListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Expand != nil {
		out.Append("expand", fmt.Sprintf("%v", *o.Expand))
	}
	if o.Filter != nil {
		out.Append("filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.IncludeAll != nil {
		out.Append("includeAll", fmt.Sprintf("%v", *o.IncludeAll))
	}
	if o.IncludeAllWithoutBillingProfiles != nil {
		out.Append("includeAllWithoutBillingProfiles", fmt.Sprintf("%v", *o.IncludeAllWithoutBillingProfiles))
	}
	if o.IncludeDeleted != nil {
		out.Append("includeDeleted", fmt.Sprintf("%v", *o.IncludeDeleted))
	}
	if o.IncludePendingAgreement != nil {
		out.Append("includePendingAgreement", fmt.Sprintf("%v", *o.IncludePendingAgreement))
	}
	if o.IncludeResellee != nil {
		out.Append("includeResellee", fmt.Sprintf("%v", *o.IncludeResellee))
	}
	if o.LegalOwnerOID != nil {
		out.Append("legalOwnerOID", fmt.Sprintf("%v", *o.LegalOwnerOID))
	}
	if o.LegalOwnerTID != nil {
		out.Append("legalOwnerTID", fmt.Sprintf("%v", *o.LegalOwnerTID))
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

type ListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// List ...
func (c BillingAccountClient) List(ctx context.Context, options ListOperationOptions) (result ListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCustomPager{},
		Path:          "/providers/Microsoft.Billing/billingAccounts",
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
		Values *[]BillingAccount `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComplete retrieves all the results into a single object
func (c BillingAccountClient) ListComplete(ctx context.Context, options ListOperationOptions) (ListCompleteResult, error) {
	return c.ListCompleteMatchingPredicate(ctx, options, BillingAccountOperationPredicate{})
}

// ListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BillingAccountClient) ListCompleteMatchingPredicate(ctx context.Context, options ListOperationOptions, predicate BillingAccountOperationPredicate) (result ListCompleteResult, err error) {
	items := make([]BillingAccount, 0)

	resp, err := c.List(ctx, options)
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

	result = ListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
