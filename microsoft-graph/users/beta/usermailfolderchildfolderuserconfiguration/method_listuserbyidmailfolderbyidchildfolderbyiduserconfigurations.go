package usermailfolderchildfolderuserconfiguration

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

type ListUserByIdMailFolderByIdChildFolderByIdUserConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UserConfigurationCollectionResponse
}

type ListUserByIdMailFolderByIdChildFolderByIdUserConfigurationsCompleteResult struct {
	Items []models.UserConfigurationCollectionResponse
}

// ListUserByIdMailFolderByIdChildFolderByIdUserConfigurations ...
func (c UserMailFolderChildFolderUserConfigurationClient) ListUserByIdMailFolderByIdChildFolderByIdUserConfigurations(ctx context.Context, id UserMailFolderChildFolderId) (result ListUserByIdMailFolderByIdChildFolderByIdUserConfigurationsOperationResponse, err error) {
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

// ListUserByIdMailFolderByIdChildFolderByIdUserConfigurationsComplete retrieves all the results into a single object
func (c UserMailFolderChildFolderUserConfigurationClient) ListUserByIdMailFolderByIdChildFolderByIdUserConfigurationsComplete(ctx context.Context, id UserMailFolderChildFolderId) (ListUserByIdMailFolderByIdChildFolderByIdUserConfigurationsCompleteResult, error) {
	return c.ListUserByIdMailFolderByIdChildFolderByIdUserConfigurationsCompleteMatchingPredicate(ctx, id, models.UserConfigurationCollectionResponseOperationPredicate{})
}

// ListUserByIdMailFolderByIdChildFolderByIdUserConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserMailFolderChildFolderUserConfigurationClient) ListUserByIdMailFolderByIdChildFolderByIdUserConfigurationsCompleteMatchingPredicate(ctx context.Context, id UserMailFolderChildFolderId, predicate models.UserConfigurationCollectionResponseOperationPredicate) (result ListUserByIdMailFolderByIdChildFolderByIdUserConfigurationsCompleteResult, err error) {
	items := make([]models.UserConfigurationCollectionResponse, 0)

	resp, err := c.ListUserByIdMailFolderByIdChildFolderByIdUserConfigurations(ctx, id)
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

	result = ListUserByIdMailFolderByIdChildFolderByIdUserConfigurationsCompleteResult{
		Items: items,
	}
	return
}
