package mecalendarcalendarpermission

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

type ListMeCalendarCalendarPermissionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.CalendarPermissionCollectionResponse
}

type ListMeCalendarCalendarPermissionsCompleteResult struct {
	Items []models.CalendarPermissionCollectionResponse
}

// ListMeCalendarCalendarPermissions ...
func (c MeCalendarCalendarPermissionClient) ListMeCalendarCalendarPermissions(ctx context.Context) (result ListMeCalendarCalendarPermissionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/calendar/calendarPermissions",
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

// ListMeCalendarCalendarPermissionsComplete retrieves all the results into a single object
func (c MeCalendarCalendarPermissionClient) ListMeCalendarCalendarPermissionsComplete(ctx context.Context) (ListMeCalendarCalendarPermissionsCompleteResult, error) {
	return c.ListMeCalendarCalendarPermissionsCompleteMatchingPredicate(ctx, models.CalendarPermissionCollectionResponseOperationPredicate{})
}

// ListMeCalendarCalendarPermissionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCalendarCalendarPermissionClient) ListMeCalendarCalendarPermissionsCompleteMatchingPredicate(ctx context.Context, predicate models.CalendarPermissionCollectionResponseOperationPredicate) (result ListMeCalendarCalendarPermissionsCompleteResult, err error) {
	items := make([]models.CalendarPermissionCollectionResponse, 0)

	resp, err := c.ListMeCalendarCalendarPermissions(ctx)
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

	result = ListMeCalendarCalendarPermissionsCompleteResult{
		Items: items,
	}
	return
}
