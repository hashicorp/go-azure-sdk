package mejoinedteamoperation

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

type ListMeJoinedTeamByIdOperationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TeamsAsyncOperationCollectionResponse
}

type ListMeJoinedTeamByIdOperationsCompleteResult struct {
	Items []models.TeamsAsyncOperationCollectionResponse
}

// ListMeJoinedTeamByIdOperations ...
func (c MeJoinedTeamOperationClient) ListMeJoinedTeamByIdOperations(ctx context.Context, id MeJoinedTeamId) (result ListMeJoinedTeamByIdOperationsOperationResponse, err error) {
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

// ListMeJoinedTeamByIdOperationsComplete retrieves all the results into a single object
func (c MeJoinedTeamOperationClient) ListMeJoinedTeamByIdOperationsComplete(ctx context.Context, id MeJoinedTeamId) (ListMeJoinedTeamByIdOperationsCompleteResult, error) {
	return c.ListMeJoinedTeamByIdOperationsCompleteMatchingPredicate(ctx, id, models.TeamsAsyncOperationCollectionResponseOperationPredicate{})
}

// ListMeJoinedTeamByIdOperationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeJoinedTeamOperationClient) ListMeJoinedTeamByIdOperationsCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamId, predicate models.TeamsAsyncOperationCollectionResponseOperationPredicate) (result ListMeJoinedTeamByIdOperationsCompleteResult, err error) {
	items := make([]models.TeamsAsyncOperationCollectionResponse, 0)

	resp, err := c.ListMeJoinedTeamByIdOperations(ctx, id)
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

	result = ListMeJoinedTeamByIdOperationsCompleteResult{
		Items: items,
	}
	return
}
