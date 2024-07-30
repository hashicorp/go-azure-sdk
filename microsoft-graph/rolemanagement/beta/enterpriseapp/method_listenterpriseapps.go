package enterpriseapp

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

type ListEnterpriseAppsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.RbacApplication
}

type ListEnterpriseAppsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.RbacApplication
}

type ListEnterpriseAppsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEnterpriseAppsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEnterpriseApps ...
func (c EnterpriseAppClient) ListEnterpriseApps(ctx context.Context) (result ListEnterpriseAppsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEnterpriseAppsCustomPager{},
		Path:       "/roleManagement/enterpriseApps",
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
		Values *[]beta.RbacApplication `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEnterpriseAppsComplete retrieves all the results into a single object
func (c EnterpriseAppClient) ListEnterpriseAppsComplete(ctx context.Context) (ListEnterpriseAppsCompleteResult, error) {
	return c.ListEnterpriseAppsCompleteMatchingPredicate(ctx, RbacApplicationOperationPredicate{})
}

// ListEnterpriseAppsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EnterpriseAppClient) ListEnterpriseAppsCompleteMatchingPredicate(ctx context.Context, predicate RbacApplicationOperationPredicate) (result ListEnterpriseAppsCompleteResult, err error) {
	items := make([]beta.RbacApplication, 0)

	resp, err := c.ListEnterpriseApps(ctx)
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

	result = ListEnterpriseAppsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
