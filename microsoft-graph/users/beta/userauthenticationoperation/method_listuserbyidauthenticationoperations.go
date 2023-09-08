package userauthenticationoperation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserByIdAuthenticationOperationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.LongRunningOperationCollectionResponse
}

type ListUserByIdAuthenticationOperationsCompleteResult struct {
	Items []models.LongRunningOperationCollectionResponse
}

// ListUserByIdAuthenticationOperations ...
func (c UserAuthenticationOperationClient) ListUserByIdAuthenticationOperations(ctx context.Context, id UserId) (result ListUserByIdAuthenticationOperationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/authentication/operations", id.ID()),
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
		Values *[]models.LongRunningOperationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdAuthenticationOperationsComplete retrieves all the results into a single object
func (c UserAuthenticationOperationClient) ListUserByIdAuthenticationOperationsComplete(ctx context.Context, id UserId) (ListUserByIdAuthenticationOperationsCompleteResult, error) {
	return c.ListUserByIdAuthenticationOperationsCompleteMatchingPredicate(ctx, id, models.LongRunningOperationCollectionResponseOperationPredicate{})
}

// ListUserByIdAuthenticationOperationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserAuthenticationOperationClient) ListUserByIdAuthenticationOperationsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.LongRunningOperationCollectionResponseOperationPredicate) (result ListUserByIdAuthenticationOperationsCompleteResult, err error) {
	items := make([]models.LongRunningOperationCollectionResponse, 0)

	resp, err := c.ListUserByIdAuthenticationOperations(ctx, id)
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

	result = ListUserByIdAuthenticationOperationsCompleteResult{
		Items: items,
	}
	return
}
