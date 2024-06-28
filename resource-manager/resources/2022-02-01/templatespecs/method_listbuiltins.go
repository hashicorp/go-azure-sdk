package templatespecs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListBuiltInsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]TemplateSpec
}

type ListBuiltInsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []TemplateSpec
}

type ListBuiltInsOperationOptions struct {
	Expand *TemplateSpecExpandKind
}

func DefaultListBuiltInsOperationOptions() ListBuiltInsOperationOptions {
	return ListBuiltInsOperationOptions{}
}

func (o ListBuiltInsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListBuiltInsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListBuiltInsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Expand != nil {
		out.Append("$expand", fmt.Sprintf("%v", *o.Expand))
	}
	return &out
}

type ListBuiltInsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListBuiltInsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListBuiltIns ...
func (c TemplateSpecsClient) ListBuiltIns(ctx context.Context, options ListBuiltInsOperationOptions) (result ListBuiltInsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListBuiltInsCustomPager{},
		Path:          "/providers/Microsoft.Resources/builtInTemplateSpecs",
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
		Values *[]TemplateSpec `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListBuiltInsComplete retrieves all the results into a single object
func (c TemplateSpecsClient) ListBuiltInsComplete(ctx context.Context, options ListBuiltInsOperationOptions) (ListBuiltInsCompleteResult, error) {
	return c.ListBuiltInsCompleteMatchingPredicate(ctx, options, TemplateSpecOperationPredicate{})
}

// ListBuiltInsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TemplateSpecsClient) ListBuiltInsCompleteMatchingPredicate(ctx context.Context, options ListBuiltInsOperationOptions, predicate TemplateSpecOperationPredicate) (result ListBuiltInsCompleteResult, err error) {
	items := make([]TemplateSpec, 0)

	resp, err := c.ListBuiltIns(ctx, options)
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

	result = ListBuiltInsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
