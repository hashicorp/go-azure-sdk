package userchatinstalledapp

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

type ListUserByIdChatByIdInstalledAppsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TeamsAppInstallationCollectionResponse
}

type ListUserByIdChatByIdInstalledAppsCompleteResult struct {
	Items []models.TeamsAppInstallationCollectionResponse
}

// ListUserByIdChatByIdInstalledApps ...
func (c UserChatInstalledAppClient) ListUserByIdChatByIdInstalledApps(ctx context.Context, id UserChatId) (result ListUserByIdChatByIdInstalledAppsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/installedApps", id.ID()),
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
		Values *[]models.TeamsAppInstallationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdChatByIdInstalledAppsComplete retrieves all the results into a single object
func (c UserChatInstalledAppClient) ListUserByIdChatByIdInstalledAppsComplete(ctx context.Context, id UserChatId) (ListUserByIdChatByIdInstalledAppsCompleteResult, error) {
	return c.ListUserByIdChatByIdInstalledAppsCompleteMatchingPredicate(ctx, id, models.TeamsAppInstallationCollectionResponseOperationPredicate{})
}

// ListUserByIdChatByIdInstalledAppsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserChatInstalledAppClient) ListUserByIdChatByIdInstalledAppsCompleteMatchingPredicate(ctx context.Context, id UserChatId, predicate models.TeamsAppInstallationCollectionResponseOperationPredicate) (result ListUserByIdChatByIdInstalledAppsCompleteResult, err error) {
	items := make([]models.TeamsAppInstallationCollectionResponse, 0)

	resp, err := c.ListUserByIdChatByIdInstalledApps(ctx, id)
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

	result = ListUserByIdChatByIdInstalledAppsCompleteResult{
		Items: items,
	}
	return
}
