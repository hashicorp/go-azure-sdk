package parametergroups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListConnectionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ParameterGroupConnectionProperties
}

type ListConnectionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ParameterGroupConnectionProperties
}

type ListConnectionsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListConnectionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConnections ...
func (c ParameterGroupsClient) ListConnections(ctx context.Context, id ProviderParameterGroupId) (result ListConnectionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConnectionsCustomPager{},
		Path:       fmt.Sprintf("%s/connections", id.ID()),
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
		Values *[]ParameterGroupConnectionProperties `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConnectionsComplete retrieves all the results into a single object
func (c ParameterGroupsClient) ListConnectionsComplete(ctx context.Context, id ProviderParameterGroupId) (ListConnectionsCompleteResult, error) {
	return c.ListConnectionsCompleteMatchingPredicate(ctx, id, ParameterGroupConnectionPropertiesOperationPredicate{})
}

// ListConnectionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ParameterGroupsClient) ListConnectionsCompleteMatchingPredicate(ctx context.Context, id ProviderParameterGroupId, predicate ParameterGroupConnectionPropertiesOperationPredicate) (result ListConnectionsCompleteResult, err error) {
	items := make([]ParameterGroupConnectionProperties, 0)

	resp, err := c.ListConnections(ctx, id)
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

	result = ListConnectionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
