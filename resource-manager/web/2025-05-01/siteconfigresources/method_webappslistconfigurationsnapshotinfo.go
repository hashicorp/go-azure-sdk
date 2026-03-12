package siteconfigresources

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

type WebAppsListConfigurationSnapshotInfoOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SiteConfigurationSnapshotInfo
}

type WebAppsListConfigurationSnapshotInfoCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SiteConfigurationSnapshotInfo
}

type WebAppsListConfigurationSnapshotInfoCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListConfigurationSnapshotInfoCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListConfigurationSnapshotInfo ...
func (c SiteConfigResourcesClient) WebAppsListConfigurationSnapshotInfo(ctx context.Context, id commonids.AppServiceId) (result WebAppsListConfigurationSnapshotInfoOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListConfigurationSnapshotInfoCustomPager{},
		Path:       fmt.Sprintf("%s/config/web/snapshots", id.ID()),
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
		Values *[]SiteConfigurationSnapshotInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListConfigurationSnapshotInfoComplete retrieves all the results into a single object
func (c SiteConfigResourcesClient) WebAppsListConfigurationSnapshotInfoComplete(ctx context.Context, id commonids.AppServiceId) (WebAppsListConfigurationSnapshotInfoCompleteResult, error) {
	return c.WebAppsListConfigurationSnapshotInfoCompleteMatchingPredicate(ctx, id, SiteConfigurationSnapshotInfoOperationPredicate{})
}

// WebAppsListConfigurationSnapshotInfoCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteConfigResourcesClient) WebAppsListConfigurationSnapshotInfoCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, predicate SiteConfigurationSnapshotInfoOperationPredicate) (result WebAppsListConfigurationSnapshotInfoCompleteResult, err error) {
	items := make([]SiteConfigurationSnapshotInfo, 0)

	resp, err := c.WebAppsListConfigurationSnapshotInfo(ctx, id)
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

	result = WebAppsListConfigurationSnapshotInfoCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
