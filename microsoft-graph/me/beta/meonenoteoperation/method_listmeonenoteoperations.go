package meonenoteoperation

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

type ListMeOnenoteOperationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnenoteOperationCollectionResponse
}

type ListMeOnenoteOperationsCompleteResult struct {
	Items []models.OnenoteOperationCollectionResponse
}

// ListMeOnenoteOperations ...
func (c MeOnenoteOperationClient) ListMeOnenoteOperations(ctx context.Context) (result ListMeOnenoteOperationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/onenote/operations",
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

// ListMeOnenoteOperationsComplete retrieves all the results into a single object
func (c MeOnenoteOperationClient) ListMeOnenoteOperationsComplete(ctx context.Context) (ListMeOnenoteOperationsCompleteResult, error) {
	return c.ListMeOnenoteOperationsCompleteMatchingPredicate(ctx, models.OnenoteOperationCollectionResponseOperationPredicate{})
}

// ListMeOnenoteOperationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOnenoteOperationClient) ListMeOnenoteOperationsCompleteMatchingPredicate(ctx context.Context, predicate models.OnenoteOperationCollectionResponseOperationPredicate) (result ListMeOnenoteOperationsCompleteResult, err error) {
	items := make([]models.OnenoteOperationCollectionResponse, 0)

	resp, err := c.ListMeOnenoteOperations(ctx)
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

	result = ListMeOnenoteOperationsCompleteResult{
		Items: items,
	}
	return
}
