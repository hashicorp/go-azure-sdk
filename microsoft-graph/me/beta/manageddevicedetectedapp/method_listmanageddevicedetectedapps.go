package manageddevicedetectedapp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListManagedDeviceDetectedAppsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DetectedApp
}

type ListManagedDeviceDetectedAppsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DetectedApp
}

type ListManagedDeviceDetectedAppsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListManagedDeviceDetectedAppsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListManagedDeviceDetectedApps ...
func (c ManagedDeviceDetectedAppClient) ListManagedDeviceDetectedApps(ctx context.Context, id MeManagedDeviceId) (result ListManagedDeviceDetectedAppsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListManagedDeviceDetectedAppsCustomPager{},
		Path:       fmt.Sprintf("%s/detectedApps", id.ID()),
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
		Values *[]beta.DetectedApp `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListManagedDeviceDetectedAppsComplete retrieves all the results into a single object
func (c ManagedDeviceDetectedAppClient) ListManagedDeviceDetectedAppsComplete(ctx context.Context, id MeManagedDeviceId) (ListManagedDeviceDetectedAppsCompleteResult, error) {
	return c.ListManagedDeviceDetectedAppsCompleteMatchingPredicate(ctx, id, DetectedAppOperationPredicate{})
}

// ListManagedDeviceDetectedAppsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedDeviceDetectedAppClient) ListManagedDeviceDetectedAppsCompleteMatchingPredicate(ctx context.Context, id MeManagedDeviceId, predicate DetectedAppOperationPredicate) (result ListManagedDeviceDetectedAppsCompleteResult, err error) {
	items := make([]beta.DetectedApp, 0)

	resp, err := c.ListManagedDeviceDetectedApps(ctx, id)
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

	result = ListManagedDeviceDetectedAppsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
