package quotarequests

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

type TatusListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]QuotaRequestDetails
}

type TatusListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []QuotaRequestDetails
}

type TatusListOperationOptions struct {
	Filter *string
	Top    *int64
}

func DefaultTatusListOperationOptions() TatusListOperationOptions {
	return TatusListOperationOptions{}
}

func (o TatusListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o TatusListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o TatusListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type TatusListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *TatusListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// TatusList ...
func (c QuotaRequestsClient) TatusList(ctx context.Context, id commonids.ScopeId, options TatusListOperationOptions) (result TatusListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &TatusListCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.Quota/quotaRequests", id.ID()),
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
		Values *[]QuotaRequestDetails `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// TatusListComplete retrieves all the results into a single object
func (c QuotaRequestsClient) TatusListComplete(ctx context.Context, id commonids.ScopeId, options TatusListOperationOptions) (TatusListCompleteResult, error) {
	return c.TatusListCompleteMatchingPredicate(ctx, id, options, QuotaRequestDetailsOperationPredicate{})
}

// TatusListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c QuotaRequestsClient) TatusListCompleteMatchingPredicate(ctx context.Context, id commonids.ScopeId, options TatusListOperationOptions, predicate QuotaRequestDetailsOperationPredicate) (result TatusListCompleteResult, err error) {
	items := make([]QuotaRequestDetails, 0)

	resp, err := c.TatusList(ctx, id, options)
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

	result = TatusListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
