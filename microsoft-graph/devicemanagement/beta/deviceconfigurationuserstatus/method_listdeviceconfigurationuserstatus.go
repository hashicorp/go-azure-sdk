package deviceconfigurationuserstatus

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

type ListDeviceConfigurationUserStatusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceConfigurationUserStatus
}

type ListDeviceConfigurationUserStatusCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceConfigurationUserStatus
}

type ListDeviceConfigurationUserStatusCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceConfigurationUserStatusCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceConfigurationUserStatus ...
func (c DeviceConfigurationUserStatusClient) ListDeviceConfigurationUserStatus(ctx context.Context, id DeviceManagementDeviceConfigurationId) (result ListDeviceConfigurationUserStatusOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceConfigurationUserStatusCustomPager{},
		Path:       fmt.Sprintf("%s/userStatuses", id.ID()),
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
		Values *[]beta.DeviceConfigurationUserStatus `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceConfigurationUserStatusComplete retrieves all the results into a single object
func (c DeviceConfigurationUserStatusClient) ListDeviceConfigurationUserStatusComplete(ctx context.Context, id DeviceManagementDeviceConfigurationId) (ListDeviceConfigurationUserStatusCompleteResult, error) {
	return c.ListDeviceConfigurationUserStatusCompleteMatchingPredicate(ctx, id, DeviceConfigurationUserStatusOperationPredicate{})
}

// ListDeviceConfigurationUserStatusCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceConfigurationUserStatusClient) ListDeviceConfigurationUserStatusCompleteMatchingPredicate(ctx context.Context, id DeviceManagementDeviceConfigurationId, predicate DeviceConfigurationUserStatusOperationPredicate) (result ListDeviceConfigurationUserStatusCompleteResult, err error) {
	items := make([]beta.DeviceConfigurationUserStatus, 0)

	resp, err := c.ListDeviceConfigurationUserStatus(ctx, id)
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

	result = ListDeviceConfigurationUserStatusCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
