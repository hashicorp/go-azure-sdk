package usercalendargroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserByIdCalendarGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.CalendarGroupCollectionResponse
}

type ListUserByIdCalendarGroupsCompleteResult struct {
	Items []models.CalendarGroupCollectionResponse
}

// ListUserByIdCalendarGroups ...
func (c UserCalendarGroupClient) ListUserByIdCalendarGroups(ctx context.Context, id UserId) (result ListUserByIdCalendarGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/calendarGroups", id.ID()),
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
		Values *[]models.CalendarGroupCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdCalendarGroupsComplete retrieves all the results into a single object
func (c UserCalendarGroupClient) ListUserByIdCalendarGroupsComplete(ctx context.Context, id UserId) (ListUserByIdCalendarGroupsCompleteResult, error) {
	return c.ListUserByIdCalendarGroupsCompleteMatchingPredicate(ctx, id, models.CalendarGroupCollectionResponseOperationPredicate{})
}

// ListUserByIdCalendarGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCalendarGroupClient) ListUserByIdCalendarGroupsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.CalendarGroupCollectionResponseOperationPredicate) (result ListUserByIdCalendarGroupsCompleteResult, err error) {
	items := make([]models.CalendarGroupCollectionResponse, 0)

	resp, err := c.ListUserByIdCalendarGroups(ctx, id)
	if err != nil {
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

	result = ListUserByIdCalendarGroupsCompleteResult{
		Items: items,
	}
	return
}
