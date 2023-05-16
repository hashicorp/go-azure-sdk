package storageaccounts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListStorageContainersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]StorageContainer
}

type ListStorageContainersCompleteResult struct {
	Items []StorageContainer
}

// ListStorageContainers ...
func (c StorageAccountsClient) ListStorageContainers(ctx context.Context, id StorageAccountId) (result ListStorageContainersOperationResponse, err error) {
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
		Values *[]StorageContainer `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListStorageContainersComplete retrieves all the results into a single object
func (c StorageAccountsClient) ListStorageContainersComplete(ctx context.Context, id StorageAccountId) (ListStorageContainersCompleteResult, error) {
	return c.ListStorageContainersCompleteMatchingPredicate(ctx, id, StorageContainerOperationPredicate{})
}

// ListStorageContainersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c StorageAccountsClient) ListStorageContainersCompleteMatchingPredicate(ctx context.Context, id StorageAccountId, predicate StorageContainerOperationPredicate) (result ListStorageContainersCompleteResult, err error) {
	items := make([]StorageContainer, 0)

	resp, err := c.ListStorageContainers(ctx, id)
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

	result = ListStorageContainersCompleteResult{
		Items: items,
	}
	return
}
