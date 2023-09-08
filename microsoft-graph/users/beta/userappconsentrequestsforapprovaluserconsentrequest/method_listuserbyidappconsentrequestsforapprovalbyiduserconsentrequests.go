package userappconsentrequestsforapprovaluserconsentrequest

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

type ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UserConsentRequestCollectionResponse
}

type ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestsCompleteResult struct {
	Items []models.UserConsentRequestCollectionResponse
}

// ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequests ...
func (c UserAppConsentRequestsForApprovalUserConsentRequestClient) ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequests(ctx context.Context, id UserAppConsentRequestsForApprovalId) (result ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/userConsentRequests", id.ID()),
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
		Values *[]models.UserConsentRequestCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestsComplete retrieves all the results into a single object
func (c UserAppConsentRequestsForApprovalUserConsentRequestClient) ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestsComplete(ctx context.Context, id UserAppConsentRequestsForApprovalId) (ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestsCompleteResult, error) {
	return c.ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestsCompleteMatchingPredicate(ctx, id, models.UserConsentRequestCollectionResponseOperationPredicate{})
}

// ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserAppConsentRequestsForApprovalUserConsentRequestClient) ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestsCompleteMatchingPredicate(ctx context.Context, id UserAppConsentRequestsForApprovalId, predicate models.UserConsentRequestCollectionResponseOperationPredicate) (result ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestsCompleteResult, err error) {
	items := make([]models.UserConsentRequestCollectionResponse, 0)

	resp, err := c.ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequests(ctx, id)
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

	result = ListUserByIdAppConsentRequestsForApprovalByIdUserConsentRequestsCompleteResult{
		Items: items,
	}
	return
}
