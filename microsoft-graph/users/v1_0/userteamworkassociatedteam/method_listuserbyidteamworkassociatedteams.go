package userteamworkassociatedteam

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

type ListUserByIdTeamworkAssociatedTeamsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AssociatedTeamInfoCollectionResponse
}

type ListUserByIdTeamworkAssociatedTeamsCompleteResult struct {
	Items []models.AssociatedTeamInfoCollectionResponse
}

// ListUserByIdTeamworkAssociatedTeams ...
func (c UserTeamworkAssociatedTeamClient) ListUserByIdTeamworkAssociatedTeams(ctx context.Context, id UserId) (result ListUserByIdTeamworkAssociatedTeamsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/teamwork/associatedTeams", id.ID()),
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
		Values *[]models.AssociatedTeamInfoCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdTeamworkAssociatedTeamsComplete retrieves all the results into a single object
func (c UserTeamworkAssociatedTeamClient) ListUserByIdTeamworkAssociatedTeamsComplete(ctx context.Context, id UserId) (ListUserByIdTeamworkAssociatedTeamsCompleteResult, error) {
	return c.ListUserByIdTeamworkAssociatedTeamsCompleteMatchingPredicate(ctx, id, models.AssociatedTeamInfoCollectionResponseOperationPredicate{})
}

// ListUserByIdTeamworkAssociatedTeamsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserTeamworkAssociatedTeamClient) ListUserByIdTeamworkAssociatedTeamsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.AssociatedTeamInfoCollectionResponseOperationPredicate) (result ListUserByIdTeamworkAssociatedTeamsCompleteResult, err error) {
	items := make([]models.AssociatedTeamInfoCollectionResponse, 0)

	resp, err := c.ListUserByIdTeamworkAssociatedTeams(ctx, id)
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

	result = ListUserByIdTeamworkAssociatedTeamsCompleteResult{
		Items: items,
	}
	return
}
