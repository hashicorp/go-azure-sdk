package userdevicecommand

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

type ListUserByIdDeviceByIdCommandsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.CommandCollectionResponse
}

type ListUserByIdDeviceByIdCommandsCompleteResult struct {
	Items []models.CommandCollectionResponse
}

// ListUserByIdDeviceByIdCommands ...
func (c UserDeviceCommandClient) ListUserByIdDeviceByIdCommands(ctx context.Context, id UserDeviceId) (result ListUserByIdDeviceByIdCommandsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/commands", id.ID()),
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
		Values *[]models.CommandCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdDeviceByIdCommandsComplete retrieves all the results into a single object
func (c UserDeviceCommandClient) ListUserByIdDeviceByIdCommandsComplete(ctx context.Context, id UserDeviceId) (ListUserByIdDeviceByIdCommandsCompleteResult, error) {
	return c.ListUserByIdDeviceByIdCommandsCompleteMatchingPredicate(ctx, id, models.CommandCollectionResponseOperationPredicate{})
}

// ListUserByIdDeviceByIdCommandsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserDeviceCommandClient) ListUserByIdDeviceByIdCommandsCompleteMatchingPredicate(ctx context.Context, id UserDeviceId, predicate models.CommandCollectionResponseOperationPredicate) (result ListUserByIdDeviceByIdCommandsCompleteResult, err error) {
	items := make([]models.CommandCollectionResponse, 0)

	resp, err := c.ListUserByIdDeviceByIdCommands(ctx, id)
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

	result = ListUserByIdDeviceByIdCommandsCompleteResult{
		Items: items,
	}
	return
}
