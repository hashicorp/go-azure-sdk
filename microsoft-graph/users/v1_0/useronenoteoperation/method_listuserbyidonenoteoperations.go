package useronenoteoperation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserByIdOnenoteOperationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnenoteOperationCollectionResponse
}

type ListUserByIdOnenoteOperationsCompleteResult struct {
	Items []models.OnenoteOperationCollectionResponse
}

// ListUserByIdOnenoteOperations ...
func (c UserOnenoteOperationClient) ListUserByIdOnenoteOperations(ctx context.Context, id UserId) (result ListUserByIdOnenoteOperationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/onenote/operations", id.ID()),
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
		Values *[]models.OnenoteOperationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdOnenoteOperationsComplete retrieves all the results into a single object
func (c UserOnenoteOperationClient) ListUserByIdOnenoteOperationsComplete(ctx context.Context, id UserId) (ListUserByIdOnenoteOperationsCompleteResult, error) {
	return c.ListUserByIdOnenoteOperationsCompleteMatchingPredicate(ctx, id, models.OnenoteOperationCollectionResponseOperationPredicate{})
}

// ListUserByIdOnenoteOperationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserOnenoteOperationClient) ListUserByIdOnenoteOperationsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.OnenoteOperationCollectionResponseOperationPredicate) (result ListUserByIdOnenoteOperationsCompleteResult, err error) {
	items := make([]models.OnenoteOperationCollectionResponse, 0)

	resp, err := c.ListUserByIdOnenoteOperations(ctx, id)
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

	result = ListUserByIdOnenoteOperationsCompleteResult{
		Items: items,
	}
	return
}
