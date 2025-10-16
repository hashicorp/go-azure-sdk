package networkclouds

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BareMetalMachineKeySetsListByClusterOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BareMetalMachineKeySet
}

type BareMetalMachineKeySetsListByClusterCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BareMetalMachineKeySet
}

type BareMetalMachineKeySetsListByClusterOperationOptions struct {
	Top *int64
}

func DefaultBareMetalMachineKeySetsListByClusterOperationOptions() BareMetalMachineKeySetsListByClusterOperationOptions {
	return BareMetalMachineKeySetsListByClusterOperationOptions{}
}

func (o BareMetalMachineKeySetsListByClusterOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o BareMetalMachineKeySetsListByClusterOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o BareMetalMachineKeySetsListByClusterOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type BareMetalMachineKeySetsListByClusterCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *BareMetalMachineKeySetsListByClusterCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// BareMetalMachineKeySetsListByCluster ...
func (c NetworkcloudsClient) BareMetalMachineKeySetsListByCluster(ctx context.Context, id ClusterId, options BareMetalMachineKeySetsListByClusterOperationOptions) (result BareMetalMachineKeySetsListByClusterOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &BareMetalMachineKeySetsListByClusterCustomPager{},
		Path:          fmt.Sprintf("%s/bareMetalMachineKeySets", id.ID()),
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
		Values *[]BareMetalMachineKeySet `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// BareMetalMachineKeySetsListByClusterComplete retrieves all the results into a single object
func (c NetworkcloudsClient) BareMetalMachineKeySetsListByClusterComplete(ctx context.Context, id ClusterId, options BareMetalMachineKeySetsListByClusterOperationOptions) (BareMetalMachineKeySetsListByClusterCompleteResult, error) {
	return c.BareMetalMachineKeySetsListByClusterCompleteMatchingPredicate(ctx, id, options, BareMetalMachineKeySetOperationPredicate{})
}

// BareMetalMachineKeySetsListByClusterCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) BareMetalMachineKeySetsListByClusterCompleteMatchingPredicate(ctx context.Context, id ClusterId, options BareMetalMachineKeySetsListByClusterOperationOptions, predicate BareMetalMachineKeySetOperationPredicate) (result BareMetalMachineKeySetsListByClusterCompleteResult, err error) {
	items := make([]BareMetalMachineKeySet, 0)

	resp, err := c.BareMetalMachineKeySetsListByCluster(ctx, id, options)
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

	result = BareMetalMachineKeySetsListByClusterCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
