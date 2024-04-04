package devops

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureDevOpsProjectsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AzureDevOpsProject
}

type AzureDevOpsProjectsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AzureDevOpsProject
}

// AzureDevOpsProjectsList ...
func (c DevOpsClient) AzureDevOpsProjectsList(ctx context.Context, id AzureDevOpsOrgId) (result AzureDevOpsProjectsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/projects", id.ID()),
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
		Values *[]AzureDevOpsProject `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AzureDevOpsProjectsListComplete retrieves all the results into a single object
func (c DevOpsClient) AzureDevOpsProjectsListComplete(ctx context.Context, id AzureDevOpsOrgId) (AzureDevOpsProjectsListCompleteResult, error) {
	return c.AzureDevOpsProjectsListCompleteMatchingPredicate(ctx, id, AzureDevOpsProjectOperationPredicate{})
}

// AzureDevOpsProjectsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DevOpsClient) AzureDevOpsProjectsListCompleteMatchingPredicate(ctx context.Context, id AzureDevOpsOrgId, predicate AzureDevOpsProjectOperationPredicate) (result AzureDevOpsProjectsListCompleteResult, err error) {
	items := make([]AzureDevOpsProject, 0)

	resp, err := c.AzureDevOpsProjectsList(ctx, id)
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

	result = AzureDevOpsProjectsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
