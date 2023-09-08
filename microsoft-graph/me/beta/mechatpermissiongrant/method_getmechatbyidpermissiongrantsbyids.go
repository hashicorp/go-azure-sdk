package mechatpermissiongrant

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

type GetMeChatByIdPermissionGrantsByIdsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObject
}

type GetMeChatByIdPermissionGrantsByIdsCompleteResult struct {
	Items []models.DirectoryObject
}

// GetMeChatByIdPermissionGrantsByIds ...
func (c MeChatPermissionGrantClient) GetMeChatByIdPermissionGrantsByIds(ctx context.Context, id MeChatId) (result GetMeChatByIdPermissionGrantsByIdsOperationResponse, err error) {
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

// GetMeChatByIdPermissionGrantsByIdsComplete retrieves all the results into a single object
func (c MeChatPermissionGrantClient) GetMeChatByIdPermissionGrantsByIdsComplete(ctx context.Context, id MeChatId) (GetMeChatByIdPermissionGrantsByIdsCompleteResult, error) {
	return c.GetMeChatByIdPermissionGrantsByIdsCompleteMatchingPredicate(ctx, id, models.DirectoryObjectOperationPredicate{})
}

// GetMeChatByIdPermissionGrantsByIdsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeChatPermissionGrantClient) GetMeChatByIdPermissionGrantsByIdsCompleteMatchingPredicate(ctx context.Context, id MeChatId, predicate models.DirectoryObjectOperationPredicate) (result GetMeChatByIdPermissionGrantsByIdsCompleteResult, err error) {
	items := make([]models.DirectoryObject, 0)

	resp, err := c.GetMeChatByIdPermissionGrantsByIds(ctx, id)
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

	result = GetMeChatByIdPermissionGrantsByIdsCompleteResult{
		Items: items,
	}
	return
}
