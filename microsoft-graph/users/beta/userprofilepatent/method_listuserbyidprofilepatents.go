package userprofilepatent

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

type ListUserByIdProfilePatentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ItemPatentCollectionResponse
}

type ListUserByIdProfilePatentsCompleteResult struct {
	Items []models.ItemPatentCollectionResponse
}

// ListUserByIdProfilePatents ...
func (c UserProfilePatentClient) ListUserByIdProfilePatents(ctx context.Context, id UserId) (result ListUserByIdProfilePatentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/profile/patents", id.ID()),
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
		Values *[]models.ItemPatentCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdProfilePatentsComplete retrieves all the results into a single object
func (c UserProfilePatentClient) ListUserByIdProfilePatentsComplete(ctx context.Context, id UserId) (ListUserByIdProfilePatentsCompleteResult, error) {
	return c.ListUserByIdProfilePatentsCompleteMatchingPredicate(ctx, id, models.ItemPatentCollectionResponseOperationPredicate{})
}

// ListUserByIdProfilePatentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserProfilePatentClient) ListUserByIdProfilePatentsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.ItemPatentCollectionResponseOperationPredicate) (result ListUserByIdProfilePatentsCompleteResult, err error) {
	items := make([]models.ItemPatentCollectionResponse, 0)

	resp, err := c.ListUserByIdProfilePatents(ctx, id)
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

	result = ListUserByIdProfilePatentsCompleteResult{
		Items: items,
	}
	return
}
