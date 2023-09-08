package grouponenoteoperation

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

type ListGroupByIdOnenoteOperationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnenoteOperationCollectionResponse
}

type ListGroupByIdOnenoteOperationsCompleteResult struct {
	Items []models.OnenoteOperationCollectionResponse
}

// ListGroupByIdOnenoteOperations ...
func (c GroupOnenoteOperationClient) ListGroupByIdOnenoteOperations(ctx context.Context, id GroupId) (result ListGroupByIdOnenoteOperationsOperationResponse, err error) {
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

// ListGroupByIdOnenoteOperationsComplete retrieves all the results into a single object
func (c GroupOnenoteOperationClient) ListGroupByIdOnenoteOperationsComplete(ctx context.Context, id GroupId) (ListGroupByIdOnenoteOperationsCompleteResult, error) {
	return c.ListGroupByIdOnenoteOperationsCompleteMatchingPredicate(ctx, id, models.OnenoteOperationCollectionResponseOperationPredicate{})
}

// ListGroupByIdOnenoteOperationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupOnenoteOperationClient) ListGroupByIdOnenoteOperationsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.OnenoteOperationCollectionResponseOperationPredicate) (result ListGroupByIdOnenoteOperationsCompleteResult, err error) {
	items := make([]models.OnenoteOperationCollectionResponse, 0)

	resp, err := c.ListGroupByIdOnenoteOperations(ctx, id)
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

	result = ListGroupByIdOnenoteOperationsCompleteResult{
		Items: items,
	}
	return
}
