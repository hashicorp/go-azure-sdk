package meteamworkassociatedteam

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

type ListMeTeamworkAssociatedTeamsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AssociatedTeamInfoCollectionResponse
}

type ListMeTeamworkAssociatedTeamsCompleteResult struct {
	Items []models.AssociatedTeamInfoCollectionResponse
}

// ListMeTeamworkAssociatedTeams ...
func (c MeTeamworkAssociatedTeamClient) ListMeTeamworkAssociatedTeams(ctx context.Context) (result ListMeTeamworkAssociatedTeamsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/teamwork/associatedTeams",
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

// ListMeTeamworkAssociatedTeamsComplete retrieves all the results into a single object
func (c MeTeamworkAssociatedTeamClient) ListMeTeamworkAssociatedTeamsComplete(ctx context.Context) (ListMeTeamworkAssociatedTeamsCompleteResult, error) {
	return c.ListMeTeamworkAssociatedTeamsCompleteMatchingPredicate(ctx, models.AssociatedTeamInfoCollectionResponseOperationPredicate{})
}

// ListMeTeamworkAssociatedTeamsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeTeamworkAssociatedTeamClient) ListMeTeamworkAssociatedTeamsCompleteMatchingPredicate(ctx context.Context, predicate models.AssociatedTeamInfoCollectionResponseOperationPredicate) (result ListMeTeamworkAssociatedTeamsCompleteResult, err error) {
	items := make([]models.AssociatedTeamInfoCollectionResponse, 0)

	resp, err := c.ListMeTeamworkAssociatedTeams(ctx)
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

	result = ListMeTeamworkAssociatedTeamsCompleteResult{
		Items: items,
	}
	return
}
