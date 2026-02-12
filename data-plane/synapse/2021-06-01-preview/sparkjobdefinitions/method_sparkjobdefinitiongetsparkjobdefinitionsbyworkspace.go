package sparkjobdefinitions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SparkJobDefinitionResource
}

type SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SparkJobDefinitionResource
}

type SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// SparkJobDefinitionGetSparkJobDefinitionsByWorkspace ...
func (c SparkJobDefinitionsClient) SparkJobDefinitionGetSparkJobDefinitionsByWorkspace(ctx context.Context) (result SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceCustomPager{},
		Path:       "/sparkJobDefinitions",
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
		Values *[]SparkJobDefinitionResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceComplete retrieves all the results into a single object
func (c SparkJobDefinitionsClient) SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceComplete(ctx context.Context) (SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceCompleteResult, error) {
	return c.SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceCompleteMatchingPredicate(ctx, SparkJobDefinitionResourceOperationPredicate{})
}

// SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SparkJobDefinitionsClient) SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceCompleteMatchingPredicate(ctx context.Context, predicate SparkJobDefinitionResourceOperationPredicate) (result SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceCompleteResult, err error) {
	items := make([]SparkJobDefinitionResource, 0)

	resp, err := c.SparkJobDefinitionGetSparkJobDefinitionsByWorkspace(ctx)
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

	result = SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
