package sqlpoolssensitivitylabels

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolSensitivityLabelsListRecommendedOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SensitivityLabel
}

type SqlPoolSensitivityLabelsListRecommendedCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SensitivityLabel
}

type SqlPoolSensitivityLabelsListRecommendedOperationOptions struct {
	Filter                         *string
	IncludeDisabledRecommendations *bool
}

func DefaultSqlPoolSensitivityLabelsListRecommendedOperationOptions() SqlPoolSensitivityLabelsListRecommendedOperationOptions {
	return SqlPoolSensitivityLabelsListRecommendedOperationOptions{}
}

func (o SqlPoolSensitivityLabelsListRecommendedOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SqlPoolSensitivityLabelsListRecommendedOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o SqlPoolSensitivityLabelsListRecommendedOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.IncludeDisabledRecommendations != nil {
		out.Append("includeDisabledRecommendations", fmt.Sprintf("%v", *o.IncludeDisabledRecommendations))
	}
	return &out
}

type SqlPoolSensitivityLabelsListRecommendedCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *SqlPoolSensitivityLabelsListRecommendedCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// SqlPoolSensitivityLabelsListRecommended ...
func (c SqlPoolsSensitivityLabelsClient) SqlPoolSensitivityLabelsListRecommended(ctx context.Context, id SqlPoolId, options SqlPoolSensitivityLabelsListRecommendedOperationOptions) (result SqlPoolSensitivityLabelsListRecommendedOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &SqlPoolSensitivityLabelsListRecommendedCustomPager{},
		Path:          fmt.Sprintf("%s/recommendedSensitivityLabels", id.ID()),
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
		Values *[]SensitivityLabel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SqlPoolSensitivityLabelsListRecommendedComplete retrieves all the results into a single object
func (c SqlPoolsSensitivityLabelsClient) SqlPoolSensitivityLabelsListRecommendedComplete(ctx context.Context, id SqlPoolId, options SqlPoolSensitivityLabelsListRecommendedOperationOptions) (SqlPoolSensitivityLabelsListRecommendedCompleteResult, error) {
	return c.SqlPoolSensitivityLabelsListRecommendedCompleteMatchingPredicate(ctx, id, options, SensitivityLabelOperationPredicate{})
}

// SqlPoolSensitivityLabelsListRecommendedCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SqlPoolsSensitivityLabelsClient) SqlPoolSensitivityLabelsListRecommendedCompleteMatchingPredicate(ctx context.Context, id SqlPoolId, options SqlPoolSensitivityLabelsListRecommendedOperationOptions, predicate SensitivityLabelOperationPredicate) (result SqlPoolSensitivityLabelsListRecommendedCompleteResult, err error) {
	items := make([]SensitivityLabel, 0)

	resp, err := c.SqlPoolSensitivityLabelsListRecommended(ctx, id, options)
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

	result = SqlPoolSensitivityLabelsListRecommendedCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
