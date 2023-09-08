package medevicecommand

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

type ListMeDeviceByIdCommandsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.CommandCollectionResponse
}

type ListMeDeviceByIdCommandsCompleteResult struct {
	Items []models.CommandCollectionResponse
}

// ListMeDeviceByIdCommands ...
func (c MeDeviceCommandClient) ListMeDeviceByIdCommands(ctx context.Context, id MeDeviceId) (result ListMeDeviceByIdCommandsOperationResponse, err error) {
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

// ListMeDeviceByIdCommandsComplete retrieves all the results into a single object
func (c MeDeviceCommandClient) ListMeDeviceByIdCommandsComplete(ctx context.Context, id MeDeviceId) (ListMeDeviceByIdCommandsCompleteResult, error) {
	return c.ListMeDeviceByIdCommandsCompleteMatchingPredicate(ctx, id, models.CommandCollectionResponseOperationPredicate{})
}

// ListMeDeviceByIdCommandsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeDeviceCommandClient) ListMeDeviceByIdCommandsCompleteMatchingPredicate(ctx context.Context, id MeDeviceId, predicate models.CommandCollectionResponseOperationPredicate) (result ListMeDeviceByIdCommandsCompleteResult, err error) {
	items := make([]models.CommandCollectionResponse, 0)

	resp, err := c.ListMeDeviceByIdCommands(ctx, id)
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

	result = ListMeDeviceByIdCommandsCompleteResult{
		Items: items,
	}
	return
}
