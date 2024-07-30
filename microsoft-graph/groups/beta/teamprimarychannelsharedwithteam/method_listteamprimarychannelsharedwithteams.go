package teamprimarychannelsharedwithteam

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

type ListTeamPrimaryChannelSharedWithTeamsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SharedWithChannelTeamInfo
}

type ListTeamPrimaryChannelSharedWithTeamsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SharedWithChannelTeamInfo
}

type ListTeamPrimaryChannelSharedWithTeamsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamPrimaryChannelSharedWithTeamsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamPrimaryChannelSharedWithTeams ...
func (c TeamPrimaryChannelSharedWithTeamClient) ListTeamPrimaryChannelSharedWithTeams(ctx context.Context, id GroupId) (result ListTeamPrimaryChannelSharedWithTeamsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTeamPrimaryChannelSharedWithTeamsCustomPager{},
		Path:       fmt.Sprintf("%s/team/primaryChannel/sharedWithTeams", id.ID()),
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
		Values *[]beta.SharedWithChannelTeamInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamPrimaryChannelSharedWithTeamsComplete retrieves all the results into a single object
func (c TeamPrimaryChannelSharedWithTeamClient) ListTeamPrimaryChannelSharedWithTeamsComplete(ctx context.Context, id GroupId) (ListTeamPrimaryChannelSharedWithTeamsCompleteResult, error) {
	return c.ListTeamPrimaryChannelSharedWithTeamsCompleteMatchingPredicate(ctx, id, SharedWithChannelTeamInfoOperationPredicate{})
}

// ListTeamPrimaryChannelSharedWithTeamsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamPrimaryChannelSharedWithTeamClient) ListTeamPrimaryChannelSharedWithTeamsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate SharedWithChannelTeamInfoOperationPredicate) (result ListTeamPrimaryChannelSharedWithTeamsCompleteResult, err error) {
	items := make([]beta.SharedWithChannelTeamInfo, 0)

	resp, err := c.ListTeamPrimaryChannelSharedWithTeams(ctx, id)
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

	result = ListTeamPrimaryChannelSharedWithTeamsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
