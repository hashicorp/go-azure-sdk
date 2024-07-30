package resourceaccessprofile

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

type AssignDeviceManagementResourceAccessProfilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementResourceAccessProfileAssignment
}

type AssignDeviceManagementResourceAccessProfilesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementResourceAccessProfileAssignment
}

type AssignDeviceManagementResourceAccessProfilesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AssignDeviceManagementResourceAccessProfilesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AssignDeviceManagementResourceAccessProfiles ...
func (c ResourceAccessProfileClient) AssignDeviceManagementResourceAccessProfiles(ctx context.Context, id DeviceManagementResourceAccessProfileId, input AssignDeviceManagementResourceAccessProfilesRequest) (result AssignDeviceManagementResourceAccessProfilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &AssignDeviceManagementResourceAccessProfilesCustomPager{},
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
		Values *[]beta.DeviceManagementResourceAccessProfileAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AssignDeviceManagementResourceAccessProfilesComplete retrieves all the results into a single object
func (c ResourceAccessProfileClient) AssignDeviceManagementResourceAccessProfilesComplete(ctx context.Context, id DeviceManagementResourceAccessProfileId, input AssignDeviceManagementResourceAccessProfilesRequest) (AssignDeviceManagementResourceAccessProfilesCompleteResult, error) {
	return c.AssignDeviceManagementResourceAccessProfilesCompleteMatchingPredicate(ctx, id, input, DeviceManagementResourceAccessProfileAssignmentOperationPredicate{})
}

// AssignDeviceManagementResourceAccessProfilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ResourceAccessProfileClient) AssignDeviceManagementResourceAccessProfilesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementResourceAccessProfileId, input AssignDeviceManagementResourceAccessProfilesRequest, predicate DeviceManagementResourceAccessProfileAssignmentOperationPredicate) (result AssignDeviceManagementResourceAccessProfilesCompleteResult, err error) {
	items := make([]beta.DeviceManagementResourceAccessProfileAssignment, 0)

	resp, err := c.AssignDeviceManagementResourceAccessProfiles(ctx, id, input)
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

	result = AssignDeviceManagementResourceAccessProfilesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
