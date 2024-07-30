package deviceconfigurationassignment

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

type ListDeviceConfigurationAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceConfigurationAssignment
}

type ListDeviceConfigurationAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceConfigurationAssignment
}

type ListDeviceConfigurationAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceConfigurationAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceConfigurationAssignments ...
func (c DeviceConfigurationAssignmentClient) ListDeviceConfigurationAssignments(ctx context.Context, id DeviceManagementDeviceConfigurationId) (result ListDeviceConfigurationAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceConfigurationAssignmentsCustomPager{},
		Path:       fmt.Sprintf("%s/assignments", id.ID()),
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
		Values *[]beta.DeviceConfigurationAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceConfigurationAssignmentsComplete retrieves all the results into a single object
func (c DeviceConfigurationAssignmentClient) ListDeviceConfigurationAssignmentsComplete(ctx context.Context, id DeviceManagementDeviceConfigurationId) (ListDeviceConfigurationAssignmentsCompleteResult, error) {
	return c.ListDeviceConfigurationAssignmentsCompleteMatchingPredicate(ctx, id, DeviceConfigurationAssignmentOperationPredicate{})
}

// ListDeviceConfigurationAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceConfigurationAssignmentClient) ListDeviceConfigurationAssignmentsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementDeviceConfigurationId, predicate DeviceConfigurationAssignmentOperationPredicate) (result ListDeviceConfigurationAssignmentsCompleteResult, err error) {
	items := make([]beta.DeviceConfigurationAssignment, 0)

	resp, err := c.ListDeviceConfigurationAssignments(ctx, id)
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

	result = ListDeviceConfigurationAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
