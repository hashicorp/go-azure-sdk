package openapis

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

type BenefitUtilizationSummariesListBySavingsPlanIdOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BenefitUtilizationSummary
}

type BenefitUtilizationSummariesListBySavingsPlanIdCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BenefitUtilizationSummary
}

type BenefitUtilizationSummariesListBySavingsPlanIdOperationOptions struct {
	Filter         *string
	GrainParameter *GrainParameter
}

func DefaultBenefitUtilizationSummariesListBySavingsPlanIdOperationOptions() BenefitUtilizationSummariesListBySavingsPlanIdOperationOptions {
	return BenefitUtilizationSummariesListBySavingsPlanIdOperationOptions{}
}

func (o BenefitUtilizationSummariesListBySavingsPlanIdOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o BenefitUtilizationSummariesListBySavingsPlanIdOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o BenefitUtilizationSummariesListBySavingsPlanIdOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.GrainParameter != nil {
		out.Append("grainParameter", fmt.Sprintf("%v", *o.GrainParameter))
	}
	return &out
}

type BenefitUtilizationSummariesListBySavingsPlanIdCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *BenefitUtilizationSummariesListBySavingsPlanIdCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// BenefitUtilizationSummariesListBySavingsPlanId ...
func (c OpenapisClient) BenefitUtilizationSummariesListBySavingsPlanId(ctx context.Context, id SavingsPlanId, options BenefitUtilizationSummariesListBySavingsPlanIdOperationOptions) (result BenefitUtilizationSummariesListBySavingsPlanIdOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &BenefitUtilizationSummariesListBySavingsPlanIdCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.CostManagement/benefitUtilizationSummaries", id.ID()),
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
			val, err := UnmarshalBenefitUtilizationSummaryImplementation(v)
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

// BenefitUtilizationSummariesListBySavingsPlanIdComplete retrieves all the results into a single object
func (c OpenapisClient) BenefitUtilizationSummariesListBySavingsPlanIdComplete(ctx context.Context, id SavingsPlanId, options BenefitUtilizationSummariesListBySavingsPlanIdOperationOptions) (BenefitUtilizationSummariesListBySavingsPlanIdCompleteResult, error) {
	return c.BenefitUtilizationSummariesListBySavingsPlanIdCompleteMatchingPredicate(ctx, id, options, BenefitUtilizationSummaryOperationPredicate{})
}

// BenefitUtilizationSummariesListBySavingsPlanIdCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) BenefitUtilizationSummariesListBySavingsPlanIdCompleteMatchingPredicate(ctx context.Context, id SavingsPlanId, options BenefitUtilizationSummariesListBySavingsPlanIdOperationOptions, predicate BenefitUtilizationSummaryOperationPredicate) (result BenefitUtilizationSummariesListBySavingsPlanIdCompleteResult, err error) {
	items := make([]BenefitUtilizationSummary, 0)

	resp, err := c.BenefitUtilizationSummariesListBySavingsPlanId(ctx, id, options)
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

	result = BenefitUtilizationSummariesListBySavingsPlanIdCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
