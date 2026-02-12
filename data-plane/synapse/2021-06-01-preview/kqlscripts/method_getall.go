package kqlscripts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAllOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]KqlScriptResource
}

type GetAllCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []KqlScriptResource
}

type GetAllCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *GetAllCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetAll ...
func (c KqlScriptsClient) GetAll(ctx context.Context) (result GetAllOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &GetAllCustomPager{},
		Path:       "/kqlScripts",
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
		Values *[]KqlScriptResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetAllComplete retrieves all the results into a single object
func (c KqlScriptsClient) GetAllComplete(ctx context.Context) (GetAllCompleteResult, error) {
	return c.GetAllCompleteMatchingPredicate(ctx, KqlScriptResourceOperationPredicate{})
}

// GetAllCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c KqlScriptsClient) GetAllCompleteMatchingPredicate(ctx context.Context, predicate KqlScriptResourceOperationPredicate) (result GetAllCompleteResult, err error) {
	items := make([]KqlScriptResource, 0)

	resp, err := c.GetAll(ctx)
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

	result = GetAllCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
