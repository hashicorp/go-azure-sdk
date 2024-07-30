package teamworkassociatedteam

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

type ListTeamworkAssociatedTeamsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AssociatedTeamInfo
}

type ListTeamworkAssociatedTeamsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AssociatedTeamInfo
}

type ListTeamworkAssociatedTeamsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamworkAssociatedTeamsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamworkAssociatedTeams ...
func (c TeamworkAssociatedTeamClient) ListTeamworkAssociatedTeams(ctx context.Context) (result ListTeamworkAssociatedTeamsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTeamworkAssociatedTeamsCustomPager{},
		Path:       "/me/teamwork/associatedTeams",
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
		Values *[]beta.AssociatedTeamInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamworkAssociatedTeamsComplete retrieves all the results into a single object
func (c TeamworkAssociatedTeamClient) ListTeamworkAssociatedTeamsComplete(ctx context.Context) (ListTeamworkAssociatedTeamsCompleteResult, error) {
	return c.ListTeamworkAssociatedTeamsCompleteMatchingPredicate(ctx, AssociatedTeamInfoOperationPredicate{})
}

// ListTeamworkAssociatedTeamsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamworkAssociatedTeamClient) ListTeamworkAssociatedTeamsCompleteMatchingPredicate(ctx context.Context, predicate AssociatedTeamInfoOperationPredicate) (result ListTeamworkAssociatedTeamsCompleteResult, err error) {
	items := make([]beta.AssociatedTeamInfo, 0)

	resp, err := c.ListTeamworkAssociatedTeams(ctx)
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

	result = ListTeamworkAssociatedTeamsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
