package meevent

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

type ListMeEventsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EventCollectionResponse
}

type ListMeEventsCompleteResult struct {
	Items []models.EventCollectionResponse
}

// ListMeEvents ...
func (c MeEventClient) ListMeEvents(ctx context.Context) (result ListMeEventsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/events",
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

// ListMeEventsComplete retrieves all the results into a single object
func (c MeEventClient) ListMeEventsComplete(ctx context.Context) (ListMeEventsCompleteResult, error) {
	return c.ListMeEventsCompleteMatchingPredicate(ctx, models.EventCollectionResponseOperationPredicate{})
}

// ListMeEventsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeEventClient) ListMeEventsCompleteMatchingPredicate(ctx context.Context, predicate models.EventCollectionResponseOperationPredicate) (result ListMeEventsCompleteResult, err error) {
	items := make([]models.EventCollectionResponse, 0)

	resp, err := c.ListMeEvents(ctx)
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

	result = ListMeEventsCompleteResult{
		Items: items,
	}
	return
}
