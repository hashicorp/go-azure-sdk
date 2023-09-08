package applicationextensionproperty

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

type ListApplicationByIdExtensionPropertiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ExtensionPropertyCollectionResponse
}

type ListApplicationByIdExtensionPropertiesCompleteResult struct {
	Items []models.ExtensionPropertyCollectionResponse
}

// ListApplicationByIdExtensionProperties ...
func (c ApplicationExtensionPropertyClient) ListApplicationByIdExtensionProperties(ctx context.Context, id ApplicationId) (result ListApplicationByIdExtensionPropertiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/extensionProperties", id.ID()),
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
		Values *[]models.ExtensionPropertyCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListApplicationByIdExtensionPropertiesComplete retrieves all the results into a single object
func (c ApplicationExtensionPropertyClient) ListApplicationByIdExtensionPropertiesComplete(ctx context.Context, id ApplicationId) (ListApplicationByIdExtensionPropertiesCompleteResult, error) {
	return c.ListApplicationByIdExtensionPropertiesCompleteMatchingPredicate(ctx, id, models.ExtensionPropertyCollectionResponseOperationPredicate{})
}

// ListApplicationByIdExtensionPropertiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationExtensionPropertyClient) ListApplicationByIdExtensionPropertiesCompleteMatchingPredicate(ctx context.Context, id ApplicationId, predicate models.ExtensionPropertyCollectionResponseOperationPredicate) (result ListApplicationByIdExtensionPropertiesCompleteResult, err error) {
	items := make([]models.ExtensionPropertyCollectionResponse, 0)

	resp, err := c.ListApplicationByIdExtensionProperties(ctx, id)
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

	result = ListApplicationByIdExtensionPropertiesCompleteResult{
		Items: items,
	}
	return
}
