package teaminstalledapp

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

type ListTeamInstalledAppsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.TeamsAppInstallation
}

type ListTeamInstalledAppsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.TeamsAppInstallation
}

type ListTeamInstalledAppsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamInstalledAppsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamInstalledApps ...
func (c TeamInstalledAppClient) ListTeamInstalledApps(ctx context.Context, id GroupId) (result ListTeamInstalledAppsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTeamInstalledAppsCustomPager{},
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
		Values *[]stable.TeamsAppInstallation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamInstalledAppsComplete retrieves all the results into a single object
func (c TeamInstalledAppClient) ListTeamInstalledAppsComplete(ctx context.Context, id GroupId) (ListTeamInstalledAppsCompleteResult, error) {
	return c.ListTeamInstalledAppsCompleteMatchingPredicate(ctx, id, TeamsAppInstallationOperationPredicate{})
}

// ListTeamInstalledAppsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamInstalledAppClient) ListTeamInstalledAppsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate TeamsAppInstallationOperationPredicate) (result ListTeamInstalledAppsCompleteResult, err error) {
	items := make([]stable.TeamsAppInstallation, 0)

	resp, err := c.ListTeamInstalledApps(ctx, id)
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

	result = ListTeamInstalledAppsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
