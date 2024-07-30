package grouplifecyclepolicy

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

type ListGroupLifecyclePoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupLifecyclePolicy
}

type ListGroupLifecyclePoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupLifecyclePolicy
}

type ListGroupLifecyclePoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupLifecyclePoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupLifecyclePolicies ...
func (c GroupLifecyclePolicyClient) ListGroupLifecyclePolicies(ctx context.Context, id GroupId) (result ListGroupLifecyclePoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListGroupLifecyclePoliciesCustomPager{},
		Path:       fmt.Sprintf("%s/groupLifecyclePolicies", id.ID()),
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
		Values *[]beta.GroupLifecyclePolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupLifecyclePoliciesComplete retrieves all the results into a single object
func (c GroupLifecyclePolicyClient) ListGroupLifecyclePoliciesComplete(ctx context.Context, id GroupId) (ListGroupLifecyclePoliciesCompleteResult, error) {
	return c.ListGroupLifecyclePoliciesCompleteMatchingPredicate(ctx, id, GroupLifecyclePolicyOperationPredicate{})
}

// ListGroupLifecyclePoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupLifecyclePolicyClient) ListGroupLifecyclePoliciesCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate GroupLifecyclePolicyOperationPredicate) (result ListGroupLifecyclePoliciesCompleteResult, err error) {
	items := make([]beta.GroupLifecyclePolicy, 0)

	resp, err := c.ListGroupLifecyclePolicies(ctx, id)
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

	result = ListGroupLifecyclePoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
