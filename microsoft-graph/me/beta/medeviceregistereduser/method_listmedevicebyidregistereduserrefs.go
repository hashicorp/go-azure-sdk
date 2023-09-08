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

type ListMeDeviceByIdRegisteredUserRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.StringCollectionResponse
}

type ListMeDeviceByIdRegisteredUserRefsCompleteResult struct {
	Items []models.StringCollectionResponse
}

// ListMeDeviceByIdRegisteredUserRefs ...
func (c MeDeviceRegisteredUserClient) ListMeDeviceByIdRegisteredUserRefs(ctx context.Context, id MeDeviceId) (result ListMeDeviceByIdRegisteredUserRefsOperationResponse, err error) {
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

// ListMeDeviceByIdRegisteredUserRefsComplete retrieves all the results into a single object
func (c MeDeviceRegisteredUserClient) ListMeDeviceByIdRegisteredUserRefsComplete(ctx context.Context, id MeDeviceId) (ListMeDeviceByIdRegisteredUserRefsCompleteResult, error) {
	return c.ListMeDeviceByIdRegisteredUserRefsCompleteMatchingPredicate(ctx, id, models.StringCollectionResponseOperationPredicate{})
}

// ListMeDeviceByIdRegisteredUserRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeDeviceRegisteredUserClient) ListMeDeviceByIdRegisteredUserRefsCompleteMatchingPredicate(ctx context.Context, id MeDeviceId, predicate models.StringCollectionResponseOperationPredicate) (result ListMeDeviceByIdRegisteredUserRefsCompleteResult, err error) {
	items := make([]models.StringCollectionResponse, 0)

	resp, err := c.ListMeDeviceByIdRegisteredUserRefs(ctx, id)
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

	result = ListMeDeviceByIdRegisteredUserRefsCompleteResult{
		Items: items,
	}
	return
}
