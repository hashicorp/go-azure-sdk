package sparkconfigurations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SparkConfigurationGetSparkConfigurationsByWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SparkConfigurationResource
}

type SparkConfigurationGetSparkConfigurationsByWorkspaceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SparkConfigurationResource
}

type SparkConfigurationGetSparkConfigurationsByWorkspaceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *SparkConfigurationGetSparkConfigurationsByWorkspaceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// SparkConfigurationGetSparkConfigurationsByWorkspace ...
func (c SparkConfigurationsClient) SparkConfigurationGetSparkConfigurationsByWorkspace(ctx context.Context) (result SparkConfigurationGetSparkConfigurationsByWorkspaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &SparkConfigurationGetSparkConfigurationsByWorkspaceCustomPager{},
		Path:       "/sparkConfigurations",
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
		Values *[]SparkConfigurationResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SparkConfigurationGetSparkConfigurationsByWorkspaceComplete retrieves all the results into a single object
func (c SparkConfigurationsClient) SparkConfigurationGetSparkConfigurationsByWorkspaceComplete(ctx context.Context) (SparkConfigurationGetSparkConfigurationsByWorkspaceCompleteResult, error) {
	return c.SparkConfigurationGetSparkConfigurationsByWorkspaceCompleteMatchingPredicate(ctx, SparkConfigurationResourceOperationPredicate{})
}

// SparkConfigurationGetSparkConfigurationsByWorkspaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SparkConfigurationsClient) SparkConfigurationGetSparkConfigurationsByWorkspaceCompleteMatchingPredicate(ctx context.Context, predicate SparkConfigurationResourceOperationPredicate) (result SparkConfigurationGetSparkConfigurationsByWorkspaceCompleteResult, err error) {
	items := make([]SparkConfigurationResource, 0)

	resp, err := c.SparkConfigurationGetSparkConfigurationsByWorkspace(ctx)
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

	result = SparkConfigurationGetSparkConfigurationsByWorkspaceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
