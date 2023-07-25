package domains

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOwnershipIdentifiersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DomainOwnershipIdentifier
}

type ListOwnershipIdentifiersCompleteResult struct {
	Items []DomainOwnershipIdentifier
}

// ListOwnershipIdentifiers ...
func (c DomainsClient) ListOwnershipIdentifiers(ctx context.Context, id DomainId) (result ListOwnershipIdentifiersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/domainOwnershipIdentifiers", id.ID()),
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
		Values *[]DomainOwnershipIdentifier `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOwnershipIdentifiersComplete retrieves all the results into a single object
func (c DomainsClient) ListOwnershipIdentifiersComplete(ctx context.Context, id DomainId) (ListOwnershipIdentifiersCompleteResult, error) {
	return c.ListOwnershipIdentifiersCompleteMatchingPredicate(ctx, id, DomainOwnershipIdentifierOperationPredicate{})
}

// ListOwnershipIdentifiersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DomainsClient) ListOwnershipIdentifiersCompleteMatchingPredicate(ctx context.Context, id DomainId, predicate DomainOwnershipIdentifierOperationPredicate) (result ListOwnershipIdentifiersCompleteResult, err error) {
	items := make([]DomainOwnershipIdentifier, 0)

	resp, err := c.ListOwnershipIdentifiers(ctx, id)
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

	result = ListOwnershipIdentifiersCompleteResult{
		Items: items,
	}
	return
}
