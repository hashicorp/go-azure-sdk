package schemaversions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListBySchemaOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SchemaVersion
}

type ListBySchemaCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SchemaVersion
}

type ListBySchemaCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListBySchemaCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListBySchema ...
func (c SchemaVersionsClient) ListBySchema(ctx context.Context, id SchemaId) (result ListBySchemaOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListBySchemaCustomPager{},
		Path:       fmt.Sprintf("%s/schemaVersions", id.ID()),
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
		Values *[]SchemaVersion `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListBySchemaComplete retrieves all the results into a single object
func (c SchemaVersionsClient) ListBySchemaComplete(ctx context.Context, id SchemaId) (ListBySchemaCompleteResult, error) {
	return c.ListBySchemaCompleteMatchingPredicate(ctx, id, SchemaVersionOperationPredicate{})
}

// ListBySchemaCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SchemaVersionsClient) ListBySchemaCompleteMatchingPredicate(ctx context.Context, id SchemaId, predicate SchemaVersionOperationPredicate) (result ListBySchemaCompleteResult, err error) {
	items := make([]SchemaVersion, 0)

	resp, err := c.ListBySchema(ctx, id)
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

	result = ListBySchemaCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
