package mecalendareventinstanceextension

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

type ListMeCalendarByIdEventByIdInstanceByIdExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ExtensionCollectionResponse
}

type ListMeCalendarByIdEventByIdInstanceByIdExtensionsCompleteResult struct {
	Items []models.ExtensionCollectionResponse
}

// ListMeCalendarByIdEventByIdInstanceByIdExtensions ...
func (c MeCalendarEventInstanceExtensionClient) ListMeCalendarByIdEventByIdInstanceByIdExtensions(ctx context.Context, id MeCalendarEventInstanceId) (result ListMeCalendarByIdEventByIdInstanceByIdExtensionsOperationResponse, err error) {
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

// ListMeCalendarByIdEventByIdInstanceByIdExtensionsComplete retrieves all the results into a single object
func (c MeCalendarEventInstanceExtensionClient) ListMeCalendarByIdEventByIdInstanceByIdExtensionsComplete(ctx context.Context, id MeCalendarEventInstanceId) (ListMeCalendarByIdEventByIdInstanceByIdExtensionsCompleteResult, error) {
	return c.ListMeCalendarByIdEventByIdInstanceByIdExtensionsCompleteMatchingPredicate(ctx, id, models.ExtensionCollectionResponseOperationPredicate{})
}

// ListMeCalendarByIdEventByIdInstanceByIdExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCalendarEventInstanceExtensionClient) ListMeCalendarByIdEventByIdInstanceByIdExtensionsCompleteMatchingPredicate(ctx context.Context, id MeCalendarEventInstanceId, predicate models.ExtensionCollectionResponseOperationPredicate) (result ListMeCalendarByIdEventByIdInstanceByIdExtensionsCompleteResult, err error) {
	items := make([]models.ExtensionCollectionResponse, 0)

	resp, err := c.ListMeCalendarByIdEventByIdInstanceByIdExtensions(ctx, id)
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

	result = ListMeCalendarByIdEventByIdInstanceByIdExtensionsCompleteResult{
		Items: items,
	}
	return
}
