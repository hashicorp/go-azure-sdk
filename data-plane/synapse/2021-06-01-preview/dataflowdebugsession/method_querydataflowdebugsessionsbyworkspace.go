package dataflowdebugsession

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryDataFlowDebugSessionsByWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DataFlowDebugSessionInfo
}

type QueryDataFlowDebugSessionsByWorkspaceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DataFlowDebugSessionInfo
}

type QueryDataFlowDebugSessionsByWorkspaceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *QueryDataFlowDebugSessionsByWorkspaceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// QueryDataFlowDebugSessionsByWorkspace ...
func (c DataFlowDebugSessionClient) QueryDataFlowDebugSessionsByWorkspace(ctx context.Context) (result QueryDataFlowDebugSessionsByWorkspaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &QueryDataFlowDebugSessionsByWorkspaceCustomPager{},
		Path:       "/queryDataFlowDebugSessions",
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
		Values *[]DataFlowDebugSessionInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// QueryDataFlowDebugSessionsByWorkspaceComplete retrieves all the results into a single object
func (c DataFlowDebugSessionClient) QueryDataFlowDebugSessionsByWorkspaceComplete(ctx context.Context) (QueryDataFlowDebugSessionsByWorkspaceCompleteResult, error) {
	return c.QueryDataFlowDebugSessionsByWorkspaceCompleteMatchingPredicate(ctx, DataFlowDebugSessionInfoOperationPredicate{})
}

// QueryDataFlowDebugSessionsByWorkspaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DataFlowDebugSessionClient) QueryDataFlowDebugSessionsByWorkspaceCompleteMatchingPredicate(ctx context.Context, predicate DataFlowDebugSessionInfoOperationPredicate) (result QueryDataFlowDebugSessionsByWorkspaceCompleteResult, err error) {
	items := make([]DataFlowDebugSessionInfo, 0)

	resp, err := c.QueryDataFlowDebugSessionsByWorkspace(ctx)
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

	result = QueryDataFlowDebugSessionsByWorkspaceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
