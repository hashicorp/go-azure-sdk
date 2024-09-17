package location

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSupportedCloudServiceSkusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SupportedSku
}

type ListSupportedCloudServiceSkusCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SupportedSku
}

type ListSupportedCloudServiceSkusOperationOptions struct {
	Filter     *string
	Maxresults *int64
}

func DefaultListSupportedCloudServiceSkusOperationOptions() ListSupportedCloudServiceSkusOperationOptions {
	return ListSupportedCloudServiceSkusOperationOptions{}
}

func (o ListSupportedCloudServiceSkusOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSupportedCloudServiceSkusOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListSupportedCloudServiceSkusOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Maxresults != nil {
		out.Append("maxresults", fmt.Sprintf("%v", *o.Maxresults))
	}
	return &out
}

type ListSupportedCloudServiceSkusCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListSupportedCloudServiceSkusCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSupportedCloudServiceSkus ...
func (c LocationClient) ListSupportedCloudServiceSkus(ctx context.Context, id LocationId, options ListSupportedCloudServiceSkusOperationOptions) (result ListSupportedCloudServiceSkusOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSupportedCloudServiceSkusCustomPager{},
		Path:          fmt.Sprintf("%s/cloudServiceSkus", id.ID()),
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
		Values *[]SupportedSku `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSupportedCloudServiceSkusComplete retrieves all the results into a single object
func (c LocationClient) ListSupportedCloudServiceSkusComplete(ctx context.Context, id LocationId, options ListSupportedCloudServiceSkusOperationOptions) (ListSupportedCloudServiceSkusCompleteResult, error) {
	return c.ListSupportedCloudServiceSkusCompleteMatchingPredicate(ctx, id, options, SupportedSkuOperationPredicate{})
}

// ListSupportedCloudServiceSkusCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LocationClient) ListSupportedCloudServiceSkusCompleteMatchingPredicate(ctx context.Context, id LocationId, options ListSupportedCloudServiceSkusOperationOptions, predicate SupportedSkuOperationPredicate) (result ListSupportedCloudServiceSkusCompleteResult, err error) {
	items := make([]SupportedSku, 0)

	resp, err := c.ListSupportedCloudServiceSkus(ctx, id, options)
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

	result = ListSupportedCloudServiceSkusCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
