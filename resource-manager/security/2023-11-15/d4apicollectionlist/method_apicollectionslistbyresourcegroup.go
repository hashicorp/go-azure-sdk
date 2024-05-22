package d4apicollectionlist

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

type APICollectionsListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ApiCollection
}

type APICollectionsListByResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ApiCollection
}

// APICollectionsListByResourceGroup ...
func (c D4APICollectionListClient) APICollectionsListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result APICollectionsListByResourceGroupOperationResponse, err error) {
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

// APICollectionsListByResourceGroupComplete retrieves all the results into a single object
func (c D4APICollectionListClient) APICollectionsListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (APICollectionsListByResourceGroupCompleteResult, error) {
	return c.APICollectionsListByResourceGroupCompleteMatchingPredicate(ctx, id, ApiCollectionOperationPredicate{})
}

// APICollectionsListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c D4APICollectionListClient) APICollectionsListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate ApiCollectionOperationPredicate) (result APICollectionsListByResourceGroupCompleteResult, err error) {
	items := make([]ApiCollection, 0)

	resp, err := c.APICollectionsListByResourceGroup(ctx, id)
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

	result = APICollectionsListByResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
