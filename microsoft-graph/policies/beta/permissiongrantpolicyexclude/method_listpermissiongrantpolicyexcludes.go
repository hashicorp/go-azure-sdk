package permissiongrantpolicyexclude

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

type ListPermissionGrantPolicyExcludesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PermissionGrantConditionSet
}

type ListPermissionGrantPolicyExcludesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PermissionGrantConditionSet
}

type ListPermissionGrantPolicyExcludesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPermissionGrantPolicyExcludesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPermissionGrantPolicyExcludes ...
func (c PermissionGrantPolicyExcludeClient) ListPermissionGrantPolicyExcludes(ctx context.Context, id PolicyPermissionGrantPolicyId) (result ListPermissionGrantPolicyExcludesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPermissionGrantPolicyExcludesCustomPager{},
		Path:       fmt.Sprintf("%s/excludes", id.ID()),
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
		Values *[]beta.PermissionGrantConditionSet `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPermissionGrantPolicyExcludesComplete retrieves all the results into a single object
func (c PermissionGrantPolicyExcludeClient) ListPermissionGrantPolicyExcludesComplete(ctx context.Context, id PolicyPermissionGrantPolicyId) (ListPermissionGrantPolicyExcludesCompleteResult, error) {
	return c.ListPermissionGrantPolicyExcludesCompleteMatchingPredicate(ctx, id, PermissionGrantConditionSetOperationPredicate{})
}

// ListPermissionGrantPolicyExcludesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PermissionGrantPolicyExcludeClient) ListPermissionGrantPolicyExcludesCompleteMatchingPredicate(ctx context.Context, id PolicyPermissionGrantPolicyId, predicate PermissionGrantConditionSetOperationPredicate) (result ListPermissionGrantPolicyExcludesCompleteResult, err error) {
	items := make([]beta.PermissionGrantConditionSet, 0)

	resp, err := c.ListPermissionGrantPolicyExcludes(ctx, id)
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

	result = ListPermissionGrantPolicyExcludesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
