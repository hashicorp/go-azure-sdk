package mecalendarcalendarpermission

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

type ListMeCalendarByIdCalendarPermissionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.CalendarPermissionCollectionResponse
}

type ListMeCalendarByIdCalendarPermissionsCompleteResult struct {
	Items []models.CalendarPermissionCollectionResponse
}

// ListMeCalendarByIdCalendarPermissions ...
func (c MeCalendarCalendarPermissionClient) ListMeCalendarByIdCalendarPermissions(ctx context.Context, id MeCalendarId) (result ListMeCalendarByIdCalendarPermissionsOperationResponse, err error) {
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

// ListMeCalendarByIdCalendarPermissionsComplete retrieves all the results into a single object
func (c MeCalendarCalendarPermissionClient) ListMeCalendarByIdCalendarPermissionsComplete(ctx context.Context, id MeCalendarId) (ListMeCalendarByIdCalendarPermissionsCompleteResult, error) {
	return c.ListMeCalendarByIdCalendarPermissionsCompleteMatchingPredicate(ctx, id, models.CalendarPermissionCollectionResponseOperationPredicate{})
}

// ListMeCalendarByIdCalendarPermissionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCalendarCalendarPermissionClient) ListMeCalendarByIdCalendarPermissionsCompleteMatchingPredicate(ctx context.Context, id MeCalendarId, predicate models.CalendarPermissionCollectionResponseOperationPredicate) (result ListMeCalendarByIdCalendarPermissionsCompleteResult, err error) {
	items := make([]models.CalendarPermissionCollectionResponse, 0)

	resp, err := c.ListMeCalendarByIdCalendarPermissions(ctx, id)
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

	result = ListMeCalendarByIdCalendarPermissionsCompleteResult{
		Items: items,
	}
	return
}
