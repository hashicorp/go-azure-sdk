package agentapplication

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAgentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AgentReference
}

type ListAgentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AgentReference
}

type ListAgentsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListAgentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAgents ...
func (c AgentApplicationClient) ListAgents(ctx context.Context, id ApplicationId) (result ListAgentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &ListAgentsCustomPager{},
		Path:       fmt.Sprintf("%s/listAgents", id.ID()),
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
		Values *[]AgentReference `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAgentsComplete retrieves all the results into a single object
func (c AgentApplicationClient) ListAgentsComplete(ctx context.Context, id ApplicationId) (ListAgentsCompleteResult, error) {
	return c.ListAgentsCompleteMatchingPredicate(ctx, id, AgentReferenceOperationPredicate{})
}

// ListAgentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AgentApplicationClient) ListAgentsCompleteMatchingPredicate(ctx context.Context, id ApplicationId, predicate AgentReferenceOperationPredicate) (result ListAgentsCompleteResult, err error) {
	items := make([]AgentReference, 0)

	resp, err := c.ListAgents(ctx, id)
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

	result = ListAgentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
