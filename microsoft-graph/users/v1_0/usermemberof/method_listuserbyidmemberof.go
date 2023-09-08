package usermemberof

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserByIdMemberOfOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListUserByIdMemberOfCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListUserByIdMemberOf ...
func (c UserMemberOfClient) ListUserByIdMemberOf(ctx context.Context, id UserId) (result ListUserByIdMemberOfOperationResponse, err error) {
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

// ListUserByIdMemberOfComplete retrieves all the results into a single object
func (c UserMemberOfClient) ListUserByIdMemberOfComplete(ctx context.Context, id UserId) (ListUserByIdMemberOfCompleteResult, error) {
	return c.ListUserByIdMemberOfCompleteMatchingPredicate(ctx, id, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListUserByIdMemberOfCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserMemberOfClient) ListUserByIdMemberOfCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListUserByIdMemberOfCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListUserByIdMemberOf(ctx, id)
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

	result = ListUserByIdMemberOfCompleteResult{
		Items: items,
	}
	return
}
