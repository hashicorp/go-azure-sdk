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

type ApiTokensListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ApiToken
}

type ApiTokensListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ApiToken
}

type ApiTokensListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ApiTokensListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ApiTokensList ...
func (c IotcentralsClient) ApiTokensList(ctx context.Context) (result ApiTokensListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ApiTokensListCustomPager{},
		Path:       "/apiTokens",
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
		Values *[]ApiToken `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ApiTokensListComplete retrieves all the results into a single object
func (c IotcentralsClient) ApiTokensListComplete(ctx context.Context) (ApiTokensListCompleteResult, error) {
	return c.ApiTokensListCompleteMatchingPredicate(ctx, ApiTokenOperationPredicate{})
}

// ApiTokensListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotcentralsClient) ApiTokensListCompleteMatchingPredicate(ctx context.Context, predicate ApiTokenOperationPredicate) (result ApiTokensListCompleteResult, err error) {
	items := make([]ApiToken, 0)

	resp, err := c.ApiTokensList(ctx)
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

	result = ApiTokensListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
