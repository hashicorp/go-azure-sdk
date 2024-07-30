package joinedteamscheduleopenshiftchangerequest

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

type ListJoinedTeamScheduleOpenShiftChangeRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.OpenShiftChangeRequest
}

type ListJoinedTeamScheduleOpenShiftChangeRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.OpenShiftChangeRequest
}

type ListJoinedTeamScheduleOpenShiftChangeRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamScheduleOpenShiftChangeRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamScheduleOpenShiftChangeRequests ...
func (c JoinedTeamScheduleOpenShiftChangeRequestClient) ListJoinedTeamScheduleOpenShiftChangeRequests(ctx context.Context, id MeJoinedTeamId) (result ListJoinedTeamScheduleOpenShiftChangeRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListJoinedTeamScheduleOpenShiftChangeRequestsCustomPager{},
		Path:       fmt.Sprintf("%s/schedule/openShiftChangeRequests", id.ID()),
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
		Values *[]stable.OpenShiftChangeRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamScheduleOpenShiftChangeRequestsComplete retrieves all the results into a single object
func (c JoinedTeamScheduleOpenShiftChangeRequestClient) ListJoinedTeamScheduleOpenShiftChangeRequestsComplete(ctx context.Context, id MeJoinedTeamId) (ListJoinedTeamScheduleOpenShiftChangeRequestsCompleteResult, error) {
	return c.ListJoinedTeamScheduleOpenShiftChangeRequestsCompleteMatchingPredicate(ctx, id, OpenShiftChangeRequestOperationPredicate{})
}

// ListJoinedTeamScheduleOpenShiftChangeRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamScheduleOpenShiftChangeRequestClient) ListJoinedTeamScheduleOpenShiftChangeRequestsCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamId, predicate OpenShiftChangeRequestOperationPredicate) (result ListJoinedTeamScheduleOpenShiftChangeRequestsCompleteResult, err error) {
	items := make([]stable.OpenShiftChangeRequest, 0)

	resp, err := c.ListJoinedTeamScheduleOpenShiftChangeRequests(ctx, id)
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

	result = ListJoinedTeamScheduleOpenShiftChangeRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
