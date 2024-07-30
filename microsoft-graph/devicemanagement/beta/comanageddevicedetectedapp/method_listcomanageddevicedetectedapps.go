package comanageddevicedetectedapp

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

type ListComanagedDeviceDetectedAppsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DetectedApp
}

type ListComanagedDeviceDetectedAppsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DetectedApp
}

type ListComanagedDeviceDetectedAppsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListComanagedDeviceDetectedAppsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListComanagedDeviceDetectedApps ...
func (c ComanagedDeviceDetectedAppClient) ListComanagedDeviceDetectedApps(ctx context.Context, id DeviceManagementComanagedDeviceId) (result ListComanagedDeviceDetectedAppsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListComanagedDeviceDetectedAppsCustomPager{},
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

// ListComanagedDeviceDetectedAppsComplete retrieves all the results into a single object
func (c ComanagedDeviceDetectedAppClient) ListComanagedDeviceDetectedAppsComplete(ctx context.Context, id DeviceManagementComanagedDeviceId) (ListComanagedDeviceDetectedAppsCompleteResult, error) {
	return c.ListComanagedDeviceDetectedAppsCompleteMatchingPredicate(ctx, id, DetectedAppOperationPredicate{})
}

// ListComanagedDeviceDetectedAppsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ComanagedDeviceDetectedAppClient) ListComanagedDeviceDetectedAppsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementComanagedDeviceId, predicate DetectedAppOperationPredicate) (result ListComanagedDeviceDetectedAppsCompleteResult, err error) {
	items := make([]beta.DetectedApp, 0)

	resp, err := c.ListComanagedDeviceDetectedApps(ctx, id)
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

	result = ListComanagedDeviceDetectedAppsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
