package grouppolicycategory

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

type ListGroupPolicyCategoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicyCategory
}

type ListGroupPolicyCategoriesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicyCategory
}

type ListGroupPolicyCategoriesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupPolicyCategoriesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupPolicyCategories ...
func (c GroupPolicyCategoryClient) ListGroupPolicyCategories(ctx context.Context) (result ListGroupPolicyCategoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListGroupPolicyCategoriesCustomPager{},
		Path:       "/deviceManagement/groupPolicyCategories",
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
		Values *[]beta.GroupPolicyCategory `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupPolicyCategoriesComplete retrieves all the results into a single object
func (c GroupPolicyCategoryClient) ListGroupPolicyCategoriesComplete(ctx context.Context) (ListGroupPolicyCategoriesCompleteResult, error) {
	return c.ListGroupPolicyCategoriesCompleteMatchingPredicate(ctx, GroupPolicyCategoryOperationPredicate{})
}

// ListGroupPolicyCategoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyCategoryClient) ListGroupPolicyCategoriesCompleteMatchingPredicate(ctx context.Context, predicate GroupPolicyCategoryOperationPredicate) (result ListGroupPolicyCategoriesCompleteResult, err error) {
	items := make([]beta.GroupPolicyCategory, 0)

	resp, err := c.ListGroupPolicyCategories(ctx)
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

	result = ListGroupPolicyCategoriesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
