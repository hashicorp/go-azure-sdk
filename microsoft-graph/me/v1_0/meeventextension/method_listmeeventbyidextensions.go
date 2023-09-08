package meeventextension

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

type ListMeEventByIdExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ExtensionCollectionResponse
}

type ListMeEventByIdExtensionsCompleteResult struct {
	Items []models.ExtensionCollectionResponse
}

// ListMeEventByIdExtensions ...
func (c MeEventExtensionClient) ListMeEventByIdExtensions(ctx context.Context, id MeEventId) (result ListMeEventByIdExtensionsOperationResponse, err error) {
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

// ListMeEventByIdExtensionsComplete retrieves all the results into a single object
func (c MeEventExtensionClient) ListMeEventByIdExtensionsComplete(ctx context.Context, id MeEventId) (ListMeEventByIdExtensionsCompleteResult, error) {
	return c.ListMeEventByIdExtensionsCompleteMatchingPredicate(ctx, id, models.ExtensionCollectionResponseOperationPredicate{})
}

// ListMeEventByIdExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeEventExtensionClient) ListMeEventByIdExtensionsCompleteMatchingPredicate(ctx context.Context, id MeEventId, predicate models.ExtensionCollectionResponseOperationPredicate) (result ListMeEventByIdExtensionsCompleteResult, err error) {
	items := make([]models.ExtensionCollectionResponse, 0)

	resp, err := c.ListMeEventByIdExtensions(ctx, id)
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

	result = ListMeEventByIdExtensionsCompleteResult{
		Items: items,
	}
	return
}
