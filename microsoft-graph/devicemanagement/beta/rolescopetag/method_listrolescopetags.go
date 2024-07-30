package rolescopetag

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

type ListRoleScopeTagsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.RoleScopeTag
}

type ListRoleScopeTagsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.RoleScopeTag
}

type ListRoleScopeTagsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRoleScopeTagsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRoleScopeTags ...
func (c RoleScopeTagClient) ListRoleScopeTags(ctx context.Context) (result ListRoleScopeTagsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListRoleScopeTagsCustomPager{},
		Path:       "/deviceManagement/roleScopeTags",
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
		Values *[]beta.RoleScopeTag `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleScopeTagsComplete retrieves all the results into a single object
func (c RoleScopeTagClient) ListRoleScopeTagsComplete(ctx context.Context) (ListRoleScopeTagsCompleteResult, error) {
	return c.ListRoleScopeTagsCompleteMatchingPredicate(ctx, RoleScopeTagOperationPredicate{})
}

// ListRoleScopeTagsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleScopeTagClient) ListRoleScopeTagsCompleteMatchingPredicate(ctx context.Context, predicate RoleScopeTagOperationPredicate) (result ListRoleScopeTagsCompleteResult, err error) {
	items := make([]beta.RoleScopeTag, 0)

	resp, err := c.ListRoleScopeTags(ctx)
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

	result = ListRoleScopeTagsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
