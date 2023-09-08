package groupteamprimarychannelmember

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

type ListGroupByIdTeamPrimaryChannelMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ConversationMemberCollectionResponse
}

type ListGroupByIdTeamPrimaryChannelMembersCompleteResult struct {
	Items []models.ConversationMemberCollectionResponse
}

// ListGroupByIdTeamPrimaryChannelMembers ...
func (c GroupTeamPrimaryChannelMemberClient) ListGroupByIdTeamPrimaryChannelMembers(ctx context.Context, id GroupId) (result ListGroupByIdTeamPrimaryChannelMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/team/primaryChannel/members", id.ID()),
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
		Values *[]models.ConversationMemberCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdTeamPrimaryChannelMembersComplete retrieves all the results into a single object
func (c GroupTeamPrimaryChannelMemberClient) ListGroupByIdTeamPrimaryChannelMembersComplete(ctx context.Context, id GroupId) (ListGroupByIdTeamPrimaryChannelMembersCompleteResult, error) {
	return c.ListGroupByIdTeamPrimaryChannelMembersCompleteMatchingPredicate(ctx, id, models.ConversationMemberCollectionResponseOperationPredicate{})
}

// ListGroupByIdTeamPrimaryChannelMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupTeamPrimaryChannelMemberClient) ListGroupByIdTeamPrimaryChannelMembersCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.ConversationMemberCollectionResponseOperationPredicate) (result ListGroupByIdTeamPrimaryChannelMembersCompleteResult, err error) {
	items := make([]models.ConversationMemberCollectionResponse, 0)

	resp, err := c.ListGroupByIdTeamPrimaryChannelMembers(ctx, id)
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

	result = ListGroupByIdTeamPrimaryChannelMembersCompleteResult{
		Items: items,
	}
	return
}
