package usercloudpc

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

type ListUserByIdCloudPCsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.CloudPCCollectionResponse
}

type ListUserByIdCloudPCsCompleteResult struct {
	Items []models.CloudPCCollectionResponse
}

// ListUserByIdCloudPCs ...
func (c UserCloudPCClient) ListUserByIdCloudPCs(ctx context.Context, id UserId) (result ListUserByIdCloudPCsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/cloudPCs", id.ID()),
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
		Values *[]models.CloudPCCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdCloudPCsComplete retrieves all the results into a single object
func (c UserCloudPCClient) ListUserByIdCloudPCsComplete(ctx context.Context, id UserId) (ListUserByIdCloudPCsCompleteResult, error) {
	return c.ListUserByIdCloudPCsCompleteMatchingPredicate(ctx, id, models.CloudPCCollectionResponseOperationPredicate{})
}

// ListUserByIdCloudPCsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCloudPCClient) ListUserByIdCloudPCsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.CloudPCCollectionResponseOperationPredicate) (result ListUserByIdCloudPCsCompleteResult, err error) {
	items := make([]models.CloudPCCollectionResponse, 0)

	resp, err := c.ListUserByIdCloudPCs(ctx, id)
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

	result = ListUserByIdCloudPCsCompleteResult{
		Items: items,
	}
	return
}
