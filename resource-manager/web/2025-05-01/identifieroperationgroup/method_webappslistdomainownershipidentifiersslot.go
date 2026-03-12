package identifieroperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListDomainOwnershipIdentifiersSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Identifier
}

type WebAppsListDomainOwnershipIdentifiersSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Identifier
}

type WebAppsListDomainOwnershipIdentifiersSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListDomainOwnershipIdentifiersSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListDomainOwnershipIdentifiersSlot ...
func (c IdentifierOperationGroupClient) WebAppsListDomainOwnershipIdentifiersSlot(ctx context.Context, id SlotId) (result WebAppsListDomainOwnershipIdentifiersSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListDomainOwnershipIdentifiersSlotCustomPager{},
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
		Values *[]Identifier `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListDomainOwnershipIdentifiersSlotComplete retrieves all the results into a single object
func (c IdentifierOperationGroupClient) WebAppsListDomainOwnershipIdentifiersSlotComplete(ctx context.Context, id SlotId) (WebAppsListDomainOwnershipIdentifiersSlotCompleteResult, error) {
	return c.WebAppsListDomainOwnershipIdentifiersSlotCompleteMatchingPredicate(ctx, id, IdentifierOperationPredicate{})
}

// WebAppsListDomainOwnershipIdentifiersSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IdentifierOperationGroupClient) WebAppsListDomainOwnershipIdentifiersSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate IdentifierOperationPredicate) (result WebAppsListDomainOwnershipIdentifiersSlotCompleteResult, err error) {
	items := make([]Identifier, 0)

	resp, err := c.WebAppsListDomainOwnershipIdentifiersSlot(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = WebAppsListDomainOwnershipIdentifiersSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
