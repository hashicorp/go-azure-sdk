package akriconnector

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByTemplateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AkriConnectorResource
}

type ListByTemplateCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AkriConnectorResource
}

type ListByTemplateCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByTemplateCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByTemplate ...
func (c AkriConnectorClient) ListByTemplate(ctx context.Context, id AkriConnectorTemplateId) (result ListByTemplateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByTemplateCustomPager{},
		Path:       fmt.Sprintf("%s/connectors", id.ID()),
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
		Values *[]AkriConnectorResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByTemplateComplete retrieves all the results into a single object
func (c AkriConnectorClient) ListByTemplateComplete(ctx context.Context, id AkriConnectorTemplateId) (ListByTemplateCompleteResult, error) {
	return c.ListByTemplateCompleteMatchingPredicate(ctx, id, AkriConnectorResourceOperationPredicate{})
}

// ListByTemplateCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AkriConnectorClient) ListByTemplateCompleteMatchingPredicate(ctx context.Context, id AkriConnectorTemplateId, predicate AkriConnectorResourceOperationPredicate) (result ListByTemplateCompleteResult, err error) {
	items := make([]AkriConnectorResource, 0)

	resp, err := c.ListByTemplate(ctx, id)
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

	result = ListByTemplateCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
