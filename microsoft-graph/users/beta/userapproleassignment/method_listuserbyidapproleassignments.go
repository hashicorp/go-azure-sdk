package userapproleassignment

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

type ListUserByIdAppRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AppRoleAssignmentCollectionResponse
}

type ListUserByIdAppRoleAssignmentsCompleteResult struct {
	Items []models.AppRoleAssignmentCollectionResponse
}

// ListUserByIdAppRoleAssignments ...
func (c UserAppRoleAssignmentClient) ListUserByIdAppRoleAssignments(ctx context.Context, id UserId) (result ListUserByIdAppRoleAssignmentsOperationResponse, err error) {
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

// ListUserByIdAppRoleAssignmentsComplete retrieves all the results into a single object
func (c UserAppRoleAssignmentClient) ListUserByIdAppRoleAssignmentsComplete(ctx context.Context, id UserId) (ListUserByIdAppRoleAssignmentsCompleteResult, error) {
	return c.ListUserByIdAppRoleAssignmentsCompleteMatchingPredicate(ctx, id, models.AppRoleAssignmentCollectionResponseOperationPredicate{})
}

// ListUserByIdAppRoleAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserAppRoleAssignmentClient) ListUserByIdAppRoleAssignmentsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.AppRoleAssignmentCollectionResponseOperationPredicate) (result ListUserByIdAppRoleAssignmentsCompleteResult, err error) {
	items := make([]models.AppRoleAssignmentCollectionResponse, 0)

	resp, err := c.ListUserByIdAppRoleAssignments(ctx, id)
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

	result = ListUserByIdAppRoleAssignmentsCompleteResult{
		Items: items,
	}
	return
}
