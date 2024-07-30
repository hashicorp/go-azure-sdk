package serviceprincipalcreationpolicyexclude

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

type ListServicePrincipalCreationPolicyExcludesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ServicePrincipalCreationConditionSet
}

type ListServicePrincipalCreationPolicyExcludesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ServicePrincipalCreationConditionSet
}

type ListServicePrincipalCreationPolicyExcludesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListServicePrincipalCreationPolicyExcludesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListServicePrincipalCreationPolicyExcludes ...
func (c ServicePrincipalCreationPolicyExcludeClient) ListServicePrincipalCreationPolicyExcludes(ctx context.Context, id PolicyServicePrincipalCreationPolicyId) (result ListServicePrincipalCreationPolicyExcludesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListServicePrincipalCreationPolicyExcludesCustomPager{},
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
		Values *[]beta.ServicePrincipalCreationConditionSet `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListServicePrincipalCreationPolicyExcludesComplete retrieves all the results into a single object
func (c ServicePrincipalCreationPolicyExcludeClient) ListServicePrincipalCreationPolicyExcludesComplete(ctx context.Context, id PolicyServicePrincipalCreationPolicyId) (ListServicePrincipalCreationPolicyExcludesCompleteResult, error) {
	return c.ListServicePrincipalCreationPolicyExcludesCompleteMatchingPredicate(ctx, id, ServicePrincipalCreationConditionSetOperationPredicate{})
}

// ListServicePrincipalCreationPolicyExcludesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServicePrincipalCreationPolicyExcludeClient) ListServicePrincipalCreationPolicyExcludesCompleteMatchingPredicate(ctx context.Context, id PolicyServicePrincipalCreationPolicyId, predicate ServicePrincipalCreationConditionSetOperationPredicate) (result ListServicePrincipalCreationPolicyExcludesCompleteResult, err error) {
	items := make([]beta.ServicePrincipalCreationConditionSet, 0)

	resp, err := c.ListServicePrincipalCreationPolicyExcludes(ctx, id)
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

	result = ListServicePrincipalCreationPolicyExcludesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
