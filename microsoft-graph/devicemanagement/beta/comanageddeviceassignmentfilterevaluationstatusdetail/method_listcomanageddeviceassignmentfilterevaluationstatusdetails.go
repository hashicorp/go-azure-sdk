package comanageddeviceassignmentfilterevaluationstatusdetail

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

type ListComanagedDeviceAssignmentFilterEvaluationStatusDetailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AssignmentFilterEvaluationStatusDetails
}

type ListComanagedDeviceAssignmentFilterEvaluationStatusDetailsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AssignmentFilterEvaluationStatusDetails
}

type ListComanagedDeviceAssignmentFilterEvaluationStatusDetailsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListComanagedDeviceAssignmentFilterEvaluationStatusDetailsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListComanagedDeviceAssignmentFilterEvaluationStatusDetails ...
func (c ComanagedDeviceAssignmentFilterEvaluationStatusDetailClient) ListComanagedDeviceAssignmentFilterEvaluationStatusDetails(ctx context.Context, id DeviceManagementComanagedDeviceId) (result ListComanagedDeviceAssignmentFilterEvaluationStatusDetailsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListComanagedDeviceAssignmentFilterEvaluationStatusDetailsCustomPager{},
		Path:       fmt.Sprintf("%s/assignmentFilterEvaluationStatusDetails", id.ID()),
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
		Values *[]beta.AssignmentFilterEvaluationStatusDetails `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComanagedDeviceAssignmentFilterEvaluationStatusDetailsComplete retrieves all the results into a single object
func (c ComanagedDeviceAssignmentFilterEvaluationStatusDetailClient) ListComanagedDeviceAssignmentFilterEvaluationStatusDetailsComplete(ctx context.Context, id DeviceManagementComanagedDeviceId) (ListComanagedDeviceAssignmentFilterEvaluationStatusDetailsCompleteResult, error) {
	return c.ListComanagedDeviceAssignmentFilterEvaluationStatusDetailsCompleteMatchingPredicate(ctx, id, AssignmentFilterEvaluationStatusDetailsOperationPredicate{})
}

// ListComanagedDeviceAssignmentFilterEvaluationStatusDetailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ComanagedDeviceAssignmentFilterEvaluationStatusDetailClient) ListComanagedDeviceAssignmentFilterEvaluationStatusDetailsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementComanagedDeviceId, predicate AssignmentFilterEvaluationStatusDetailsOperationPredicate) (result ListComanagedDeviceAssignmentFilterEvaluationStatusDetailsCompleteResult, err error) {
	items := make([]beta.AssignmentFilterEvaluationStatusDetails, 0)

	resp, err := c.ListComanagedDeviceAssignmentFilterEvaluationStatusDetails(ctx, id)
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

	result = ListComanagedDeviceAssignmentFilterEvaluationStatusDetailsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
