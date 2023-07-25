package appserviceenvironments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMultiRolePoolsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WorkerPoolResource
}

type ListMultiRolePoolsCompleteResult struct {
	Items []WorkerPoolResource
}

// ListMultiRolePools ...
func (c AppServiceEnvironmentsClient) ListMultiRolePools(ctx context.Context, id HostingEnvironmentId) (result ListMultiRolePoolsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/multiRolePools", id.ID()),
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
		Values *[]WorkerPoolResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMultiRolePoolsComplete retrieves all the results into a single object
func (c AppServiceEnvironmentsClient) ListMultiRolePoolsComplete(ctx context.Context, id HostingEnvironmentId) (ListMultiRolePoolsCompleteResult, error) {
	return c.ListMultiRolePoolsCompleteMatchingPredicate(ctx, id, WorkerPoolResourceOperationPredicate{})
}

// ListMultiRolePoolsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppServiceEnvironmentsClient) ListMultiRolePoolsCompleteMatchingPredicate(ctx context.Context, id HostingEnvironmentId, predicate WorkerPoolResourceOperationPredicate) (result ListMultiRolePoolsCompleteResult, err error) {
	items := make([]WorkerPoolResource, 0)

	resp, err := c.ListMultiRolePools(ctx, id)
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

	result = ListMultiRolePoolsCompleteResult{
		Items: items,
	}
	return
}
