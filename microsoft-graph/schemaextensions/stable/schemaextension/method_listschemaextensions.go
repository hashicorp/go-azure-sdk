package schemaextension

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSchemaExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.SchemaExtension
}

type ListSchemaExtensionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.SchemaExtension
}

type ListSchemaExtensionsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListSchemaExtensionsOperationOptions() ListSchemaExtensionsOperationOptions {
	return ListSchemaExtensionsOperationOptions{}
}

func (o ListSchemaExtensionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSchemaExtensionsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListSchemaExtensionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSchemaExtensionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSchemaExtensionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSchemaExtensions - List schemaExtensions. Get a list of schemaExtension objects in your tenant. The schema
// extensions can be InDevelopment, Available, or Deprecated and includes schema extensions
func (c SchemaExtensionClient) ListSchemaExtensions(ctx context.Context, options ListSchemaExtensionsOperationOptions) (result ListSchemaExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSchemaExtensionsCustomPager{},
		Path:          "/schemaExtensions",
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
		Values *[]stable.SchemaExtension `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSchemaExtensionsComplete retrieves all the results into a single object
func (c SchemaExtensionClient) ListSchemaExtensionsComplete(ctx context.Context, options ListSchemaExtensionsOperationOptions) (ListSchemaExtensionsCompleteResult, error) {
	return c.ListSchemaExtensionsCompleteMatchingPredicate(ctx, options, SchemaExtensionOperationPredicate{})
}

// ListSchemaExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SchemaExtensionClient) ListSchemaExtensionsCompleteMatchingPredicate(ctx context.Context, options ListSchemaExtensionsOperationOptions, predicate SchemaExtensionOperationPredicate) (result ListSchemaExtensionsCompleteResult, err error) {
	items := make([]stable.SchemaExtension, 0)

	resp, err := c.ListSchemaExtensions(ctx, options)
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

	result = ListSchemaExtensionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
