package usermanageddeviceuser

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

type ListUserByIdManagedDeviceByIdUsersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UserCollectionResponse
}

type ListUserByIdManagedDeviceByIdUsersCompleteResult struct {
	Items []models.UserCollectionResponse
}

// ListUserByIdManagedDeviceByIdUsers ...
func (c UserManagedDeviceUserClient) ListUserByIdManagedDeviceByIdUsers(ctx context.Context, id UserManagedDeviceId) (result ListUserByIdManagedDeviceByIdUsersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/users", id.ID()),
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
		Values *[]models.UserCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdManagedDeviceByIdUsersComplete retrieves all the results into a single object
func (c UserManagedDeviceUserClient) ListUserByIdManagedDeviceByIdUsersComplete(ctx context.Context, id UserManagedDeviceId) (ListUserByIdManagedDeviceByIdUsersCompleteResult, error) {
	return c.ListUserByIdManagedDeviceByIdUsersCompleteMatchingPredicate(ctx, id, models.UserCollectionResponseOperationPredicate{})
}

// ListUserByIdManagedDeviceByIdUsersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserManagedDeviceUserClient) ListUserByIdManagedDeviceByIdUsersCompleteMatchingPredicate(ctx context.Context, id UserManagedDeviceId, predicate models.UserCollectionResponseOperationPredicate) (result ListUserByIdManagedDeviceByIdUsersCompleteResult, err error) {
	items := make([]models.UserCollectionResponse, 0)

	resp, err := c.ListUserByIdManagedDeviceByIdUsers(ctx, id)
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

	result = ListUserByIdManagedDeviceByIdUsersCompleteResult{
		Items: items,
	}
	return
}
