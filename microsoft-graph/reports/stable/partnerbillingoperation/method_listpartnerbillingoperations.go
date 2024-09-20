package partnerbillingoperation

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPartnerBillingOperationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.PartnersBillingOperation
}

type ListPartnerBillingOperationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.PartnersBillingOperation
}

type ListPartnerBillingOperationsOperationOptions struct {
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

func DefaultListPartnerBillingOperationsOperationOptions() ListPartnerBillingOperationsOperationOptions {
	return ListPartnerBillingOperationsOperationOptions{}
}

func (o ListPartnerBillingOperationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPartnerBillingOperationsOperationOptions) ToOData() *odata.Query {
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

func (o ListPartnerBillingOperationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPartnerBillingOperationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPartnerBillingOperationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPartnerBillingOperations - Get operation. Read the properties and relationships of an operation object.
func (c PartnerBillingOperationClient) ListPartnerBillingOperations(ctx context.Context, options ListPartnerBillingOperationsOperationOptions) (result ListPartnerBillingOperationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPartnerBillingOperationsCustomPager{},
		Path:          "/reports/partners/billing/operations",
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]stable.PartnersBillingOperation, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalPartnersBillingOperationImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.PartnersBillingOperation (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListPartnerBillingOperationsComplete retrieves all the results into a single object
func (c PartnerBillingOperationClient) ListPartnerBillingOperationsComplete(ctx context.Context, options ListPartnerBillingOperationsOperationOptions) (ListPartnerBillingOperationsCompleteResult, error) {
	return c.ListPartnerBillingOperationsCompleteMatchingPredicate(ctx, options, PartnersBillingOperationOperationPredicate{})
}

// ListPartnerBillingOperationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PartnerBillingOperationClient) ListPartnerBillingOperationsCompleteMatchingPredicate(ctx context.Context, options ListPartnerBillingOperationsOperationOptions, predicate PartnersBillingOperationOperationPredicate) (result ListPartnerBillingOperationsCompleteResult, err error) {
	items := make([]stable.PartnersBillingOperation, 0)

	resp, err := c.ListPartnerBillingOperations(ctx, options)
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

	result = ListPartnerBillingOperationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
