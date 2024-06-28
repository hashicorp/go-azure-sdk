package resourcegroups

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

type ListResourcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]GenericResourceExpanded
}

type ListResourcesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []GenericResourceExpanded
}

type ListResourcesOperationOptions struct {
	Expand *string
	Filter *string
	Top    *int64
}

func DefaultListResourcesOperationOptions() ListResourcesOperationOptions {
	return ListResourcesOperationOptions{}
}

func (o ListResourcesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListResourcesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListResourcesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Expand != nil {
		out.Append("$expand", fmt.Sprintf("%v", *o.Expand))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type ListResourcesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListResourcesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListResources ...
func (c ResourceGroupsClient) ListResources(ctx context.Context, id commonids.ResourceGroupId, options ListResourcesOperationOptions) (result ListResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Pager:         &ListResourcesCustomPager{},
		Path:          fmt.Sprintf("%s/resources", id.ID()),
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
		Values *[]GenericResourceExpanded `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListResourcesComplete retrieves all the results into a single object
func (c ResourceGroupsClient) ListResourcesComplete(ctx context.Context, id commonids.ResourceGroupId, options ListResourcesOperationOptions) (ListResourcesCompleteResult, error) {
	return c.ListResourcesCompleteMatchingPredicate(ctx, id, options, GenericResourceExpandedOperationPredicate{})
}

// ListResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ResourceGroupsClient) ListResourcesCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, options ListResourcesOperationOptions, predicate GenericResourceExpandedOperationPredicate) (result ListResourcesCompleteResult, err error) {
	items := make([]GenericResourceExpanded, 0)

	resp, err := c.ListResources(ctx, id, options)
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

	result = ListResourcesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
