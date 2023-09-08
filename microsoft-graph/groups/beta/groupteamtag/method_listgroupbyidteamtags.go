package groupteamtag

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

type ListGroupByIdTeamTagsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TeamworkTagCollectionResponse
}

type ListGroupByIdTeamTagsCompleteResult struct {
	Items []models.TeamworkTagCollectionResponse
}

// ListGroupByIdTeamTags ...
func (c GroupTeamTagClient) ListGroupByIdTeamTags(ctx context.Context, id GroupId) (result ListGroupByIdTeamTagsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/team/tags", id.ID()),
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
		Values *[]models.TeamworkTagCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdTeamTagsComplete retrieves all the results into a single object
func (c GroupTeamTagClient) ListGroupByIdTeamTagsComplete(ctx context.Context, id GroupId) (ListGroupByIdTeamTagsCompleteResult, error) {
	return c.ListGroupByIdTeamTagsCompleteMatchingPredicate(ctx, id, models.TeamworkTagCollectionResponseOperationPredicate{})
}

// ListGroupByIdTeamTagsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupTeamTagClient) ListGroupByIdTeamTagsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.TeamworkTagCollectionResponseOperationPredicate) (result ListGroupByIdTeamTagsCompleteResult, err error) {
	items := make([]models.TeamworkTagCollectionResponse, 0)

	resp, err := c.ListGroupByIdTeamTags(ctx, id)
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

	result = ListGroupByIdTeamTagsCompleteResult{
		Items: items,
	}
	return
}
