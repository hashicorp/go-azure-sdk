package userprofilephone

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

type ListUserByIdProfilePhonesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ItemPhoneCollectionResponse
}

type ListUserByIdProfilePhonesCompleteResult struct {
	Items []models.ItemPhoneCollectionResponse
}

// ListUserByIdProfilePhones ...
func (c UserProfilePhoneClient) ListUserByIdProfilePhones(ctx context.Context, id UserId) (result ListUserByIdProfilePhonesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/profile/phones", id.ID()),
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
		Values *[]models.ItemPhoneCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdProfilePhonesComplete retrieves all the results into a single object
func (c UserProfilePhoneClient) ListUserByIdProfilePhonesComplete(ctx context.Context, id UserId) (ListUserByIdProfilePhonesCompleteResult, error) {
	return c.ListUserByIdProfilePhonesCompleteMatchingPredicate(ctx, id, models.ItemPhoneCollectionResponseOperationPredicate{})
}

// ListUserByIdProfilePhonesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserProfilePhoneClient) ListUserByIdProfilePhonesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.ItemPhoneCollectionResponseOperationPredicate) (result ListUserByIdProfilePhonesCompleteResult, err error) {
	items := make([]models.ItemPhoneCollectionResponse, 0)

	resp, err := c.ListUserByIdProfilePhones(ctx, id)
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

	result = ListUserByIdProfilePhonesCompleteResult{
		Items: items,
	}
	return
}
