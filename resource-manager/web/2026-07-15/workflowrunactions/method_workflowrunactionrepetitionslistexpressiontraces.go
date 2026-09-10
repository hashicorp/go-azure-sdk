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

type WorkflowRunActionRepetitionsListExpressionTracesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]interface{}
}

type WorkflowRunActionRepetitionsListExpressionTracesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []interface{}
}

type WorkflowRunActionRepetitionsListExpressionTracesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WorkflowRunActionRepetitionsListExpressionTracesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WorkflowRunActionRepetitionsListExpressionTraces ...
func (c WorkflowRunActionsClient) WorkflowRunActionRepetitionsListExpressionTraces(ctx context.Context, id RepetitionId) (result WorkflowRunActionRepetitionsListExpressionTracesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &WorkflowRunActionRepetitionsListExpressionTracesCustomPager{},
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

// WorkflowRunActionRepetitionsListExpressionTracesComplete retrieves all the results into a single object
func (c WorkflowRunActionsClient) WorkflowRunActionRepetitionsListExpressionTracesComplete(ctx context.Context, id RepetitionId) (result WorkflowRunActionRepetitionsListExpressionTracesCompleteResult, err error) {
	items := make([]interface{}, 0)

	resp, err := c.WorkflowRunActionRepetitionsListExpressionTraces(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			items = append(items, v)
		}
	}

	result = WorkflowRunActionRepetitionsListExpressionTracesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
