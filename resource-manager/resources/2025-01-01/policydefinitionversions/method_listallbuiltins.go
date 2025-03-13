package policydefinitionversions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAllBuiltinsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PolicyDefinitionVersion
}

type ListAllBuiltinsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PolicyDefinitionVersion
}

type ListAllBuiltinsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListAllBuiltinsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAllBuiltins ...
func (c PolicyDefinitionVersionsClient) ListAllBuiltins(ctx context.Context) (result ListAllBuiltinsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &ListAllBuiltinsCustomPager{},
		Path:       "/providers/Microsoft.Authorization/listPolicyDefinitionVersions",
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
		Values *[]PolicyDefinitionVersion `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAllBuiltinsComplete retrieves all the results into a single object
func (c PolicyDefinitionVersionsClient) ListAllBuiltinsComplete(ctx context.Context) (ListAllBuiltinsCompleteResult, error) {
	return c.ListAllBuiltinsCompleteMatchingPredicate(ctx, PolicyDefinitionVersionOperationPredicate{})
}

// ListAllBuiltinsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyDefinitionVersionsClient) ListAllBuiltinsCompleteMatchingPredicate(ctx context.Context, predicate PolicyDefinitionVersionOperationPredicate) (result ListAllBuiltinsCompleteResult, err error) {
	items := make([]PolicyDefinitionVersion, 0)

	resp, err := c.ListAllBuiltins(ctx)
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

	result = ListAllBuiltinsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
