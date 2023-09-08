package userapprovalstep

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

type ListUserByIdApprovalByIdStepsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ApprovalStepCollectionResponse
}

type ListUserByIdApprovalByIdStepsCompleteResult struct {
	Items []models.ApprovalStepCollectionResponse
}

// ListUserByIdApprovalByIdSteps ...
func (c UserApprovalStepClient) ListUserByIdApprovalByIdSteps(ctx context.Context, id UserApprovalId) (result ListUserByIdApprovalByIdStepsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/steps", id.ID()),
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
		Values *[]models.ApprovalStepCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdApprovalByIdStepsComplete retrieves all the results into a single object
func (c UserApprovalStepClient) ListUserByIdApprovalByIdStepsComplete(ctx context.Context, id UserApprovalId) (ListUserByIdApprovalByIdStepsCompleteResult, error) {
	return c.ListUserByIdApprovalByIdStepsCompleteMatchingPredicate(ctx, id, models.ApprovalStepCollectionResponseOperationPredicate{})
}

// ListUserByIdApprovalByIdStepsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserApprovalStepClient) ListUserByIdApprovalByIdStepsCompleteMatchingPredicate(ctx context.Context, id UserApprovalId, predicate models.ApprovalStepCollectionResponseOperationPredicate) (result ListUserByIdApprovalByIdStepsCompleteResult, err error) {
	items := make([]models.ApprovalStepCollectionResponse, 0)

	resp, err := c.ListUserByIdApprovalByIdSteps(ctx, id)
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

	result = ListUserByIdApprovalByIdStepsCompleteResult{
		Items: items,
	}
	return
}
