package teamprimarychanneltab

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

type ListTeamPrimaryChannelTabsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.TeamsTab
}

type ListTeamPrimaryChannelTabsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.TeamsTab
}

type ListTeamPrimaryChannelTabsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamPrimaryChannelTabsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamPrimaryChannelTabs ...
func (c TeamPrimaryChannelTabClient) ListTeamPrimaryChannelTabs(ctx context.Context, id GroupId) (result ListTeamPrimaryChannelTabsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTeamPrimaryChannelTabsCustomPager{},
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
		Values *[]stable.TeamsTab `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamPrimaryChannelTabsComplete retrieves all the results into a single object
func (c TeamPrimaryChannelTabClient) ListTeamPrimaryChannelTabsComplete(ctx context.Context, id GroupId) (ListTeamPrimaryChannelTabsCompleteResult, error) {
	return c.ListTeamPrimaryChannelTabsCompleteMatchingPredicate(ctx, id, TeamsTabOperationPredicate{})
}

// ListTeamPrimaryChannelTabsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamPrimaryChannelTabClient) ListTeamPrimaryChannelTabsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate TeamsTabOperationPredicate) (result ListTeamPrimaryChannelTabsCompleteResult, err error) {
	items := make([]stable.TeamsTab, 0)

	resp, err := c.ListTeamPrimaryChannelTabs(ctx, id)
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

	result = ListTeamPrimaryChannelTabsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
