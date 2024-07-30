package joinedteamscheduleoffershiftrequest

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

type ListJoinedTeamScheduleOfferShiftRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.OfferShiftRequest
}

type ListJoinedTeamScheduleOfferShiftRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.OfferShiftRequest
}

type ListJoinedTeamScheduleOfferShiftRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamScheduleOfferShiftRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamScheduleOfferShiftRequests ...
func (c JoinedTeamScheduleOfferShiftRequestClient) ListJoinedTeamScheduleOfferShiftRequests(ctx context.Context, id MeJoinedTeamId) (result ListJoinedTeamScheduleOfferShiftRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListJoinedTeamScheduleOfferShiftRequestsCustomPager{},
		Path:       fmt.Sprintf("%s/schedule/offerShiftRequests", id.ID()),
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
		Values *[]stable.OfferShiftRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamScheduleOfferShiftRequestsComplete retrieves all the results into a single object
func (c JoinedTeamScheduleOfferShiftRequestClient) ListJoinedTeamScheduleOfferShiftRequestsComplete(ctx context.Context, id MeJoinedTeamId) (ListJoinedTeamScheduleOfferShiftRequestsCompleteResult, error) {
	return c.ListJoinedTeamScheduleOfferShiftRequestsCompleteMatchingPredicate(ctx, id, OfferShiftRequestOperationPredicate{})
}

// ListJoinedTeamScheduleOfferShiftRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamScheduleOfferShiftRequestClient) ListJoinedTeamScheduleOfferShiftRequestsCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamId, predicate OfferShiftRequestOperationPredicate) (result ListJoinedTeamScheduleOfferShiftRequestsCompleteResult, err error) {
	items := make([]stable.OfferShiftRequest, 0)

	resp, err := c.ListJoinedTeamScheduleOfferShiftRequests(ctx, id)
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

	result = ListJoinedTeamScheduleOfferShiftRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
