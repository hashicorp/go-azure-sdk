package userprofileaddress

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

type ListUserByIdProfileAddresOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ItemAddressCollectionResponse
}

type ListUserByIdProfileAddresCompleteResult struct {
	Items []models.ItemAddressCollectionResponse
}

// ListUserByIdProfileAddres ...
func (c UserProfileAddressClient) ListUserByIdProfileAddres(ctx context.Context, id UserId) (result ListUserByIdProfileAddresOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/profile/addresses", id.ID()),
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
		Values *[]models.ItemAddressCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdProfileAddresComplete retrieves all the results into a single object
func (c UserProfileAddressClient) ListUserByIdProfileAddresComplete(ctx context.Context, id UserId) (ListUserByIdProfileAddresCompleteResult, error) {
	return c.ListUserByIdProfileAddresCompleteMatchingPredicate(ctx, id, models.ItemAddressCollectionResponseOperationPredicate{})
}

// ListUserByIdProfileAddresCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserProfileAddressClient) ListUserByIdProfileAddresCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.ItemAddressCollectionResponseOperationPredicate) (result ListUserByIdProfileAddresCompleteResult, err error) {
	items := make([]models.ItemAddressCollectionResponse, 0)

	resp, err := c.ListUserByIdProfileAddres(ctx, id)
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

	result = ListUserByIdProfileAddresCompleteResult{
		Items: items,
	}
	return
}
