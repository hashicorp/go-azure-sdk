package usertodolisttasklinkedresource

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

type ListUserByIdTodoListByIdTaskByIdLinkedResourcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.LinkedResourceCollectionResponse
}

type ListUserByIdTodoListByIdTaskByIdLinkedResourcesCompleteResult struct {
	Items []models.LinkedResourceCollectionResponse
}

// ListUserByIdTodoListByIdTaskByIdLinkedResources ...
func (c UserTodoListTaskLinkedResourceClient) ListUserByIdTodoListByIdTaskByIdLinkedResources(ctx context.Context, id UserTodoListTaskId) (result ListUserByIdTodoListByIdTaskByIdLinkedResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/linkedResources", id.ID()),
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
		Values *[]models.LinkedResourceCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdTodoListByIdTaskByIdLinkedResourcesComplete retrieves all the results into a single object
func (c UserTodoListTaskLinkedResourceClient) ListUserByIdTodoListByIdTaskByIdLinkedResourcesComplete(ctx context.Context, id UserTodoListTaskId) (ListUserByIdTodoListByIdTaskByIdLinkedResourcesCompleteResult, error) {
	return c.ListUserByIdTodoListByIdTaskByIdLinkedResourcesCompleteMatchingPredicate(ctx, id, models.LinkedResourceCollectionResponseOperationPredicate{})
}

// ListUserByIdTodoListByIdTaskByIdLinkedResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserTodoListTaskLinkedResourceClient) ListUserByIdTodoListByIdTaskByIdLinkedResourcesCompleteMatchingPredicate(ctx context.Context, id UserTodoListTaskId, predicate models.LinkedResourceCollectionResponseOperationPredicate) (result ListUserByIdTodoListByIdTaskByIdLinkedResourcesCompleteResult, err error) {
	items := make([]models.LinkedResourceCollectionResponse, 0)

	resp, err := c.ListUserByIdTodoListByIdTaskByIdLinkedResources(ctx, id)
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

	result = ListUserByIdTodoListByIdTaskByIdLinkedResourcesCompleteResult{
		Items: items,
	}
	return
}
