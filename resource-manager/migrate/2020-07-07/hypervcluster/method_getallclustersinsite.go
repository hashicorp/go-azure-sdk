package hypervcluster

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAllClustersInSiteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]HyperVCluster
}

type GetAllClustersInSiteCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []HyperVCluster
}

type GetAllClustersInSiteOperationOptions struct {
	Filter *string
}

func DefaultGetAllClustersInSiteOperationOptions() GetAllClustersInSiteOperationOptions {
	return GetAllClustersInSiteOperationOptions{}
}

func (o GetAllClustersInSiteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAllClustersInSiteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o GetAllClustersInSiteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// GetAllClustersInSite ...
func (c HyperVClusterClient) GetAllClustersInSite(ctx context.Context, id HyperVSiteId, options GetAllClustersInSiteOperationOptions) (result GetAllClustersInSiteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/clusters", id.ID()),
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
		Values *[]HyperVCluster `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetAllClustersInSiteComplete retrieves all the results into a single object
func (c HyperVClusterClient) GetAllClustersInSiteComplete(ctx context.Context, id HyperVSiteId, options GetAllClustersInSiteOperationOptions) (GetAllClustersInSiteCompleteResult, error) {
	return c.GetAllClustersInSiteCompleteMatchingPredicate(ctx, id, options, HyperVClusterOperationPredicate{})
}

// GetAllClustersInSiteCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HyperVClusterClient) GetAllClustersInSiteCompleteMatchingPredicate(ctx context.Context, id HyperVSiteId, options GetAllClustersInSiteOperationOptions, predicate HyperVClusterOperationPredicate) (result GetAllClustersInSiteCompleteResult, err error) {
	items := make([]HyperVCluster, 0)

	resp, err := c.GetAllClustersInSite(ctx, id, options)
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

	result = GetAllClustersInSiteCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
