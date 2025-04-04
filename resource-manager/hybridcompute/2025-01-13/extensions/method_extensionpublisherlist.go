package extensions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtensionPublisherListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ExtensionPublisher
}

type ExtensionPublisherListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ExtensionPublisher
}

type ExtensionPublisherListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ExtensionPublisherListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ExtensionPublisherList ...
func (c ExtensionsClient) ExtensionPublisherList(ctx context.Context, id LocationId) (result ExtensionPublisherListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ExtensionPublisherListCustomPager{},
		Path:       fmt.Sprintf("%s/publishers", id.ID()),
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
		Values *[]ExtensionPublisher `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ExtensionPublisherListComplete retrieves all the results into a single object
func (c ExtensionsClient) ExtensionPublisherListComplete(ctx context.Context, id LocationId) (ExtensionPublisherListCompleteResult, error) {
	return c.ExtensionPublisherListCompleteMatchingPredicate(ctx, id, ExtensionPublisherOperationPredicate{})
}

// ExtensionPublisherListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ExtensionsClient) ExtensionPublisherListCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate ExtensionPublisherOperationPredicate) (result ExtensionPublisherListCompleteResult, err error) {
	items := make([]ExtensionPublisher, 0)

	resp, err := c.ExtensionPublisherList(ctx, id)
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

	result = ExtensionPublisherListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
