package networksecurityperimeter

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]NetworkSecurityPerimeterConfiguration
}

type ListConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []NetworkSecurityPerimeterConfiguration
}

type ListConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConfigurations ...
func (c NetworkSecurityPerimeterClient) ListConfigurations(ctx context.Context, id BatchAccountId) (result ListConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConfigurationsCustomPager{},
		Path:       fmt.Sprintf("%s/networkSecurityPerimeterConfigurations", id.ID()),
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
		Values *[]NetworkSecurityPerimeterConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConfigurationsComplete retrieves all the results into a single object
func (c NetworkSecurityPerimeterClient) ListConfigurationsComplete(ctx context.Context, id BatchAccountId) (ListConfigurationsCompleteResult, error) {
	return c.ListConfigurationsCompleteMatchingPredicate(ctx, id, NetworkSecurityPerimeterConfigurationOperationPredicate{})
}

// ListConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkSecurityPerimeterClient) ListConfigurationsCompleteMatchingPredicate(ctx context.Context, id BatchAccountId, predicate NetworkSecurityPerimeterConfigurationOperationPredicate) (result ListConfigurationsCompleteResult, err error) {
	items := make([]NetworkSecurityPerimeterConfiguration, 0)

	resp, err := c.ListConfigurations(ctx, id)
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

	result = ListConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
