package identityb2xuserflowuserattributeassignment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListIdentityB2xUserFlowByIdUserAttributeAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.IdentityUserFlowAttributeAssignmentCollectionResponse
}

type ListIdentityB2xUserFlowByIdUserAttributeAssignmentsCompleteResult struct {
	Items []models.IdentityUserFlowAttributeAssignmentCollectionResponse
}

// ListIdentityB2xUserFlowByIdUserAttributeAssignments ...
func (c IdentityB2xUserFlowUserAttributeAssignmentClient) ListIdentityB2xUserFlowByIdUserAttributeAssignments(ctx context.Context, id IdentityB2xUserFlowId) (result ListIdentityB2xUserFlowByIdUserAttributeAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/userAttributeAssignments", id.ID()),
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
		Values *[]models.IdentityUserFlowAttributeAssignmentCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIdentityB2xUserFlowByIdUserAttributeAssignmentsComplete retrieves all the results into a single object
func (c IdentityB2xUserFlowUserAttributeAssignmentClient) ListIdentityB2xUserFlowByIdUserAttributeAssignmentsComplete(ctx context.Context, id IdentityB2xUserFlowId) (ListIdentityB2xUserFlowByIdUserAttributeAssignmentsCompleteResult, error) {
	return c.ListIdentityB2xUserFlowByIdUserAttributeAssignmentsCompleteMatchingPredicate(ctx, id, models.IdentityUserFlowAttributeAssignmentCollectionResponseOperationPredicate{})
}

// ListIdentityB2xUserFlowByIdUserAttributeAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IdentityB2xUserFlowUserAttributeAssignmentClient) ListIdentityB2xUserFlowByIdUserAttributeAssignmentsCompleteMatchingPredicate(ctx context.Context, id IdentityB2xUserFlowId, predicate models.IdentityUserFlowAttributeAssignmentCollectionResponseOperationPredicate) (result ListIdentityB2xUserFlowByIdUserAttributeAssignmentsCompleteResult, err error) {
	items := make([]models.IdentityUserFlowAttributeAssignmentCollectionResponse, 0)

	resp, err := c.ListIdentityB2xUserFlowByIdUserAttributeAssignments(ctx, id)
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

	result = ListIdentityB2xUserFlowByIdUserAttributeAssignmentsCompleteResult{
		Items: items,
	}
	return
}
