package mecloudpc

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

type ListMeCloudPCsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.CloudPCCollectionResponse
}

type ListMeCloudPCsCompleteResult struct {
	Items []models.CloudPCCollectionResponse
}

// ListMeCloudPCs ...
func (c MeCloudPCClient) ListMeCloudPCs(ctx context.Context) (result ListMeCloudPCsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/cloudPCs",
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

// ListMeCloudPCsComplete retrieves all the results into a single object
func (c MeCloudPCClient) ListMeCloudPCsComplete(ctx context.Context) (ListMeCloudPCsCompleteResult, error) {
	return c.ListMeCloudPCsCompleteMatchingPredicate(ctx, models.CloudPCCollectionResponseOperationPredicate{})
}

// ListMeCloudPCsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCloudPCClient) ListMeCloudPCsCompleteMatchingPredicate(ctx context.Context, predicate models.CloudPCCollectionResponseOperationPredicate) (result ListMeCloudPCsCompleteResult, err error) {
	items := make([]models.CloudPCCollectionResponse, 0)

	resp, err := c.ListMeCloudPCs(ctx)
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

	result = ListMeCloudPCsCompleteResult{
		Items: items,
	}
	return
}
