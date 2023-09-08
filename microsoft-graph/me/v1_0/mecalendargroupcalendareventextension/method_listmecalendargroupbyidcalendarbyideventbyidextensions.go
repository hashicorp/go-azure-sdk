package mecalendargroupcalendareventextension

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

type ListMeCalendarGroupByIdCalendarByIdEventByIdExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ExtensionCollectionResponse
}

type ListMeCalendarGroupByIdCalendarByIdEventByIdExtensionsCompleteResult struct {
	Items []models.ExtensionCollectionResponse
}

// ListMeCalendarGroupByIdCalendarByIdEventByIdExtensions ...
func (c MeCalendarGroupCalendarEventExtensionClient) ListMeCalendarGroupByIdCalendarByIdEventByIdExtensions(ctx context.Context, id MeCalendarGroupCalendarEventId) (result ListMeCalendarGroupByIdCalendarByIdEventByIdExtensionsOperationResponse, err error) {
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

// ListMeCalendarGroupByIdCalendarByIdEventByIdExtensionsComplete retrieves all the results into a single object
func (c MeCalendarGroupCalendarEventExtensionClient) ListMeCalendarGroupByIdCalendarByIdEventByIdExtensionsComplete(ctx context.Context, id MeCalendarGroupCalendarEventId) (ListMeCalendarGroupByIdCalendarByIdEventByIdExtensionsCompleteResult, error) {
	return c.ListMeCalendarGroupByIdCalendarByIdEventByIdExtensionsCompleteMatchingPredicate(ctx, id, models.ExtensionCollectionResponseOperationPredicate{})
}

// ListMeCalendarGroupByIdCalendarByIdEventByIdExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCalendarGroupCalendarEventExtensionClient) ListMeCalendarGroupByIdCalendarByIdEventByIdExtensionsCompleteMatchingPredicate(ctx context.Context, id MeCalendarGroupCalendarEventId, predicate models.ExtensionCollectionResponseOperationPredicate) (result ListMeCalendarGroupByIdCalendarByIdEventByIdExtensionsCompleteResult, err error) {
	items := make([]models.ExtensionCollectionResponse, 0)

	resp, err := c.ListMeCalendarGroupByIdCalendarByIdEventByIdExtensions(ctx, id)
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

	result = ListMeCalendarGroupByIdCalendarByIdEventByIdExtensionsCompleteResult{
		Items: items,
	}
	return
}
