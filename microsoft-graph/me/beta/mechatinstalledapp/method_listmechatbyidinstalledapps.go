package mechatinstalledapp

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

type ListMeChatByIdInstalledAppsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TeamsAppInstallationCollectionResponse
}

type ListMeChatByIdInstalledAppsCompleteResult struct {
	Items []models.TeamsAppInstallationCollectionResponse
}

// ListMeChatByIdInstalledApps ...
func (c MeChatInstalledAppClient) ListMeChatByIdInstalledApps(ctx context.Context, id MeChatId) (result ListMeChatByIdInstalledAppsOperationResponse, err error) {
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

// ListMeChatByIdInstalledAppsComplete retrieves all the results into a single object
func (c MeChatInstalledAppClient) ListMeChatByIdInstalledAppsComplete(ctx context.Context, id MeChatId) (ListMeChatByIdInstalledAppsCompleteResult, error) {
	return c.ListMeChatByIdInstalledAppsCompleteMatchingPredicate(ctx, id, models.TeamsAppInstallationCollectionResponseOperationPredicate{})
}

// ListMeChatByIdInstalledAppsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeChatInstalledAppClient) ListMeChatByIdInstalledAppsCompleteMatchingPredicate(ctx context.Context, id MeChatId, predicate models.TeamsAppInstallationCollectionResponseOperationPredicate) (result ListMeChatByIdInstalledAppsCompleteResult, err error) {
	items := make([]models.TeamsAppInstallationCollectionResponse, 0)

	resp, err := c.ListMeChatByIdInstalledApps(ctx, id)
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

	result = ListMeChatByIdInstalledAppsCompleteResult{
		Items: items,
	}
	return
}
