package siteonenoteresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSiteOnenoteResourcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.OnenoteResource
}

type ListSiteOnenoteResourcesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.OnenoteResource
}

type ListSiteOnenoteResourcesOperationOptions struct {
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

func DefaultListSiteOnenoteResourcesOperationOptions() ListSiteOnenoteResourcesOperationOptions {
	return ListSiteOnenoteResourcesOperationOptions{}
}

func (o ListSiteOnenoteResourcesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteOnenoteResourcesOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteOnenoteResourcesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteOnenoteResourcesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteOnenoteResourcesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteOnenoteResources - Get resources from groups. The image and other file resources in OneNote pages. Getting a
// resources collection isn't supported, but you can get the binary content of a specific resource. Read-only. Nullable.
func (c SiteOnenoteResourceClient) ListSiteOnenoteResources(ctx context.Context, id stable.GroupIdSiteId, options ListSiteOnenoteResourcesOperationOptions) (result ListSiteOnenoteResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteOnenoteResourcesCustomPager{},
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
		Values *[]stable.OnenoteResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteOnenoteResourcesComplete retrieves all the results into a single object
func (c SiteOnenoteResourceClient) ListSiteOnenoteResourcesComplete(ctx context.Context, id stable.GroupIdSiteId, options ListSiteOnenoteResourcesOperationOptions) (ListSiteOnenoteResourcesCompleteResult, error) {
	return c.ListSiteOnenoteResourcesCompleteMatchingPredicate(ctx, id, options, OnenoteResourceOperationPredicate{})
}

// ListSiteOnenoteResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteOnenoteResourceClient) ListSiteOnenoteResourcesCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdSiteId, options ListSiteOnenoteResourcesOperationOptions, predicate OnenoteResourceOperationPredicate) (result ListSiteOnenoteResourcesCompleteResult, err error) {
	items := make([]stable.OnenoteResource, 0)

	resp, err := c.ListSiteOnenoteResources(ctx, id, options)
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

	result = ListSiteOnenoteResourcesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
