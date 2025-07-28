package createresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreationSupportedListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]CreateResourceSupportedProperties
}

type CreationSupportedListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []CreateResourceSupportedProperties
}

type CreationSupportedListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *CreationSupportedListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// CreationSupportedList ...
func (c CreateResourceClient) CreationSupportedList(ctx context.Context, id SubscriptionStatusId) (result CreationSupportedListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &CreationSupportedListCustomPager{},
		Path:       id.ID(),
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

// CreationSupportedListComplete retrieves all the results into a single object
func (c CreateResourceClient) CreationSupportedListComplete(ctx context.Context, id SubscriptionStatusId) (CreationSupportedListCompleteResult, error) {
	return c.CreationSupportedListCompleteMatchingPredicate(ctx, id, CreateResourceSupportedPropertiesOperationPredicate{})
}

// CreationSupportedListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CreateResourceClient) CreationSupportedListCompleteMatchingPredicate(ctx context.Context, id SubscriptionStatusId, predicate CreateResourceSupportedPropertiesOperationPredicate) (result CreationSupportedListCompleteResult, err error) {
	items := make([]CreateResourceSupportedProperties, 0)

	resp, err := c.CreationSupportedList(ctx, id)
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

	result = CreationSupportedListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
