package openapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyMetadataListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SlimPolicyMetadata
}

type PolicyMetadataListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SlimPolicyMetadata
}

type PolicyMetadataListOperationOptions struct {
	Top *int64
}

func DefaultPolicyMetadataListOperationOptions() PolicyMetadataListOperationOptions {
	return PolicyMetadataListOperationOptions{}
}

func (o PolicyMetadataListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o PolicyMetadataListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o PolicyMetadataListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type PolicyMetadataListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *PolicyMetadataListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// PolicyMetadataList ...
func (c OpenapisClient) PolicyMetadataList(ctx context.Context, options PolicyMetadataListOperationOptions) (result PolicyMetadataListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &PolicyMetadataListCustomPager{},
		Path:          "/providers/Microsoft.PolicyInsights/policyMetadata",
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
		Values *[]SlimPolicyMetadata `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// PolicyMetadataListComplete retrieves all the results into a single object
func (c OpenapisClient) PolicyMetadataListComplete(ctx context.Context, options PolicyMetadataListOperationOptions) (PolicyMetadataListCompleteResult, error) {
	return c.PolicyMetadataListCompleteMatchingPredicate(ctx, options, SlimPolicyMetadataOperationPredicate{})
}

// PolicyMetadataListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) PolicyMetadataListCompleteMatchingPredicate(ctx context.Context, options PolicyMetadataListOperationOptions, predicate SlimPolicyMetadataOperationPredicate) (result PolicyMetadataListCompleteResult, err error) {
	items := make([]SlimPolicyMetadata, 0)

	resp, err := c.PolicyMetadataList(ctx, options)
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

	result = PolicyMetadataListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
