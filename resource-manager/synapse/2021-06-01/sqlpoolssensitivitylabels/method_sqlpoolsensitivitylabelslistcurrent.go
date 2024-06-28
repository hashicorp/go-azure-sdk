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

type SqlPoolSensitivityLabelsListCurrentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SensitivityLabel
}

type SqlPoolSensitivityLabelsListCurrentCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SensitivityLabel
}

type SqlPoolSensitivityLabelsListCurrentOperationOptions struct {
	Filter *string
}

func DefaultSqlPoolSensitivityLabelsListCurrentOperationOptions() SqlPoolSensitivityLabelsListCurrentOperationOptions {
	return SqlPoolSensitivityLabelsListCurrentOperationOptions{}
}

func (o SqlPoolSensitivityLabelsListCurrentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SqlPoolSensitivityLabelsListCurrentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o SqlPoolSensitivityLabelsListCurrentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type SqlPoolSensitivityLabelsListCurrentCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *SqlPoolSensitivityLabelsListCurrentCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// SqlPoolSensitivityLabelsListCurrent ...
func (c SqlPoolsSensitivityLabelsClient) SqlPoolSensitivityLabelsListCurrent(ctx context.Context, id SqlPoolId, options SqlPoolSensitivityLabelsListCurrentOperationOptions) (result SqlPoolSensitivityLabelsListCurrentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Pager:         &SqlPoolSensitivityLabelsListCurrentCustomPager{},
		Path:          fmt.Sprintf("%s/currentSensitivityLabels", id.ID()),
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
		Values *[]SensitivityLabel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SqlPoolSensitivityLabelsListCurrentComplete retrieves all the results into a single object
func (c SqlPoolsSensitivityLabelsClient) SqlPoolSensitivityLabelsListCurrentComplete(ctx context.Context, id SqlPoolId, options SqlPoolSensitivityLabelsListCurrentOperationOptions) (SqlPoolSensitivityLabelsListCurrentCompleteResult, error) {
	return c.SqlPoolSensitivityLabelsListCurrentCompleteMatchingPredicate(ctx, id, options, SensitivityLabelOperationPredicate{})
}

// SqlPoolSensitivityLabelsListCurrentCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SqlPoolsSensitivityLabelsClient) SqlPoolSensitivityLabelsListCurrentCompleteMatchingPredicate(ctx context.Context, id SqlPoolId, options SqlPoolSensitivityLabelsListCurrentOperationOptions, predicate SensitivityLabelOperationPredicate) (result SqlPoolSensitivityLabelsListCurrentCompleteResult, err error) {
	items := make([]SensitivityLabel, 0)

	resp, err := c.SqlPoolSensitivityLabelsListCurrent(ctx, id, options)
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

	result = SqlPoolSensitivityLabelsListCurrentCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
