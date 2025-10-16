package fileservice

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

type ListServiceUsagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]FileServiceUsage
}

type ListServiceUsagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []FileServiceUsage
}

type ListServiceUsagesOperationOptions struct {
	Maxpagesize *int64
}

func DefaultListServiceUsagesOperationOptions() ListServiceUsagesOperationOptions {
	return ListServiceUsagesOperationOptions{}
}

func (o ListServiceUsagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListServiceUsagesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListServiceUsagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Maxpagesize != nil {
		out.Append("$maxpagesize", fmt.Sprintf("%v", *o.Maxpagesize))
	}
	return &out
}

type ListServiceUsagesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListServiceUsagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListServiceUsages ...
func (c FileServiceClient) ListServiceUsages(ctx context.Context, id commonids.StorageAccountId, options ListServiceUsagesOperationOptions) (result ListServiceUsagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListServiceUsagesCustomPager{},
		Path:          fmt.Sprintf("%s/fileServices/default/usages", id.ID()),
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
		Values *[]FileServiceUsage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListServiceUsagesComplete retrieves all the results into a single object
func (c FileServiceClient) ListServiceUsagesComplete(ctx context.Context, id commonids.StorageAccountId, options ListServiceUsagesOperationOptions) (ListServiceUsagesCompleteResult, error) {
	return c.ListServiceUsagesCompleteMatchingPredicate(ctx, id, options, FileServiceUsageOperationPredicate{})
}

// ListServiceUsagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c FileServiceClient) ListServiceUsagesCompleteMatchingPredicate(ctx context.Context, id commonids.StorageAccountId, options ListServiceUsagesOperationOptions, predicate FileServiceUsageOperationPredicate) (result ListServiceUsagesCompleteResult, err error) {
	items := make([]FileServiceUsage, 0)

	resp, err := c.ListServiceUsages(ctx, id, options)
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

	result = ListServiceUsagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
