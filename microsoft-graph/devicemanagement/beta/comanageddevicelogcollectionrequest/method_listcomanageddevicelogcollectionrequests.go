package comanageddevicelogcollectionrequest

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

type ListComanagedDeviceLogCollectionRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceLogCollectionResponse
}

type ListComanagedDeviceLogCollectionRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceLogCollectionResponse
}

type ListComanagedDeviceLogCollectionRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListComanagedDeviceLogCollectionRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListComanagedDeviceLogCollectionRequests ...
func (c ComanagedDeviceLogCollectionRequestClient) ListComanagedDeviceLogCollectionRequests(ctx context.Context, id DeviceManagementComanagedDeviceId) (result ListComanagedDeviceLogCollectionRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListComanagedDeviceLogCollectionRequestsCustomPager{},
		Path:       fmt.Sprintf("%s/logCollectionRequests", id.ID()),
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
		Values *[]beta.DeviceLogCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComanagedDeviceLogCollectionRequestsComplete retrieves all the results into a single object
func (c ComanagedDeviceLogCollectionRequestClient) ListComanagedDeviceLogCollectionRequestsComplete(ctx context.Context, id DeviceManagementComanagedDeviceId) (ListComanagedDeviceLogCollectionRequestsCompleteResult, error) {
	return c.ListComanagedDeviceLogCollectionRequestsCompleteMatchingPredicate(ctx, id, DeviceLogCollectionResponseOperationPredicate{})
}

// ListComanagedDeviceLogCollectionRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ComanagedDeviceLogCollectionRequestClient) ListComanagedDeviceLogCollectionRequestsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementComanagedDeviceId, predicate DeviceLogCollectionResponseOperationPredicate) (result ListComanagedDeviceLogCollectionRequestsCompleteResult, err error) {
	items := make([]beta.DeviceLogCollectionResponse, 0)

	resp, err := c.ListComanagedDeviceLogCollectionRequests(ctx, id)
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

	result = ListComanagedDeviceLogCollectionRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
