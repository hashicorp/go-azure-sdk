package networksecurityperimeterconfigurations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByServiceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]NetworkSecurityPerimeterConfiguration
}

type ListByServiceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []NetworkSecurityPerimeterConfiguration
}

type ListByServiceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByServiceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByService ...
func (c NetworkSecurityPerimeterConfigurationsClient) ListByService(ctx context.Context, id SearchServiceId) (result ListByServiceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByServiceCustomPager{},
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

// ListByServiceComplete retrieves all the results into a single object
func (c NetworkSecurityPerimeterConfigurationsClient) ListByServiceComplete(ctx context.Context, id SearchServiceId) (ListByServiceCompleteResult, error) {
	return c.ListByServiceCompleteMatchingPredicate(ctx, id, NetworkSecurityPerimeterConfigurationOperationPredicate{})
}

// ListByServiceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkSecurityPerimeterConfigurationsClient) ListByServiceCompleteMatchingPredicate(ctx context.Context, id SearchServiceId, predicate NetworkSecurityPerimeterConfigurationOperationPredicate) (result ListByServiceCompleteResult, err error) {
	items := make([]NetworkSecurityPerimeterConfiguration, 0)

	resp, err := c.ListByService(ctx, id)
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

	result = ListByServiceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
