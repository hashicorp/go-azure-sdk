package agentsconnectors

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListWithSecretsByAgentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AgentConnectorWithSecrets
}

type ListWithSecretsByAgentCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AgentConnectorWithSecrets
}

type ListWithSecretsByAgentCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListWithSecretsByAgentCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListWithSecretsByAgent ...
func (c AgentsConnectorsClient) ListWithSecretsByAgent(ctx context.Context, id AgentId) (result ListWithSecretsByAgentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &ListWithSecretsByAgentCustomPager{},
		Path:       fmt.Sprintf("%s/listConnectorsWithSecrets", id.ID()),
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
		Values *[]AgentConnectorWithSecrets `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListWithSecretsByAgentComplete retrieves all the results into a single object
func (c AgentsConnectorsClient) ListWithSecretsByAgentComplete(ctx context.Context, id AgentId) (ListWithSecretsByAgentCompleteResult, error) {
	return c.ListWithSecretsByAgentCompleteMatchingPredicate(ctx, id, AgentConnectorWithSecretsOperationPredicate{})
}

// ListWithSecretsByAgentCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AgentsConnectorsClient) ListWithSecretsByAgentCompleteMatchingPredicate(ctx context.Context, id AgentId, predicate AgentConnectorWithSecretsOperationPredicate) (result ListWithSecretsByAgentCompleteResult, err error) {
	items := make([]AgentConnectorWithSecrets, 0)

	resp, err := c.ListWithSecretsByAgent(ctx, id)
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

	result = ListWithSecretsByAgentCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
