package devicecompliancepolicydevicestatus

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

type ListDeviceCompliancePolicyDeviceStatusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DeviceComplianceDeviceStatus
}

type ListDeviceCompliancePolicyDeviceStatusCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DeviceComplianceDeviceStatus
}

type ListDeviceCompliancePolicyDeviceStatusCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceCompliancePolicyDeviceStatusCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceCompliancePolicyDeviceStatus ...
func (c DeviceCompliancePolicyDeviceStatusClient) ListDeviceCompliancePolicyDeviceStatus(ctx context.Context, id DeviceManagementDeviceCompliancePolicyId) (result ListDeviceCompliancePolicyDeviceStatusOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceCompliancePolicyDeviceStatusCustomPager{},
		Path:       fmt.Sprintf("%s/deviceStatuses", id.ID()),
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
		Values *[]stable.DeviceComplianceDeviceStatus `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceCompliancePolicyDeviceStatusComplete retrieves all the results into a single object
func (c DeviceCompliancePolicyDeviceStatusClient) ListDeviceCompliancePolicyDeviceStatusComplete(ctx context.Context, id DeviceManagementDeviceCompliancePolicyId) (ListDeviceCompliancePolicyDeviceStatusCompleteResult, error) {
	return c.ListDeviceCompliancePolicyDeviceStatusCompleteMatchingPredicate(ctx, id, DeviceComplianceDeviceStatusOperationPredicate{})
}

// ListDeviceCompliancePolicyDeviceStatusCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCompliancePolicyDeviceStatusClient) ListDeviceCompliancePolicyDeviceStatusCompleteMatchingPredicate(ctx context.Context, id DeviceManagementDeviceCompliancePolicyId, predicate DeviceComplianceDeviceStatusOperationPredicate) (result ListDeviceCompliancePolicyDeviceStatusCompleteResult, err error) {
	items := make([]stable.DeviceComplianceDeviceStatus, 0)

	resp, err := c.ListDeviceCompliancePolicyDeviceStatus(ctx, id)
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

	result = ListDeviceCompliancePolicyDeviceStatusCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
