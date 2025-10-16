package schemas

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListBySchemaRegistryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Schema
}

type ListBySchemaRegistryCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Schema
}

type ListBySchemaRegistryCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListBySchemaRegistryCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListBySchemaRegistry ...
func (c SchemasClient) ListBySchemaRegistry(ctx context.Context, id SchemaRegistryId) (result ListBySchemaRegistryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListBySchemaRegistryCustomPager{},
		Path:       fmt.Sprintf("%s/schemas", id.ID()),
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
		Values *[]Schema `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListBySchemaRegistryComplete retrieves all the results into a single object
func (c SchemasClient) ListBySchemaRegistryComplete(ctx context.Context, id SchemaRegistryId) (ListBySchemaRegistryCompleteResult, error) {
	return c.ListBySchemaRegistryCompleteMatchingPredicate(ctx, id, SchemaOperationPredicate{})
}

// ListBySchemaRegistryCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SchemasClient) ListBySchemaRegistryCompleteMatchingPredicate(ctx context.Context, id SchemaRegistryId, predicate SchemaOperationPredicate) (result ListBySchemaRegistryCompleteResult, err error) {
	items := make([]Schema, 0)

	resp, err := c.ListBySchemaRegistry(ctx, id)
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

	result = ListBySchemaRegistryCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
