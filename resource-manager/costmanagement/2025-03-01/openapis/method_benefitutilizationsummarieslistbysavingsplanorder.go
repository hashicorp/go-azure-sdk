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

type BenefitUtilizationSummariesListBySavingsPlanOrderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BenefitUtilizationSummary
}

type BenefitUtilizationSummariesListBySavingsPlanOrderCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BenefitUtilizationSummary
}

type BenefitUtilizationSummariesListBySavingsPlanOrderOperationOptions struct {
	Filter         *string
	GrainParameter *GrainParameter
}

func DefaultBenefitUtilizationSummariesListBySavingsPlanOrderOperationOptions() BenefitUtilizationSummariesListBySavingsPlanOrderOperationOptions {
	return BenefitUtilizationSummariesListBySavingsPlanOrderOperationOptions{}
}

func (o BenefitUtilizationSummariesListBySavingsPlanOrderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o BenefitUtilizationSummariesListBySavingsPlanOrderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o BenefitUtilizationSummariesListBySavingsPlanOrderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.GrainParameter != nil {
		out.Append("grainParameter", fmt.Sprintf("%v", *o.GrainParameter))
	}
	return &out
}

type BenefitUtilizationSummariesListBySavingsPlanOrderCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *BenefitUtilizationSummariesListBySavingsPlanOrderCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// BenefitUtilizationSummariesListBySavingsPlanOrder ...
func (c OpenapisClient) BenefitUtilizationSummariesListBySavingsPlanOrder(ctx context.Context, id SavingsPlanOrderId, options BenefitUtilizationSummariesListBySavingsPlanOrderOperationOptions) (result BenefitUtilizationSummariesListBySavingsPlanOrderOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &BenefitUtilizationSummariesListBySavingsPlanOrderCustomPager{},
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

// BenefitUtilizationSummariesListBySavingsPlanOrderComplete retrieves all the results into a single object
func (c OpenapisClient) BenefitUtilizationSummariesListBySavingsPlanOrderComplete(ctx context.Context, id SavingsPlanOrderId, options BenefitUtilizationSummariesListBySavingsPlanOrderOperationOptions) (BenefitUtilizationSummariesListBySavingsPlanOrderCompleteResult, error) {
	return c.BenefitUtilizationSummariesListBySavingsPlanOrderCompleteMatchingPredicate(ctx, id, options, BenefitUtilizationSummaryOperationPredicate{})
}

// BenefitUtilizationSummariesListBySavingsPlanOrderCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) BenefitUtilizationSummariesListBySavingsPlanOrderCompleteMatchingPredicate(ctx context.Context, id SavingsPlanOrderId, options BenefitUtilizationSummariesListBySavingsPlanOrderOperationOptions, predicate BenefitUtilizationSummaryOperationPredicate) (result BenefitUtilizationSummariesListBySavingsPlanOrderCompleteResult, err error) {
	items := make([]BenefitUtilizationSummary, 0)

	resp, err := c.BenefitUtilizationSummariesListBySavingsPlanOrder(ctx, id, options)
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

	result = BenefitUtilizationSummariesListBySavingsPlanOrderCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
