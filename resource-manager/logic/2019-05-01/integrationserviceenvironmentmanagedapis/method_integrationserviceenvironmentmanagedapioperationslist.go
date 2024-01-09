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

type IntegrationServiceEnvironmentManagedApiOperationsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ApiOperation
}

type IntegrationServiceEnvironmentManagedApiOperationsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ApiOperation
}

// IntegrationServiceEnvironmentManagedApiOperationsList ...
func (c IntegrationServiceEnvironmentManagedApisClient) IntegrationServiceEnvironmentManagedApiOperationsList(ctx context.Context, id ManagedApiId) (result IntegrationServiceEnvironmentManagedApiOperationsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/apiOperations", id.ID()),
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
		Values *[]ApiOperation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// IntegrationServiceEnvironmentManagedApiOperationsListComplete retrieves all the results into a single object
func (c IntegrationServiceEnvironmentManagedApisClient) IntegrationServiceEnvironmentManagedApiOperationsListComplete(ctx context.Context, id ManagedApiId) (IntegrationServiceEnvironmentManagedApiOperationsListCompleteResult, error) {
	return c.IntegrationServiceEnvironmentManagedApiOperationsListCompleteMatchingPredicate(ctx, id, ApiOperationOperationPredicate{})
}

// IntegrationServiceEnvironmentManagedApiOperationsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IntegrationServiceEnvironmentManagedApisClient) IntegrationServiceEnvironmentManagedApiOperationsListCompleteMatchingPredicate(ctx context.Context, id ManagedApiId, predicate ApiOperationOperationPredicate) (result IntegrationServiceEnvironmentManagedApiOperationsListCompleteResult, err error) {
	items := make([]ApiOperation, 0)

	resp, err := c.IntegrationServiceEnvironmentManagedApiOperationsList(ctx, id)
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

	result = IntegrationServiceEnvironmentManagedApiOperationsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
