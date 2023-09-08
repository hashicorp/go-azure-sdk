package mecalendarcalendarviewextension

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

type ListMeCalendarByIdCalendarViewByIdExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ExtensionCollectionResponse
}

type ListMeCalendarByIdCalendarViewByIdExtensionsCompleteResult struct {
	Items []models.ExtensionCollectionResponse
}

// ListMeCalendarByIdCalendarViewByIdExtensions ...
func (c MeCalendarCalendarViewExtensionClient) ListMeCalendarByIdCalendarViewByIdExtensions(ctx context.Context, id MeCalendarCalendarViewId) (result ListMeCalendarByIdCalendarViewByIdExtensionsOperationResponse, err error) {
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

// ListMeCalendarByIdCalendarViewByIdExtensionsComplete retrieves all the results into a single object
func (c MeCalendarCalendarViewExtensionClient) ListMeCalendarByIdCalendarViewByIdExtensionsComplete(ctx context.Context, id MeCalendarCalendarViewId) (ListMeCalendarByIdCalendarViewByIdExtensionsCompleteResult, error) {
	return c.ListMeCalendarByIdCalendarViewByIdExtensionsCompleteMatchingPredicate(ctx, id, models.ExtensionCollectionResponseOperationPredicate{})
}

// ListMeCalendarByIdCalendarViewByIdExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCalendarCalendarViewExtensionClient) ListMeCalendarByIdCalendarViewByIdExtensionsCompleteMatchingPredicate(ctx context.Context, id MeCalendarCalendarViewId, predicate models.ExtensionCollectionResponseOperationPredicate) (result ListMeCalendarByIdCalendarViewByIdExtensionsCompleteResult, err error) {
	items := make([]models.ExtensionCollectionResponse, 0)

	resp, err := c.ListMeCalendarByIdCalendarViewByIdExtensions(ctx, id)
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

	result = ListMeCalendarByIdCalendarViewByIdExtensionsCompleteResult{
		Items: items,
	}
	return
}
