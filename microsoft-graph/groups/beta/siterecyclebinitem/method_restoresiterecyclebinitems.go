package siterecyclebinitem

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

type RestoreSiteRecycleBinItemsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.RecycleBinItem
}

type RestoreSiteRecycleBinItemsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.RecycleBinItem
}

type RestoreSiteRecycleBinItemsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultRestoreSiteRecycleBinItemsOperationOptions() RestoreSiteRecycleBinItemsOperationOptions {
	return RestoreSiteRecycleBinItemsOperationOptions{}
}

func (o RestoreSiteRecycleBinItemsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RestoreSiteRecycleBinItemsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o RestoreSiteRecycleBinItemsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type RestoreSiteRecycleBinItemsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *RestoreSiteRecycleBinItemsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// RestoreSiteRecycleBinItems - Invoke action restore
func (c SiteRecycleBinItemClient) RestoreSiteRecycleBinItems(ctx context.Context, id beta.GroupIdSiteId, input RestoreSiteRecycleBinItemsRequest, options RestoreSiteRecycleBinItemsOperationOptions) (result RestoreSiteRecycleBinItemsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &RestoreSiteRecycleBinItemsCustomPager{},
		Path:          fmt.Sprintf("%s/recycleBin/items/restore", id.ID()),
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
		Values *[]beta.RecycleBinItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// RestoreSiteRecycleBinItemsComplete retrieves all the results into a single object
func (c SiteRecycleBinItemClient) RestoreSiteRecycleBinItemsComplete(ctx context.Context, id beta.GroupIdSiteId, input RestoreSiteRecycleBinItemsRequest, options RestoreSiteRecycleBinItemsOperationOptions) (RestoreSiteRecycleBinItemsCompleteResult, error) {
	return c.RestoreSiteRecycleBinItemsCompleteMatchingPredicate(ctx, id, input, options, RecycleBinItemOperationPredicate{})
}

// RestoreSiteRecycleBinItemsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteRecycleBinItemClient) RestoreSiteRecycleBinItemsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdSiteId, input RestoreSiteRecycleBinItemsRequest, options RestoreSiteRecycleBinItemsOperationOptions, predicate RecycleBinItemOperationPredicate) (result RestoreSiteRecycleBinItemsCompleteResult, err error) {
	items := make([]beta.RecycleBinItem, 0)

	resp, err := c.RestoreSiteRecycleBinItems(ctx, id, input, options)
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

	result = RestoreSiteRecycleBinItemsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
