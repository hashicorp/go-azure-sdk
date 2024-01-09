package recommendations

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

type ListHistoryForWebAppOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Recommendation
}

type ListHistoryForWebAppCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Recommendation
}

type ListHistoryForWebAppOperationOptions struct {
	ExpiredOnly *bool
	Filter      *string
}

func DefaultListHistoryForWebAppOperationOptions() ListHistoryForWebAppOperationOptions {
	return ListHistoryForWebAppOperationOptions{}
}

func (o ListHistoryForWebAppOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListHistoryForWebAppOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListHistoryForWebAppOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.ExpiredOnly != nil {
		out.Append("expiredOnly", fmt.Sprintf("%v", *o.ExpiredOnly))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// ListHistoryForWebApp ...
func (c RecommendationsClient) ListHistoryForWebApp(ctx context.Context, id commonids.AppServiceId, options ListHistoryForWebAppOperationOptions) (result ListHistoryForWebAppOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/recommendationHistory", id.ID()),
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
		Values *[]Recommendation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListHistoryForWebAppComplete retrieves all the results into a single object
func (c RecommendationsClient) ListHistoryForWebAppComplete(ctx context.Context, id commonids.AppServiceId, options ListHistoryForWebAppOperationOptions) (ListHistoryForWebAppCompleteResult, error) {
	return c.ListHistoryForWebAppCompleteMatchingPredicate(ctx, id, options, RecommendationOperationPredicate{})
}

// ListHistoryForWebAppCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RecommendationsClient) ListHistoryForWebAppCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, options ListHistoryForWebAppOperationOptions, predicate RecommendationOperationPredicate) (result ListHistoryForWebAppCompleteResult, err error) {
	items := make([]Recommendation, 0)

	resp, err := c.ListHistoryForWebApp(ctx, id, options)
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

	result = ListHistoryForWebAppCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
