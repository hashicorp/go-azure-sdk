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

type ListBySavingsPlanOrderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BenefitUtilizationSummary
}

type ListBySavingsPlanOrderCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BenefitUtilizationSummary
}

type ListBySavingsPlanOrderOperationOptions struct {
	Filter         *string
	GrainParameter *GrainParameter
}

func DefaultListBySavingsPlanOrderOperationOptions() ListBySavingsPlanOrderOperationOptions {
	return ListBySavingsPlanOrderOperationOptions{}
}

func (o ListBySavingsPlanOrderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListBySavingsPlanOrderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListBySavingsPlanOrderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.GrainParameter != nil {
		out.Append("grainParameter", fmt.Sprintf("%v", *o.GrainParameter))
	}
	return &out
}

// ListBySavingsPlanOrder ...
func (c BenefitUtilizationSummariesClient) ListBySavingsPlanOrder(ctx context.Context, id SavingsPlanOrderId, options ListBySavingsPlanOrderOperationOptions) (result ListBySavingsPlanOrderOperationResponse, err error) {
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

// ListBySavingsPlanOrderComplete retrieves all the results into a single object
func (c BenefitUtilizationSummariesClient) ListBySavingsPlanOrderComplete(ctx context.Context, id SavingsPlanOrderId, options ListBySavingsPlanOrderOperationOptions) (ListBySavingsPlanOrderCompleteResult, error) {
	return c.ListBySavingsPlanOrderCompleteMatchingPredicate(ctx, id, options, BenefitUtilizationSummaryOperationPredicate{})
}

// ListBySavingsPlanOrderCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BenefitUtilizationSummariesClient) ListBySavingsPlanOrderCompleteMatchingPredicate(ctx context.Context, id SavingsPlanOrderId, options ListBySavingsPlanOrderOperationOptions, predicate BenefitUtilizationSummaryOperationPredicate) (result ListBySavingsPlanOrderCompleteResult, err error) {
	items := make([]BenefitUtilizationSummary, 0)

	resp, err := c.ListBySavingsPlanOrder(ctx, id, options)
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

	result = ListBySavingsPlanOrderCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
