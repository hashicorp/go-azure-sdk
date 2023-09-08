package userdeviceregisteredowner

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

type ListUserByIdDeviceByIdRegisteredOwnerRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.StringCollectionResponse
}

type ListUserByIdDeviceByIdRegisteredOwnerRefsCompleteResult struct {
	Items []models.StringCollectionResponse
}

// ListUserByIdDeviceByIdRegisteredOwnerRefs ...
func (c UserDeviceRegisteredOwnerClient) ListUserByIdDeviceByIdRegisteredOwnerRefs(ctx context.Context, id UserDeviceId) (result ListUserByIdDeviceByIdRegisteredOwnerRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/registeredOwners/$ref", id.ID()),
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

// ListUserByIdDeviceByIdRegisteredOwnerRefsComplete retrieves all the results into a single object
func (c UserDeviceRegisteredOwnerClient) ListUserByIdDeviceByIdRegisteredOwnerRefsComplete(ctx context.Context, id UserDeviceId) (ListUserByIdDeviceByIdRegisteredOwnerRefsCompleteResult, error) {
	return c.ListUserByIdDeviceByIdRegisteredOwnerRefsCompleteMatchingPredicate(ctx, id, models.StringCollectionResponseOperationPredicate{})
}

// ListUserByIdDeviceByIdRegisteredOwnerRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserDeviceRegisteredOwnerClient) ListUserByIdDeviceByIdRegisteredOwnerRefsCompleteMatchingPredicate(ctx context.Context, id UserDeviceId, predicate models.StringCollectionResponseOperationPredicate) (result ListUserByIdDeviceByIdRegisteredOwnerRefsCompleteResult, err error) {
	items := make([]models.StringCollectionResponse, 0)

	resp, err := c.ListUserByIdDeviceByIdRegisteredOwnerRefs(ctx, id)
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

	result = ListUserByIdDeviceByIdRegisteredOwnerRefsCompleteResult{
		Items: items,
	}
	return
}
