package grouppolicyconfiguration

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

type AssignDeviceManagementGroupPolicyConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicyConfigurationAssignment
}

type AssignDeviceManagementGroupPolicyConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicyConfigurationAssignment
}

type AssignDeviceManagementGroupPolicyConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AssignDeviceManagementGroupPolicyConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AssignDeviceManagementGroupPolicyConfigurations ...
func (c GroupPolicyConfigurationClient) AssignDeviceManagementGroupPolicyConfigurations(ctx context.Context, id DeviceManagementGroupPolicyConfigurationId, input AssignDeviceManagementGroupPolicyConfigurationsRequest) (result AssignDeviceManagementGroupPolicyConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &AssignDeviceManagementGroupPolicyConfigurationsCustomPager{},
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
		Values *[]beta.GroupPolicyConfigurationAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AssignDeviceManagementGroupPolicyConfigurationsComplete retrieves all the results into a single object
func (c GroupPolicyConfigurationClient) AssignDeviceManagementGroupPolicyConfigurationsComplete(ctx context.Context, id DeviceManagementGroupPolicyConfigurationId, input AssignDeviceManagementGroupPolicyConfigurationsRequest) (AssignDeviceManagementGroupPolicyConfigurationsCompleteResult, error) {
	return c.AssignDeviceManagementGroupPolicyConfigurationsCompleteMatchingPredicate(ctx, id, input, GroupPolicyConfigurationAssignmentOperationPredicate{})
}

// AssignDeviceManagementGroupPolicyConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyConfigurationClient) AssignDeviceManagementGroupPolicyConfigurationsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementGroupPolicyConfigurationId, input AssignDeviceManagementGroupPolicyConfigurationsRequest, predicate GroupPolicyConfigurationAssignmentOperationPredicate) (result AssignDeviceManagementGroupPolicyConfigurationsCompleteResult, err error) {
	items := make([]beta.GroupPolicyConfigurationAssignment, 0)

	resp, err := c.AssignDeviceManagementGroupPolicyConfigurations(ctx, id, input)
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

	result = AssignDeviceManagementGroupPolicyConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
