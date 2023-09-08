package usercalendarcalendarpermission

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

type ListUserByIdCalendarByIdCalendarPermissionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.CalendarPermissionCollectionResponse
}

type ListUserByIdCalendarByIdCalendarPermissionsCompleteResult struct {
	Items []models.CalendarPermissionCollectionResponse
}

// ListUserByIdCalendarByIdCalendarPermissions ...
func (c UserCalendarCalendarPermissionClient) ListUserByIdCalendarByIdCalendarPermissions(ctx context.Context, id UserCalendarId) (result ListUserByIdCalendarByIdCalendarPermissionsOperationResponse, err error) {
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

// ListUserByIdCalendarByIdCalendarPermissionsComplete retrieves all the results into a single object
func (c UserCalendarCalendarPermissionClient) ListUserByIdCalendarByIdCalendarPermissionsComplete(ctx context.Context, id UserCalendarId) (ListUserByIdCalendarByIdCalendarPermissionsCompleteResult, error) {
	return c.ListUserByIdCalendarByIdCalendarPermissionsCompleteMatchingPredicate(ctx, id, models.CalendarPermissionCollectionResponseOperationPredicate{})
}

// ListUserByIdCalendarByIdCalendarPermissionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCalendarCalendarPermissionClient) ListUserByIdCalendarByIdCalendarPermissionsCompleteMatchingPredicate(ctx context.Context, id UserCalendarId, predicate models.CalendarPermissionCollectionResponseOperationPredicate) (result ListUserByIdCalendarByIdCalendarPermissionsCompleteResult, err error) {
	items := make([]models.CalendarPermissionCollectionResponse, 0)

	resp, err := c.ListUserByIdCalendarByIdCalendarPermissions(ctx, id)
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

	result = ListUserByIdCalendarByIdCalendarPermissionsCompleteResult{
		Items: items,
	}
	return
}
