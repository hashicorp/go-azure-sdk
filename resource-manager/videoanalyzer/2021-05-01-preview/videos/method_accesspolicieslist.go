package videos

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPoliciesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AccessPolicyEntity
}

type AccessPoliciesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AccessPolicyEntity
}

type AccessPoliciesListOperationOptions struct {
	Top *int64
}

func DefaultAccessPoliciesListOperationOptions() AccessPoliciesListOperationOptions {
	return AccessPoliciesListOperationOptions{}
}

func (o AccessPoliciesListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AccessPoliciesListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o AccessPoliciesListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type AccessPoliciesListCustomPager struct {
	NextLink *odata.Link `json:"@nextLink"`
}

func (p *AccessPoliciesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AccessPoliciesList ...
func (c VideosClient) AccessPoliciesList(ctx context.Context, id VideoAnalyzerId, options AccessPoliciesListOperationOptions) (result AccessPoliciesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &AccessPoliciesListCustomPager{},
		Path:          fmt.Sprintf("%s/accessPolicies", id.ID()),
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
		Values *[]AccessPolicyEntity `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AccessPoliciesListComplete retrieves all the results into a single object
func (c VideosClient) AccessPoliciesListComplete(ctx context.Context, id VideoAnalyzerId, options AccessPoliciesListOperationOptions) (AccessPoliciesListCompleteResult, error) {
	return c.AccessPoliciesListCompleteMatchingPredicate(ctx, id, options, AccessPolicyEntityOperationPredicate{})
}

// AccessPoliciesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VideosClient) AccessPoliciesListCompleteMatchingPredicate(ctx context.Context, id VideoAnalyzerId, options AccessPoliciesListOperationOptions, predicate AccessPolicyEntityOperationPredicate) (result AccessPoliciesListCompleteResult, err error) {
	items := make([]AccessPolicyEntity, 0)

	resp, err := c.AccessPoliciesList(ctx, id, options)
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

	result = AccessPoliciesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
