package schemaextension

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSchemaExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SchemaExtension
}

type ListSchemaExtensionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SchemaExtension
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

// ListSchemaExtensions ...
func (c SchemaExtensionClient) ListSchemaExtensions(ctx context.Context) (result ListSchemaExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSchemaExtensionsCustomPager{},
		Path:       "/schemaExtensions",
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
		Values *[]beta.SchemaExtension `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSchemaExtensionsComplete retrieves all the results into a single object
func (c SchemaExtensionClient) ListSchemaExtensionsComplete(ctx context.Context) (ListSchemaExtensionsCompleteResult, error) {
	return c.ListSchemaExtensionsCompleteMatchingPredicate(ctx, SchemaExtensionOperationPredicate{})
}

// ListSchemaExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SchemaExtensionClient) ListSchemaExtensionsCompleteMatchingPredicate(ctx context.Context, predicate SchemaExtensionOperationPredicate) (result ListSchemaExtensionsCompleteResult, err error) {
	items := make([]beta.SchemaExtension, 0)

	resp, err := c.ListSchemaExtensions(ctx)
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
