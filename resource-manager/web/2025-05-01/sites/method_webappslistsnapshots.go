package sites

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

type WebAppsListSnapshotsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Snapshot
}

type WebAppsListSnapshotsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Snapshot
}

type WebAppsListSnapshotsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListSnapshotsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListSnapshots ...
func (c SitesClient) WebAppsListSnapshots(ctx context.Context, id commonids.AppServiceId) (result WebAppsListSnapshotsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListSnapshotsCustomPager{},
		Path:       fmt.Sprintf("%s/snapshots", id.ID()),
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
		Values *[]Snapshot `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListSnapshotsComplete retrieves all the results into a single object
func (c SitesClient) WebAppsListSnapshotsComplete(ctx context.Context, id commonids.AppServiceId) (WebAppsListSnapshotsCompleteResult, error) {
	return c.WebAppsListSnapshotsCompleteMatchingPredicate(ctx, id, SnapshotOperationPredicate{})
}

// WebAppsListSnapshotsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SitesClient) WebAppsListSnapshotsCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, predicate SnapshotOperationPredicate) (result WebAppsListSnapshotsCompleteResult, err error) {
	items := make([]Snapshot, 0)

	resp, err := c.WebAppsListSnapshots(ctx, id)
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

	result = WebAppsListSnapshotsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
