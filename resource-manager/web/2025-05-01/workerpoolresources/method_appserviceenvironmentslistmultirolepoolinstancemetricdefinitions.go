package workerpoolresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppServiceEnvironmentsListMultiRolePoolInstanceMetricDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ResourceMetricDefinition
}

type AppServiceEnvironmentsListMultiRolePoolInstanceMetricDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ResourceMetricDefinition
}

type AppServiceEnvironmentsListMultiRolePoolInstanceMetricDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *AppServiceEnvironmentsListMultiRolePoolInstanceMetricDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AppServiceEnvironmentsListMultiRolePoolInstanceMetricDefinitions ...
func (c WorkerPoolResourcesClient) AppServiceEnvironmentsListMultiRolePoolInstanceMetricDefinitions(ctx context.Context, id DefaultInstanceId) (result AppServiceEnvironmentsListMultiRolePoolInstanceMetricDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &AppServiceEnvironmentsListMultiRolePoolInstanceMetricDefinitionsCustomPager{},
		Path:       fmt.Sprintf("%s/metricdefinitions", id.ID()),
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
		Values *[]ResourceMetricDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AppServiceEnvironmentsListMultiRolePoolInstanceMetricDefinitionsComplete retrieves all the results into a single object
func (c WorkerPoolResourcesClient) AppServiceEnvironmentsListMultiRolePoolInstanceMetricDefinitionsComplete(ctx context.Context, id DefaultInstanceId) (AppServiceEnvironmentsListMultiRolePoolInstanceMetricDefinitionsCompleteResult, error) {
	return c.AppServiceEnvironmentsListMultiRolePoolInstanceMetricDefinitionsCompleteMatchingPredicate(ctx, id, ResourceMetricDefinitionOperationPredicate{})
}

// AppServiceEnvironmentsListMultiRolePoolInstanceMetricDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WorkerPoolResourcesClient) AppServiceEnvironmentsListMultiRolePoolInstanceMetricDefinitionsCompleteMatchingPredicate(ctx context.Context, id DefaultInstanceId, predicate ResourceMetricDefinitionOperationPredicate) (result AppServiceEnvironmentsListMultiRolePoolInstanceMetricDefinitionsCompleteResult, err error) {
	items := make([]ResourceMetricDefinition, 0)

	resp, err := c.AppServiceEnvironmentsListMultiRolePoolInstanceMetricDefinitions(ctx, id)
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

	result = AppServiceEnvironmentsListMultiRolePoolInstanceMetricDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
