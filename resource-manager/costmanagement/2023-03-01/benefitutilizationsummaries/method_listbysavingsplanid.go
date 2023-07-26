package benefitutilizationsummaries

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListBySavingsPlanIdOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BenefitUtilizationSummary
}

type ListBySavingsPlanIdCompleteResult struct {
	Items []BenefitUtilizationSummary
}

type ListBySavingsPlanIdOperationOptions struct {
	Filter         *string
	GrainParameter *GrainParameter
}

func DefaultListBySavingsPlanIdOperationOptions() ListBySavingsPlanIdOperationOptions {
	return ListBySavingsPlanIdOperationOptions{}
}

func (o ListBySavingsPlanIdOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListBySavingsPlanIdOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListBySavingsPlanIdOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.GrainParameter != nil {
		out.Append("grainParameter", fmt.Sprintf("%v", *o.GrainParameter))
	}
	return &out
}

// ListBySavingsPlanId ...
func (c BenefitUtilizationSummariesClient) ListBySavingsPlanId(ctx context.Context, id SavingsPlanId, options ListBySavingsPlanIdOperationOptions) (result ListBySavingsPlanIdOperationResponse, err error) {
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
		Values *[]BenefitUtilizationSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListBySavingsPlanIdComplete retrieves all the results into a single object
func (c BenefitUtilizationSummariesClient) ListBySavingsPlanIdComplete(ctx context.Context, id SavingsPlanId, options ListBySavingsPlanIdOperationOptions) (ListBySavingsPlanIdCompleteResult, error) {
	return c.ListBySavingsPlanIdCompleteMatchingPredicate(ctx, id, options, BenefitUtilizationSummaryOperationPredicate{})
}

// ListBySavingsPlanIdCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BenefitUtilizationSummariesClient) ListBySavingsPlanIdCompleteMatchingPredicate(ctx context.Context, id SavingsPlanId, options ListBySavingsPlanIdOperationOptions, predicate BenefitUtilizationSummaryOperationPredicate) (result ListBySavingsPlanIdCompleteResult, err error) {
	items := make([]BenefitUtilizationSummary, 0)

	resp, err := c.ListBySavingsPlanId(ctx, id, options)
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

	result = ListBySavingsPlanIdCompleteResult{
		Items: items,
	}
	return
}
