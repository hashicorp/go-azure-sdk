package capacityreservations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByCapacityReservationGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]CapacityReservation
}

type ListByCapacityReservationGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []CapacityReservation
}

type ListByCapacityReservationGroupCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByCapacityReservationGroupCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByCapacityReservationGroup ...
func (c CapacityReservationsClient) ListByCapacityReservationGroup(ctx context.Context, id CapacityReservationGroupId) (result ListByCapacityReservationGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByCapacityReservationGroupCustomPager{},
		Path:       fmt.Sprintf("%s/capacityReservations", id.ID()),
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
		Values *[]CapacityReservation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByCapacityReservationGroupComplete retrieves all the results into a single object
func (c CapacityReservationsClient) ListByCapacityReservationGroupComplete(ctx context.Context, id CapacityReservationGroupId) (ListByCapacityReservationGroupCompleteResult, error) {
	return c.ListByCapacityReservationGroupCompleteMatchingPredicate(ctx, id, CapacityReservationOperationPredicate{})
}

// ListByCapacityReservationGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CapacityReservationsClient) ListByCapacityReservationGroupCompleteMatchingPredicate(ctx context.Context, id CapacityReservationGroupId, predicate CapacityReservationOperationPredicate) (result ListByCapacityReservationGroupCompleteResult, err error) {
	items := make([]CapacityReservation, 0)

	resp, err := c.ListByCapacityReservationGroup(ctx, id)
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

	result = ListByCapacityReservationGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
