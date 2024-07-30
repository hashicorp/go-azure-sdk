package grouppolicycategorychildren

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

type ListGroupPolicyCategoryChildrensOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicyCategory
}

type ListGroupPolicyCategoryChildrensCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicyCategory
}

type ListGroupPolicyCategoryChildrensCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupPolicyCategoryChildrensCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupPolicyCategoryChildrens ...
func (c GroupPolicyCategoryChildrenClient) ListGroupPolicyCategoryChildrens(ctx context.Context, id DeviceManagementGroupPolicyCategoryId) (result ListGroupPolicyCategoryChildrensOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListGroupPolicyCategoryChildrensCustomPager{},
		Path:       fmt.Sprintf("%s/children", id.ID()),
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

// ListGroupPolicyCategoryChildrensComplete retrieves all the results into a single object
func (c GroupPolicyCategoryChildrenClient) ListGroupPolicyCategoryChildrensComplete(ctx context.Context, id DeviceManagementGroupPolicyCategoryId) (ListGroupPolicyCategoryChildrensCompleteResult, error) {
	return c.ListGroupPolicyCategoryChildrensCompleteMatchingPredicate(ctx, id, GroupPolicyCategoryOperationPredicate{})
}

// ListGroupPolicyCategoryChildrensCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyCategoryChildrenClient) ListGroupPolicyCategoryChildrensCompleteMatchingPredicate(ctx context.Context, id DeviceManagementGroupPolicyCategoryId, predicate GroupPolicyCategoryOperationPredicate) (result ListGroupPolicyCategoryChildrensCompleteResult, err error) {
	items := make([]beta.GroupPolicyCategory, 0)

	resp, err := c.ListGroupPolicyCategoryChildrens(ctx, id)
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

	result = ListGroupPolicyCategoryChildrensCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
