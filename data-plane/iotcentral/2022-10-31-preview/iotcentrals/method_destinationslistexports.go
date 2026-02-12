package iotcentrals

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DestinationsListExportsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Export
}

type DestinationsListExportsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Export
}

type DestinationsListExportsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DestinationsListExportsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DestinationsListExports ...
func (c IotcentralsClient) DestinationsListExports(ctx context.Context, id DestinationIdId) (result DestinationsListExportsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &DestinationsListExportsCustomPager{},
		Path:       fmt.Sprintf("%s/exports", id.Path()),
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
		Values *[]Export `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// DestinationsListExportsComplete retrieves all the results into a single object
func (c IotcentralsClient) DestinationsListExportsComplete(ctx context.Context, id DestinationIdId) (DestinationsListExportsCompleteResult, error) {
	return c.DestinationsListExportsCompleteMatchingPredicate(ctx, id, ExportOperationPredicate{})
}

// DestinationsListExportsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotcentralsClient) DestinationsListExportsCompleteMatchingPredicate(ctx context.Context, id DestinationIdId, predicate ExportOperationPredicate) (result DestinationsListExportsCompleteResult, err error) {
	items := make([]Export, 0)

	resp, err := c.DestinationsListExports(ctx, id)
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

	result = DestinationsListExportsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
