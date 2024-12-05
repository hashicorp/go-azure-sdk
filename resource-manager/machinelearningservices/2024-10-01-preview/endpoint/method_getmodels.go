package endpoint

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetModelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]EndpointModelProperties
}

type GetModelsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []EndpointModelProperties
}

type GetModelsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *GetModelsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetModels ...
func (c EndpointClient) GetModels(ctx context.Context, id EndpointId) (result GetModelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &GetModelsCustomPager{},
		Path:       fmt.Sprintf("%s/models", id.ID()),
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
		Values *[]EndpointModelProperties `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetModelsComplete retrieves all the results into a single object
func (c EndpointClient) GetModelsComplete(ctx context.Context, id EndpointId) (GetModelsCompleteResult, error) {
	return c.GetModelsCompleteMatchingPredicate(ctx, id, EndpointModelPropertiesOperationPredicate{})
}

// GetModelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EndpointClient) GetModelsCompleteMatchingPredicate(ctx context.Context, id EndpointId, predicate EndpointModelPropertiesOperationPredicate) (result GetModelsCompleteResult, err error) {
	items := make([]EndpointModelProperties, 0)

	resp, err := c.GetModels(ctx, id)
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

	result = GetModelsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
