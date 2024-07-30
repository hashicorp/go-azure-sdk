package siteinformationprotectiondatalosspreventionpolicy

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

type ListSiteInformationProtectionDataLossPreventionPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DataLossPreventionPolicy
}

type ListSiteInformationProtectionDataLossPreventionPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DataLossPreventionPolicy
}

type ListSiteInformationProtectionDataLossPreventionPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteInformationProtectionDataLossPreventionPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteInformationProtectionDataLossPreventionPolicies ...
func (c SiteInformationProtectionDataLossPreventionPolicyClient) ListSiteInformationProtectionDataLossPreventionPolicies(ctx context.Context, id GroupIdSiteId) (result ListSiteInformationProtectionDataLossPreventionPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSiteInformationProtectionDataLossPreventionPoliciesCustomPager{},
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

// ListSiteInformationProtectionDataLossPreventionPoliciesComplete retrieves all the results into a single object
func (c SiteInformationProtectionDataLossPreventionPolicyClient) ListSiteInformationProtectionDataLossPreventionPoliciesComplete(ctx context.Context, id GroupIdSiteId) (ListSiteInformationProtectionDataLossPreventionPoliciesCompleteResult, error) {
	return c.ListSiteInformationProtectionDataLossPreventionPoliciesCompleteMatchingPredicate(ctx, id, DataLossPreventionPolicyOperationPredicate{})
}

// ListSiteInformationProtectionDataLossPreventionPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteInformationProtectionDataLossPreventionPolicyClient) ListSiteInformationProtectionDataLossPreventionPoliciesCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteId, predicate DataLossPreventionPolicyOperationPredicate) (result ListSiteInformationProtectionDataLossPreventionPoliciesCompleteResult, err error) {
	items := make([]beta.DataLossPreventionPolicy, 0)

	resp, err := c.ListSiteInformationProtectionDataLossPreventionPolicies(ctx, id)
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

	result = ListSiteInformationProtectionDataLossPreventionPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
