package claimsmappingpolicyappliesto

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

type ListClaimsMappingPolicyAppliesTosOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type ListClaimsMappingPolicyAppliesTosCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type ListClaimsMappingPolicyAppliesTosCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListClaimsMappingPolicyAppliesTosCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListClaimsMappingPolicyAppliesTos ...
func (c ClaimsMappingPolicyAppliesToClient) ListClaimsMappingPolicyAppliesTos(ctx context.Context, id PolicyClaimsMappingPolicyId) (result ListClaimsMappingPolicyAppliesTosOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListClaimsMappingPolicyAppliesTosCustomPager{},
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

// ListClaimsMappingPolicyAppliesTosComplete retrieves all the results into a single object
func (c ClaimsMappingPolicyAppliesToClient) ListClaimsMappingPolicyAppliesTosComplete(ctx context.Context, id PolicyClaimsMappingPolicyId) (ListClaimsMappingPolicyAppliesTosCompleteResult, error) {
	return c.ListClaimsMappingPolicyAppliesTosCompleteMatchingPredicate(ctx, id, DirectoryObjectOperationPredicate{})
}

// ListClaimsMappingPolicyAppliesTosCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ClaimsMappingPolicyAppliesToClient) ListClaimsMappingPolicyAppliesTosCompleteMatchingPredicate(ctx context.Context, id PolicyClaimsMappingPolicyId, predicate DirectoryObjectOperationPredicate) (result ListClaimsMappingPolicyAppliesTosCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.ListClaimsMappingPolicyAppliesTos(ctx, id)
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

	result = ListClaimsMappingPolicyAppliesTosCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
