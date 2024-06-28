package sqlpoolssecurityalertpolicies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolSecurityAlertPoliciesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SqlPoolSecurityAlertPolicy
}

type SqlPoolSecurityAlertPoliciesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SqlPoolSecurityAlertPolicy
}

type SqlPoolSecurityAlertPoliciesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *SqlPoolSecurityAlertPoliciesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// SqlPoolSecurityAlertPoliciesList ...
func (c SqlPoolsSecurityAlertPoliciesClient) SqlPoolSecurityAlertPoliciesList(ctx context.Context, id SqlPoolId) (result SqlPoolSecurityAlertPoliciesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &SqlPoolSecurityAlertPoliciesListCustomPager{},
		Path:       fmt.Sprintf("%s/securityAlertPolicies", id.ID()),
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
		Values *[]SqlPoolSecurityAlertPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SqlPoolSecurityAlertPoliciesListComplete retrieves all the results into a single object
func (c SqlPoolsSecurityAlertPoliciesClient) SqlPoolSecurityAlertPoliciesListComplete(ctx context.Context, id SqlPoolId) (SqlPoolSecurityAlertPoliciesListCompleteResult, error) {
	return c.SqlPoolSecurityAlertPoliciesListCompleteMatchingPredicate(ctx, id, SqlPoolSecurityAlertPolicyOperationPredicate{})
}

// SqlPoolSecurityAlertPoliciesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SqlPoolsSecurityAlertPoliciesClient) SqlPoolSecurityAlertPoliciesListCompleteMatchingPredicate(ctx context.Context, id SqlPoolId, predicate SqlPoolSecurityAlertPolicyOperationPredicate) (result SqlPoolSecurityAlertPoliciesListCompleteResult, err error) {
	items := make([]SqlPoolSecurityAlertPolicy, 0)

	resp, err := c.SqlPoolSecurityAlertPoliciesList(ctx, id)
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

	result = SqlPoolSecurityAlertPoliciesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
