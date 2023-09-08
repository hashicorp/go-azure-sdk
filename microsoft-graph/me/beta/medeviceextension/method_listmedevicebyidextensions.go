package medeviceextension

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

type ListMeDeviceByIdExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ExtensionCollectionResponse
}

type ListMeDeviceByIdExtensionsCompleteResult struct {
	Items []models.ExtensionCollectionResponse
}

// ListMeDeviceByIdExtensions ...
func (c MeDeviceExtensionClient) ListMeDeviceByIdExtensions(ctx context.Context, id MeDeviceId) (result ListMeDeviceByIdExtensionsOperationResponse, err error) {
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

// ListMeDeviceByIdExtensionsComplete retrieves all the results into a single object
func (c MeDeviceExtensionClient) ListMeDeviceByIdExtensionsComplete(ctx context.Context, id MeDeviceId) (ListMeDeviceByIdExtensionsCompleteResult, error) {
	return c.ListMeDeviceByIdExtensionsCompleteMatchingPredicate(ctx, id, models.ExtensionCollectionResponseOperationPredicate{})
}

// ListMeDeviceByIdExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeDeviceExtensionClient) ListMeDeviceByIdExtensionsCompleteMatchingPredicate(ctx context.Context, id MeDeviceId, predicate models.ExtensionCollectionResponseOperationPredicate) (result ListMeDeviceByIdExtensionsCompleteResult, err error) {
	items := make([]models.ExtensionCollectionResponse, 0)

	resp, err := c.ListMeDeviceByIdExtensions(ctx, id)
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

	result = ListMeDeviceByIdExtensionsCompleteResult{
		Items: items,
	}
	return
}
