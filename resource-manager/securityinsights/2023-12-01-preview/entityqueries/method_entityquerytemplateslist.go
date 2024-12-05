package entityqueries

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityQueryTemplatesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]EntityQueryTemplate
}

type EntityQueryTemplatesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []EntityQueryTemplate
}

type EntityQueryTemplatesListOperationOptions struct {
	Kind *Kind
}

func DefaultEntityQueryTemplatesListOperationOptions() EntityQueryTemplatesListOperationOptions {
	return EntityQueryTemplatesListOperationOptions{}
}

func (o EntityQueryTemplatesListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o EntityQueryTemplatesListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o EntityQueryTemplatesListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Kind != nil {
		out.Append("kind", fmt.Sprintf("%v", *o.Kind))
	}
	return &out
}

type EntityQueryTemplatesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *EntityQueryTemplatesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// EntityQueryTemplatesList ...
func (c EntityQueriesClient) EntityQueryTemplatesList(ctx context.Context, id WorkspaceId, options EntityQueryTemplatesListOperationOptions) (result EntityQueryTemplatesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &EntityQueryTemplatesListCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/entityQueryTemplates", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]EntityQueryTemplate, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := UnmarshalEntityQueryTemplateImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for EntityQueryTemplate (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// EntityQueryTemplatesListComplete retrieves all the results into a single object
func (c EntityQueriesClient) EntityQueryTemplatesListComplete(ctx context.Context, id WorkspaceId, options EntityQueryTemplatesListOperationOptions) (EntityQueryTemplatesListCompleteResult, error) {
	return c.EntityQueryTemplatesListCompleteMatchingPredicate(ctx, id, options, EntityQueryTemplateOperationPredicate{})
}

// EntityQueryTemplatesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntityQueriesClient) EntityQueryTemplatesListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, options EntityQueryTemplatesListOperationOptions, predicate EntityQueryTemplateOperationPredicate) (result EntityQueryTemplatesListCompleteResult, err error) {
	items := make([]EntityQueryTemplate, 0)

	resp, err := c.EntityQueryTemplatesList(ctx, id, options)
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

	result = EntityQueryTemplatesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
