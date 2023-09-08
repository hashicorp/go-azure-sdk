package userjoinedteamoperation

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

type ListUserByIdJoinedTeamByIdOperationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TeamsAsyncOperationCollectionResponse
}

type ListUserByIdJoinedTeamByIdOperationsCompleteResult struct {
	Items []models.TeamsAsyncOperationCollectionResponse
}

// ListUserByIdJoinedTeamByIdOperations ...
func (c UserJoinedTeamOperationClient) ListUserByIdJoinedTeamByIdOperations(ctx context.Context, id UserJoinedTeamId) (result ListUserByIdJoinedTeamByIdOperationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
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
		Values *[]models.TeamsAsyncOperationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdJoinedTeamByIdOperationsComplete retrieves all the results into a single object
func (c UserJoinedTeamOperationClient) ListUserByIdJoinedTeamByIdOperationsComplete(ctx context.Context, id UserJoinedTeamId) (ListUserByIdJoinedTeamByIdOperationsCompleteResult, error) {
	return c.ListUserByIdJoinedTeamByIdOperationsCompleteMatchingPredicate(ctx, id, models.TeamsAsyncOperationCollectionResponseOperationPredicate{})
}

// ListUserByIdJoinedTeamByIdOperationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserJoinedTeamOperationClient) ListUserByIdJoinedTeamByIdOperationsCompleteMatchingPredicate(ctx context.Context, id UserJoinedTeamId, predicate models.TeamsAsyncOperationCollectionResponseOperationPredicate) (result ListUserByIdJoinedTeamByIdOperationsCompleteResult, err error) {
	items := make([]models.TeamsAsyncOperationCollectionResponse, 0)

	resp, err := c.ListUserByIdJoinedTeamByIdOperations(ctx, id)
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

	result = ListUserByIdJoinedTeamByIdOperationsCompleteResult{
		Items: items,
	}
	return
}
