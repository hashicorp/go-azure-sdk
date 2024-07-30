package teamtagmember

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListTeamTagMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.TeamworkTagMember
}

type ListTeamTagMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.TeamworkTagMember
}

type ListTeamTagMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamTagMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamTagMembers ...
func (c TeamTagMemberClient) ListTeamTagMembers(ctx context.Context, id GroupIdTeamTagId) (result ListTeamTagMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTeamTagMembersCustomPager{},
		Path:       fmt.Sprintf("%s/members", id.ID()),
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
		Values *[]stable.TeamworkTagMember `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamTagMembersComplete retrieves all the results into a single object
func (c TeamTagMemberClient) ListTeamTagMembersComplete(ctx context.Context, id GroupIdTeamTagId) (ListTeamTagMembersCompleteResult, error) {
	return c.ListTeamTagMembersCompleteMatchingPredicate(ctx, id, TeamworkTagMemberOperationPredicate{})
}

// ListTeamTagMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamTagMemberClient) ListTeamTagMembersCompleteMatchingPredicate(ctx context.Context, id GroupIdTeamTagId, predicate TeamworkTagMemberOperationPredicate) (result ListTeamTagMembersCompleteResult, err error) {
	items := make([]stable.TeamworkTagMember, 0)

	resp, err := c.ListTeamTagMembers(ctx, id)
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

	result = ListTeamTagMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
