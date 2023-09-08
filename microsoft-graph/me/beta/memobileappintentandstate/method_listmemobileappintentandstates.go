package memobileappintentandstate

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

type ListMeMobileAppIntentAndStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.MobileAppIntentAndStateCollectionResponse
}

type ListMeMobileAppIntentAndStatesCompleteResult struct {
	Items []models.MobileAppIntentAndStateCollectionResponse
}

// ListMeMobileAppIntentAndStates ...
func (c MeMobileAppIntentAndStateClient) ListMeMobileAppIntentAndStates(ctx context.Context) (result ListMeMobileAppIntentAndStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/mobileAppIntentAndStates",
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
		Values *[]models.MobileAppIntentAndStateCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeMobileAppIntentAndStatesComplete retrieves all the results into a single object
func (c MeMobileAppIntentAndStateClient) ListMeMobileAppIntentAndStatesComplete(ctx context.Context) (ListMeMobileAppIntentAndStatesCompleteResult, error) {
	return c.ListMeMobileAppIntentAndStatesCompleteMatchingPredicate(ctx, models.MobileAppIntentAndStateCollectionResponseOperationPredicate{})
}

// ListMeMobileAppIntentAndStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeMobileAppIntentAndStateClient) ListMeMobileAppIntentAndStatesCompleteMatchingPredicate(ctx context.Context, predicate models.MobileAppIntentAndStateCollectionResponseOperationPredicate) (result ListMeMobileAppIntentAndStatesCompleteResult, err error) {
	items := make([]models.MobileAppIntentAndStateCollectionResponse, 0)

	resp, err := c.ListMeMobileAppIntentAndStates(ctx)
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

	result = ListMeMobileAppIntentAndStatesCompleteResult{
		Items: items,
	}
	return
}
