package joinedteamscheduleopenshift

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

type ListJoinedTeamScheduleOpenShiftsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.OpenShift
}

type ListJoinedTeamScheduleOpenShiftsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.OpenShift
}

type ListJoinedTeamScheduleOpenShiftsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamScheduleOpenShiftsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamScheduleOpenShifts ...
func (c JoinedTeamScheduleOpenShiftClient) ListJoinedTeamScheduleOpenShifts(ctx context.Context, id UserIdJoinedTeamId) (result ListJoinedTeamScheduleOpenShiftsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListJoinedTeamScheduleOpenShiftsCustomPager{},
		Path:       fmt.Sprintf("%s/schedule/openShifts", id.ID()),
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
		Values *[]stable.OpenShift `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamScheduleOpenShiftsComplete retrieves all the results into a single object
func (c JoinedTeamScheduleOpenShiftClient) ListJoinedTeamScheduleOpenShiftsComplete(ctx context.Context, id UserIdJoinedTeamId) (ListJoinedTeamScheduleOpenShiftsCompleteResult, error) {
	return c.ListJoinedTeamScheduleOpenShiftsCompleteMatchingPredicate(ctx, id, OpenShiftOperationPredicate{})
}

// ListJoinedTeamScheduleOpenShiftsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamScheduleOpenShiftClient) ListJoinedTeamScheduleOpenShiftsCompleteMatchingPredicate(ctx context.Context, id UserIdJoinedTeamId, predicate OpenShiftOperationPredicate) (result ListJoinedTeamScheduleOpenShiftsCompleteResult, err error) {
	items := make([]stable.OpenShift, 0)

	resp, err := c.ListJoinedTeamScheduleOpenShifts(ctx, id)
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

	result = ListJoinedTeamScheduleOpenShiftsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
