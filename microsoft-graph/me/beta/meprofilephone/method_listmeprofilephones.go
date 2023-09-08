package meprofilephone

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

type ListMeProfilePhonesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ItemPhoneCollectionResponse
}

type ListMeProfilePhonesCompleteResult struct {
	Items []models.ItemPhoneCollectionResponse
}

// ListMeProfilePhones ...
func (c MeProfilePhoneClient) ListMeProfilePhones(ctx context.Context) (result ListMeProfilePhonesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/profile/phones",
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

// ListMeProfilePhonesComplete retrieves all the results into a single object
func (c MeProfilePhoneClient) ListMeProfilePhonesComplete(ctx context.Context) (ListMeProfilePhonesCompleteResult, error) {
	return c.ListMeProfilePhonesCompleteMatchingPredicate(ctx, models.ItemPhoneCollectionResponseOperationPredicate{})
}

// ListMeProfilePhonesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeProfilePhoneClient) ListMeProfilePhonesCompleteMatchingPredicate(ctx context.Context, predicate models.ItemPhoneCollectionResponseOperationPredicate) (result ListMeProfilePhonesCompleteResult, err error) {
	items := make([]models.ItemPhoneCollectionResponse, 0)

	resp, err := c.ListMeProfilePhones(ctx)
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

	result = ListMeProfilePhonesCompleteResult{
		Items: items,
	}
	return
}
