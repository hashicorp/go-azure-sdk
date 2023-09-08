package usermailfolderuserconfiguration

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

type ListUserByIdMailFolderByIdUserConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UserConfigurationCollectionResponse
}

type ListUserByIdMailFolderByIdUserConfigurationsCompleteResult struct {
	Items []models.UserConfigurationCollectionResponse
}

// ListUserByIdMailFolderByIdUserConfigurations ...
func (c UserMailFolderUserConfigurationClient) ListUserByIdMailFolderByIdUserConfigurations(ctx context.Context, id UserMailFolderId) (result ListUserByIdMailFolderByIdUserConfigurationsOperationResponse, err error) {
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

// ListUserByIdMailFolderByIdUserConfigurationsComplete retrieves all the results into a single object
func (c UserMailFolderUserConfigurationClient) ListUserByIdMailFolderByIdUserConfigurationsComplete(ctx context.Context, id UserMailFolderId) (ListUserByIdMailFolderByIdUserConfigurationsCompleteResult, error) {
	return c.ListUserByIdMailFolderByIdUserConfigurationsCompleteMatchingPredicate(ctx, id, models.UserConfigurationCollectionResponseOperationPredicate{})
}

// ListUserByIdMailFolderByIdUserConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserMailFolderUserConfigurationClient) ListUserByIdMailFolderByIdUserConfigurationsCompleteMatchingPredicate(ctx context.Context, id UserMailFolderId, predicate models.UserConfigurationCollectionResponseOperationPredicate) (result ListUserByIdMailFolderByIdUserConfigurationsCompleteResult, err error) {
	items := make([]models.UserConfigurationCollectionResponse, 0)

	resp, err := c.ListUserByIdMailFolderByIdUserConfigurations(ctx, id)
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

	result = ListUserByIdMailFolderByIdUserConfigurationsCompleteResult{
		Items: items,
	}
	return
}
