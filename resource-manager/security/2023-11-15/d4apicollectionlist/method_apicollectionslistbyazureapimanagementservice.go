package d4apicollectionlist

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type APICollectionsListByAzureApiManagementServiceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ApiCollection
}

type APICollectionsListByAzureApiManagementServiceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ApiCollection
}

// APICollectionsListByAzureApiManagementService ...
func (c D4APICollectionListClient) APICollectionsListByAzureApiManagementService(ctx context.Context, id ServiceId) (result APICollectionsListByAzureApiManagementServiceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.Security/apiCollections", id.ID()),
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
		Values *[]ApiCollection `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// APICollectionsListByAzureApiManagementServiceComplete retrieves all the results into a single object
func (c D4APICollectionListClient) APICollectionsListByAzureApiManagementServiceComplete(ctx context.Context, id ServiceId) (APICollectionsListByAzureApiManagementServiceCompleteResult, error) {
	return c.APICollectionsListByAzureApiManagementServiceCompleteMatchingPredicate(ctx, id, ApiCollectionOperationPredicate{})
}

// APICollectionsListByAzureApiManagementServiceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c D4APICollectionListClient) APICollectionsListByAzureApiManagementServiceCompleteMatchingPredicate(ctx context.Context, id ServiceId, predicate ApiCollectionOperationPredicate) (result APICollectionsListByAzureApiManagementServiceCompleteResult, err error) {
	items := make([]ApiCollection, 0)

	resp, err := c.APICollectionsListByAzureApiManagementService(ctx, id)
	if err != nil {
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

	result = APICollectionsListByAzureApiManagementServiceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
