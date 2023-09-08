package meprofileaddress

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

type ListMeProfileAddresOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ItemAddressCollectionResponse
}

type ListMeProfileAddresCompleteResult struct {
	Items []models.ItemAddressCollectionResponse
}

// ListMeProfileAddres ...
func (c MeProfileAddressClient) ListMeProfileAddres(ctx context.Context) (result ListMeProfileAddresOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/profile/addresses",
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

// ListMeProfileAddresComplete retrieves all the results into a single object
func (c MeProfileAddressClient) ListMeProfileAddresComplete(ctx context.Context) (ListMeProfileAddresCompleteResult, error) {
	return c.ListMeProfileAddresCompleteMatchingPredicate(ctx, models.ItemAddressCollectionResponseOperationPredicate{})
}

// ListMeProfileAddresCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeProfileAddressClient) ListMeProfileAddresCompleteMatchingPredicate(ctx context.Context, predicate models.ItemAddressCollectionResponseOperationPredicate) (result ListMeProfileAddresCompleteResult, err error) {
	items := make([]models.ItemAddressCollectionResponse, 0)

	resp, err := c.ListMeProfileAddres(ctx)
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

	result = ListMeProfileAddresCompleteResult{
		Items: items,
	}
	return
}
