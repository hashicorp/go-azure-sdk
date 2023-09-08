package medeviceregisteredowner

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

type ListMeDeviceByIdRegisteredOwnerRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.StringCollectionResponse
}

type ListMeDeviceByIdRegisteredOwnerRefsCompleteResult struct {
	Items []models.StringCollectionResponse
}

// ListMeDeviceByIdRegisteredOwnerRefs ...
func (c MeDeviceRegisteredOwnerClient) ListMeDeviceByIdRegisteredOwnerRefs(ctx context.Context, id MeDeviceId) (result ListMeDeviceByIdRegisteredOwnerRefsOperationResponse, err error) {
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

// ListMeDeviceByIdRegisteredOwnerRefsComplete retrieves all the results into a single object
func (c MeDeviceRegisteredOwnerClient) ListMeDeviceByIdRegisteredOwnerRefsComplete(ctx context.Context, id MeDeviceId) (ListMeDeviceByIdRegisteredOwnerRefsCompleteResult, error) {
	return c.ListMeDeviceByIdRegisteredOwnerRefsCompleteMatchingPredicate(ctx, id, models.StringCollectionResponseOperationPredicate{})
}

// ListMeDeviceByIdRegisteredOwnerRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeDeviceRegisteredOwnerClient) ListMeDeviceByIdRegisteredOwnerRefsCompleteMatchingPredicate(ctx context.Context, id MeDeviceId, predicate models.StringCollectionResponseOperationPredicate) (result ListMeDeviceByIdRegisteredOwnerRefsCompleteResult, err error) {
	items := make([]models.StringCollectionResponse, 0)

	resp, err := c.ListMeDeviceByIdRegisteredOwnerRefs(ctx, id)
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

	result = ListMeDeviceByIdRegisteredOwnerRefsCompleteResult{
		Items: items,
	}
	return
}
