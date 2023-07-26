package hypervhost

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAllHostsInSiteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]HyperVHost
}

type GetAllHostsInSiteCompleteResult struct {
	Items []HyperVHost
}

type GetAllHostsInSiteOperationOptions struct {
	Filter *string
}

func DefaultGetAllHostsInSiteOperationOptions() GetAllHostsInSiteOperationOptions {
	return GetAllHostsInSiteOperationOptions{}
}

func (o GetAllHostsInSiteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAllHostsInSiteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o GetAllHostsInSiteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// GetAllHostsInSite ...
func (c HyperVHostClient) GetAllHostsInSite(ctx context.Context, id HyperVSiteId, options GetAllHostsInSiteOperationOptions) (result GetAllHostsInSiteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/hosts", id.ID()),
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
		Values *[]HyperVHost `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetAllHostsInSiteComplete retrieves all the results into a single object
func (c HyperVHostClient) GetAllHostsInSiteComplete(ctx context.Context, id HyperVSiteId, options GetAllHostsInSiteOperationOptions) (GetAllHostsInSiteCompleteResult, error) {
	return c.GetAllHostsInSiteCompleteMatchingPredicate(ctx, id, options, HyperVHostOperationPredicate{})
}

// GetAllHostsInSiteCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HyperVHostClient) GetAllHostsInSiteCompleteMatchingPredicate(ctx context.Context, id HyperVSiteId, options GetAllHostsInSiteOperationOptions, predicate HyperVHostOperationPredicate) (result GetAllHostsInSiteCompleteResult, err error) {
	items := make([]HyperVHost, 0)

	resp, err := c.GetAllHostsInSite(ctx, id, options)
	if err != nil {
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

	result = GetAllHostsInSiteCompleteResult{
		Items: items,
	}
	return
}
