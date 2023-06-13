package databases

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseOperationsListByDatabaseOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DatabaseOperation
}

type DatabaseOperationsListByDatabaseCompleteResult struct {
	Items []DatabaseOperation
}

// DatabaseOperationsListByDatabase ...
func (c DatabasesClient) DatabaseOperationsListByDatabase(ctx context.Context, id DatabaseId) (result DatabaseOperationsListByDatabaseOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/operations", id.ID()),
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
		Values *[]DatabaseOperation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// DatabaseOperationsListByDatabaseComplete retrieves all the results into a single object
func (c DatabasesClient) DatabaseOperationsListByDatabaseComplete(ctx context.Context, id DatabaseId) (DatabaseOperationsListByDatabaseCompleteResult, error) {
	return c.DatabaseOperationsListByDatabaseCompleteMatchingPredicate(ctx, id, DatabaseOperationOperationPredicate{})
}

// DatabaseOperationsListByDatabaseCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DatabasesClient) DatabaseOperationsListByDatabaseCompleteMatchingPredicate(ctx context.Context, id DatabaseId, predicate DatabaseOperationOperationPredicate) (result DatabaseOperationsListByDatabaseCompleteResult, err error) {
	items := make([]DatabaseOperation, 0)

	resp, err := c.DatabaseOperationsListByDatabase(ctx, id)
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

	result = DatabaseOperationsListByDatabaseCompleteResult{
		Items: items,
	}
	return
}
