package meteamworkinstalledapp

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

type ListMeTeamworkInstalledAppsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UserScopeTeamsAppInstallationCollectionResponse
}

type ListMeTeamworkInstalledAppsCompleteResult struct {
	Items []models.UserScopeTeamsAppInstallationCollectionResponse
}

// ListMeTeamworkInstalledApps ...
func (c MeTeamworkInstalledAppClient) ListMeTeamworkInstalledApps(ctx context.Context) (result ListMeTeamworkInstalledAppsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/teamwork/installedApps",
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

// ListMeTeamworkInstalledAppsComplete retrieves all the results into a single object
func (c MeTeamworkInstalledAppClient) ListMeTeamworkInstalledAppsComplete(ctx context.Context) (ListMeTeamworkInstalledAppsCompleteResult, error) {
	return c.ListMeTeamworkInstalledAppsCompleteMatchingPredicate(ctx, models.UserScopeTeamsAppInstallationCollectionResponseOperationPredicate{})
}

// ListMeTeamworkInstalledAppsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeTeamworkInstalledAppClient) ListMeTeamworkInstalledAppsCompleteMatchingPredicate(ctx context.Context, predicate models.UserScopeTeamsAppInstallationCollectionResponseOperationPredicate) (result ListMeTeamworkInstalledAppsCompleteResult, err error) {
	items := make([]models.UserScopeTeamsAppInstallationCollectionResponse, 0)

	resp, err := c.ListMeTeamworkInstalledApps(ctx)
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

	result = ListMeTeamworkInstalledAppsCompleteResult{
		Items: items,
	}
	return
}
