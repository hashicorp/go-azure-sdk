package deviceconfiguration

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

type AssignDeviceManagementDeviceConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DeviceConfigurationAssignment
}

type AssignDeviceManagementDeviceConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DeviceConfigurationAssignment
}

type AssignDeviceManagementDeviceConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AssignDeviceManagementDeviceConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AssignDeviceManagementDeviceConfigurations ...
func (c DeviceConfigurationClient) AssignDeviceManagementDeviceConfigurations(ctx context.Context, id DeviceManagementDeviceConfigurationId, input AssignDeviceManagementDeviceConfigurationsRequest) (result AssignDeviceManagementDeviceConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &AssignDeviceManagementDeviceConfigurationsCustomPager{},
		Path:       fmt.Sprintf("%s/assign", id.ID()),
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
		Values *[]stable.DeviceConfigurationAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AssignDeviceManagementDeviceConfigurationsComplete retrieves all the results into a single object
func (c DeviceConfigurationClient) AssignDeviceManagementDeviceConfigurationsComplete(ctx context.Context, id DeviceManagementDeviceConfigurationId, input AssignDeviceManagementDeviceConfigurationsRequest) (AssignDeviceManagementDeviceConfigurationsCompleteResult, error) {
	return c.AssignDeviceManagementDeviceConfigurationsCompleteMatchingPredicate(ctx, id, input, DeviceConfigurationAssignmentOperationPredicate{})
}

// AssignDeviceManagementDeviceConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceConfigurationClient) AssignDeviceManagementDeviceConfigurationsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementDeviceConfigurationId, input AssignDeviceManagementDeviceConfigurationsRequest, predicate DeviceConfigurationAssignmentOperationPredicate) (result AssignDeviceManagementDeviceConfigurationsCompleteResult, err error) {
	items := make([]stable.DeviceConfigurationAssignment, 0)

	resp, err := c.AssignDeviceManagementDeviceConfigurations(ctx, id, input)
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

	result = AssignDeviceManagementDeviceConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
