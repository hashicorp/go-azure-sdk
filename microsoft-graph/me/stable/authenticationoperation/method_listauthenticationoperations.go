package authenticationoperation

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

type ListAuthenticationOperationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.LongRunningOperation
}

type ListAuthenticationOperationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.LongRunningOperation
}

type ListAuthenticationOperationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationOperationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationOperations ...
func (c AuthenticationOperationClient) ListAuthenticationOperations(ctx context.Context) (result ListAuthenticationOperationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAuthenticationOperationsCustomPager{},
		Path:       "/me/authentication/operations",
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
		Values *[]stable.LongRunningOperation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthenticationOperationsComplete retrieves all the results into a single object
func (c AuthenticationOperationClient) ListAuthenticationOperationsComplete(ctx context.Context) (ListAuthenticationOperationsCompleteResult, error) {
	return c.ListAuthenticationOperationsCompleteMatchingPredicate(ctx, LongRunningOperationOperationPredicate{})
}

// ListAuthenticationOperationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationOperationClient) ListAuthenticationOperationsCompleteMatchingPredicate(ctx context.Context, predicate LongRunningOperationOperationPredicate) (result ListAuthenticationOperationsCompleteResult, err error) {
	items := make([]stable.LongRunningOperation, 0)

	resp, err := c.ListAuthenticationOperations(ctx)
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

	result = ListAuthenticationOperationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
