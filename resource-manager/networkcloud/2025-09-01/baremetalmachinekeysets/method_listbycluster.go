package baremetalmachinekeysets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByClusterOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BareMetalMachineKeySet
}

type ListByClusterCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BareMetalMachineKeySet
}

type ListByClusterOperationOptions struct {
	Top *int64
}

func DefaultListByClusterOperationOptions() ListByClusterOperationOptions {
	return ListByClusterOperationOptions{}
}

func (o ListByClusterOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByClusterOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListByClusterOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type ListByClusterCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByClusterCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByCluster ...
func (c BareMetalMachineKeySetsClient) ListByCluster(ctx context.Context, id ClusterId, options ListByClusterOperationOptions) (result ListByClusterOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListByClusterCustomPager{},
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

// ListByClusterComplete retrieves all the results into a single object
func (c BareMetalMachineKeySetsClient) ListByClusterComplete(ctx context.Context, id ClusterId, options ListByClusterOperationOptions) (ListByClusterCompleteResult, error) {
	return c.ListByClusterCompleteMatchingPredicate(ctx, id, options, BareMetalMachineKeySetOperationPredicate{})
}

// ListByClusterCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BareMetalMachineKeySetsClient) ListByClusterCompleteMatchingPredicate(ctx context.Context, id ClusterId, options ListByClusterOperationOptions, predicate BareMetalMachineKeySetOperationPredicate) (result ListByClusterCompleteResult, err error) {
	items := make([]BareMetalMachineKeySet, 0)

	resp, err := c.ListByCluster(ctx, id, options)
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

	result = ListByClusterCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
