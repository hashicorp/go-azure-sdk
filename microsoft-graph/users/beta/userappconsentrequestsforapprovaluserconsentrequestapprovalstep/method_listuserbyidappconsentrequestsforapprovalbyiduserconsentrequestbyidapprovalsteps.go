package userappconsentrequestsforapprovaluserconsentrequestapprovalstep

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

type ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalStepsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ApprovalStepCollectionResponse
}

type ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalStepsCompleteResult struct {
	Items []models.ApprovalStepCollectionResponse
}

// ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalSteps ...
func (c UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepClient) ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalSteps(ctx context.Context, id UserAppConsentRequestsForApprovalUserConsentRequestId) (result ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalStepsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/approval/steps", id.ID()),
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

// ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalStepsComplete retrieves all the results into a single object
func (c UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepClient) ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalStepsComplete(ctx context.Context, id UserAppConsentRequestsForApprovalUserConsentRequestId) (ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalStepsCompleteResult, error) {
	return c.ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalStepsCompleteMatchingPredicate(ctx, id, models.ApprovalStepCollectionResponseOperationPredicate{})
}

// ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalStepsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepClient) ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalStepsCompleteMatchingPredicate(ctx context.Context, id UserAppConsentRequestsForApprovalUserConsentRequestId, predicate models.ApprovalStepCollectionResponseOperationPredicate) (result ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalStepsCompleteResult, err error) {
	items := make([]models.ApprovalStepCollectionResponse, 0)

	resp, err := c.ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalSteps(ctx, id)
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

	result = ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalStepsCompleteResult{
		Items: items,
	}
	return
}
