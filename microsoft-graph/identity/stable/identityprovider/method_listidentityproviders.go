package identityprovider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListIdentityProvidersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityProviderBase
}

type ListIdentityProvidersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityProviderBase
}

type ListIdentityProvidersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListIdentityProvidersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListIdentityProviders ...
func (c IdentityProviderClient) ListIdentityProviders(ctx context.Context) (result ListIdentityProvidersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListIdentityProvidersCustomPager{},
		Path:       "/identity/identityProviders",
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
		Values *[]stable.IdentityProviderBase `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIdentityProvidersComplete retrieves all the results into a single object
func (c IdentityProviderClient) ListIdentityProvidersComplete(ctx context.Context) (ListIdentityProvidersCompleteResult, error) {
	return c.ListIdentityProvidersCompleteMatchingPredicate(ctx, IdentityProviderBaseOperationPredicate{})
}

// ListIdentityProvidersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IdentityProviderClient) ListIdentityProvidersCompleteMatchingPredicate(ctx context.Context, predicate IdentityProviderBaseOperationPredicate) (result ListIdentityProvidersCompleteResult, err error) {
	items := make([]stable.IdentityProviderBase, 0)

	resp, err := c.ListIdentityProviders(ctx)
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

	result = ListIdentityProvidersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
