package meauthenticationoperation

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

type ListMeAuthenticationOperationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.LongRunningOperationCollectionResponse
}

type ListMeAuthenticationOperationsCompleteResult struct {
	Items []models.LongRunningOperationCollectionResponse
}

// ListMeAuthenticationOperations ...
func (c MeAuthenticationOperationClient) ListMeAuthenticationOperations(ctx context.Context) (result ListMeAuthenticationOperationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/authentication/operations",
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

// ListMeAuthenticationOperationsComplete retrieves all the results into a single object
func (c MeAuthenticationOperationClient) ListMeAuthenticationOperationsComplete(ctx context.Context) (ListMeAuthenticationOperationsCompleteResult, error) {
	return c.ListMeAuthenticationOperationsCompleteMatchingPredicate(ctx, models.LongRunningOperationCollectionResponseOperationPredicate{})
}

// ListMeAuthenticationOperationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeAuthenticationOperationClient) ListMeAuthenticationOperationsCompleteMatchingPredicate(ctx context.Context, predicate models.LongRunningOperationCollectionResponseOperationPredicate) (result ListMeAuthenticationOperationsCompleteResult, err error) {
	items := make([]models.LongRunningOperationCollectionResponse, 0)

	resp, err := c.ListMeAuthenticationOperations(ctx)
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

	result = ListMeAuthenticationOperationsCompleteResult{
		Items: items,
	}
	return
}
