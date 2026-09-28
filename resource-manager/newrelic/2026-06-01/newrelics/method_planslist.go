package newrelics

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

type PlansListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PlanDataResource
}

type PlansListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PlanDataResource
}

type PlansListOperationOptions struct {
	AccountId      *string
	OrganizationId *string
}

func DefaultPlansListOperationOptions() PlansListOperationOptions {
	return PlansListOperationOptions{}
}

func (o PlansListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o PlansListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o PlansListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.AccountId != nil {
		out.Append("accountId", fmt.Sprintf("%v", *o.AccountId))
	}
	if o.OrganizationId != nil {
		out.Append("organizationId", fmt.Sprintf("%v", *o.OrganizationId))
	}
	return &out
}

type PlansListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *PlansListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// PlansList ...
func (c NewRelicsClient) PlansList(ctx context.Context, id commonids.SubscriptionId, options PlansListOperationOptions) (result PlansListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &PlansListCustomPager{},
		Path:          fmt.Sprintf("%s/providers/NewRelic.Observability/plans", id.ID()),
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
		Values *[]PlanDataResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// PlansListComplete retrieves all the results into a single object
func (c NewRelicsClient) PlansListComplete(ctx context.Context, id commonids.SubscriptionId, options PlansListOperationOptions) (PlansListCompleteResult, error) {
	return c.PlansListCompleteMatchingPredicate(ctx, id, options, PlanDataResourceOperationPredicate{})
}

// PlansListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NewRelicsClient) PlansListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options PlansListOperationOptions, predicate PlanDataResourceOperationPredicate) (result PlansListCompleteResult, err error) {
	items := make([]PlanDataResource, 0)

	resp, err := c.PlansList(ctx, id, options)
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

	result = PlansListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
