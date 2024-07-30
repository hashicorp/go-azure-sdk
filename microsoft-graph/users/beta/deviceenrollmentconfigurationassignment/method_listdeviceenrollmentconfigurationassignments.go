package deviceenrollmentconfigurationassignment

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

type ListDeviceEnrollmentConfigurationAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.EnrollmentConfigurationAssignment
}

type ListDeviceEnrollmentConfigurationAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.EnrollmentConfigurationAssignment
}

type ListDeviceEnrollmentConfigurationAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceEnrollmentConfigurationAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceEnrollmentConfigurationAssignments ...
func (c DeviceEnrollmentConfigurationAssignmentClient) ListDeviceEnrollmentConfigurationAssignments(ctx context.Context, id UserIdDeviceEnrollmentConfigurationId) (result ListDeviceEnrollmentConfigurationAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceEnrollmentConfigurationAssignmentsCustomPager{},
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
		Values *[]beta.EnrollmentConfigurationAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceEnrollmentConfigurationAssignmentsComplete retrieves all the results into a single object
func (c DeviceEnrollmentConfigurationAssignmentClient) ListDeviceEnrollmentConfigurationAssignmentsComplete(ctx context.Context, id UserIdDeviceEnrollmentConfigurationId) (ListDeviceEnrollmentConfigurationAssignmentsCompleteResult, error) {
	return c.ListDeviceEnrollmentConfigurationAssignmentsCompleteMatchingPredicate(ctx, id, EnrollmentConfigurationAssignmentOperationPredicate{})
}

// ListDeviceEnrollmentConfigurationAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceEnrollmentConfigurationAssignmentClient) ListDeviceEnrollmentConfigurationAssignmentsCompleteMatchingPredicate(ctx context.Context, id UserIdDeviceEnrollmentConfigurationId, predicate EnrollmentConfigurationAssignmentOperationPredicate) (result ListDeviceEnrollmentConfigurationAssignmentsCompleteResult, err error) {
	items := make([]beta.EnrollmentConfigurationAssignment, 0)

	resp, err := c.ListDeviceEnrollmentConfigurationAssignments(ctx, id)
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

	result = ListDeviceEnrollmentConfigurationAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
