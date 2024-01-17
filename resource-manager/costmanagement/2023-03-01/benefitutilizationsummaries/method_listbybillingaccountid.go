package benefitutilizationsummaries

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByBillingAccountIdOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BenefitUtilizationSummary
}

type ListByBillingAccountIdCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BenefitUtilizationSummary
}

type ListByBillingAccountIdOperationOptions struct {
	Filter         *string
	GrainParameter *GrainParameter
}

func DefaultListByBillingAccountIdOperationOptions() ListByBillingAccountIdOperationOptions {
	return ListByBillingAccountIdOperationOptions{}
}

func (o ListByBillingAccountIdOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByBillingAccountIdOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByBillingAccountIdOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.GrainParameter != nil {
		out.Append("grainParameter", fmt.Sprintf("%v", *o.GrainParameter))
	}
	return &out
}

// ListByBillingAccountId ...
func (c BenefitUtilizationSummariesClient) ListByBillingAccountId(ctx context.Context, id BillingAccountId, options ListByBillingAccountIdOperationOptions) (result ListByBillingAccountIdOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/providers/Microsoft.CostManagement/benefitUtilizationSummaries", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]BenefitUtilizationSummary, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := unmarshalBenefitUtilizationSummaryImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for BenefitUtilizationSummary (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListByBillingAccountIdComplete retrieves all the results into a single object
func (c BenefitUtilizationSummariesClient) ListByBillingAccountIdComplete(ctx context.Context, id BillingAccountId, options ListByBillingAccountIdOperationOptions) (ListByBillingAccountIdCompleteResult, error) {
	return c.ListByBillingAccountIdCompleteMatchingPredicate(ctx, id, options, BenefitUtilizationSummaryOperationPredicate{})
}

// ListByBillingAccountIdCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BenefitUtilizationSummariesClient) ListByBillingAccountIdCompleteMatchingPredicate(ctx context.Context, id BillingAccountId, options ListByBillingAccountIdOperationOptions, predicate BenefitUtilizationSummaryOperationPredicate) (result ListByBillingAccountIdCompleteResult, err error) {
	items := make([]BenefitUtilizationSummary, 0)

	resp, err := c.ListByBillingAccountId(ctx, id, options)
	if err != nil {
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

	result = ListByBillingAccountIdCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
