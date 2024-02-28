package iotsecuritysolutionsanalytics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAllOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]IoTSecuritySolutionAnalyticsModel
}

type GetAllCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []IoTSecuritySolutionAnalyticsModel
}

// GetAll ...
func (c IoTSecuritySolutionsAnalyticsClient) GetAll(ctx context.Context, id IotSecuritySolutionId) (result GetAllOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/analyticsModels", id.ID()),
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
		Values *[]IoTSecuritySolutionAnalyticsModel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetAllComplete retrieves all the results into a single object
func (c IoTSecuritySolutionsAnalyticsClient) GetAllComplete(ctx context.Context, id IotSecuritySolutionId) (GetAllCompleteResult, error) {
	return c.GetAllCompleteMatchingPredicate(ctx, id, IoTSecuritySolutionAnalyticsModelOperationPredicate{})
}

// GetAllCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IoTSecuritySolutionsAnalyticsClient) GetAllCompleteMatchingPredicate(ctx context.Context, id IotSecuritySolutionId, predicate IoTSecuritySolutionAnalyticsModelOperationPredicate) (result GetAllCompleteResult, err error) {
	items := make([]IoTSecuritySolutionAnalyticsModel, 0)

	resp, err := c.GetAll(ctx, id)
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

	result = GetAllCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
