package scripts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScriptCmdletsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ScriptCmdlet
}

type ScriptCmdletsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ScriptCmdlet
}

// ScriptCmdletsList ...
func (c ScriptsClient) ScriptCmdletsList(ctx context.Context, id ScriptPackageId) (result ScriptCmdletsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/scriptCmdlets", id.ID()),
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
		Values *[]ScriptCmdlet `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ScriptCmdletsListComplete retrieves all the results into a single object
func (c ScriptsClient) ScriptCmdletsListComplete(ctx context.Context, id ScriptPackageId) (ScriptCmdletsListCompleteResult, error) {
	return c.ScriptCmdletsListCompleteMatchingPredicate(ctx, id, ScriptCmdletOperationPredicate{})
}

// ScriptCmdletsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ScriptsClient) ScriptCmdletsListCompleteMatchingPredicate(ctx context.Context, id ScriptPackageId, predicate ScriptCmdletOperationPredicate) (result ScriptCmdletsListCompleteResult, err error) {
	items := make([]ScriptCmdlet, 0)

	resp, err := c.ScriptCmdletsList(ctx, id)
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

	result = ScriptCmdletsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
