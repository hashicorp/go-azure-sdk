package discoveryrules

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByHealthModelOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DiscoveryRule
}

type ListByHealthModelCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DiscoveryRule
}

type ListByHealthModelOperationOptions struct {
	Timestamp *string
}

func DefaultListByHealthModelOperationOptions() ListByHealthModelOperationOptions {
	return ListByHealthModelOperationOptions{}
}

func (o ListByHealthModelOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByHealthModelOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListByHealthModelOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Timestamp != nil {
		out.Append("timestamp", fmt.Sprintf("%v", *o.Timestamp))
	}
	return &out
}

type ListByHealthModelCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByHealthModelCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByHealthModel ...
func (c DiscoveryRulesClient) ListByHealthModel(ctx context.Context, id HealthmodelId, options ListByHealthModelOperationOptions) (result ListByHealthModelOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListByHealthModelCustomPager{},
		Path:          fmt.Sprintf("%s/discoveryrules", id.ID()),
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
		Values *[]DiscoveryRule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByHealthModelComplete retrieves all the results into a single object
func (c DiscoveryRulesClient) ListByHealthModelComplete(ctx context.Context, id HealthmodelId, options ListByHealthModelOperationOptions) (ListByHealthModelCompleteResult, error) {
	return c.ListByHealthModelCompleteMatchingPredicate(ctx, id, options, DiscoveryRuleOperationPredicate{})
}

// ListByHealthModelCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DiscoveryRulesClient) ListByHealthModelCompleteMatchingPredicate(ctx context.Context, id HealthmodelId, options ListByHealthModelOperationOptions, predicate DiscoveryRuleOperationPredicate) (result ListByHealthModelCompleteResult, err error) {
	items := make([]DiscoveryRule, 0)

	resp, err := c.ListByHealthModel(ctx, id, options)
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

	result = ListByHealthModelCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
