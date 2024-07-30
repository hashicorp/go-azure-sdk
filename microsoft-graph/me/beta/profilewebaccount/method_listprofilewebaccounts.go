package profilewebaccount

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListProfileWebAccountsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.WebAccount
}

type ListProfileWebAccountsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.WebAccount
}

type ListProfileWebAccountsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfileWebAccountsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfileWebAccounts ...
func (c ProfileWebAccountClient) ListProfileWebAccounts(ctx context.Context) (result ListProfileWebAccountsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListProfileWebAccountsCustomPager{},
		Path:       "/me/profile/webAccounts",
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
		Values *[]beta.WebAccount `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfileWebAccountsComplete retrieves all the results into a single object
func (c ProfileWebAccountClient) ListProfileWebAccountsComplete(ctx context.Context) (ListProfileWebAccountsCompleteResult, error) {
	return c.ListProfileWebAccountsCompleteMatchingPredicate(ctx, WebAccountOperationPredicate{})
}

// ListProfileWebAccountsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfileWebAccountClient) ListProfileWebAccountsCompleteMatchingPredicate(ctx context.Context, predicate WebAccountOperationPredicate) (result ListProfileWebAccountsCompleteResult, err error) {
	items := make([]beta.WebAccount, 0)

	resp, err := c.ListProfileWebAccounts(ctx)
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

	result = ListProfileWebAccountsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
