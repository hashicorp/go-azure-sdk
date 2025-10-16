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

type AgentPoolsListByKubernetesClusterOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AgentPool
}

type AgentPoolsListByKubernetesClusterCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AgentPool
}

type AgentPoolsListByKubernetesClusterOperationOptions struct {
	Top *int64
}

func DefaultAgentPoolsListByKubernetesClusterOperationOptions() AgentPoolsListByKubernetesClusterOperationOptions {
	return AgentPoolsListByKubernetesClusterOperationOptions{}
}

func (o AgentPoolsListByKubernetesClusterOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AgentPoolsListByKubernetesClusterOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o AgentPoolsListByKubernetesClusterOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type AgentPoolsListByKubernetesClusterCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *AgentPoolsListByKubernetesClusterCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AgentPoolsListByKubernetesCluster ...
func (c NetworkcloudsClient) AgentPoolsListByKubernetesCluster(ctx context.Context, id KubernetesClusterId, options AgentPoolsListByKubernetesClusterOperationOptions) (result AgentPoolsListByKubernetesClusterOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &AgentPoolsListByKubernetesClusterCustomPager{},
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

// AgentPoolsListByKubernetesClusterComplete retrieves all the results into a single object
func (c NetworkcloudsClient) AgentPoolsListByKubernetesClusterComplete(ctx context.Context, id KubernetesClusterId, options AgentPoolsListByKubernetesClusterOperationOptions) (AgentPoolsListByKubernetesClusterCompleteResult, error) {
	return c.AgentPoolsListByKubernetesClusterCompleteMatchingPredicate(ctx, id, options, AgentPoolOperationPredicate{})
}

// AgentPoolsListByKubernetesClusterCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) AgentPoolsListByKubernetesClusterCompleteMatchingPredicate(ctx context.Context, id KubernetesClusterId, options AgentPoolsListByKubernetesClusterOperationOptions, predicate AgentPoolOperationPredicate) (result AgentPoolsListByKubernetesClusterCompleteResult, err error) {
	items := make([]AgentPool, 0)

	resp, err := c.AgentPoolsListByKubernetesCluster(ctx, id, options)
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

	result = AgentPoolsListByKubernetesClusterCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
