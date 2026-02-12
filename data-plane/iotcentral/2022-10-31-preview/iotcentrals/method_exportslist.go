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

type ExportsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Export
}

type ExportsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Export
}

type ExportsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ExportsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ExportsList ...
func (c IotcentralsClient) ExportsList(ctx context.Context) (result ExportsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ExportsListCustomPager{},
		Path:       "/dataExport/exports",
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

// ExportsListComplete retrieves all the results into a single object
func (c IotcentralsClient) ExportsListComplete(ctx context.Context) (ExportsListCompleteResult, error) {
	return c.ExportsListCompleteMatchingPredicate(ctx, ExportOperationPredicate{})
}

// ExportsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotcentralsClient) ExportsListCompleteMatchingPredicate(ctx context.Context, predicate ExportOperationPredicate) (result ExportsListCompleteResult, err error) {
	items := make([]Export, 0)

	resp, err := c.ExportsList(ctx)
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

	result = ExportsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
