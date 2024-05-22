package workspaceazureadonlyauthentications

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureADOnlyAuthenticationsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AzureADOnlyAuthentication
}

type AzureADOnlyAuthenticationsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AzureADOnlyAuthentication
}

// AzureADOnlyAuthenticationsList ...
func (c WorkspaceAzureADOnlyAuthenticationsClient) AzureADOnlyAuthenticationsList(ctx context.Context, id WorkspaceId) (result AzureADOnlyAuthenticationsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/azureADOnlyAuthentications", id.ID()),
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
		Values *[]AzureADOnlyAuthentication `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AzureADOnlyAuthenticationsListComplete retrieves all the results into a single object
func (c WorkspaceAzureADOnlyAuthenticationsClient) AzureADOnlyAuthenticationsListComplete(ctx context.Context, id WorkspaceId) (AzureADOnlyAuthenticationsListCompleteResult, error) {
	return c.AzureADOnlyAuthenticationsListCompleteMatchingPredicate(ctx, id, AzureADOnlyAuthenticationOperationPredicate{})
}

// AzureADOnlyAuthenticationsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WorkspaceAzureADOnlyAuthenticationsClient) AzureADOnlyAuthenticationsListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate AzureADOnlyAuthenticationOperationPredicate) (result AzureADOnlyAuthenticationsListCompleteResult, err error) {
	items := make([]AzureADOnlyAuthentication, 0)

	resp, err := c.AzureADOnlyAuthenticationsList(ctx, id)
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

	result = AzureADOnlyAuthenticationsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
