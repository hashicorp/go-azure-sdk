package usagesinformation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsagesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]CurrentUsagesBase
}

type UsagesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []CurrentUsagesBase
}

// UsagesList ...
func (c UsagesInformationClient) UsagesList(ctx context.Context, id commonids.ScopeId) (result UsagesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.Quota/usages", id.ID()),
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
		Values *[]CurrentUsagesBase `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// UsagesListComplete retrieves all the results into a single object
func (c UsagesInformationClient) UsagesListComplete(ctx context.Context, id commonids.ScopeId) (UsagesListCompleteResult, error) {
	return c.UsagesListCompleteMatchingPredicate(ctx, id, CurrentUsagesBaseOperationPredicate{})
}

// UsagesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UsagesInformationClient) UsagesListCompleteMatchingPredicate(ctx context.Context, id commonids.ScopeId, predicate CurrentUsagesBaseOperationPredicate) (result UsagesListCompleteResult, err error) {
	items := make([]CurrentUsagesBase, 0)

	resp, err := c.UsagesList(ctx, id)
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

	result = UsagesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
