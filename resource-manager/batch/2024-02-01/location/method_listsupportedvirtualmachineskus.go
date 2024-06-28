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

type ListSupportedVirtualMachineSkusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SupportedSku
}

type ListSupportedVirtualMachineSkusCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SupportedSku
}

type ListSupportedVirtualMachineSkusOperationOptions struct {
	Filter     *string
	Maxresults *int64
}

func DefaultListSupportedVirtualMachineSkusOperationOptions() ListSupportedVirtualMachineSkusOperationOptions {
	return ListSupportedVirtualMachineSkusOperationOptions{}
}

func (o ListSupportedVirtualMachineSkusOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSupportedVirtualMachineSkusOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListSupportedVirtualMachineSkusOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Maxresults != nil {
		out.Append("maxresults", fmt.Sprintf("%v", *o.Maxresults))
	}
	return &out
}

type ListSupportedVirtualMachineSkusCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListSupportedVirtualMachineSkusCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSupportedVirtualMachineSkus ...
func (c LocationClient) ListSupportedVirtualMachineSkus(ctx context.Context, id LocationId, options ListSupportedVirtualMachineSkusOperationOptions) (result ListSupportedVirtualMachineSkusOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSupportedVirtualMachineSkusCustomPager{},
		Path:          fmt.Sprintf("%s/virtualMachineSkus", id.ID()),
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

// ListSupportedVirtualMachineSkusComplete retrieves all the results into a single object
func (c LocationClient) ListSupportedVirtualMachineSkusComplete(ctx context.Context, id LocationId, options ListSupportedVirtualMachineSkusOperationOptions) (ListSupportedVirtualMachineSkusCompleteResult, error) {
	return c.ListSupportedVirtualMachineSkusCompleteMatchingPredicate(ctx, id, options, SupportedSkuOperationPredicate{})
}

// ListSupportedVirtualMachineSkusCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LocationClient) ListSupportedVirtualMachineSkusCompleteMatchingPredicate(ctx context.Context, id LocationId, options ListSupportedVirtualMachineSkusOperationOptions, predicate SupportedSkuOperationPredicate) (result ListSupportedVirtualMachineSkusCompleteResult, err error) {
	items := make([]SupportedSku, 0)

	resp, err := c.ListSupportedVirtualMachineSkus(ctx, id, options)
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

	result = ListSupportedVirtualMachineSkusCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
