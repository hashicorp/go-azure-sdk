package onenoteresource

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

type ListOnenoteResourcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.OnenoteResource
}

type ListOnenoteResourcesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.OnenoteResource
}

type ListOnenoteResourcesOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListOnenoteResourcesOperationOptions() ListOnenoteResourcesOperationOptions {
	return ListOnenoteResourcesOperationOptions{}
}

func (o ListOnenoteResourcesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOnenoteResourcesOperationOptions) ToOData() *odata.Query {
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

func (o ListOnenoteResourcesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOnenoteResourcesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOnenoteResourcesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOnenoteResources - Get resources from users. The image and other file resources in OneNote pages. Getting a
// resources collection isn't supported, but you can get the binary content of a specific resource. Read-only. Nullable.
func (c OnenoteResourceClient) ListOnenoteResources(ctx context.Context, id beta.UserId, options ListOnenoteResourcesOperationOptions) (result ListOnenoteResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOnenoteResourcesCustomPager{},
		Path:          fmt.Sprintf("%s/onenote/resources", id.ID()),
		RetryFunc:     options.RetryFunc,
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
		Values *[]beta.OnenoteResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOnenoteResourcesComplete retrieves all the results into a single object
func (c OnenoteResourceClient) ListOnenoteResourcesComplete(ctx context.Context, id beta.UserId, options ListOnenoteResourcesOperationOptions) (ListOnenoteResourcesCompleteResult, error) {
	return c.ListOnenoteResourcesCompleteMatchingPredicate(ctx, id, options, OnenoteResourceOperationPredicate{})
}

// ListOnenoteResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OnenoteResourceClient) ListOnenoteResourcesCompleteMatchingPredicate(ctx context.Context, id beta.UserId, options ListOnenoteResourcesOperationOptions, predicate OnenoteResourceOperationPredicate) (result ListOnenoteResourcesCompleteResult, err error) {
	items := make([]beta.OnenoteResource, 0)

	resp, err := c.ListOnenoteResources(ctx, id, options)
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

	result = ListOnenoteResourcesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
