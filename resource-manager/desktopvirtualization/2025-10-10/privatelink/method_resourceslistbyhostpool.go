package privatelink

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourcesListByHostPoolOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateLinkResource
}

type ResourcesListByHostPoolCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PrivateLinkResource
}

type ResourcesListByHostPoolOperationOptions struct {
	InitialSkip  *int64
	IsDescending *bool
	PageSize     *int64
}

func DefaultResourcesListByHostPoolOperationOptions() ResourcesListByHostPoolOperationOptions {
	return ResourcesListByHostPoolOperationOptions{}
}

func (o ResourcesListByHostPoolOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ResourcesListByHostPoolOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ResourcesListByHostPoolOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.InitialSkip != nil {
		out.Append("initialSkip", fmt.Sprintf("%v", *o.InitialSkip))
	}
	if o.IsDescending != nil {
		out.Append("isDescending", fmt.Sprintf("%v", *o.IsDescending))
	}
	if o.PageSize != nil {
		out.Append("pageSize", fmt.Sprintf("%v", *o.PageSize))
	}
	return &out
}

type ResourcesListByHostPoolCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ResourcesListByHostPoolCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ResourcesListByHostPool ...
func (c PrivateLinkClient) ResourcesListByHostPool(ctx context.Context, id HostPoolId, options ResourcesListByHostPoolOperationOptions) (result ResourcesListByHostPoolOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ResourcesListByHostPoolCustomPager{},
		Path:          fmt.Sprintf("%s/privateLinkResources", id.ID()),
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
		Values *[]PrivateLinkResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ResourcesListByHostPoolComplete retrieves all the results into a single object
func (c PrivateLinkClient) ResourcesListByHostPoolComplete(ctx context.Context, id HostPoolId, options ResourcesListByHostPoolOperationOptions) (ResourcesListByHostPoolCompleteResult, error) {
	return c.ResourcesListByHostPoolCompleteMatchingPredicate(ctx, id, options, PrivateLinkResourceOperationPredicate{})
}

// ResourcesListByHostPoolCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateLinkClient) ResourcesListByHostPoolCompleteMatchingPredicate(ctx context.Context, id HostPoolId, options ResourcesListByHostPoolOperationOptions, predicate PrivateLinkResourceOperationPredicate) (result ResourcesListByHostPoolCompleteResult, err error) {
	items := make([]PrivateLinkResource, 0)

	resp, err := c.ResourcesListByHostPool(ctx, id, options)
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

	result = ResourcesListByHostPoolCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
