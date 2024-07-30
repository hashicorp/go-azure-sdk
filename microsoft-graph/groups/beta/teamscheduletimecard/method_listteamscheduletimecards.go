package teamscheduletimecard

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

type ListTeamScheduleTimeCardsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.TimeCard
}

type ListTeamScheduleTimeCardsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.TimeCard
}

type ListTeamScheduleTimeCardsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamScheduleTimeCardsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamScheduleTimeCards ...
func (c TeamScheduleTimeCardClient) ListTeamScheduleTimeCards(ctx context.Context, id GroupId) (result ListTeamScheduleTimeCardsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTeamScheduleTimeCardsCustomPager{},
		Path:       fmt.Sprintf("%s/team/schedule/timeCards", id.ID()),
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
		Values *[]beta.TimeCard `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamScheduleTimeCardsComplete retrieves all the results into a single object
func (c TeamScheduleTimeCardClient) ListTeamScheduleTimeCardsComplete(ctx context.Context, id GroupId) (ListTeamScheduleTimeCardsCompleteResult, error) {
	return c.ListTeamScheduleTimeCardsCompleteMatchingPredicate(ctx, id, TimeCardOperationPredicate{})
}

// ListTeamScheduleTimeCardsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamScheduleTimeCardClient) ListTeamScheduleTimeCardsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate TimeCardOperationPredicate) (result ListTeamScheduleTimeCardsCompleteResult, err error) {
	items := make([]beta.TimeCard, 0)

	resp, err := c.ListTeamScheduleTimeCards(ctx, id)
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

	result = ListTeamScheduleTimeCardsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
