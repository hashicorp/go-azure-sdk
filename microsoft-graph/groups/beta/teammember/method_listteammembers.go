package teammember

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListTeamMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ConversationMember
}

type ListTeamMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ConversationMember
}

type ListTeamMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamMembers ...
func (c TeamMemberClient) ListTeamMembers(ctx context.Context, id GroupId) (result ListTeamMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTeamMembersCustomPager{},
		Path:       fmt.Sprintf("%s/team/members", id.ID()),
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
		Values *[]beta.ConversationMember `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamMembersComplete retrieves all the results into a single object
func (c TeamMemberClient) ListTeamMembersComplete(ctx context.Context, id GroupId) (ListTeamMembersCompleteResult, error) {
	return c.ListTeamMembersCompleteMatchingPredicate(ctx, id, ConversationMemberOperationPredicate{})
}

// ListTeamMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamMemberClient) ListTeamMembersCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate ConversationMemberOperationPredicate) (result ListTeamMembersCompleteResult, err error) {
	items := make([]beta.ConversationMember, 0)

	resp, err := c.ListTeamMembers(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = ListTeamMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
