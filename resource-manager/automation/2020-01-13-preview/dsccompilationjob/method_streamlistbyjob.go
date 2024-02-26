package dsccompilationjob

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

type StreamListByJobOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]JobStream
}

type StreamListByJobCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []JobStream
}

// StreamListByJob ...
func (c DscCompilationJobClient) StreamListByJob(ctx context.Context, id commonids.AutomationCompilationJobId) (result StreamListByJobOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/streams", id.ID()),
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
		Values *[]JobStream `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// StreamListByJobComplete retrieves all the results into a single object
func (c DscCompilationJobClient) StreamListByJobComplete(ctx context.Context, id commonids.AutomationCompilationJobId) (StreamListByJobCompleteResult, error) {
	return c.StreamListByJobCompleteMatchingPredicate(ctx, id, JobStreamOperationPredicate{})
}

// StreamListByJobCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DscCompilationJobClient) StreamListByJobCompleteMatchingPredicate(ctx context.Context, id commonids.AutomationCompilationJobId, predicate JobStreamOperationPredicate) (result StreamListByJobCompleteResult, err error) {
	items := make([]JobStream, 0)

	resp, err := c.StreamListByJob(ctx, id)
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

	result = StreamListByJobCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
