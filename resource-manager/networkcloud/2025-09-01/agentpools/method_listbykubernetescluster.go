package agentpools

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByKubernetesClusterOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AgentPool
}

type ListByKubernetesClusterCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AgentPool
}

type ListByKubernetesClusterOperationOptions struct {
	Top *int64
}

func DefaultListByKubernetesClusterOperationOptions() ListByKubernetesClusterOperationOptions {
	return ListByKubernetesClusterOperationOptions{}
}

func (o ListByKubernetesClusterOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByKubernetesClusterOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListByKubernetesClusterOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type ListByKubernetesClusterCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByKubernetesClusterCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByKubernetesCluster ...
func (c AgentPoolsClient) ListByKubernetesCluster(ctx context.Context, id KubernetesClusterId, options ListByKubernetesClusterOperationOptions) (result ListByKubernetesClusterOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListByKubernetesClusterCustomPager{},
		Path:          fmt.Sprintf("%s/agentPools", id.ID()),
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
		Values *[]AgentPool `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByKubernetesClusterComplete retrieves all the results into a single object
func (c AgentPoolsClient) ListByKubernetesClusterComplete(ctx context.Context, id KubernetesClusterId, options ListByKubernetesClusterOperationOptions) (ListByKubernetesClusterCompleteResult, error) {
	return c.ListByKubernetesClusterCompleteMatchingPredicate(ctx, id, options, AgentPoolOperationPredicate{})
}

// ListByKubernetesClusterCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AgentPoolsClient) ListByKubernetesClusterCompleteMatchingPredicate(ctx context.Context, id KubernetesClusterId, options ListByKubernetesClusterOperationOptions, predicate AgentPoolOperationPredicate) (result ListByKubernetesClusterCompleteResult, err error) {
	items := make([]AgentPool, 0)

	resp, err := c.ListByKubernetesCluster(ctx, id, options)
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

	result = ListByKubernetesClusterCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
