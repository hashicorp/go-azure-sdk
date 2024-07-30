package teamscheduletimesoff

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

type ListTeamScheduleTimesOffsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.TimeOff
}

type ListTeamScheduleTimesOffsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.TimeOff
}

type ListTeamScheduleTimesOffsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamScheduleTimesOffsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamScheduleTimesOffs ...
func (c TeamScheduleTimesOffClient) ListTeamScheduleTimesOffs(ctx context.Context, id GroupId) (result ListTeamScheduleTimesOffsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTeamScheduleTimesOffsCustomPager{},
		Path:       fmt.Sprintf("%s/team/schedule/timesOff", id.ID()),
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
		Values *[]stable.TimeOff `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamScheduleTimesOffsComplete retrieves all the results into a single object
func (c TeamScheduleTimesOffClient) ListTeamScheduleTimesOffsComplete(ctx context.Context, id GroupId) (ListTeamScheduleTimesOffsCompleteResult, error) {
	return c.ListTeamScheduleTimesOffsCompleteMatchingPredicate(ctx, id, TimeOffOperationPredicate{})
}

// ListTeamScheduleTimesOffsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamScheduleTimesOffClient) ListTeamScheduleTimesOffsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate TimeOffOperationPredicate) (result ListTeamScheduleTimesOffsCompleteResult, err error) {
	items := make([]stable.TimeOff, 0)

	resp, err := c.ListTeamScheduleTimesOffs(ctx, id)
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

	result = ListTeamScheduleTimesOffsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
