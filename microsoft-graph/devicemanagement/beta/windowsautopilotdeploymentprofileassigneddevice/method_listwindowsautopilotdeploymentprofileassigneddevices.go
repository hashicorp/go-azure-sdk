package windowsautopilotdeploymentprofileassigneddevice

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

type ListWindowsAutopilotDeploymentProfileAssignedDevicesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.WindowsAutopilotDeviceIdentity
}

type ListWindowsAutopilotDeploymentProfileAssignedDevicesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.WindowsAutopilotDeviceIdentity
}

type ListWindowsAutopilotDeploymentProfileAssignedDevicesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListWindowsAutopilotDeploymentProfileAssignedDevicesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListWindowsAutopilotDeploymentProfileAssignedDevices ...
func (c WindowsAutopilotDeploymentProfileAssignedDeviceClient) ListWindowsAutopilotDeploymentProfileAssignedDevices(ctx context.Context, id DeviceManagementWindowsAutopilotDeploymentProfileId) (result ListWindowsAutopilotDeploymentProfileAssignedDevicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListWindowsAutopilotDeploymentProfileAssignedDevicesCustomPager{},
		Path:       fmt.Sprintf("%s/assignedDevices", id.ID()),
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
		Values *[]beta.WindowsAutopilotDeviceIdentity `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListWindowsAutopilotDeploymentProfileAssignedDevicesComplete retrieves all the results into a single object
func (c WindowsAutopilotDeploymentProfileAssignedDeviceClient) ListWindowsAutopilotDeploymentProfileAssignedDevicesComplete(ctx context.Context, id DeviceManagementWindowsAutopilotDeploymentProfileId) (ListWindowsAutopilotDeploymentProfileAssignedDevicesCompleteResult, error) {
	return c.ListWindowsAutopilotDeploymentProfileAssignedDevicesCompleteMatchingPredicate(ctx, id, WindowsAutopilotDeviceIdentityOperationPredicate{})
}

// ListWindowsAutopilotDeploymentProfileAssignedDevicesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WindowsAutopilotDeploymentProfileAssignedDeviceClient) ListWindowsAutopilotDeploymentProfileAssignedDevicesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementWindowsAutopilotDeploymentProfileId, predicate WindowsAutopilotDeviceIdentityOperationPredicate) (result ListWindowsAutopilotDeploymentProfileAssignedDevicesCompleteResult, err error) {
	items := make([]beta.WindowsAutopilotDeviceIdentity, 0)

	resp, err := c.ListWindowsAutopilotDeploymentProfileAssignedDevices(ctx, id)
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

	result = ListWindowsAutopilotDeploymentProfileAssignedDevicesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
