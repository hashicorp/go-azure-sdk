package joinedteamtag

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

type ListJoinedTeamTagsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.TeamworkTag
}

type ListJoinedTeamTagsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.TeamworkTag
}

type ListJoinedTeamTagsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamTagsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamTags ...
func (c JoinedTeamTagClient) ListJoinedTeamTags(ctx context.Context, id MeJoinedTeamId) (result ListJoinedTeamTagsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListJoinedTeamTagsCustomPager{},
		Path:       fmt.Sprintf("%s/tags", id.ID()),
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
		Values *[]stable.TeamworkTag `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamTagsComplete retrieves all the results into a single object
func (c JoinedTeamTagClient) ListJoinedTeamTagsComplete(ctx context.Context, id MeJoinedTeamId) (ListJoinedTeamTagsCompleteResult, error) {
	return c.ListJoinedTeamTagsCompleteMatchingPredicate(ctx, id, TeamworkTagOperationPredicate{})
}

// ListJoinedTeamTagsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamTagClient) ListJoinedTeamTagsCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamId, predicate TeamworkTagOperationPredicate) (result ListJoinedTeamTagsCompleteResult, err error) {
	items := make([]stable.TeamworkTag, 0)

	resp, err := c.ListJoinedTeamTags(ctx, id)
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

	result = ListJoinedTeamTagsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
