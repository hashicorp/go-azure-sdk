package exchangeonpremisespolicy

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

type ListExchangeOnPremisesPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementExchangeOnPremisesPolicy
}

type ListExchangeOnPremisesPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementExchangeOnPremisesPolicy
}

type ListExchangeOnPremisesPoliciesOperationOptions struct {
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

func DefaultListExchangeOnPremisesPoliciesOperationOptions() ListExchangeOnPremisesPoliciesOperationOptions {
	return ListExchangeOnPremisesPoliciesOperationOptions{}
}

func (o ListExchangeOnPremisesPoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListExchangeOnPremisesPoliciesOperationOptions) ToOData() *odata.Query {
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

func (o ListExchangeOnPremisesPoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListExchangeOnPremisesPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListExchangeOnPremisesPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListExchangeOnPremisesPolicies - Get exchangeOnPremisesPolicies from deviceManagement. The list of Exchange On
// Premisis policies configured by the tenant.
func (c ExchangeOnPremisesPolicyClient) ListExchangeOnPremisesPolicies(ctx context.Context, options ListExchangeOnPremisesPoliciesOperationOptions) (result ListExchangeOnPremisesPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListExchangeOnPremisesPoliciesCustomPager{},
		Path:          "/deviceManagement/exchangeOnPremisesPolicies",
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
		Values *[]beta.DeviceManagementExchangeOnPremisesPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListExchangeOnPremisesPoliciesComplete retrieves all the results into a single object
func (c ExchangeOnPremisesPolicyClient) ListExchangeOnPremisesPoliciesComplete(ctx context.Context, options ListExchangeOnPremisesPoliciesOperationOptions) (ListExchangeOnPremisesPoliciesCompleteResult, error) {
	return c.ListExchangeOnPremisesPoliciesCompleteMatchingPredicate(ctx, options, DeviceManagementExchangeOnPremisesPolicyOperationPredicate{})
}

// ListExchangeOnPremisesPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ExchangeOnPremisesPolicyClient) ListExchangeOnPremisesPoliciesCompleteMatchingPredicate(ctx context.Context, options ListExchangeOnPremisesPoliciesOperationOptions, predicate DeviceManagementExchangeOnPremisesPolicyOperationPredicate) (result ListExchangeOnPremisesPoliciesCompleteResult, err error) {
	items := make([]beta.DeviceManagementExchangeOnPremisesPolicy, 0)

	resp, err := c.ListExchangeOnPremisesPolicies(ctx, options)
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

	result = ListExchangeOnPremisesPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
