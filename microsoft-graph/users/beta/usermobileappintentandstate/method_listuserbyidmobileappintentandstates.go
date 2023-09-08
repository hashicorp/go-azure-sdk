package usermobileappintentandstate

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

type ListUserByIdMobileAppIntentAndStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.MobileAppIntentAndStateCollectionResponse
}

type ListUserByIdMobileAppIntentAndStatesCompleteResult struct {
	Items []models.MobileAppIntentAndStateCollectionResponse
}

// ListUserByIdMobileAppIntentAndStates ...
func (c UserMobileAppIntentAndStateClient) ListUserByIdMobileAppIntentAndStates(ctx context.Context, id UserId) (result ListUserByIdMobileAppIntentAndStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/mobileAppIntentAndStates", id.ID()),
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

// ListUserByIdMobileAppIntentAndStatesComplete retrieves all the results into a single object
func (c UserMobileAppIntentAndStateClient) ListUserByIdMobileAppIntentAndStatesComplete(ctx context.Context, id UserId) (ListUserByIdMobileAppIntentAndStatesCompleteResult, error) {
	return c.ListUserByIdMobileAppIntentAndStatesCompleteMatchingPredicate(ctx, id, models.MobileAppIntentAndStateCollectionResponseOperationPredicate{})
}

// ListUserByIdMobileAppIntentAndStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserMobileAppIntentAndStateClient) ListUserByIdMobileAppIntentAndStatesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.MobileAppIntentAndStateCollectionResponseOperationPredicate) (result ListUserByIdMobileAppIntentAndStatesCompleteResult, err error) {
	items := make([]models.MobileAppIntentAndStateCollectionResponse, 0)

	resp, err := c.ListUserByIdMobileAppIntentAndStates(ctx, id)
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

	result = ListUserByIdMobileAppIntentAndStatesCompleteResult{
		Items: items,
	}
	return
}
