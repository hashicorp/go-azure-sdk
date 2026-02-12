package sqlscripts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlScriptGetSqlScriptsByWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SqlScriptResource
}

type SqlScriptGetSqlScriptsByWorkspaceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SqlScriptResource
}

type SqlScriptGetSqlScriptsByWorkspaceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *SqlScriptGetSqlScriptsByWorkspaceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// SqlScriptGetSqlScriptsByWorkspace ...
func (c SqlScriptsClient) SqlScriptGetSqlScriptsByWorkspace(ctx context.Context) (result SqlScriptGetSqlScriptsByWorkspaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &SqlScriptGetSqlScriptsByWorkspaceCustomPager{},
		Path:       "/sqlScripts",
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
		Values *[]SqlScriptResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SqlScriptGetSqlScriptsByWorkspaceComplete retrieves all the results into a single object
func (c SqlScriptsClient) SqlScriptGetSqlScriptsByWorkspaceComplete(ctx context.Context) (SqlScriptGetSqlScriptsByWorkspaceCompleteResult, error) {
	return c.SqlScriptGetSqlScriptsByWorkspaceCompleteMatchingPredicate(ctx, SqlScriptResourceOperationPredicate{})
}

// SqlScriptGetSqlScriptsByWorkspaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SqlScriptsClient) SqlScriptGetSqlScriptsByWorkspaceCompleteMatchingPredicate(ctx context.Context, predicate SqlScriptResourceOperationPredicate) (result SqlScriptGetSqlScriptsByWorkspaceCompleteResult, err error) {
	items := make([]SqlScriptResource, 0)

	resp, err := c.SqlScriptGetSqlScriptsByWorkspace(ctx)
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

	result = SqlScriptGetSqlScriptsByWorkspaceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
