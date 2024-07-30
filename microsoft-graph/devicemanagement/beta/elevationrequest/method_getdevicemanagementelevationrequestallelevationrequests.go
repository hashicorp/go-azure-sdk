package elevationrequest

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

type GetDeviceManagementElevationRequestAllElevationRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PrivilegeManagementElevationRequest
}

type GetDeviceManagementElevationRequestAllElevationRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PrivilegeManagementElevationRequest
}

type GetDeviceManagementElevationRequestAllElevationRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetDeviceManagementElevationRequestAllElevationRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetDeviceManagementElevationRequestAllElevationRequests ...
func (c ElevationRequestClient) GetDeviceManagementElevationRequestAllElevationRequests(ctx context.Context, id DeviceManagementElevationRequestId) (result GetDeviceManagementElevationRequestAllElevationRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &GetDeviceManagementElevationRequestAllElevationRequestsCustomPager{},
		Path:       fmt.Sprintf("%s/getAllElevationRequests", id.ID()),
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
		Values *[]beta.PrivilegeManagementElevationRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetDeviceManagementElevationRequestAllElevationRequestsComplete retrieves all the results into a single object
func (c ElevationRequestClient) GetDeviceManagementElevationRequestAllElevationRequestsComplete(ctx context.Context, id DeviceManagementElevationRequestId) (GetDeviceManagementElevationRequestAllElevationRequestsCompleteResult, error) {
	return c.GetDeviceManagementElevationRequestAllElevationRequestsCompleteMatchingPredicate(ctx, id, PrivilegeManagementElevationRequestOperationPredicate{})
}

// GetDeviceManagementElevationRequestAllElevationRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ElevationRequestClient) GetDeviceManagementElevationRequestAllElevationRequestsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementElevationRequestId, predicate PrivilegeManagementElevationRequestOperationPredicate) (result GetDeviceManagementElevationRequestAllElevationRequestsCompleteResult, err error) {
	items := make([]beta.PrivilegeManagementElevationRequest, 0)

	resp, err := c.GetDeviceManagementElevationRequestAllElevationRequests(ctx, id)
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

	result = GetDeviceManagementElevationRequestAllElevationRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
