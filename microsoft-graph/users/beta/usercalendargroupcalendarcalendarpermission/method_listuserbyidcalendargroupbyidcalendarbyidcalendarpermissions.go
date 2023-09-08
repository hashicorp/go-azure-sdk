package usercalendargroupcalendarcalendarpermission

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

type ListUserByIdCalendarGroupByIdCalendarByIdCalendarPermissionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.CalendarPermissionCollectionResponse
}

type ListUserByIdCalendarGroupByIdCalendarByIdCalendarPermissionsCompleteResult struct {
	Items []models.CalendarPermissionCollectionResponse
}

// ListUserByIdCalendarGroupByIdCalendarByIdCalendarPermissions ...
func (c UserCalendarGroupCalendarCalendarPermissionClient) ListUserByIdCalendarGroupByIdCalendarByIdCalendarPermissions(ctx context.Context, id UserCalendarGroupCalendarId) (result ListUserByIdCalendarGroupByIdCalendarByIdCalendarPermissionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/calendarPermissions", id.ID()),
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

// ListUserByIdCalendarGroupByIdCalendarByIdCalendarPermissionsComplete retrieves all the results into a single object
func (c UserCalendarGroupCalendarCalendarPermissionClient) ListUserByIdCalendarGroupByIdCalendarByIdCalendarPermissionsComplete(ctx context.Context, id UserCalendarGroupCalendarId) (ListUserByIdCalendarGroupByIdCalendarByIdCalendarPermissionsCompleteResult, error) {
	return c.ListUserByIdCalendarGroupByIdCalendarByIdCalendarPermissionsCompleteMatchingPredicate(ctx, id, models.CalendarPermissionCollectionResponseOperationPredicate{})
}

// ListUserByIdCalendarGroupByIdCalendarByIdCalendarPermissionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCalendarGroupCalendarCalendarPermissionClient) ListUserByIdCalendarGroupByIdCalendarByIdCalendarPermissionsCompleteMatchingPredicate(ctx context.Context, id UserCalendarGroupCalendarId, predicate models.CalendarPermissionCollectionResponseOperationPredicate) (result ListUserByIdCalendarGroupByIdCalendarByIdCalendarPermissionsCompleteResult, err error) {
	items := make([]models.CalendarPermissionCollectionResponse, 0)

	resp, err := c.ListUserByIdCalendarGroupByIdCalendarByIdCalendarPermissions(ctx, id)
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

	result = ListUserByIdCalendarGroupByIdCalendarByIdCalendarPermissionsCompleteResult{
		Items: items,
	}
	return
}
