package useronenotenotebook

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

type ListUserByIdOnenoteNotebooksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.NotebookCollectionResponse
}

type ListUserByIdOnenoteNotebooksCompleteResult struct {
	Items []models.NotebookCollectionResponse
}

// ListUserByIdOnenoteNotebooks ...
func (c UserOnenoteNotebookClient) ListUserByIdOnenoteNotebooks(ctx context.Context, id UserId) (result ListUserByIdOnenoteNotebooksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/onenote/notebooks", id.ID()),
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
		Values *[]models.NotebookCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdOnenoteNotebooksComplete retrieves all the results into a single object
func (c UserOnenoteNotebookClient) ListUserByIdOnenoteNotebooksComplete(ctx context.Context, id UserId) (ListUserByIdOnenoteNotebooksCompleteResult, error) {
	return c.ListUserByIdOnenoteNotebooksCompleteMatchingPredicate(ctx, id, models.NotebookCollectionResponseOperationPredicate{})
}

// ListUserByIdOnenoteNotebooksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserOnenoteNotebookClient) ListUserByIdOnenoteNotebooksCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.NotebookCollectionResponseOperationPredicate) (result ListUserByIdOnenoteNotebooksCompleteResult, err error) {
	items := make([]models.NotebookCollectionResponse, 0)

	resp, err := c.ListUserByIdOnenoteNotebooks(ctx, id)
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

	result = ListUserByIdOnenoteNotebooksCompleteResult{
		Items: items,
	}
	return
}
