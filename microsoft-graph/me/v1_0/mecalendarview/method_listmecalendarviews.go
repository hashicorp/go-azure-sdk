package mecalendarview

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

type ListMeCalendarViewsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EventCollectionResponse
}

type ListMeCalendarViewsCompleteResult struct {
	Items []models.EventCollectionResponse
}

// ListMeCalendarViews ...
func (c MeCalendarViewClient) ListMeCalendarViews(ctx context.Context) (result ListMeCalendarViewsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/calendarView",
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
		Values *[]models.EventCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeCalendarViewsComplete retrieves all the results into a single object
func (c MeCalendarViewClient) ListMeCalendarViewsComplete(ctx context.Context) (ListMeCalendarViewsCompleteResult, error) {
	return c.ListMeCalendarViewsCompleteMatchingPredicate(ctx, models.EventCollectionResponseOperationPredicate{})
}

// ListMeCalendarViewsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCalendarViewClient) ListMeCalendarViewsCompleteMatchingPredicate(ctx context.Context, predicate models.EventCollectionResponseOperationPredicate) (result ListMeCalendarViewsCompleteResult, err error) {
	items := make([]models.EventCollectionResponse, 0)

	resp, err := c.ListMeCalendarViews(ctx)
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

	result = ListMeCalendarViewsCompleteResult{
		Items: items,
	}
	return
}
