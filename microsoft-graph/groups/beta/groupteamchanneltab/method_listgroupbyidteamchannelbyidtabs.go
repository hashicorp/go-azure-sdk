package groupteamchanneltab

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

type ListGroupByIdTeamChannelByIdTabsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TeamsTabCollectionResponse
}

type ListGroupByIdTeamChannelByIdTabsCompleteResult struct {
	Items []models.TeamsTabCollectionResponse
}

// ListGroupByIdTeamChannelByIdTabs ...
func (c GroupTeamChannelTabClient) ListGroupByIdTeamChannelByIdTabs(ctx context.Context, id GroupTeamChannelId) (result ListGroupByIdTeamChannelByIdTabsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/tabs", id.ID()),
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
		Values *[]models.TeamsTabCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdTeamChannelByIdTabsComplete retrieves all the results into a single object
func (c GroupTeamChannelTabClient) ListGroupByIdTeamChannelByIdTabsComplete(ctx context.Context, id GroupTeamChannelId) (ListGroupByIdTeamChannelByIdTabsCompleteResult, error) {
	return c.ListGroupByIdTeamChannelByIdTabsCompleteMatchingPredicate(ctx, id, models.TeamsTabCollectionResponseOperationPredicate{})
}

// ListGroupByIdTeamChannelByIdTabsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupTeamChannelTabClient) ListGroupByIdTeamChannelByIdTabsCompleteMatchingPredicate(ctx context.Context, id GroupTeamChannelId, predicate models.TeamsTabCollectionResponseOperationPredicate) (result ListGroupByIdTeamChannelByIdTabsCompleteResult, err error) {
	items := make([]models.TeamsTabCollectionResponse, 0)

	resp, err := c.ListGroupByIdTeamChannelByIdTabs(ctx, id)
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

	result = ListGroupByIdTeamChannelByIdTabsCompleteResult{
		Items: items,
	}
	return
}
