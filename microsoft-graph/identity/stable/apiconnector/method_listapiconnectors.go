package apiconnector

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListApiConnectorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityApiConnector
}

type ListApiConnectorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityApiConnector
}

type ListApiConnectorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListApiConnectorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListApiConnectors ...
func (c ApiConnectorClient) ListApiConnectors(ctx context.Context) (result ListApiConnectorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListApiConnectorsCustomPager{},
		Path:       "/identity/apiConnectors",
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
		Values *[]stable.IdentityApiConnector `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListApiConnectorsComplete retrieves all the results into a single object
func (c ApiConnectorClient) ListApiConnectorsComplete(ctx context.Context) (ListApiConnectorsCompleteResult, error) {
	return c.ListApiConnectorsCompleteMatchingPredicate(ctx, IdentityApiConnectorOperationPredicate{})
}

// ListApiConnectorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApiConnectorClient) ListApiConnectorsCompleteMatchingPredicate(ctx context.Context, predicate IdentityApiConnectorOperationPredicate) (result ListApiConnectorsCompleteResult, err error) {
	items := make([]stable.IdentityApiConnector, 0)

	resp, err := c.ListApiConnectors(ctx)
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

	result = ListApiConnectorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
