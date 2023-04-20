package containers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByStorageAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Container
}

type ListByStorageAccountCompleteResult struct {
	Items []Container
}

// ListByStorageAccount ...
func (c ContainersClient) ListByStorageAccount(ctx context.Context, id StorageAccountId) (result ListByStorageAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/containers", id.ID()),
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
		Values *[]Container `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByStorageAccountComplete retrieves all the results into a single object
func (c ContainersClient) ListByStorageAccountComplete(ctx context.Context, id StorageAccountId) (ListByStorageAccountCompleteResult, error) {
	return c.ListByStorageAccountCompleteMatchingPredicate(ctx, id, ContainerOperationPredicate{})
}

// ListByStorageAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ContainersClient) ListByStorageAccountCompleteMatchingPredicate(ctx context.Context, id StorageAccountId, predicate ContainerOperationPredicate) (result ListByStorageAccountCompleteResult, err error) {
	items := make([]Container, 0)

	resp, err := c.ListByStorageAccount(ctx, id)
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

	result = ListByStorageAccountCompleteResult{
		Items: items,
	}
	return
}
