package informationprotectiondatalosspreventionpolicy

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

type ListInformationProtectionDataLossPreventionPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DataLossPreventionPolicy
}

type ListInformationProtectionDataLossPreventionPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DataLossPreventionPolicy
}

type ListInformationProtectionDataLossPreventionPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListInformationProtectionDataLossPreventionPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListInformationProtectionDataLossPreventionPolicies ...
func (c InformationProtectionDataLossPreventionPolicyClient) ListInformationProtectionDataLossPreventionPolicies(ctx context.Context, id UserId) (result ListInformationProtectionDataLossPreventionPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListInformationProtectionDataLossPreventionPoliciesCustomPager{},
		Path:       fmt.Sprintf("%s/informationProtection/dataLossPreventionPolicies", id.ID()),
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
		Values *[]beta.DataLossPreventionPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListInformationProtectionDataLossPreventionPoliciesComplete retrieves all the results into a single object
func (c InformationProtectionDataLossPreventionPolicyClient) ListInformationProtectionDataLossPreventionPoliciesComplete(ctx context.Context, id UserId) (ListInformationProtectionDataLossPreventionPoliciesCompleteResult, error) {
	return c.ListInformationProtectionDataLossPreventionPoliciesCompleteMatchingPredicate(ctx, id, DataLossPreventionPolicyOperationPredicate{})
}

// ListInformationProtectionDataLossPreventionPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InformationProtectionDataLossPreventionPolicyClient) ListInformationProtectionDataLossPreventionPoliciesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate DataLossPreventionPolicyOperationPredicate) (result ListInformationProtectionDataLossPreventionPoliciesCompleteResult, err error) {
	items := make([]beta.DataLossPreventionPolicy, 0)

	resp, err := c.ListInformationProtectionDataLossPreventionPolicies(ctx, id)
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

	result = ListInformationProtectionDataLossPreventionPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
