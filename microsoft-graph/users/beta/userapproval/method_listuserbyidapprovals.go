package userapproval

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

type ListUserByIdApprovalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ApprovalCollectionResponse
}

type ListUserByIdApprovalsCompleteResult struct {
	Items []models.ApprovalCollectionResponse
}

// ListUserByIdApprovals ...
func (c UserApprovalClient) ListUserByIdApprovals(ctx context.Context, id UserId) (result ListUserByIdApprovalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/approvals", id.ID()),
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
		Values *[]models.ApprovalCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdApprovalsComplete retrieves all the results into a single object
func (c UserApprovalClient) ListUserByIdApprovalsComplete(ctx context.Context, id UserId) (ListUserByIdApprovalsCompleteResult, error) {
	return c.ListUserByIdApprovalsCompleteMatchingPredicate(ctx, id, models.ApprovalCollectionResponseOperationPredicate{})
}

// ListUserByIdApprovalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserApprovalClient) ListUserByIdApprovalsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.ApprovalCollectionResponseOperationPredicate) (result ListUserByIdApprovalsCompleteResult, err error) {
	items := make([]models.ApprovalCollectionResponse, 0)

	resp, err := c.ListUserByIdApprovals(ctx, id)
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

	result = ListUserByIdApprovalsCompleteResult{
		Items: items,
	}
	return
}
