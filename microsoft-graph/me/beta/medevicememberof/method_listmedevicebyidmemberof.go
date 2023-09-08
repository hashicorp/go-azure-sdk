package medevicememberof

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

type ListMeDeviceByIdMemberOfOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListMeDeviceByIdMemberOfCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListMeDeviceByIdMemberOf ...
func (c MeDeviceMemberOfClient) ListMeDeviceByIdMemberOf(ctx context.Context, id MeDeviceId) (result ListMeDeviceByIdMemberOfOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/memberOf", id.ID()),
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

// ListMeDeviceByIdMemberOfComplete retrieves all the results into a single object
func (c MeDeviceMemberOfClient) ListMeDeviceByIdMemberOfComplete(ctx context.Context, id MeDeviceId) (ListMeDeviceByIdMemberOfCompleteResult, error) {
	return c.ListMeDeviceByIdMemberOfCompleteMatchingPredicate(ctx, id, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListMeDeviceByIdMemberOfCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeDeviceMemberOfClient) ListMeDeviceByIdMemberOfCompleteMatchingPredicate(ctx context.Context, id MeDeviceId, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListMeDeviceByIdMemberOfCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListMeDeviceByIdMemberOf(ctx, id)
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

	result = ListMeDeviceByIdMemberOfCompleteResult{
		Items: items,
	}
	return
}
