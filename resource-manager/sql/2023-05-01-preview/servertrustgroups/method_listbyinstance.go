package servertrustgroups

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

type ListByInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ServerTrustGroup
}

type ListByInstanceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ServerTrustGroup
}

// ListByInstance ...
func (c ServerTrustGroupsClient) ListByInstance(ctx context.Context, id commonids.SqlManagedInstanceId) (result ListByInstanceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/serverTrustGroups", id.ID()),
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
		Values *[]ServerTrustGroup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByInstanceComplete retrieves all the results into a single object
func (c ServerTrustGroupsClient) ListByInstanceComplete(ctx context.Context, id commonids.SqlManagedInstanceId) (ListByInstanceCompleteResult, error) {
	return c.ListByInstanceCompleteMatchingPredicate(ctx, id, ServerTrustGroupOperationPredicate{})
}

// ListByInstanceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServerTrustGroupsClient) ListByInstanceCompleteMatchingPredicate(ctx context.Context, id commonids.SqlManagedInstanceId, predicate ServerTrustGroupOperationPredicate) (result ListByInstanceCompleteResult, err error) {
	items := make([]ServerTrustGroup, 0)

	resp, err := c.ListByInstance(ctx, id)
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

	result = ListByInstanceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
