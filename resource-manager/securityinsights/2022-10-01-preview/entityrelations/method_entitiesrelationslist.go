package entityrelations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitiesRelationsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Relation
}

type EntitiesRelationsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Relation
}

type EntitiesRelationsListOperationOptions struct {
	Filter  *string
	Orderby *string
	Top     *int64
}

func DefaultEntitiesRelationsListOperationOptions() EntitiesRelationsListOperationOptions {
	return EntitiesRelationsListOperationOptions{}
}

func (o EntitiesRelationsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o EntitiesRelationsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o EntitiesRelationsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Orderby != nil {
		out.Append("$orderby", fmt.Sprintf("%v", *o.Orderby))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// EntitiesRelationsList ...
func (c EntityRelationsClient) EntitiesRelationsList(ctx context.Context, id EntityId, options EntitiesRelationsListOperationOptions) (result EntitiesRelationsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/relations", id.ID()),
		OptionsObject: options,
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
		Values *[]Relation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// EntitiesRelationsListComplete retrieves all the results into a single object
func (c EntityRelationsClient) EntitiesRelationsListComplete(ctx context.Context, id EntityId, options EntitiesRelationsListOperationOptions) (EntitiesRelationsListCompleteResult, error) {
	return c.EntitiesRelationsListCompleteMatchingPredicate(ctx, id, options, RelationOperationPredicate{})
}

// EntitiesRelationsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntityRelationsClient) EntitiesRelationsListCompleteMatchingPredicate(ctx context.Context, id EntityId, options EntitiesRelationsListOperationOptions, predicate RelationOperationPredicate) (result EntitiesRelationsListCompleteResult, err error) {
	items := make([]Relation, 0)

	resp, err := c.EntitiesRelationsList(ctx, id, options)
	if err != nil {
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

	result = EntitiesRelationsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
