package joinedteamprimarychannelsharedwithteam

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

type ListJoinedTeamPrimaryChannelSharedWithTeamsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.SharedWithChannelTeamInfo
}

type ListJoinedTeamPrimaryChannelSharedWithTeamsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.SharedWithChannelTeamInfo
}

type ListJoinedTeamPrimaryChannelSharedWithTeamsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamPrimaryChannelSharedWithTeamsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamPrimaryChannelSharedWithTeams ...
func (c JoinedTeamPrimaryChannelSharedWithTeamClient) ListJoinedTeamPrimaryChannelSharedWithTeams(ctx context.Context, id MeJoinedTeamId) (result ListJoinedTeamPrimaryChannelSharedWithTeamsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListJoinedTeamPrimaryChannelSharedWithTeamsCustomPager{},
		Path:       fmt.Sprintf("%s/primaryChannel/sharedWithTeams", id.ID()),
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
		Values *[]stable.SharedWithChannelTeamInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamPrimaryChannelSharedWithTeamsComplete retrieves all the results into a single object
func (c JoinedTeamPrimaryChannelSharedWithTeamClient) ListJoinedTeamPrimaryChannelSharedWithTeamsComplete(ctx context.Context, id MeJoinedTeamId) (ListJoinedTeamPrimaryChannelSharedWithTeamsCompleteResult, error) {
	return c.ListJoinedTeamPrimaryChannelSharedWithTeamsCompleteMatchingPredicate(ctx, id, SharedWithChannelTeamInfoOperationPredicate{})
}

// ListJoinedTeamPrimaryChannelSharedWithTeamsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamPrimaryChannelSharedWithTeamClient) ListJoinedTeamPrimaryChannelSharedWithTeamsCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamId, predicate SharedWithChannelTeamInfoOperationPredicate) (result ListJoinedTeamPrimaryChannelSharedWithTeamsCompleteResult, err error) {
	items := make([]stable.SharedWithChannelTeamInfo, 0)

	resp, err := c.ListJoinedTeamPrimaryChannelSharedWithTeams(ctx, id)
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

	result = ListJoinedTeamPrimaryChannelSharedWithTeamsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
