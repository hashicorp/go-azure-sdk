package meeventexceptionoccurrenceinstance

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

type ListMeEventByIdExceptionOccurrenceByIdInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EventCollectionResponse
}

type ListMeEventByIdExceptionOccurrenceByIdInstancesCompleteResult struct {
	Items []models.EventCollectionResponse
}

// ListMeEventByIdExceptionOccurrenceByIdInstances ...
func (c MeEventExceptionOccurrenceInstanceClient) ListMeEventByIdExceptionOccurrenceByIdInstances(ctx context.Context, id MeEventExceptionOccurrenceId) (result ListMeEventByIdExceptionOccurrenceByIdInstancesOperationResponse, err error) {
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

// ListMeEventByIdExceptionOccurrenceByIdInstancesComplete retrieves all the results into a single object
func (c MeEventExceptionOccurrenceInstanceClient) ListMeEventByIdExceptionOccurrenceByIdInstancesComplete(ctx context.Context, id MeEventExceptionOccurrenceId) (ListMeEventByIdExceptionOccurrenceByIdInstancesCompleteResult, error) {
	return c.ListMeEventByIdExceptionOccurrenceByIdInstancesCompleteMatchingPredicate(ctx, id, models.EventCollectionResponseOperationPredicate{})
}

// ListMeEventByIdExceptionOccurrenceByIdInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeEventExceptionOccurrenceInstanceClient) ListMeEventByIdExceptionOccurrenceByIdInstancesCompleteMatchingPredicate(ctx context.Context, id MeEventExceptionOccurrenceId, predicate models.EventCollectionResponseOperationPredicate) (result ListMeEventByIdExceptionOccurrenceByIdInstancesCompleteResult, err error) {
	items := make([]models.EventCollectionResponse, 0)

	resp, err := c.ListMeEventByIdExceptionOccurrenceByIdInstances(ctx, id)
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

	result = ListMeEventByIdExceptionOccurrenceByIdInstancesCompleteResult{
		Items: items,
	}
	return
}
