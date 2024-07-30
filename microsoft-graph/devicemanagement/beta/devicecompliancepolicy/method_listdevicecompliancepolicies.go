package devicecompliancepolicy

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

type ListDeviceCompliancePoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceCompliancePolicy
}

type ListDeviceCompliancePoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceCompliancePolicy
}

type ListDeviceCompliancePoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceCompliancePoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceCompliancePolicies ...
func (c DeviceCompliancePolicyClient) ListDeviceCompliancePolicies(ctx context.Context) (result ListDeviceCompliancePoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceCompliancePoliciesCustomPager{},
		Path:       "/deviceManagement/deviceCompliancePolicies",
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
		Values *[]beta.DeviceCompliancePolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceCompliancePoliciesComplete retrieves all the results into a single object
func (c DeviceCompliancePolicyClient) ListDeviceCompliancePoliciesComplete(ctx context.Context) (ListDeviceCompliancePoliciesCompleteResult, error) {
	return c.ListDeviceCompliancePoliciesCompleteMatchingPredicate(ctx, DeviceCompliancePolicyOperationPredicate{})
}

// ListDeviceCompliancePoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCompliancePolicyClient) ListDeviceCompliancePoliciesCompleteMatchingPredicate(ctx context.Context, predicate DeviceCompliancePolicyOperationPredicate) (result ListDeviceCompliancePoliciesCompleteResult, err error) {
	items := make([]beta.DeviceCompliancePolicy, 0)

	resp, err := c.ListDeviceCompliancePolicies(ctx)
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

	result = ListDeviceCompliancePoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
