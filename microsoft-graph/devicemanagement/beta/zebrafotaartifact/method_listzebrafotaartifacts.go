package zebrafotaartifact

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

type ListZebraFotaArtifactsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ZebraFotaArtifact
}

type ListZebraFotaArtifactsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ZebraFotaArtifact
}

type ListZebraFotaArtifactsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListZebraFotaArtifactsOperationOptions() ListZebraFotaArtifactsOperationOptions {
	return ListZebraFotaArtifactsOperationOptions{}
}

func (o ListZebraFotaArtifactsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListZebraFotaArtifactsOperationOptions) ToOData() *odata.Query {
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

func (o ListZebraFotaArtifactsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListZebraFotaArtifactsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListZebraFotaArtifactsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListZebraFotaArtifacts - Get zebraFotaArtifacts from deviceManagement. The Collection of ZebraFotaArtifacts.
func (c ZebraFotaArtifactClient) ListZebraFotaArtifacts(ctx context.Context, options ListZebraFotaArtifactsOperationOptions) (result ListZebraFotaArtifactsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListZebraFotaArtifactsCustomPager{},
		Path:          "/deviceManagement/zebraFotaArtifacts",
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
		Values *[]beta.ZebraFotaArtifact `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListZebraFotaArtifactsComplete retrieves all the results into a single object
func (c ZebraFotaArtifactClient) ListZebraFotaArtifactsComplete(ctx context.Context, options ListZebraFotaArtifactsOperationOptions) (ListZebraFotaArtifactsCompleteResult, error) {
	return c.ListZebraFotaArtifactsCompleteMatchingPredicate(ctx, options, ZebraFotaArtifactOperationPredicate{})
}

// ListZebraFotaArtifactsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ZebraFotaArtifactClient) ListZebraFotaArtifactsCompleteMatchingPredicate(ctx context.Context, options ListZebraFotaArtifactsOperationOptions, predicate ZebraFotaArtifactOperationPredicate) (result ListZebraFotaArtifactsCompleteResult, err error) {
	items := make([]beta.ZebraFotaArtifact, 0)

	resp, err := c.ListZebraFotaArtifacts(ctx, options)
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

	result = ListZebraFotaArtifactsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
