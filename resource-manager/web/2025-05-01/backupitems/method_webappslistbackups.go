package backupitems

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListBackupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BackupItem
}

type WebAppsListBackupsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BackupItem
}

type WebAppsListBackupsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListBackupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListBackups ...
func (c BackupItemsClient) WebAppsListBackups(ctx context.Context, id commonids.AppServiceId) (result WebAppsListBackupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListBackupsCustomPager{},
		Path:       fmt.Sprintf("%s/backups", id.ID()),
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
		Values *[]BackupItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListBackupsComplete retrieves all the results into a single object
func (c BackupItemsClient) WebAppsListBackupsComplete(ctx context.Context, id commonids.AppServiceId) (WebAppsListBackupsCompleteResult, error) {
	return c.WebAppsListBackupsCompleteMatchingPredicate(ctx, id, BackupItemOperationPredicate{})
}

// WebAppsListBackupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BackupItemsClient) WebAppsListBackupsCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, predicate BackupItemOperationPredicate) (result WebAppsListBackupsCompleteResult, err error) {
	items := make([]BackupItem, 0)

	resp, err := c.WebAppsListBackups(ctx, id)
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

	result = WebAppsListBackupsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
