package activesessionhostconfiguration

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByHostPoolOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ActiveSessionHostConfiguration
}

type ListByHostPoolCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ActiveSessionHostConfiguration
}

type ListByHostPoolCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByHostPoolCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByHostPool ...
func (c ActiveSessionHostConfigurationClient) ListByHostPool(ctx context.Context, id HostPoolId) (result ListByHostPoolOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByHostPoolCustomPager{},
		Path:       fmt.Sprintf("%s/activeSessionHostConfigurations", id.ID()),
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
		Values *[]ActiveSessionHostConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByHostPoolComplete retrieves all the results into a single object
func (c ActiveSessionHostConfigurationClient) ListByHostPoolComplete(ctx context.Context, id HostPoolId) (ListByHostPoolCompleteResult, error) {
	return c.ListByHostPoolCompleteMatchingPredicate(ctx, id, ActiveSessionHostConfigurationOperationPredicate{})
}

// ListByHostPoolCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ActiveSessionHostConfigurationClient) ListByHostPoolCompleteMatchingPredicate(ctx context.Context, id HostPoolId, predicate ActiveSessionHostConfigurationOperationPredicate) (result ListByHostPoolCompleteResult, err error) {
	items := make([]ActiveSessionHostConfiguration, 0)

	resp, err := c.ListByHostPool(ctx, id)
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

	result = ListByHostPoolCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
