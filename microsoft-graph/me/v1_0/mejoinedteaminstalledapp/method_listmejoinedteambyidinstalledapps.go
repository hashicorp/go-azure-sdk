package mejoinedteaminstalledapp

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

type ListMeJoinedTeamByIdInstalledAppsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TeamsAppInstallationCollectionResponse
}

type ListMeJoinedTeamByIdInstalledAppsCompleteResult struct {
	Items []models.TeamsAppInstallationCollectionResponse
}

// ListMeJoinedTeamByIdInstalledApps ...
func (c MeJoinedTeamInstalledAppClient) ListMeJoinedTeamByIdInstalledApps(ctx context.Context, id MeJoinedTeamId) (result ListMeJoinedTeamByIdInstalledAppsOperationResponse, err error) {
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

// ListMeJoinedTeamByIdInstalledAppsComplete retrieves all the results into a single object
func (c MeJoinedTeamInstalledAppClient) ListMeJoinedTeamByIdInstalledAppsComplete(ctx context.Context, id MeJoinedTeamId) (ListMeJoinedTeamByIdInstalledAppsCompleteResult, error) {
	return c.ListMeJoinedTeamByIdInstalledAppsCompleteMatchingPredicate(ctx, id, models.TeamsAppInstallationCollectionResponseOperationPredicate{})
}

// ListMeJoinedTeamByIdInstalledAppsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeJoinedTeamInstalledAppClient) ListMeJoinedTeamByIdInstalledAppsCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamId, predicate models.TeamsAppInstallationCollectionResponseOperationPredicate) (result ListMeJoinedTeamByIdInstalledAppsCompleteResult, err error) {
	items := make([]models.TeamsAppInstallationCollectionResponse, 0)

	resp, err := c.ListMeJoinedTeamByIdInstalledApps(ctx, id)
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

	result = ListMeJoinedTeamByIdInstalledAppsCompleteResult{
		Items: items,
	}
	return
}
