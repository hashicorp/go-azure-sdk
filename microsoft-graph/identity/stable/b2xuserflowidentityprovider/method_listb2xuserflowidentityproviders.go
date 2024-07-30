package b2xuserflowidentityprovider

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

type ListB2xUserFlowIdentityProvidersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityProvider
}

type ListB2xUserFlowIdentityProvidersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityProvider
}

type ListB2xUserFlowIdentityProvidersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListB2xUserFlowIdentityProvidersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListB2xUserFlowIdentityProviders ...
func (c B2xUserFlowIdentityProviderClient) ListB2xUserFlowIdentityProviders(ctx context.Context, id IdentityB2xUserFlowId) (result ListB2xUserFlowIdentityProvidersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListB2xUserFlowIdentityProvidersCustomPager{},
		Path:       fmt.Sprintf("%s/identityProviders", id.ID()),
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
		Values *[]stable.IdentityProvider `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListB2xUserFlowIdentityProvidersComplete retrieves all the results into a single object
func (c B2xUserFlowIdentityProviderClient) ListB2xUserFlowIdentityProvidersComplete(ctx context.Context, id IdentityB2xUserFlowId) (ListB2xUserFlowIdentityProvidersCompleteResult, error) {
	return c.ListB2xUserFlowIdentityProvidersCompleteMatchingPredicate(ctx, id, IdentityProviderOperationPredicate{})
}

// ListB2xUserFlowIdentityProvidersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c B2xUserFlowIdentityProviderClient) ListB2xUserFlowIdentityProvidersCompleteMatchingPredicate(ctx context.Context, id IdentityB2xUserFlowId, predicate IdentityProviderOperationPredicate) (result ListB2xUserFlowIdentityProvidersCompleteResult, err error) {
	items := make([]stable.IdentityProvider, 0)

	resp, err := c.ListB2xUserFlowIdentityProviders(ctx, id)
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

	result = ListB2xUserFlowIdentityProvidersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
