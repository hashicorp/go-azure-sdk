package appmanagementpolicy

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

type ListAppManagementPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AppManagementPolicy
}

type ListAppManagementPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AppManagementPolicy
}

type ListAppManagementPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAppManagementPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppManagementPolicies ...
func (c AppManagementPolicyClient) ListAppManagementPolicies(ctx context.Context, id ApplicationId) (result ListAppManagementPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAppManagementPoliciesCustomPager{},
		Path:       fmt.Sprintf("%s/appManagementPolicies", id.ID()),
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
		Values *[]stable.AppManagementPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAppManagementPoliciesComplete retrieves all the results into a single object
func (c AppManagementPolicyClient) ListAppManagementPoliciesComplete(ctx context.Context, id ApplicationId) (ListAppManagementPoliciesCompleteResult, error) {
	return c.ListAppManagementPoliciesCompleteMatchingPredicate(ctx, id, AppManagementPolicyOperationPredicate{})
}

// ListAppManagementPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppManagementPolicyClient) ListAppManagementPoliciesCompleteMatchingPredicate(ctx context.Context, id ApplicationId, predicate AppManagementPolicyOperationPredicate) (result ListAppManagementPoliciesCompleteResult, err error) {
	items := make([]stable.AppManagementPolicy, 0)

	resp, err := c.ListAppManagementPolicies(ctx, id)
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

	result = ListAppManagementPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
