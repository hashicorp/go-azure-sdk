package endpoint

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RaiPoliciesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RaiPolicyPropertiesBasicResource
}

type RaiPoliciesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []RaiPolicyPropertiesBasicResource
}

type RaiPoliciesListOperationOptions struct {
	ProxyApiVersion *string
}

func DefaultRaiPoliciesListOperationOptions() RaiPoliciesListOperationOptions {
	return RaiPoliciesListOperationOptions{}
}

func (o RaiPoliciesListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RaiPoliciesListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o RaiPoliciesListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.ProxyApiVersion != nil {
		out.Append("proxy-api-version", fmt.Sprintf("%v", *o.ProxyApiVersion))
	}
	return &out
}

type RaiPoliciesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *RaiPoliciesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// RaiPoliciesList ...
func (c EndpointClient) RaiPoliciesList(ctx context.Context, id EndpointId, options RaiPoliciesListOperationOptions) (result RaiPoliciesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &RaiPoliciesListCustomPager{},
		Path:          fmt.Sprintf("%s/raiPolicies", id.ID()),
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
		Values *[]RaiPolicyPropertiesBasicResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// RaiPoliciesListComplete retrieves all the results into a single object
func (c EndpointClient) RaiPoliciesListComplete(ctx context.Context, id EndpointId, options RaiPoliciesListOperationOptions) (RaiPoliciesListCompleteResult, error) {
	return c.RaiPoliciesListCompleteMatchingPredicate(ctx, id, options, RaiPolicyPropertiesBasicResourceOperationPredicate{})
}

// RaiPoliciesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EndpointClient) RaiPoliciesListCompleteMatchingPredicate(ctx context.Context, id EndpointId, options RaiPoliciesListOperationOptions, predicate RaiPolicyPropertiesBasicResourceOperationPredicate) (result RaiPoliciesListCompleteResult, err error) {
	items := make([]RaiPolicyPropertiesBasicResource, 0)

	resp, err := c.RaiPoliciesList(ctx, id, options)
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

	result = RaiPoliciesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
