package userprofileanniversary

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

type ListUserByIdProfileAnniversariesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PersonAnnualEventCollectionResponse
}

type ListUserByIdProfileAnniversariesCompleteResult struct {
	Items []models.PersonAnnualEventCollectionResponse
}

// ListUserByIdProfileAnniversaries ...
func (c UserProfileAnniversaryClient) ListUserByIdProfileAnniversaries(ctx context.Context, id UserId) (result ListUserByIdProfileAnniversariesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/profile/anniversaries", id.ID()),
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
		Values *[]models.PersonAnnualEventCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdProfileAnniversariesComplete retrieves all the results into a single object
func (c UserProfileAnniversaryClient) ListUserByIdProfileAnniversariesComplete(ctx context.Context, id UserId) (ListUserByIdProfileAnniversariesCompleteResult, error) {
	return c.ListUserByIdProfileAnniversariesCompleteMatchingPredicate(ctx, id, models.PersonAnnualEventCollectionResponseOperationPredicate{})
}

// ListUserByIdProfileAnniversariesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserProfileAnniversaryClient) ListUserByIdProfileAnniversariesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.PersonAnnualEventCollectionResponseOperationPredicate) (result ListUserByIdProfileAnniversariesCompleteResult, err error) {
	items := make([]models.PersonAnnualEventCollectionResponse, 0)

	resp, err := c.ListUserByIdProfileAnniversaries(ctx, id)
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

	result = ListUserByIdProfileAnniversariesCompleteResult{
		Items: items,
	}
	return
}
