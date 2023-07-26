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

type QueryByFactoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DataFlowDebugSessionInfo
}

type QueryByFactoryCompleteResult struct {
	Items []DataFlowDebugSessionInfo
}

// QueryByFactory ...
func (c DataFlowDebugSessionClient) QueryByFactory(ctx context.Context, id FactoryId) (result QueryByFactoryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/queryDataFlowDebugSessions", id.ID()),
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

// QueryByFactoryComplete retrieves all the results into a single object
func (c DataFlowDebugSessionClient) QueryByFactoryComplete(ctx context.Context, id FactoryId) (QueryByFactoryCompleteResult, error) {
	return c.QueryByFactoryCompleteMatchingPredicate(ctx, id, DataFlowDebugSessionInfoOperationPredicate{})
}

// QueryByFactoryCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DataFlowDebugSessionClient) QueryByFactoryCompleteMatchingPredicate(ctx context.Context, id FactoryId, predicate DataFlowDebugSessionInfoOperationPredicate) (result QueryByFactoryCompleteResult, err error) {
	items := make([]DataFlowDebugSessionInfo, 0)

	resp, err := c.QueryByFactory(ctx, id)
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

	result = QueryByFactoryCompleteResult{
		Items: items,
	}
	return
}
