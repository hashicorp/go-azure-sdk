package mecalendareventinstance

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

type ListMeCalendarEventByIdInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EventCollectionResponse
}

type ListMeCalendarEventByIdInstancesCompleteResult struct {
	Items []models.EventCollectionResponse
}

// ListMeCalendarEventByIdInstances ...
func (c MeCalendarEventInstanceClient) ListMeCalendarEventByIdInstances(ctx context.Context, id MeCalendarEventId) (result ListMeCalendarEventByIdInstancesOperationResponse, err error) {
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

// ListMeCalendarEventByIdInstancesComplete retrieves all the results into a single object
func (c MeCalendarEventInstanceClient) ListMeCalendarEventByIdInstancesComplete(ctx context.Context, id MeCalendarEventId) (ListMeCalendarEventByIdInstancesCompleteResult, error) {
	return c.ListMeCalendarEventByIdInstancesCompleteMatchingPredicate(ctx, id, models.EventCollectionResponseOperationPredicate{})
}

// ListMeCalendarEventByIdInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCalendarEventInstanceClient) ListMeCalendarEventByIdInstancesCompleteMatchingPredicate(ctx context.Context, id MeCalendarEventId, predicate models.EventCollectionResponseOperationPredicate) (result ListMeCalendarEventByIdInstancesCompleteResult, err error) {
	items := make([]models.EventCollectionResponse, 0)

	resp, err := c.ListMeCalendarEventByIdInstances(ctx, id)
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

	result = ListMeCalendarEventByIdInstancesCompleteResult{
		Items: items,
	}
	return
}
