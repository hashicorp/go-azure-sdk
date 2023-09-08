package meapproval

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

type ListMeApprovalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ApprovalCollectionResponse
}

type ListMeApprovalsCompleteResult struct {
	Items []models.ApprovalCollectionResponse
}

// ListMeApprovals ...
func (c MeApprovalClient) ListMeApprovals(ctx context.Context) (result ListMeApprovalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/approvals",
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

// ListMeApprovalsComplete retrieves all the results into a single object
func (c MeApprovalClient) ListMeApprovalsComplete(ctx context.Context) (ListMeApprovalsCompleteResult, error) {
	return c.ListMeApprovalsCompleteMatchingPredicate(ctx, models.ApprovalCollectionResponseOperationPredicate{})
}

// ListMeApprovalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeApprovalClient) ListMeApprovalsCompleteMatchingPredicate(ctx context.Context, predicate models.ApprovalCollectionResponseOperationPredicate) (result ListMeApprovalsCompleteResult, err error) {
	items := make([]models.ApprovalCollectionResponse, 0)

	resp, err := c.ListMeApprovals(ctx)
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

	result = ListMeApprovalsCompleteResult{
		Items: items,
	}
	return
}
