package permissiongrantpolicyinclude

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

type ListPermissionGrantPolicyIncludesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PermissionGrantConditionSet
}

type ListPermissionGrantPolicyIncludesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PermissionGrantConditionSet
}

type ListPermissionGrantPolicyIncludesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPermissionGrantPolicyIncludesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPermissionGrantPolicyIncludes ...
func (c PermissionGrantPolicyIncludeClient) ListPermissionGrantPolicyIncludes(ctx context.Context, id PolicyPermissionGrantPolicyId) (result ListPermissionGrantPolicyIncludesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPermissionGrantPolicyIncludesCustomPager{},
		Path:       fmt.Sprintf("%s/includes", id.ID()),
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

// ListPermissionGrantPolicyIncludesComplete retrieves all the results into a single object
func (c PermissionGrantPolicyIncludeClient) ListPermissionGrantPolicyIncludesComplete(ctx context.Context, id PolicyPermissionGrantPolicyId) (ListPermissionGrantPolicyIncludesCompleteResult, error) {
	return c.ListPermissionGrantPolicyIncludesCompleteMatchingPredicate(ctx, id, PermissionGrantConditionSetOperationPredicate{})
}

// ListPermissionGrantPolicyIncludesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PermissionGrantPolicyIncludeClient) ListPermissionGrantPolicyIncludesCompleteMatchingPredicate(ctx context.Context, id PolicyPermissionGrantPolicyId, predicate PermissionGrantConditionSetOperationPredicate) (result ListPermissionGrantPolicyIncludesCompleteResult, err error) {
	items := make([]beta.PermissionGrantConditionSet, 0)

	resp, err := c.ListPermissionGrantPolicyIncludes(ctx, id)
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

	result = ListPermissionGrantPolicyIncludesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
