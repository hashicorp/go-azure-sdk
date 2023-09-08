package userjoinedteam

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

type ListUserByIdJoinedTeamsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TeamCollectionResponse
}

type ListUserByIdJoinedTeamsCompleteResult struct {
	Items []models.TeamCollectionResponse
}

// ListUserByIdJoinedTeams ...
func (c UserJoinedTeamClient) ListUserByIdJoinedTeams(ctx context.Context, id UserId) (result ListUserByIdJoinedTeamsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/joinedTeams", id.ID()),
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
		Values *[]models.TeamCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdJoinedTeamsComplete retrieves all the results into a single object
func (c UserJoinedTeamClient) ListUserByIdJoinedTeamsComplete(ctx context.Context, id UserId) (ListUserByIdJoinedTeamsCompleteResult, error) {
	return c.ListUserByIdJoinedTeamsCompleteMatchingPredicate(ctx, id, models.TeamCollectionResponseOperationPredicate{})
}

// ListUserByIdJoinedTeamsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserJoinedTeamClient) ListUserByIdJoinedTeamsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.TeamCollectionResponseOperationPredicate) (result ListUserByIdJoinedTeamsCompleteResult, err error) {
	items := make([]models.TeamCollectionResponse, 0)

	resp, err := c.ListUserByIdJoinedTeams(ctx, id)
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

	result = ListUserByIdJoinedTeamsCompleteResult{
		Items: items,
	}
	return
}
