package openaiintegration

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenAIListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]OpenAIIntegrationRPModel
}

type OpenAIListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []OpenAIIntegrationRPModel
}

type OpenAIListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *OpenAIListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// OpenAIList ...
func (c OpenAIIntegrationClient) OpenAIList(ctx context.Context, id MonitorId) (result OpenAIListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &OpenAIListCustomPager{},
		Path:       fmt.Sprintf("%s/openAIIntegrations", id.ID()),
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
		Values *[]OpenAIIntegrationRPModel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// OpenAIListComplete retrieves all the results into a single object
func (c OpenAIIntegrationClient) OpenAIListComplete(ctx context.Context, id MonitorId) (OpenAIListCompleteResult, error) {
	return c.OpenAIListCompleteMatchingPredicate(ctx, id, OpenAIIntegrationRPModelOperationPredicate{})
}

// OpenAIListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenAIIntegrationClient) OpenAIListCompleteMatchingPredicate(ctx context.Context, id MonitorId, predicate OpenAIIntegrationRPModelOperationPredicate) (result OpenAIListCompleteResult, err error) {
	items := make([]OpenAIIntegrationRPModel, 0)

	resp, err := c.OpenAIList(ctx, id)
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

	result = OpenAIListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
