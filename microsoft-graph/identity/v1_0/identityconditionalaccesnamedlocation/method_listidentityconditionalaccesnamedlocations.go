package identityconditionalaccesnamedlocation

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

type ListIdentityConditionalAccesNamedLocationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.NamedLocationCollectionResponse
}

type ListIdentityConditionalAccesNamedLocationsCompleteResult struct {
	Items []models.NamedLocationCollectionResponse
}

// ListIdentityConditionalAccesNamedLocations ...
func (c IdentityConditionalAccesNamedLocationClient) ListIdentityConditionalAccesNamedLocations(ctx context.Context) (result ListIdentityConditionalAccesNamedLocationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/identity/conditionalAccess/namedLocations",
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
		Values *[]models.NamedLocationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIdentityConditionalAccesNamedLocationsComplete retrieves all the results into a single object
func (c IdentityConditionalAccesNamedLocationClient) ListIdentityConditionalAccesNamedLocationsComplete(ctx context.Context) (ListIdentityConditionalAccesNamedLocationsCompleteResult, error) {
	return c.ListIdentityConditionalAccesNamedLocationsCompleteMatchingPredicate(ctx, models.NamedLocationCollectionResponseOperationPredicate{})
}

// ListIdentityConditionalAccesNamedLocationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IdentityConditionalAccesNamedLocationClient) ListIdentityConditionalAccesNamedLocationsCompleteMatchingPredicate(ctx context.Context, predicate models.NamedLocationCollectionResponseOperationPredicate) (result ListIdentityConditionalAccesNamedLocationsCompleteResult, err error) {
	items := make([]models.NamedLocationCollectionResponse, 0)

	resp, err := c.ListIdentityConditionalAccesNamedLocations(ctx)
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

	result = ListIdentityConditionalAccesNamedLocationsCompleteResult{
		Items: items,
	}
	return
}
