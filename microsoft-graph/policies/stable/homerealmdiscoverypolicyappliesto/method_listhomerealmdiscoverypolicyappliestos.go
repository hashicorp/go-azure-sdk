package homerealmdiscoverypolicyappliesto

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

type ListHomeRealmDiscoveryPolicyAppliesTosOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type ListHomeRealmDiscoveryPolicyAppliesTosCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type ListHomeRealmDiscoveryPolicyAppliesTosCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListHomeRealmDiscoveryPolicyAppliesTosCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListHomeRealmDiscoveryPolicyAppliesTos ...
func (c HomeRealmDiscoveryPolicyAppliesToClient) ListHomeRealmDiscoveryPolicyAppliesTos(ctx context.Context, id PolicyHomeRealmDiscoveryPolicyId) (result ListHomeRealmDiscoveryPolicyAppliesTosOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListHomeRealmDiscoveryPolicyAppliesTosCustomPager{},
		Path:       fmt.Sprintf("%s/appliesTo", id.ID()),
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
		Values *[]stable.DirectoryObject `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListHomeRealmDiscoveryPolicyAppliesTosComplete retrieves all the results into a single object
func (c HomeRealmDiscoveryPolicyAppliesToClient) ListHomeRealmDiscoveryPolicyAppliesTosComplete(ctx context.Context, id PolicyHomeRealmDiscoveryPolicyId) (ListHomeRealmDiscoveryPolicyAppliesTosCompleteResult, error) {
	return c.ListHomeRealmDiscoveryPolicyAppliesTosCompleteMatchingPredicate(ctx, id, DirectoryObjectOperationPredicate{})
}

// ListHomeRealmDiscoveryPolicyAppliesTosCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HomeRealmDiscoveryPolicyAppliesToClient) ListHomeRealmDiscoveryPolicyAppliesTosCompleteMatchingPredicate(ctx context.Context, id PolicyHomeRealmDiscoveryPolicyId, predicate DirectoryObjectOperationPredicate) (result ListHomeRealmDiscoveryPolicyAppliesTosCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.ListHomeRealmDiscoveryPolicyAppliesTos(ctx, id)
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

	result = ListHomeRealmDiscoveryPolicyAppliesTosCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
