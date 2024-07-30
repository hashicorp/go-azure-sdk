package joinedteaminstalledapp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListJoinedTeamInstalledAppsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.TeamsAppInstallation
}

type ListJoinedTeamInstalledAppsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.TeamsAppInstallation
}

type ListJoinedTeamInstalledAppsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamInstalledAppsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamInstalledApps ...
func (c JoinedTeamInstalledAppClient) ListJoinedTeamInstalledApps(ctx context.Context, id MeJoinedTeamId) (result ListJoinedTeamInstalledAppsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListJoinedTeamInstalledAppsCustomPager{},
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
		Values *[]stable.TeamsAppInstallation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamInstalledAppsComplete retrieves all the results into a single object
func (c JoinedTeamInstalledAppClient) ListJoinedTeamInstalledAppsComplete(ctx context.Context, id MeJoinedTeamId) (ListJoinedTeamInstalledAppsCompleteResult, error) {
	return c.ListJoinedTeamInstalledAppsCompleteMatchingPredicate(ctx, id, TeamsAppInstallationOperationPredicate{})
}

// ListJoinedTeamInstalledAppsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamInstalledAppClient) ListJoinedTeamInstalledAppsCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamId, predicate TeamsAppInstallationOperationPredicate) (result ListJoinedTeamInstalledAppsCompleteResult, err error) {
	items := make([]stable.TeamsAppInstallation, 0)

	resp, err := c.ListJoinedTeamInstalledApps(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = ListJoinedTeamInstalledAppsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
