package groupapproleassignment

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

type ListGroupByIdAppRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AppRoleAssignmentCollectionResponse
}

type ListGroupByIdAppRoleAssignmentsCompleteResult struct {
	Items []models.AppRoleAssignmentCollectionResponse
}

// ListGroupByIdAppRoleAssignments ...
func (c GroupAppRoleAssignmentClient) ListGroupByIdAppRoleAssignments(ctx context.Context, id GroupId) (result ListGroupByIdAppRoleAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/appRoleAssignments", id.ID()),
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
		Values *[]models.AppRoleAssignmentCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdAppRoleAssignmentsComplete retrieves all the results into a single object
func (c GroupAppRoleAssignmentClient) ListGroupByIdAppRoleAssignmentsComplete(ctx context.Context, id GroupId) (ListGroupByIdAppRoleAssignmentsCompleteResult, error) {
	return c.ListGroupByIdAppRoleAssignmentsCompleteMatchingPredicate(ctx, id, models.AppRoleAssignmentCollectionResponseOperationPredicate{})
}

// ListGroupByIdAppRoleAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupAppRoleAssignmentClient) ListGroupByIdAppRoleAssignmentsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.AppRoleAssignmentCollectionResponseOperationPredicate) (result ListGroupByIdAppRoleAssignmentsCompleteResult, err error) {
	items := make([]models.AppRoleAssignmentCollectionResponse, 0)

	resp, err := c.ListGroupByIdAppRoleAssignments(ctx, id)
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

	result = ListGroupByIdAppRoleAssignmentsCompleteResult{
		Items: items,
	}
	return
}
