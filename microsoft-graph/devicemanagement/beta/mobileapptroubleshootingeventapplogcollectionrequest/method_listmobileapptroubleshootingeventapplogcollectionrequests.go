package mobileapptroubleshootingeventapplogcollectionrequest

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

type ListMobileAppTroubleshootingEventAppLogCollectionRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AppLogCollectionRequest
}

type ListMobileAppTroubleshootingEventAppLogCollectionRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AppLogCollectionRequest
}

type ListMobileAppTroubleshootingEventAppLogCollectionRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMobileAppTroubleshootingEventAppLogCollectionRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMobileAppTroubleshootingEventAppLogCollectionRequests ...
func (c MobileAppTroubleshootingEventAppLogCollectionRequestClient) ListMobileAppTroubleshootingEventAppLogCollectionRequests(ctx context.Context, id DeviceManagementMobileAppTroubleshootingEventId) (result ListMobileAppTroubleshootingEventAppLogCollectionRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMobileAppTroubleshootingEventAppLogCollectionRequestsCustomPager{},
		Path:       fmt.Sprintf("%s/appLogCollectionRequests", id.ID()),
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
		Values *[]beta.AppLogCollectionRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMobileAppTroubleshootingEventAppLogCollectionRequestsComplete retrieves all the results into a single object
func (c MobileAppTroubleshootingEventAppLogCollectionRequestClient) ListMobileAppTroubleshootingEventAppLogCollectionRequestsComplete(ctx context.Context, id DeviceManagementMobileAppTroubleshootingEventId) (ListMobileAppTroubleshootingEventAppLogCollectionRequestsCompleteResult, error) {
	return c.ListMobileAppTroubleshootingEventAppLogCollectionRequestsCompleteMatchingPredicate(ctx, id, AppLogCollectionRequestOperationPredicate{})
}

// ListMobileAppTroubleshootingEventAppLogCollectionRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MobileAppTroubleshootingEventAppLogCollectionRequestClient) ListMobileAppTroubleshootingEventAppLogCollectionRequestsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementMobileAppTroubleshootingEventId, predicate AppLogCollectionRequestOperationPredicate) (result ListMobileAppTroubleshootingEventAppLogCollectionRequestsCompleteResult, err error) {
	items := make([]beta.AppLogCollectionRequest, 0)

	resp, err := c.ListMobileAppTroubleshootingEventAppLogCollectionRequests(ctx, id)
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

	result = ListMobileAppTroubleshootingEventAppLogCollectionRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
