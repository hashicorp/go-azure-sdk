package meprofileanniversary

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

type ListMeProfileAnniversariesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PersonAnnualEventCollectionResponse
}

type ListMeProfileAnniversariesCompleteResult struct {
	Items []models.PersonAnnualEventCollectionResponse
}

// ListMeProfileAnniversaries ...
func (c MeProfileAnniversaryClient) ListMeProfileAnniversaries(ctx context.Context) (result ListMeProfileAnniversariesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/profile/anniversaries",
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

// ListMeProfileAnniversariesComplete retrieves all the results into a single object
func (c MeProfileAnniversaryClient) ListMeProfileAnniversariesComplete(ctx context.Context) (ListMeProfileAnniversariesCompleteResult, error) {
	return c.ListMeProfileAnniversariesCompleteMatchingPredicate(ctx, models.PersonAnnualEventCollectionResponseOperationPredicate{})
}

// ListMeProfileAnniversariesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeProfileAnniversaryClient) ListMeProfileAnniversariesCompleteMatchingPredicate(ctx context.Context, predicate models.PersonAnnualEventCollectionResponseOperationPredicate) (result ListMeProfileAnniversariesCompleteResult, err error) {
	items := make([]models.PersonAnnualEventCollectionResponse, 0)

	resp, err := c.ListMeProfileAnniversaries(ctx)
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

	result = ListMeProfileAnniversariesCompleteResult{
		Items: items,
	}
	return
}
