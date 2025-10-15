package networkclouds

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

type ClusterManagersListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ClusterManager
}

type ClusterManagersListByResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ClusterManager
}

type ClusterManagersListByResourceGroupOperationOptions struct {
	Top *int64
}

func DefaultClusterManagersListByResourceGroupOperationOptions() ClusterManagersListByResourceGroupOperationOptions {
	return ClusterManagersListByResourceGroupOperationOptions{}
}

func (o ClusterManagersListByResourceGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ClusterManagersListByResourceGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ClusterManagersListByResourceGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type ClusterManagersListByResourceGroupCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ClusterManagersListByResourceGroupCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ClusterManagersListByResourceGroup ...
func (c NetworkcloudsClient) ClusterManagersListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId, options ClusterManagersListByResourceGroupOperationOptions) (result ClusterManagersListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ClusterManagersListByResourceGroupCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.NetworkCloud/clusterManagers", id.ID()),
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
		Values *[]ClusterManager `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ClusterManagersListByResourceGroupComplete retrieves all the results into a single object
func (c NetworkcloudsClient) ClusterManagersListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId, options ClusterManagersListByResourceGroupOperationOptions) (ClusterManagersListByResourceGroupCompleteResult, error) {
	return c.ClusterManagersListByResourceGroupCompleteMatchingPredicate(ctx, id, options, ClusterManagerOperationPredicate{})
}

// ClusterManagersListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) ClusterManagersListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, options ClusterManagersListByResourceGroupOperationOptions, predicate ClusterManagerOperationPredicate) (result ClusterManagersListByResourceGroupCompleteResult, err error) {
	items := make([]ClusterManager, 0)

	resp, err := c.ClusterManagersListByResourceGroup(ctx, id, options)
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

	result = ClusterManagersListByResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
