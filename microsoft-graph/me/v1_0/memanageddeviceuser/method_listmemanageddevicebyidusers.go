package memanageddeviceuser

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

type ListMeManagedDeviceByIdUsersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UserCollectionResponse
}

type ListMeManagedDeviceByIdUsersCompleteResult struct {
	Items []models.UserCollectionResponse
}

// ListMeManagedDeviceByIdUsers ...
func (c MeManagedDeviceUserClient) ListMeManagedDeviceByIdUsers(ctx context.Context, id MeManagedDeviceId) (result ListMeManagedDeviceByIdUsersOperationResponse, err error) {
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

// ListMeManagedDeviceByIdUsersComplete retrieves all the results into a single object
func (c MeManagedDeviceUserClient) ListMeManagedDeviceByIdUsersComplete(ctx context.Context, id MeManagedDeviceId) (ListMeManagedDeviceByIdUsersCompleteResult, error) {
	return c.ListMeManagedDeviceByIdUsersCompleteMatchingPredicate(ctx, id, models.UserCollectionResponseOperationPredicate{})
}

// ListMeManagedDeviceByIdUsersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeManagedDeviceUserClient) ListMeManagedDeviceByIdUsersCompleteMatchingPredicate(ctx context.Context, id MeManagedDeviceId, predicate models.UserCollectionResponseOperationPredicate) (result ListMeManagedDeviceByIdUsersCompleteResult, err error) {
	items := make([]models.UserCollectionResponse, 0)

	resp, err := c.ListMeManagedDeviceByIdUsers(ctx, id)
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

	result = ListMeManagedDeviceByIdUsersCompleteResult{
		Items: items,
	}
	return
}
