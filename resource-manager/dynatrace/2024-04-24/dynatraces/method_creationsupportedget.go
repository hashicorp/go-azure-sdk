package dynatraces

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreationSupportedGetOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]CreateResourceSupportedProperties
}

type CreationSupportedGetCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []CreateResourceSupportedProperties
}

type CreationSupportedGetCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *CreationSupportedGetCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// CreationSupportedGet ...
func (c DynatracesClient) CreationSupportedGet(ctx context.Context, id SubscriptionStatusId) (result CreationSupportedGetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &CreationSupportedGetCustomPager{},
		Path:       fmt.Sprintf("%s/default", id.ID()),
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
		Values *[]CreateResourceSupportedProperties `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// CreationSupportedGetComplete retrieves all the results into a single object
func (c DynatracesClient) CreationSupportedGetComplete(ctx context.Context, id SubscriptionStatusId) (CreationSupportedGetCompleteResult, error) {
	return c.CreationSupportedGetCompleteMatchingPredicate(ctx, id, CreateResourceSupportedPropertiesOperationPredicate{})
}

// CreationSupportedGetCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DynatracesClient) CreationSupportedGetCompleteMatchingPredicate(ctx context.Context, id SubscriptionStatusId, predicate CreateResourceSupportedPropertiesOperationPredicate) (result CreationSupportedGetCompleteResult, err error) {
	items := make([]CreateResourceSupportedProperties, 0)

	resp, err := c.CreationSupportedGet(ctx, id)
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

	result = CreationSupportedGetCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
