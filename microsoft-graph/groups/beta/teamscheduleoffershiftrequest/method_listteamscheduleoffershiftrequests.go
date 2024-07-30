package teamscheduleoffershiftrequest

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

type ListTeamScheduleOfferShiftRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.OfferShiftRequest
}

type ListTeamScheduleOfferShiftRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.OfferShiftRequest
}

type ListTeamScheduleOfferShiftRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamScheduleOfferShiftRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamScheduleOfferShiftRequests ...
func (c TeamScheduleOfferShiftRequestClient) ListTeamScheduleOfferShiftRequests(ctx context.Context, id GroupId) (result ListTeamScheduleOfferShiftRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTeamScheduleOfferShiftRequestsCustomPager{},
		Path:       fmt.Sprintf("%s/team/schedule/offerShiftRequests", id.ID()),
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
		Values *[]beta.OfferShiftRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamScheduleOfferShiftRequestsComplete retrieves all the results into a single object
func (c TeamScheduleOfferShiftRequestClient) ListTeamScheduleOfferShiftRequestsComplete(ctx context.Context, id GroupId) (ListTeamScheduleOfferShiftRequestsCompleteResult, error) {
	return c.ListTeamScheduleOfferShiftRequestsCompleteMatchingPredicate(ctx, id, OfferShiftRequestOperationPredicate{})
}

// ListTeamScheduleOfferShiftRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamScheduleOfferShiftRequestClient) ListTeamScheduleOfferShiftRequestsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate OfferShiftRequestOperationPredicate) (result ListTeamScheduleOfferShiftRequestsCompleteResult, err error) {
	items := make([]beta.OfferShiftRequest, 0)

	resp, err := c.ListTeamScheduleOfferShiftRequests(ctx, id)
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

	result = ListTeamScheduleOfferShiftRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
