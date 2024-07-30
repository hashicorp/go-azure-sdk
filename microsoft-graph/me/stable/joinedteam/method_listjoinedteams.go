package joinedteam

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

type ListJoinedTeamsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Team
}

type ListJoinedTeamsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Team
}

type ListJoinedTeamsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeams ...
func (c JoinedTeamClient) ListJoinedTeams(ctx context.Context) (result ListJoinedTeamsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListJoinedTeamsCustomPager{},
		Path:       "/me/joinedTeams",
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
		Values *[]stable.Team `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamsComplete retrieves all the results into a single object
func (c JoinedTeamClient) ListJoinedTeamsComplete(ctx context.Context) (ListJoinedTeamsCompleteResult, error) {
	return c.ListJoinedTeamsCompleteMatchingPredicate(ctx, TeamOperationPredicate{})
}

// ListJoinedTeamsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamClient) ListJoinedTeamsCompleteMatchingPredicate(ctx context.Context, predicate TeamOperationPredicate) (result ListJoinedTeamsCompleteResult, err error) {
	items := make([]stable.Team, 0)

	resp, err := c.ListJoinedTeams(ctx)
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

	result = ListJoinedTeamsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
