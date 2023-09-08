package userteamworkinstalledapp

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

type ListUserByIdTeamworkInstalledAppsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UserScopeTeamsAppInstallationCollectionResponse
}

type ListUserByIdTeamworkInstalledAppsCompleteResult struct {
	Items []models.UserScopeTeamsAppInstallationCollectionResponse
}

// ListUserByIdTeamworkInstalledApps ...
func (c UserTeamworkInstalledAppClient) ListUserByIdTeamworkInstalledApps(ctx context.Context, id UserId) (result ListUserByIdTeamworkInstalledAppsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/teamwork/installedApps", id.ID()),
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
		Values *[]models.UserScopeTeamsAppInstallationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdTeamworkInstalledAppsComplete retrieves all the results into a single object
func (c UserTeamworkInstalledAppClient) ListUserByIdTeamworkInstalledAppsComplete(ctx context.Context, id UserId) (ListUserByIdTeamworkInstalledAppsCompleteResult, error) {
	return c.ListUserByIdTeamworkInstalledAppsCompleteMatchingPredicate(ctx, id, models.UserScopeTeamsAppInstallationCollectionResponseOperationPredicate{})
}

// ListUserByIdTeamworkInstalledAppsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserTeamworkInstalledAppClient) ListUserByIdTeamworkInstalledAppsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.UserScopeTeamsAppInstallationCollectionResponseOperationPredicate) (result ListUserByIdTeamworkInstalledAppsCompleteResult, err error) {
	items := make([]models.UserScopeTeamsAppInstallationCollectionResponse, 0)

	resp, err := c.ListUserByIdTeamworkInstalledApps(ctx, id)
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

	result = ListUserByIdTeamworkInstalledAppsCompleteResult{
		Items: items,
	}
	return
}
