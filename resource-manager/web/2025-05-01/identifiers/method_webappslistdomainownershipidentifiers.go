package identifiers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListDomainOwnershipIdentifiersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Identifier
}

type WebAppsListDomainOwnershipIdentifiersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Identifier
}

type WebAppsListDomainOwnershipIdentifiersCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListDomainOwnershipIdentifiersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListDomainOwnershipIdentifiers ...
func (c IdentifiersClient) WebAppsListDomainOwnershipIdentifiers(ctx context.Context, id commonids.AppServiceId) (result WebAppsListDomainOwnershipIdentifiersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListDomainOwnershipIdentifiersCustomPager{},
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

// WebAppsListDomainOwnershipIdentifiersComplete retrieves all the results into a single object
func (c IdentifiersClient) WebAppsListDomainOwnershipIdentifiersComplete(ctx context.Context, id commonids.AppServiceId) (WebAppsListDomainOwnershipIdentifiersCompleteResult, error) {
	return c.WebAppsListDomainOwnershipIdentifiersCompleteMatchingPredicate(ctx, id, IdentifierOperationPredicate{})
}

// WebAppsListDomainOwnershipIdentifiersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IdentifiersClient) WebAppsListDomainOwnershipIdentifiersCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, predicate IdentifierOperationPredicate) (result WebAppsListDomainOwnershipIdentifiersCompleteResult, err error) {
	items := make([]Identifier, 0)

	resp, err := c.WebAppsListDomainOwnershipIdentifiers(ctx, id)
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

	result = WebAppsListDomainOwnershipIdentifiersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
