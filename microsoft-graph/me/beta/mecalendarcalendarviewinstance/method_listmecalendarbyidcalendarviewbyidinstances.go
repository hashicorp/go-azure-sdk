package mecalendarcalendarviewinstance

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

type ListMeCalendarByIdCalendarViewByIdInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EventCollectionResponse
}

type ListMeCalendarByIdCalendarViewByIdInstancesCompleteResult struct {
	Items []models.EventCollectionResponse
}

// ListMeCalendarByIdCalendarViewByIdInstances ...
func (c MeCalendarCalendarViewInstanceClient) ListMeCalendarByIdCalendarViewByIdInstances(ctx context.Context, id MeCalendarCalendarViewId) (result ListMeCalendarByIdCalendarViewByIdInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/instances", id.ID()),
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

// ListMeCalendarByIdCalendarViewByIdInstancesComplete retrieves all the results into a single object
func (c MeCalendarCalendarViewInstanceClient) ListMeCalendarByIdCalendarViewByIdInstancesComplete(ctx context.Context, id MeCalendarCalendarViewId) (ListMeCalendarByIdCalendarViewByIdInstancesCompleteResult, error) {
	return c.ListMeCalendarByIdCalendarViewByIdInstancesCompleteMatchingPredicate(ctx, id, models.EventCollectionResponseOperationPredicate{})
}

// ListMeCalendarByIdCalendarViewByIdInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCalendarCalendarViewInstanceClient) ListMeCalendarByIdCalendarViewByIdInstancesCompleteMatchingPredicate(ctx context.Context, id MeCalendarCalendarViewId, predicate models.EventCollectionResponseOperationPredicate) (result ListMeCalendarByIdCalendarViewByIdInstancesCompleteResult, err error) {
	items := make([]models.EventCollectionResponse, 0)

	resp, err := c.ListMeCalendarByIdCalendarViewByIdInstances(ctx, id)
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

	result = ListMeCalendarByIdCalendarViewByIdInstancesCompleteResult{
		Items: items,
	}
	return
}
