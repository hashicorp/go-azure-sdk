package openapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkItemConfigurationsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WorkItemConfiguration
}

type WorkItemConfigurationsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WorkItemConfiguration
}

type WorkItemConfigurationsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WorkItemConfigurationsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WorkItemConfigurationsList ...
func (c OpenapisClient) WorkItemConfigurationsList(ctx context.Context, id ComponentId) (result WorkItemConfigurationsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WorkItemConfigurationsListCustomPager{},
		Path:       fmt.Sprintf("%s/workItemConfigs", id.ID()),
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
		Values *[]WorkItemConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WorkItemConfigurationsListComplete retrieves all the results into a single object
func (c OpenapisClient) WorkItemConfigurationsListComplete(ctx context.Context, id ComponentId) (WorkItemConfigurationsListCompleteResult, error) {
	return c.WorkItemConfigurationsListCompleteMatchingPredicate(ctx, id, WorkItemConfigurationOperationPredicate{})
}

// WorkItemConfigurationsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) WorkItemConfigurationsListCompleteMatchingPredicate(ctx context.Context, id ComponentId, predicate WorkItemConfigurationOperationPredicate) (result WorkItemConfigurationsListCompleteResult, err error) {
	items := make([]WorkItemConfiguration, 0)

	resp, err := c.WorkItemConfigurationsList(ctx, id)
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

	result = WorkItemConfigurationsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
