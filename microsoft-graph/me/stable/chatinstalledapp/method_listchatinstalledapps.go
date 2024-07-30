package chatinstalledapp

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

type ListChatInstalledAppsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.TeamsAppInstallation
}

type ListChatInstalledAppsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.TeamsAppInstallation
}

type ListChatInstalledAppsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListChatInstalledAppsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListChatInstalledApps ...
func (c ChatInstalledAppClient) ListChatInstalledApps(ctx context.Context, id MeChatId) (result ListChatInstalledAppsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListChatInstalledAppsCustomPager{},
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

// ListChatInstalledAppsComplete retrieves all the results into a single object
func (c ChatInstalledAppClient) ListChatInstalledAppsComplete(ctx context.Context, id MeChatId) (ListChatInstalledAppsCompleteResult, error) {
	return c.ListChatInstalledAppsCompleteMatchingPredicate(ctx, id, TeamsAppInstallationOperationPredicate{})
}

// ListChatInstalledAppsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ChatInstalledAppClient) ListChatInstalledAppsCompleteMatchingPredicate(ctx context.Context, id MeChatId, predicate TeamsAppInstallationOperationPredicate) (result ListChatInstalledAppsCompleteResult, err error) {
	items := make([]stable.TeamsAppInstallation, 0)

	resp, err := c.ListChatInstalledApps(ctx, id)
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

	result = ListChatInstalledAppsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
