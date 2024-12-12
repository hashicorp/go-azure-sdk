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

type ResourcesListByWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateLinkResource
}

type ResourcesListByWorkspaceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PrivateLinkResource
}

type ResourcesListByWorkspaceOperationOptions struct {
	InitialSkip  *int64
	IsDescending *bool
	PageSize     *int64
}

func DefaultResourcesListByWorkspaceOperationOptions() ResourcesListByWorkspaceOperationOptions {
	return ResourcesListByWorkspaceOperationOptions{}
}

func (o ResourcesListByWorkspaceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ResourcesListByWorkspaceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ResourcesListByWorkspaceOperationOptions) ToQuery() *client.QueryParams {
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

type ResourcesListByWorkspaceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ResourcesListByWorkspaceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ResourcesListByWorkspace ...
func (c PrivateLinkClient) ResourcesListByWorkspace(ctx context.Context, id WorkspaceId, options ResourcesListByWorkspaceOperationOptions) (result ResourcesListByWorkspaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ResourcesListByWorkspaceCustomPager{},
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

// ResourcesListByWorkspaceComplete retrieves all the results into a single object
func (c PrivateLinkClient) ResourcesListByWorkspaceComplete(ctx context.Context, id WorkspaceId, options ResourcesListByWorkspaceOperationOptions) (ResourcesListByWorkspaceCompleteResult, error) {
	return c.ResourcesListByWorkspaceCompleteMatchingPredicate(ctx, id, options, PrivateLinkResourceOperationPredicate{})
}

// ResourcesListByWorkspaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateLinkClient) ResourcesListByWorkspaceCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, options ResourcesListByWorkspaceOperationOptions, predicate PrivateLinkResourceOperationPredicate) (result ResourcesListByWorkspaceCompleteResult, err error) {
	items := make([]PrivateLinkResource, 0)

	resp, err := c.ResourcesListByWorkspace(ctx, id, options)
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

	result = ResourcesListByWorkspaceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
