package b2xuserflowuserflowidentityprovider

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

type ListB2xUserFlowUserFlowIdentityProvidersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityProviderBase
}

type ListB2xUserFlowUserFlowIdentityProvidersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityProviderBase
}

type ListB2xUserFlowUserFlowIdentityProvidersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListB2xUserFlowUserFlowIdentityProvidersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListB2xUserFlowUserFlowIdentityProviders ...
func (c B2xUserFlowUserFlowIdentityProviderClient) ListB2xUserFlowUserFlowIdentityProviders(ctx context.Context, id IdentityB2xUserFlowId) (result ListB2xUserFlowUserFlowIdentityProvidersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListB2xUserFlowUserFlowIdentityProvidersCustomPager{},
		Path:       fmt.Sprintf("%s/userFlowIdentityProviders", id.ID()),
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

// ListB2xUserFlowUserFlowIdentityProvidersComplete retrieves all the results into a single object
func (c B2xUserFlowUserFlowIdentityProviderClient) ListB2xUserFlowUserFlowIdentityProvidersComplete(ctx context.Context, id IdentityB2xUserFlowId) (ListB2xUserFlowUserFlowIdentityProvidersCompleteResult, error) {
	return c.ListB2xUserFlowUserFlowIdentityProvidersCompleteMatchingPredicate(ctx, id, IdentityProviderBaseOperationPredicate{})
}

// ListB2xUserFlowUserFlowIdentityProvidersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c B2xUserFlowUserFlowIdentityProviderClient) ListB2xUserFlowUserFlowIdentityProvidersCompleteMatchingPredicate(ctx context.Context, id IdentityB2xUserFlowId, predicate IdentityProviderBaseOperationPredicate) (result ListB2xUserFlowUserFlowIdentityProvidersCompleteResult, err error) {
	items := make([]stable.IdentityProviderBase, 0)

	resp, err := c.ListB2xUserFlowUserFlowIdentityProviders(ctx, id)
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

	result = ListB2xUserFlowUserFlowIdentityProvidersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
