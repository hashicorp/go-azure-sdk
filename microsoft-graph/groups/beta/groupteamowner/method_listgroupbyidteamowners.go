package groupteamowner

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

type ListGroupByIdTeamOwnersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UserCollectionResponse
}

type ListGroupByIdTeamOwnersCompleteResult struct {
	Items []models.UserCollectionResponse
}

// ListGroupByIdTeamOwners ...
func (c GroupTeamOwnerClient) ListGroupByIdTeamOwners(ctx context.Context, id GroupId) (result ListGroupByIdTeamOwnersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/team/owners", id.ID()),
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
		Values *[]models.UserCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdTeamOwnersComplete retrieves all the results into a single object
func (c GroupTeamOwnerClient) ListGroupByIdTeamOwnersComplete(ctx context.Context, id GroupId) (ListGroupByIdTeamOwnersCompleteResult, error) {
	return c.ListGroupByIdTeamOwnersCompleteMatchingPredicate(ctx, id, models.UserCollectionResponseOperationPredicate{})
}

// ListGroupByIdTeamOwnersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupTeamOwnerClient) ListGroupByIdTeamOwnersCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.UserCollectionResponseOperationPredicate) (result ListGroupByIdTeamOwnersCompleteResult, err error) {
	items := make([]models.UserCollectionResponse, 0)

	resp, err := c.ListGroupByIdTeamOwners(ctx, id)
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

	result = ListGroupByIdTeamOwnersCompleteResult{
		Items: items,
	}
	return
}
