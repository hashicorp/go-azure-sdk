package serviceprincipalcreationpolicyinclude

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

type ListServicePrincipalCreationPolicyIncludesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ServicePrincipalCreationConditionSet
}

type ListServicePrincipalCreationPolicyIncludesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ServicePrincipalCreationConditionSet
}

type ListServicePrincipalCreationPolicyIncludesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListServicePrincipalCreationPolicyIncludesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListServicePrincipalCreationPolicyIncludes ...
func (c ServicePrincipalCreationPolicyIncludeClient) ListServicePrincipalCreationPolicyIncludes(ctx context.Context, id PolicyServicePrincipalCreationPolicyId) (result ListServicePrincipalCreationPolicyIncludesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListServicePrincipalCreationPolicyIncludesCustomPager{},
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
		Values *[]beta.ServicePrincipalCreationConditionSet `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListServicePrincipalCreationPolicyIncludesComplete retrieves all the results into a single object
func (c ServicePrincipalCreationPolicyIncludeClient) ListServicePrincipalCreationPolicyIncludesComplete(ctx context.Context, id PolicyServicePrincipalCreationPolicyId) (ListServicePrincipalCreationPolicyIncludesCompleteResult, error) {
	return c.ListServicePrincipalCreationPolicyIncludesCompleteMatchingPredicate(ctx, id, ServicePrincipalCreationConditionSetOperationPredicate{})
}

// ListServicePrincipalCreationPolicyIncludesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServicePrincipalCreationPolicyIncludeClient) ListServicePrincipalCreationPolicyIncludesCompleteMatchingPredicate(ctx context.Context, id PolicyServicePrincipalCreationPolicyId, predicate ServicePrincipalCreationConditionSetOperationPredicate) (result ListServicePrincipalCreationPolicyIncludesCompleteResult, err error) {
	items := make([]beta.ServicePrincipalCreationConditionSet, 0)

	resp, err := c.ListServicePrincipalCreationPolicyIncludes(ctx, id)
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

	result = ListServicePrincipalCreationPolicyIncludesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
