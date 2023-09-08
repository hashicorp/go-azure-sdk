package meappconsentrequestsforapprovaluserconsentrequestapprovalstep

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

type ListMeAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalStepsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ApprovalStepCollectionResponse
}

type ListMeAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalStepsCompleteResult struct {
	Items []models.ApprovalStepCollectionResponse
}

// ListMeAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalSteps ...
func (c MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepClient) ListMeAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalSteps(ctx context.Context, id MeAppConsentRequestsForApprovalUserConsentRequestId) (result ListMeAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalStepsOperationResponse, err error) {
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

// ListMeAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalStepsComplete retrieves all the results into a single object
func (c MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepClient) ListMeAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalStepsComplete(ctx context.Context, id MeAppConsentRequestsForApprovalUserConsentRequestId) (ListMeAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalStepsCompleteResult, error) {
	return c.ListMeAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalStepsCompleteMatchingPredicate(ctx, id, models.ApprovalStepCollectionResponseOperationPredicate{})
}

// ListMeAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalStepsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepClient) ListMeAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalStepsCompleteMatchingPredicate(ctx context.Context, id MeAppConsentRequestsForApprovalUserConsentRequestId, predicate models.ApprovalStepCollectionResponseOperationPredicate) (result ListMeAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalStepsCompleteResult, err error) {
	items := make([]models.ApprovalStepCollectionResponse, 0)

	resp, err := c.ListMeAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalSteps(ctx, id)
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

	result = ListMeAppConsentRequestsForApprovalByIdUserConsentRequestByIdApprovalStepsCompleteResult{
		Items: items,
	}
	return
}
