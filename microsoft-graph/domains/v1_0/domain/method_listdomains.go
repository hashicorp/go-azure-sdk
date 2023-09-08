package domain

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

type ListDomainsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DomainCollectionResponse
}

type ListDomainsCompleteResult struct {
	Items []models.DomainCollectionResponse
}

// ListDomains ...
func (c DomainClient) ListDomains(ctx context.Context) (result ListDomainsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/domains",
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
		Values *[]models.DomainCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDomainsComplete retrieves all the results into a single object
func (c DomainClient) ListDomainsComplete(ctx context.Context) (ListDomainsCompleteResult, error) {
	return c.ListDomainsCompleteMatchingPredicate(ctx, models.DomainCollectionResponseOperationPredicate{})
}

// ListDomainsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DomainClient) ListDomainsCompleteMatchingPredicate(ctx context.Context, predicate models.DomainCollectionResponseOperationPredicate) (result ListDomainsCompleteResult, err error) {
	items := make([]models.DomainCollectionResponse, 0)

	resp, err := c.ListDomains(ctx)
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

	result = ListDomainsCompleteResult{
		Items: items,
	}
	return
}
