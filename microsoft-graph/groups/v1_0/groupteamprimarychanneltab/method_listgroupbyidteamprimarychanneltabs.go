package groupteamprimarychanneltab

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

type ListGroupByIdTeamPrimaryChannelTabsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TeamsTabCollectionResponse
}

type ListGroupByIdTeamPrimaryChannelTabsCompleteResult struct {
	Items []models.TeamsTabCollectionResponse
}

// ListGroupByIdTeamPrimaryChannelTabs ...
func (c GroupTeamPrimaryChannelTabClient) ListGroupByIdTeamPrimaryChannelTabs(ctx context.Context, id GroupId) (result ListGroupByIdTeamPrimaryChannelTabsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/team/primaryChannel/tabs", id.ID()),
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

// ListGroupByIdTeamPrimaryChannelTabsComplete retrieves all the results into a single object
func (c GroupTeamPrimaryChannelTabClient) ListGroupByIdTeamPrimaryChannelTabsComplete(ctx context.Context, id GroupId) (ListGroupByIdTeamPrimaryChannelTabsCompleteResult, error) {
	return c.ListGroupByIdTeamPrimaryChannelTabsCompleteMatchingPredicate(ctx, id, models.TeamsTabCollectionResponseOperationPredicate{})
}

// ListGroupByIdTeamPrimaryChannelTabsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupTeamPrimaryChannelTabClient) ListGroupByIdTeamPrimaryChannelTabsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.TeamsTabCollectionResponseOperationPredicate) (result ListGroupByIdTeamPrimaryChannelTabsCompleteResult, err error) {
	items := make([]models.TeamsTabCollectionResponse, 0)

	resp, err := c.ListGroupByIdTeamPrimaryChannelTabs(ctx, id)
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

	result = ListGroupByIdTeamPrimaryChannelTabsCompleteResult{
		Items: items,
	}
	return
}
