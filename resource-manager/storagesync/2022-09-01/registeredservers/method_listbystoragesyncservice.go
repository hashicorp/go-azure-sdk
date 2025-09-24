package registeredservers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByStorageSyncServiceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RegisteredServer
}

type ListByStorageSyncServiceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []RegisteredServer
}

type ListByStorageSyncServiceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByStorageSyncServiceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByStorageSyncService ...
func (c RegisteredServersClient) ListByStorageSyncService(ctx context.Context, id StorageSyncServiceId) (result ListByStorageSyncServiceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByStorageSyncServiceCustomPager{},
		Path:       fmt.Sprintf("%s/registeredServers", id.ID()),
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
		Values *[]RegisteredServer `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByStorageSyncServiceComplete retrieves all the results into a single object
func (c RegisteredServersClient) ListByStorageSyncServiceComplete(ctx context.Context, id StorageSyncServiceId) (ListByStorageSyncServiceCompleteResult, error) {
	return c.ListByStorageSyncServiceCompleteMatchingPredicate(ctx, id, RegisteredServerOperationPredicate{})
}

// ListByStorageSyncServiceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RegisteredServersClient) ListByStorageSyncServiceCompleteMatchingPredicate(ctx context.Context, id StorageSyncServiceId, predicate RegisteredServerOperationPredicate) (result ListByStorageSyncServiceCompleteResult, err error) {
	items := make([]RegisteredServer, 0)

	resp, err := c.ListByStorageSyncService(ctx, id)
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

	result = ListByStorageSyncServiceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
