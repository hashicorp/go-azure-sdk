package exchangeconnector

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

type ListExchangeConnectorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementExchangeConnector
}

type ListExchangeConnectorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementExchangeConnector
}

type ListExchangeConnectorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListExchangeConnectorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListExchangeConnectors ...
func (c ExchangeConnectorClient) ListExchangeConnectors(ctx context.Context) (result ListExchangeConnectorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListExchangeConnectorsCustomPager{},
		Path:       "/deviceManagement/exchangeConnectors",
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
		Values *[]beta.DeviceManagementExchangeConnector `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListExchangeConnectorsComplete retrieves all the results into a single object
func (c ExchangeConnectorClient) ListExchangeConnectorsComplete(ctx context.Context) (ListExchangeConnectorsCompleteResult, error) {
	return c.ListExchangeConnectorsCompleteMatchingPredicate(ctx, DeviceManagementExchangeConnectorOperationPredicate{})
}

// ListExchangeConnectorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ExchangeConnectorClient) ListExchangeConnectorsCompleteMatchingPredicate(ctx context.Context, predicate DeviceManagementExchangeConnectorOperationPredicate) (result ListExchangeConnectorsCompleteResult, err error) {
	items := make([]beta.DeviceManagementExchangeConnector, 0)

	resp, err := c.ListExchangeConnectors(ctx)
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

	result = ListExchangeConnectorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
