package meextension

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

type ListMeExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ExtensionCollectionResponse
}

type ListMeExtensionsCompleteResult struct {
	Items []models.ExtensionCollectionResponse
}

// ListMeExtensions ...
func (c MeExtensionClient) ListMeExtensions(ctx context.Context) (result ListMeExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/extensions",
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

// ListMeExtensionsComplete retrieves all the results into a single object
func (c MeExtensionClient) ListMeExtensionsComplete(ctx context.Context) (ListMeExtensionsCompleteResult, error) {
	return c.ListMeExtensionsCompleteMatchingPredicate(ctx, models.ExtensionCollectionResponseOperationPredicate{})
}

// ListMeExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeExtensionClient) ListMeExtensionsCompleteMatchingPredicate(ctx context.Context, predicate models.ExtensionCollectionResponseOperationPredicate) (result ListMeExtensionsCompleteResult, err error) {
	items := make([]models.ExtensionCollectionResponse, 0)

	resp, err := c.ListMeExtensions(ctx)
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

	result = ListMeExtensionsCompleteResult{
		Items: items,
	}
	return
}
