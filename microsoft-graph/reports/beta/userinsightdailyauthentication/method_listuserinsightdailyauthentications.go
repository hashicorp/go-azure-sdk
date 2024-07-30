package userinsightdailyauthentication

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

type ListUserInsightDailyAuthenticationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AuthenticationsMetric
}

type ListUserInsightDailyAuthenticationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AuthenticationsMetric
}

type ListUserInsightDailyAuthenticationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserInsightDailyAuthenticationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserInsightDailyAuthentications ...
func (c UserInsightDailyAuthenticationClient) ListUserInsightDailyAuthentications(ctx context.Context) (result ListUserInsightDailyAuthenticationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserInsightDailyAuthenticationsCustomPager{},
		Path:       "/reports/userInsights/daily/authentications",
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
		Values *[]beta.AuthenticationsMetric `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserInsightDailyAuthenticationsComplete retrieves all the results into a single object
func (c UserInsightDailyAuthenticationClient) ListUserInsightDailyAuthenticationsComplete(ctx context.Context) (ListUserInsightDailyAuthenticationsCompleteResult, error) {
	return c.ListUserInsightDailyAuthenticationsCompleteMatchingPredicate(ctx, AuthenticationsMetricOperationPredicate{})
}

// ListUserInsightDailyAuthenticationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightDailyAuthenticationClient) ListUserInsightDailyAuthenticationsCompleteMatchingPredicate(ctx context.Context, predicate AuthenticationsMetricOperationPredicate) (result ListUserInsightDailyAuthenticationsCompleteResult, err error) {
	items := make([]beta.AuthenticationsMetric, 0)

	resp, err := c.ListUserInsightDailyAuthentications(ctx)
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

	result = ListUserInsightDailyAuthenticationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
