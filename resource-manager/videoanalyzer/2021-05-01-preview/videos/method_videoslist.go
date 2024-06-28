package videos

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VideosListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]VideoEntity
}

type VideosListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []VideoEntity
}

type VideosListOperationOptions struct {
	Top *int64
}

func DefaultVideosListOperationOptions() VideosListOperationOptions {
	return VideosListOperationOptions{}
}

func (o VideosListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o VideosListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o VideosListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type VideosListCustomPager struct {
	NextLink *odata.Link `json:"@nextLink"`
}

func (p *VideosListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// VideosList ...
func (c VideosClient) VideosList(ctx context.Context, id VideoAnalyzerId, options VideosListOperationOptions) (result VideosListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Pager:         &VideosListCustomPager{},
		Path:          fmt.Sprintf("%s/videos", id.ID()),
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
		Values *[]VideoEntity `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// VideosListComplete retrieves all the results into a single object
func (c VideosClient) VideosListComplete(ctx context.Context, id VideoAnalyzerId, options VideosListOperationOptions) (VideosListCompleteResult, error) {
	return c.VideosListCompleteMatchingPredicate(ctx, id, options, VideoEntityOperationPredicate{})
}

// VideosListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VideosClient) VideosListCompleteMatchingPredicate(ctx context.Context, id VideoAnalyzerId, options VideosListOperationOptions, predicate VideoEntityOperationPredicate) (result VideosListCompleteResult, err error) {
	items := make([]VideoEntity, 0)

	resp, err := c.VideosList(ctx, id, options)
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

	result = VideosListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
