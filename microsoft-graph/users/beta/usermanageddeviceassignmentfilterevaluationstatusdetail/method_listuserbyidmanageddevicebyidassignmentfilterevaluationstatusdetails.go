package usermanageddeviceassignmentfilterevaluationstatusdetail

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

type ListUserByIdManagedDeviceByIdAssignmentFilterEvaluationStatusDetailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AssignmentFilterEvaluationStatusDetailsCollectionResponse
}

type ListUserByIdManagedDeviceByIdAssignmentFilterEvaluationStatusDetailsCompleteResult struct {
	Items []models.AssignmentFilterEvaluationStatusDetailsCollectionResponse
}

// ListUserByIdManagedDeviceByIdAssignmentFilterEvaluationStatusDetails ...
func (c UserManagedDeviceAssignmentFilterEvaluationStatusDetailClient) ListUserByIdManagedDeviceByIdAssignmentFilterEvaluationStatusDetails(ctx context.Context, id UserManagedDeviceId) (result ListUserByIdManagedDeviceByIdAssignmentFilterEvaluationStatusDetailsOperationResponse, err error) {
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

// ListUserByIdManagedDeviceByIdAssignmentFilterEvaluationStatusDetailsComplete retrieves all the results into a single object
func (c UserManagedDeviceAssignmentFilterEvaluationStatusDetailClient) ListUserByIdManagedDeviceByIdAssignmentFilterEvaluationStatusDetailsComplete(ctx context.Context, id UserManagedDeviceId) (ListUserByIdManagedDeviceByIdAssignmentFilterEvaluationStatusDetailsCompleteResult, error) {
	return c.ListUserByIdManagedDeviceByIdAssignmentFilterEvaluationStatusDetailsCompleteMatchingPredicate(ctx, id, models.AssignmentFilterEvaluationStatusDetailsCollectionResponseOperationPredicate{})
}

// ListUserByIdManagedDeviceByIdAssignmentFilterEvaluationStatusDetailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserManagedDeviceAssignmentFilterEvaluationStatusDetailClient) ListUserByIdManagedDeviceByIdAssignmentFilterEvaluationStatusDetailsCompleteMatchingPredicate(ctx context.Context, id UserManagedDeviceId, predicate models.AssignmentFilterEvaluationStatusDetailsCollectionResponseOperationPredicate) (result ListUserByIdManagedDeviceByIdAssignmentFilterEvaluationStatusDetailsCompleteResult, err error) {
	items := make([]models.AssignmentFilterEvaluationStatusDetailsCollectionResponse, 0)

	resp, err := c.ListUserByIdManagedDeviceByIdAssignmentFilterEvaluationStatusDetails(ctx, id)
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

	result = ListUserByIdManagedDeviceByIdAssignmentFilterEvaluationStatusDetailsCompleteResult{
		Items: items,
	}
	return
}
