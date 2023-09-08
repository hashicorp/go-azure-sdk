package meprofileinterest

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

type ListMeProfileInterestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PersonInterestCollectionResponse
}

type ListMeProfileInterestsCompleteResult struct {
	Items []models.PersonInterestCollectionResponse
}

// ListMeProfileInterests ...
func (c MeProfileInterestClient) ListMeProfileInterests(ctx context.Context) (result ListMeProfileInterestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/profile/interests",
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
		Values *[]models.PersonInterestCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeProfileInterestsComplete retrieves all the results into a single object
func (c MeProfileInterestClient) ListMeProfileInterestsComplete(ctx context.Context) (ListMeProfileInterestsCompleteResult, error) {
	return c.ListMeProfileInterestsCompleteMatchingPredicate(ctx, models.PersonInterestCollectionResponseOperationPredicate{})
}

// ListMeProfileInterestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeProfileInterestClient) ListMeProfileInterestsCompleteMatchingPredicate(ctx context.Context, predicate models.PersonInterestCollectionResponseOperationPredicate) (result ListMeProfileInterestsCompleteResult, err error) {
	items := make([]models.PersonInterestCollectionResponse, 0)

	resp, err := c.ListMeProfileInterests(ctx)
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

	result = ListMeProfileInterestsCompleteResult{
		Items: items,
	}
	return
}
