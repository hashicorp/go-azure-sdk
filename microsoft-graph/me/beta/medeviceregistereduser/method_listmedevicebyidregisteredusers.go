package medeviceregistereduser

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

type ListMeDeviceByIdRegisteredUsersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListMeDeviceByIdRegisteredUsersCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListMeDeviceByIdRegisteredUsers ...
func (c MeDeviceRegisteredUserClient) ListMeDeviceByIdRegisteredUsers(ctx context.Context, id MeDeviceId) (result ListMeDeviceByIdRegisteredUsersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/registeredUsers", id.ID()),
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
		Values *[]models.DirectoryObjectCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeDeviceByIdRegisteredUsersComplete retrieves all the results into a single object
func (c MeDeviceRegisteredUserClient) ListMeDeviceByIdRegisteredUsersComplete(ctx context.Context, id MeDeviceId) (ListMeDeviceByIdRegisteredUsersCompleteResult, error) {
	return c.ListMeDeviceByIdRegisteredUsersCompleteMatchingPredicate(ctx, id, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListMeDeviceByIdRegisteredUsersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeDeviceRegisteredUserClient) ListMeDeviceByIdRegisteredUsersCompleteMatchingPredicate(ctx context.Context, id MeDeviceId, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListMeDeviceByIdRegisteredUsersCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListMeDeviceByIdRegisteredUsers(ctx, id)
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

	result = ListMeDeviceByIdRegisteredUsersCompleteResult{
		Items: items,
	}
	return
}
