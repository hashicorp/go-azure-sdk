package groupteaminstalledapp

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

type ListGroupByIdTeamInstalledAppsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TeamsAppInstallationCollectionResponse
}

type ListGroupByIdTeamInstalledAppsCompleteResult struct {
	Items []models.TeamsAppInstallationCollectionResponse
}

// ListGroupByIdTeamInstalledApps ...
func (c GroupTeamInstalledAppClient) ListGroupByIdTeamInstalledApps(ctx context.Context, id GroupId) (result ListGroupByIdTeamInstalledAppsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/team/installedApps", id.ID()),
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

// ListGroupByIdTeamInstalledAppsComplete retrieves all the results into a single object
func (c GroupTeamInstalledAppClient) ListGroupByIdTeamInstalledAppsComplete(ctx context.Context, id GroupId) (ListGroupByIdTeamInstalledAppsCompleteResult, error) {
	return c.ListGroupByIdTeamInstalledAppsCompleteMatchingPredicate(ctx, id, models.TeamsAppInstallationCollectionResponseOperationPredicate{})
}

// ListGroupByIdTeamInstalledAppsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupTeamInstalledAppClient) ListGroupByIdTeamInstalledAppsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.TeamsAppInstallationCollectionResponseOperationPredicate) (result ListGroupByIdTeamInstalledAppsCompleteResult, err error) {
	items := make([]models.TeamsAppInstallationCollectionResponse, 0)

	resp, err := c.ListGroupByIdTeamInstalledApps(ctx, id)
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

	result = ListGroupByIdTeamInstalledAppsCompleteResult{
		Items: items,
	}
	return
}
