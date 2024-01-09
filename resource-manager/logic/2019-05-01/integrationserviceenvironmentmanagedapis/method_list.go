package integrationserviceenvironmentmanagedapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]IntegrationServiceEnvironmentManagedApi
}

type ListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []IntegrationServiceEnvironmentManagedApi
}

// List ...
func (c IntegrationServiceEnvironmentManagedApisClient) List(ctx context.Context, id IntegrationServiceEnvironmentId) (result ListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/managedApis", id.ID()),
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
		Values *[]IntegrationServiceEnvironmentManagedApi `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComplete retrieves all the results into a single object
func (c IntegrationServiceEnvironmentManagedApisClient) ListComplete(ctx context.Context, id IntegrationServiceEnvironmentId) (ListCompleteResult, error) {
	return c.ListCompleteMatchingPredicate(ctx, id, IntegrationServiceEnvironmentManagedApiOperationPredicate{})
}

// ListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IntegrationServiceEnvironmentManagedApisClient) ListCompleteMatchingPredicate(ctx context.Context, id IntegrationServiceEnvironmentId, predicate IntegrationServiceEnvironmentManagedApiOperationPredicate) (result ListCompleteResult, err error) {
	items := make([]IntegrationServiceEnvironmentManagedApi, 0)

	resp, err := c.List(ctx, id)
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

	result = ListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
