package groupteampermissiongrant

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

type ListGroupByIdTeamPermissionGrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ResourceSpecificPermissionGrantCollectionResponse
}

type ListGroupByIdTeamPermissionGrantsCompleteResult struct {
	Items []models.ResourceSpecificPermissionGrantCollectionResponse
}

// ListGroupByIdTeamPermissionGrants ...
func (c GroupTeamPermissionGrantClient) ListGroupByIdTeamPermissionGrants(ctx context.Context, id GroupId) (result ListGroupByIdTeamPermissionGrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/team/permissionGrants", id.ID()),
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
		Values *[]models.ResourceSpecificPermissionGrantCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdTeamPermissionGrantsComplete retrieves all the results into a single object
func (c GroupTeamPermissionGrantClient) ListGroupByIdTeamPermissionGrantsComplete(ctx context.Context, id GroupId) (ListGroupByIdTeamPermissionGrantsCompleteResult, error) {
	return c.ListGroupByIdTeamPermissionGrantsCompleteMatchingPredicate(ctx, id, models.ResourceSpecificPermissionGrantCollectionResponseOperationPredicate{})
}

// ListGroupByIdTeamPermissionGrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupTeamPermissionGrantClient) ListGroupByIdTeamPermissionGrantsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.ResourceSpecificPermissionGrantCollectionResponseOperationPredicate) (result ListGroupByIdTeamPermissionGrantsCompleteResult, err error) {
	items := make([]models.ResourceSpecificPermissionGrantCollectionResponse, 0)

	resp, err := c.ListGroupByIdTeamPermissionGrants(ctx, id)
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

	result = ListGroupByIdTeamPermissionGrantsCompleteResult{
		Items: items,
	}
	return
}
