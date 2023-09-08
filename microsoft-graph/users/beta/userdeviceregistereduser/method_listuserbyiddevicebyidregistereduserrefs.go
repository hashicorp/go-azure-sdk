package userdeviceregistereduser

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

type ListUserByIdDeviceByIdRegisteredUserRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.StringCollectionResponse
}

type ListUserByIdDeviceByIdRegisteredUserRefsCompleteResult struct {
	Items []models.StringCollectionResponse
}

// ListUserByIdDeviceByIdRegisteredUserRefs ...
func (c UserDeviceRegisteredUserClient) ListUserByIdDeviceByIdRegisteredUserRefs(ctx context.Context, id UserDeviceId) (result ListUserByIdDeviceByIdRegisteredUserRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/registeredUsers/$ref", id.ID()),
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
		Values *[]models.StringCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdDeviceByIdRegisteredUserRefsComplete retrieves all the results into a single object
func (c UserDeviceRegisteredUserClient) ListUserByIdDeviceByIdRegisteredUserRefsComplete(ctx context.Context, id UserDeviceId) (ListUserByIdDeviceByIdRegisteredUserRefsCompleteResult, error) {
	return c.ListUserByIdDeviceByIdRegisteredUserRefsCompleteMatchingPredicate(ctx, id, models.StringCollectionResponseOperationPredicate{})
}

// ListUserByIdDeviceByIdRegisteredUserRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserDeviceRegisteredUserClient) ListUserByIdDeviceByIdRegisteredUserRefsCompleteMatchingPredicate(ctx context.Context, id UserDeviceId, predicate models.StringCollectionResponseOperationPredicate) (result ListUserByIdDeviceByIdRegisteredUserRefsCompleteResult, err error) {
	items := make([]models.StringCollectionResponse, 0)

	resp, err := c.ListUserByIdDeviceByIdRegisteredUserRefs(ctx, id)
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

	result = ListUserByIdDeviceByIdRegisteredUserRefsCompleteResult{
		Items: items,
	}
	return
}
