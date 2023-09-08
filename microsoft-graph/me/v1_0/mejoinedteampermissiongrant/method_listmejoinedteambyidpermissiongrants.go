package mejoinedteampermissiongrant

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

type ListMeJoinedTeamByIdPermissionGrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ResourceSpecificPermissionGrantCollectionResponse
}

type ListMeJoinedTeamByIdPermissionGrantsCompleteResult struct {
	Items []models.ResourceSpecificPermissionGrantCollectionResponse
}

// ListMeJoinedTeamByIdPermissionGrants ...
func (c MeJoinedTeamPermissionGrantClient) ListMeJoinedTeamByIdPermissionGrants(ctx context.Context, id MeJoinedTeamId) (result ListMeJoinedTeamByIdPermissionGrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/permissionGrants", id.ID()),
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

// ListMeJoinedTeamByIdPermissionGrantsComplete retrieves all the results into a single object
func (c MeJoinedTeamPermissionGrantClient) ListMeJoinedTeamByIdPermissionGrantsComplete(ctx context.Context, id MeJoinedTeamId) (ListMeJoinedTeamByIdPermissionGrantsCompleteResult, error) {
	return c.ListMeJoinedTeamByIdPermissionGrantsCompleteMatchingPredicate(ctx, id, models.ResourceSpecificPermissionGrantCollectionResponseOperationPredicate{})
}

// ListMeJoinedTeamByIdPermissionGrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeJoinedTeamPermissionGrantClient) ListMeJoinedTeamByIdPermissionGrantsCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamId, predicate models.ResourceSpecificPermissionGrantCollectionResponseOperationPredicate) (result ListMeJoinedTeamByIdPermissionGrantsCompleteResult, err error) {
	items := make([]models.ResourceSpecificPermissionGrantCollectionResponse, 0)

	resp, err := c.ListMeJoinedTeamByIdPermissionGrants(ctx, id)
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

	result = ListMeJoinedTeamByIdPermissionGrantsCompleteResult{
		Items: items,
	}
	return
}
