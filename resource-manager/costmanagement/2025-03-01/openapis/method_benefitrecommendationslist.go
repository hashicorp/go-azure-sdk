package openapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BenefitRecommendationsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BenefitRecommendationModel
}

type BenefitRecommendationsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BenefitRecommendationModel
}

type BenefitRecommendationsListOperationOptions struct {
	Expand  *string
	Filter  *string
	Orderby *string
}

func DefaultBenefitRecommendationsListOperationOptions() BenefitRecommendationsListOperationOptions {
	return BenefitRecommendationsListOperationOptions{}
}

func (o BenefitRecommendationsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o BenefitRecommendationsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o BenefitRecommendationsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Expand != nil {
		out.Append("$expand", fmt.Sprintf("%v", *o.Expand))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Orderby != nil {
		out.Append("$orderby", fmt.Sprintf("%v", *o.Orderby))
	}
	return &out
}

type BenefitRecommendationsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *BenefitRecommendationsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// BenefitRecommendationsList ...
func (c OpenapisClient) BenefitRecommendationsList(ctx context.Context, id commonids.ScopeId, options BenefitRecommendationsListOperationOptions) (result BenefitRecommendationsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &BenefitRecommendationsListCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.CostManagement/benefitRecommendations", id.ID()),
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
		Values *[]BenefitRecommendationModel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// BenefitRecommendationsListComplete retrieves all the results into a single object
func (c OpenapisClient) BenefitRecommendationsListComplete(ctx context.Context, id commonids.ScopeId, options BenefitRecommendationsListOperationOptions) (BenefitRecommendationsListCompleteResult, error) {
	return c.BenefitRecommendationsListCompleteMatchingPredicate(ctx, id, options, BenefitRecommendationModelOperationPredicate{})
}

// BenefitRecommendationsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) BenefitRecommendationsListCompleteMatchingPredicate(ctx context.Context, id commonids.ScopeId, options BenefitRecommendationsListOperationOptions, predicate BenefitRecommendationModelOperationPredicate) (result BenefitRecommendationsListCompleteResult, err error) {
	items := make([]BenefitRecommendationModel, 0)

	resp, err := c.BenefitRecommendationsList(ctx, id, options)
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

	result = BenefitRecommendationsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
