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

type GetMeJoinedTeamByIdPermissionGrantsByIdsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObject
}

type GetMeJoinedTeamByIdPermissionGrantsByIdsCompleteResult struct {
	Items []models.DirectoryObject
}

// GetMeJoinedTeamByIdPermissionGrantsByIds ...
func (c MeJoinedTeamPermissionGrantClient) GetMeJoinedTeamByIdPermissionGrantsByIds(ctx context.Context, id MeJoinedTeamId) (result GetMeJoinedTeamByIdPermissionGrantsByIdsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/permissionGrants/getByIds", id.ID()),
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
		Values *[]models.DirectoryObject `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetMeJoinedTeamByIdPermissionGrantsByIdsComplete retrieves all the results into a single object
func (c MeJoinedTeamPermissionGrantClient) GetMeJoinedTeamByIdPermissionGrantsByIdsComplete(ctx context.Context, id MeJoinedTeamId) (GetMeJoinedTeamByIdPermissionGrantsByIdsCompleteResult, error) {
	return c.GetMeJoinedTeamByIdPermissionGrantsByIdsCompleteMatchingPredicate(ctx, id, models.DirectoryObjectOperationPredicate{})
}

// GetMeJoinedTeamByIdPermissionGrantsByIdsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeJoinedTeamPermissionGrantClient) GetMeJoinedTeamByIdPermissionGrantsByIdsCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamId, predicate models.DirectoryObjectOperationPredicate) (result GetMeJoinedTeamByIdPermissionGrantsByIdsCompleteResult, err error) {
	items := make([]models.DirectoryObject, 0)

	resp, err := c.GetMeJoinedTeamByIdPermissionGrantsByIds(ctx, id)
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

	result = GetMeJoinedTeamByIdPermissionGrantsByIdsCompleteResult{
		Items: items,
	}
	return
}
