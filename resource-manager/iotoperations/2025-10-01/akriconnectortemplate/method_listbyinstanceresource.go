package akriconnectortemplate

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByInstanceResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AkriConnectorTemplateResource
}

type ListByInstanceResourceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AkriConnectorTemplateResource
}

type ListByInstanceResourceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByInstanceResourceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByInstanceResource ...
func (c AkriConnectorTemplateClient) ListByInstanceResource(ctx context.Context, id InstanceId) (result ListByInstanceResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByInstanceResourceCustomPager{},
		Path:       fmt.Sprintf("%s/akriConnectorTemplates", id.ID()),
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
		Values *[]AkriConnectorTemplateResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByInstanceResourceComplete retrieves all the results into a single object
func (c AkriConnectorTemplateClient) ListByInstanceResourceComplete(ctx context.Context, id InstanceId) (ListByInstanceResourceCompleteResult, error) {
	return c.ListByInstanceResourceCompleteMatchingPredicate(ctx, id, AkriConnectorTemplateResourceOperationPredicate{})
}

// ListByInstanceResourceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AkriConnectorTemplateClient) ListByInstanceResourceCompleteMatchingPredicate(ctx context.Context, id InstanceId, predicate AkriConnectorTemplateResourceOperationPredicate) (result ListByInstanceResourceCompleteResult, err error) {
	items := make([]AkriConnectorTemplateResource, 0)

	resp, err := c.ListByInstanceResource(ctx, id)
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

	result = ListByInstanceResourceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
