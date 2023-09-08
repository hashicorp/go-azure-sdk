package memailfolderuserconfiguration

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

type ListMeMailFolderByIdUserConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UserConfigurationCollectionResponse
}

type ListMeMailFolderByIdUserConfigurationsCompleteResult struct {
	Items []models.UserConfigurationCollectionResponse
}

// ListMeMailFolderByIdUserConfigurations ...
func (c MeMailFolderUserConfigurationClient) ListMeMailFolderByIdUserConfigurations(ctx context.Context, id MeMailFolderId) (result ListMeMailFolderByIdUserConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/userConfigurations", id.ID()),
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
		Values *[]models.UserConfigurationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeMailFolderByIdUserConfigurationsComplete retrieves all the results into a single object
func (c MeMailFolderUserConfigurationClient) ListMeMailFolderByIdUserConfigurationsComplete(ctx context.Context, id MeMailFolderId) (ListMeMailFolderByIdUserConfigurationsCompleteResult, error) {
	return c.ListMeMailFolderByIdUserConfigurationsCompleteMatchingPredicate(ctx, id, models.UserConfigurationCollectionResponseOperationPredicate{})
}

// ListMeMailFolderByIdUserConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeMailFolderUserConfigurationClient) ListMeMailFolderByIdUserConfigurationsCompleteMatchingPredicate(ctx context.Context, id MeMailFolderId, predicate models.UserConfigurationCollectionResponseOperationPredicate) (result ListMeMailFolderByIdUserConfigurationsCompleteResult, err error) {
	items := make([]models.UserConfigurationCollectionResponse, 0)

	resp, err := c.ListMeMailFolderByIdUserConfigurations(ctx, id)
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

	result = ListMeMailFolderByIdUserConfigurationsCompleteResult{
		Items: items,
	}
	return
}
