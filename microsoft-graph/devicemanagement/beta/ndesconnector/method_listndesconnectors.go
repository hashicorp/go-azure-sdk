package ndesconnector

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

type ListNdesConnectorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.NdesConnector
}

type ListNdesConnectorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.NdesConnector
}

type ListNdesConnectorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListNdesConnectorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListNdesConnectors ...
func (c NdesConnectorClient) ListNdesConnectors(ctx context.Context) (result ListNdesConnectorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListNdesConnectorsCustomPager{},
		Path:       "/deviceManagement/ndesConnectors",
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
		Values *[]beta.NdesConnector `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListNdesConnectorsComplete retrieves all the results into a single object
func (c NdesConnectorClient) ListNdesConnectorsComplete(ctx context.Context) (ListNdesConnectorsCompleteResult, error) {
	return c.ListNdesConnectorsCompleteMatchingPredicate(ctx, NdesConnectorOperationPredicate{})
}

// ListNdesConnectorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NdesConnectorClient) ListNdesConnectorsCompleteMatchingPredicate(ctx context.Context, predicate NdesConnectorOperationPredicate) (result ListNdesConnectorsCompleteResult, err error) {
	items := make([]beta.NdesConnector, 0)

	resp, err := c.ListNdesConnectors(ctx)
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

	result = ListNdesConnectorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
