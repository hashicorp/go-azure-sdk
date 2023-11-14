package workflowrunactions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListExpressionTracesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]interface{}
}

type ListExpressionTracesCompleteResult struct {
	Items []interface{}
}

// ListExpressionTraces ...
func (c WorkflowRunActionsClient) ListExpressionTraces(ctx context.Context, id ActionId) (result ListExpressionTracesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/listExpressionTraces", id.ID()),
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
		Values *[]interface{} `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListExpressionTracesComplete retrieves all the results into a single object
func (c WorkflowRunActionsClient) ListExpressionTracesComplete(ctx context.Context, id ActionId) (result ListExpressionTracesCompleteResult, err error) {
	items := make([]interface{}, 0)

	resp, err := c.ListExpressionTraces(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			items = append(items, v)
		}
	}

	result = ListExpressionTracesCompleteResult{
		Items: items,
	}
	return
}
