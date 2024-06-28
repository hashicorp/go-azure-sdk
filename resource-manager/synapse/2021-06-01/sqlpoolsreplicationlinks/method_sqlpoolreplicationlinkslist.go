package sqlpoolsreplicationlinks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolReplicationLinksListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ReplicationLink
}

type SqlPoolReplicationLinksListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ReplicationLink
}

type SqlPoolReplicationLinksListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *SqlPoolReplicationLinksListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// SqlPoolReplicationLinksList ...
func (c SqlPoolsReplicationLinksClient) SqlPoolReplicationLinksList(ctx context.Context, id SqlPoolId) (result SqlPoolReplicationLinksListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &SqlPoolReplicationLinksListCustomPager{},
		Path:       fmt.Sprintf("%s/replicationLinks", id.ID()),
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
		Values *[]ReplicationLink `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SqlPoolReplicationLinksListComplete retrieves all the results into a single object
func (c SqlPoolsReplicationLinksClient) SqlPoolReplicationLinksListComplete(ctx context.Context, id SqlPoolId) (SqlPoolReplicationLinksListCompleteResult, error) {
	return c.SqlPoolReplicationLinksListCompleteMatchingPredicate(ctx, id, ReplicationLinkOperationPredicate{})
}

// SqlPoolReplicationLinksListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SqlPoolsReplicationLinksClient) SqlPoolReplicationLinksListCompleteMatchingPredicate(ctx context.Context, id SqlPoolId, predicate ReplicationLinkOperationPredicate) (result SqlPoolReplicationLinksListCompleteResult, err error) {
	items := make([]ReplicationLink, 0)

	resp, err := c.SqlPoolReplicationLinksList(ctx, id)
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

	result = SqlPoolReplicationLinksListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
