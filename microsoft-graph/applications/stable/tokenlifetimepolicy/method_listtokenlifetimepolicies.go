package tokenlifetimepolicy

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

type ListTokenLifetimePoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.TokenLifetimePolicy
}

type ListTokenLifetimePoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.TokenLifetimePolicy
}

type ListTokenLifetimePoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTokenLifetimePoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTokenLifetimePolicies ...
func (c TokenLifetimePolicyClient) ListTokenLifetimePolicies(ctx context.Context, id ApplicationId) (result ListTokenLifetimePoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTokenLifetimePoliciesCustomPager{},
		Path:       fmt.Sprintf("%s/tokenLifetimePolicies", id.ID()),
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
		Values *[]stable.TokenLifetimePolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTokenLifetimePoliciesComplete retrieves all the results into a single object
func (c TokenLifetimePolicyClient) ListTokenLifetimePoliciesComplete(ctx context.Context, id ApplicationId) (ListTokenLifetimePoliciesCompleteResult, error) {
	return c.ListTokenLifetimePoliciesCompleteMatchingPredicate(ctx, id, TokenLifetimePolicyOperationPredicate{})
}

// ListTokenLifetimePoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TokenLifetimePolicyClient) ListTokenLifetimePoliciesCompleteMatchingPredicate(ctx context.Context, id ApplicationId, predicate TokenLifetimePolicyOperationPredicate) (result ListTokenLifetimePoliciesCompleteResult, err error) {
	items := make([]stable.TokenLifetimePolicy, 0)

	resp, err := c.ListTokenLifetimePolicies(ctx, id)
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

	result = ListTokenLifetimePoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
