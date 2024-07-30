package manageddeviceassignmentfilterevaluationstatusdetail

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

type ListManagedDeviceAssignmentFilterEvaluationStatusDetailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AssignmentFilterEvaluationStatusDetails
}

type ListManagedDeviceAssignmentFilterEvaluationStatusDetailsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AssignmentFilterEvaluationStatusDetails
}

type ListManagedDeviceAssignmentFilterEvaluationStatusDetailsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListManagedDeviceAssignmentFilterEvaluationStatusDetailsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListManagedDeviceAssignmentFilterEvaluationStatusDetails ...
func (c ManagedDeviceAssignmentFilterEvaluationStatusDetailClient) ListManagedDeviceAssignmentFilterEvaluationStatusDetails(ctx context.Context, id UserIdManagedDeviceId) (result ListManagedDeviceAssignmentFilterEvaluationStatusDetailsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListManagedDeviceAssignmentFilterEvaluationStatusDetailsCustomPager{},
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

// ListManagedDeviceAssignmentFilterEvaluationStatusDetailsComplete retrieves all the results into a single object
func (c ManagedDeviceAssignmentFilterEvaluationStatusDetailClient) ListManagedDeviceAssignmentFilterEvaluationStatusDetailsComplete(ctx context.Context, id UserIdManagedDeviceId) (ListManagedDeviceAssignmentFilterEvaluationStatusDetailsCompleteResult, error) {
	return c.ListManagedDeviceAssignmentFilterEvaluationStatusDetailsCompleteMatchingPredicate(ctx, id, AssignmentFilterEvaluationStatusDetailsOperationPredicate{})
}

// ListManagedDeviceAssignmentFilterEvaluationStatusDetailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedDeviceAssignmentFilterEvaluationStatusDetailClient) ListManagedDeviceAssignmentFilterEvaluationStatusDetailsCompleteMatchingPredicate(ctx context.Context, id UserIdManagedDeviceId, predicate AssignmentFilterEvaluationStatusDetailsOperationPredicate) (result ListManagedDeviceAssignmentFilterEvaluationStatusDetailsCompleteResult, err error) {
	items := make([]beta.AssignmentFilterEvaluationStatusDetails, 0)

	resp, err := c.ListManagedDeviceAssignmentFilterEvaluationStatusDetails(ctx, id)
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

	result = ListManagedDeviceAssignmentFilterEvaluationStatusDetailsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
