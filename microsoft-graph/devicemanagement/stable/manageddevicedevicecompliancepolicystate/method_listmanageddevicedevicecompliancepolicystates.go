package manageddevicedevicecompliancepolicystate

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

type ListManagedDeviceDeviceCompliancePolicyStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DeviceCompliancePolicyState
}

type ListManagedDeviceDeviceCompliancePolicyStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DeviceCompliancePolicyState
}

type ListManagedDeviceDeviceCompliancePolicyStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListManagedDeviceDeviceCompliancePolicyStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListManagedDeviceDeviceCompliancePolicyStates ...
func (c ManagedDeviceDeviceCompliancePolicyStateClient) ListManagedDeviceDeviceCompliancePolicyStates(ctx context.Context, id DeviceManagementManagedDeviceId) (result ListManagedDeviceDeviceCompliancePolicyStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListManagedDeviceDeviceCompliancePolicyStatesCustomPager{},
		Path:       fmt.Sprintf("%s/deviceCompliancePolicyStates", id.ID()),
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
		Values *[]stable.DeviceCompliancePolicyState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListManagedDeviceDeviceCompliancePolicyStatesComplete retrieves all the results into a single object
func (c ManagedDeviceDeviceCompliancePolicyStateClient) ListManagedDeviceDeviceCompliancePolicyStatesComplete(ctx context.Context, id DeviceManagementManagedDeviceId) (ListManagedDeviceDeviceCompliancePolicyStatesCompleteResult, error) {
	return c.ListManagedDeviceDeviceCompliancePolicyStatesCompleteMatchingPredicate(ctx, id, DeviceCompliancePolicyStateOperationPredicate{})
}

// ListManagedDeviceDeviceCompliancePolicyStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedDeviceDeviceCompliancePolicyStateClient) ListManagedDeviceDeviceCompliancePolicyStatesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementManagedDeviceId, predicate DeviceCompliancePolicyStateOperationPredicate) (result ListManagedDeviceDeviceCompliancePolicyStatesCompleteResult, err error) {
	items := make([]stable.DeviceCompliancePolicyState, 0)

	resp, err := c.ListManagedDeviceDeviceCompliancePolicyStates(ctx, id)
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

	result = ListManagedDeviceDeviceCompliancePolicyStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
