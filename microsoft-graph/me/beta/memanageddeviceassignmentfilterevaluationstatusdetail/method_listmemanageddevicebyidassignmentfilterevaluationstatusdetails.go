package memanageddeviceassignmentfilterevaluationstatusdetail

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMeManagedDeviceByIdAssignmentFilterEvaluationStatusDetailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AssignmentFilterEvaluationStatusDetailsCollectionResponse
}

type ListMeManagedDeviceByIdAssignmentFilterEvaluationStatusDetailsCompleteResult struct {
	Items []models.AssignmentFilterEvaluationStatusDetailsCollectionResponse
}

// ListMeManagedDeviceByIdAssignmentFilterEvaluationStatusDetails ...
func (c MeManagedDeviceAssignmentFilterEvaluationStatusDetailClient) ListMeManagedDeviceByIdAssignmentFilterEvaluationStatusDetails(ctx context.Context, id MeManagedDeviceId) (result ListMeManagedDeviceByIdAssignmentFilterEvaluationStatusDetailsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]models.AssignmentFilterEvaluationStatusDetailsCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeManagedDeviceByIdAssignmentFilterEvaluationStatusDetailsComplete retrieves all the results into a single object
func (c MeManagedDeviceAssignmentFilterEvaluationStatusDetailClient) ListMeManagedDeviceByIdAssignmentFilterEvaluationStatusDetailsComplete(ctx context.Context, id MeManagedDeviceId) (ListMeManagedDeviceByIdAssignmentFilterEvaluationStatusDetailsCompleteResult, error) {
	return c.ListMeManagedDeviceByIdAssignmentFilterEvaluationStatusDetailsCompleteMatchingPredicate(ctx, id, models.AssignmentFilterEvaluationStatusDetailsCollectionResponseOperationPredicate{})
}

// ListMeManagedDeviceByIdAssignmentFilterEvaluationStatusDetailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeManagedDeviceAssignmentFilterEvaluationStatusDetailClient) ListMeManagedDeviceByIdAssignmentFilterEvaluationStatusDetailsCompleteMatchingPredicate(ctx context.Context, id MeManagedDeviceId, predicate models.AssignmentFilterEvaluationStatusDetailsCollectionResponseOperationPredicate) (result ListMeManagedDeviceByIdAssignmentFilterEvaluationStatusDetailsCompleteResult, err error) {
	items := make([]models.AssignmentFilterEvaluationStatusDetailsCollectionResponse, 0)

	resp, err := c.ListMeManagedDeviceByIdAssignmentFilterEvaluationStatusDetails(ctx, id)
	if err != nil {
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

	result = ListMeManagedDeviceByIdAssignmentFilterEvaluationStatusDetailsCompleteResult{
		Items: items,
	}
	return
}
