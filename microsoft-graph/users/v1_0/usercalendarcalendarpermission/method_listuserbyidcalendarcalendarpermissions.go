package usercalendarcalendarpermission

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserByIdCalendarCalendarPermissionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.CalendarPermissionCollectionResponse
}

type ListUserByIdCalendarCalendarPermissionsCompleteResult struct {
	Items []models.CalendarPermissionCollectionResponse
}

// ListUserByIdCalendarCalendarPermissions ...
func (c UserCalendarCalendarPermissionClient) ListUserByIdCalendarCalendarPermissions(ctx context.Context, id UserId) (result ListUserByIdCalendarCalendarPermissionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/calendar/calendarPermissions", id.ID()),
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
		Values *[]models.CalendarPermissionCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdCalendarCalendarPermissionsComplete retrieves all the results into a single object
func (c UserCalendarCalendarPermissionClient) ListUserByIdCalendarCalendarPermissionsComplete(ctx context.Context, id UserId) (ListUserByIdCalendarCalendarPermissionsCompleteResult, error) {
	return c.ListUserByIdCalendarCalendarPermissionsCompleteMatchingPredicate(ctx, id, models.CalendarPermissionCollectionResponseOperationPredicate{})
}

// ListUserByIdCalendarCalendarPermissionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCalendarCalendarPermissionClient) ListUserByIdCalendarCalendarPermissionsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.CalendarPermissionCollectionResponseOperationPredicate) (result ListUserByIdCalendarCalendarPermissionsCompleteResult, err error) {
	items := make([]models.CalendarPermissionCollectionResponse, 0)

	resp, err := c.ListUserByIdCalendarCalendarPermissions(ctx, id)
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

	result = ListUserByIdCalendarCalendarPermissionsCompleteResult{
		Items: items,
	}
	return
}
