package usercalendarviewextension

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

type ListUserByIdCalendarViewByIdExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ExtensionCollectionResponse
}

type ListUserByIdCalendarViewByIdExtensionsCompleteResult struct {
	Items []models.ExtensionCollectionResponse
}

// ListUserByIdCalendarViewByIdExtensions ...
func (c UserCalendarViewExtensionClient) ListUserByIdCalendarViewByIdExtensions(ctx context.Context, id UserCalendarViewId) (result ListUserByIdCalendarViewByIdExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/extensions", id.ID()),
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
		Values *[]models.ExtensionCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdCalendarViewByIdExtensionsComplete retrieves all the results into a single object
func (c UserCalendarViewExtensionClient) ListUserByIdCalendarViewByIdExtensionsComplete(ctx context.Context, id UserCalendarViewId) (ListUserByIdCalendarViewByIdExtensionsCompleteResult, error) {
	return c.ListUserByIdCalendarViewByIdExtensionsCompleteMatchingPredicate(ctx, id, models.ExtensionCollectionResponseOperationPredicate{})
}

// ListUserByIdCalendarViewByIdExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCalendarViewExtensionClient) ListUserByIdCalendarViewByIdExtensionsCompleteMatchingPredicate(ctx context.Context, id UserCalendarViewId, predicate models.ExtensionCollectionResponseOperationPredicate) (result ListUserByIdCalendarViewByIdExtensionsCompleteResult, err error) {
	items := make([]models.ExtensionCollectionResponse, 0)

	resp, err := c.ListUserByIdCalendarViewByIdExtensions(ctx, id)
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

	result = ListUserByIdCalendarViewByIdExtensionsCompleteResult{
		Items: items,
	}
	return
}
