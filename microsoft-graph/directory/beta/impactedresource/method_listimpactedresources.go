package impactedresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListImpactedResourcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ImpactedResource
}

type ListImpactedResourcesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ImpactedResource
}

type ListImpactedResourcesOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListImpactedResourcesOperationOptions() ListImpactedResourcesOperationOptions {
	return ListImpactedResourcesOperationOptions{}
}

func (o ListImpactedResourcesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListImpactedResourcesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListImpactedResourcesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListImpactedResourcesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListImpactedResourcesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListImpactedResources - Get impactedResources from directory
func (c ImpactedResourceClient) ListImpactedResources(ctx context.Context, options ListImpactedResourcesOperationOptions) (result ListImpactedResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListImpactedResourcesCustomPager{},
		Path:          "/directory/impactedResources",
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
		Values *[]beta.ImpactedResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListImpactedResourcesComplete retrieves all the results into a single object
func (c ImpactedResourceClient) ListImpactedResourcesComplete(ctx context.Context, options ListImpactedResourcesOperationOptions) (ListImpactedResourcesCompleteResult, error) {
	return c.ListImpactedResourcesCompleteMatchingPredicate(ctx, options, ImpactedResourceOperationPredicate{})
}

// ListImpactedResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ImpactedResourceClient) ListImpactedResourcesCompleteMatchingPredicate(ctx context.Context, options ListImpactedResourcesOperationOptions, predicate ImpactedResourceOperationPredicate) (result ListImpactedResourcesCompleteResult, err error) {
	items := make([]beta.ImpactedResource, 0)

	resp, err := c.ListImpactedResources(ctx, options)
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

	result = ListImpactedResourcesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
