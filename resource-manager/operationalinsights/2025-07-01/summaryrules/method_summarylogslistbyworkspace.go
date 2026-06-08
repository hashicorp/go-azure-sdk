package summaryrules

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SummaryLogsListByWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SummaryLogs
}

type SummaryLogsListByWorkspaceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SummaryLogs
}

type SummaryLogsListByWorkspaceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *SummaryLogsListByWorkspaceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// SummaryLogsListByWorkspace ...
func (c SummaryRulesClient) SummaryLogsListByWorkspace(ctx context.Context, id WorkspaceId) (result SummaryLogsListByWorkspaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &SummaryLogsListByWorkspaceCustomPager{},
		Path:       fmt.Sprintf("%s/summaryLogs", id.ID()),
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
		Values *[]SummaryLogs `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SummaryLogsListByWorkspaceComplete retrieves all the results into a single object
func (c SummaryRulesClient) SummaryLogsListByWorkspaceComplete(ctx context.Context, id WorkspaceId) (SummaryLogsListByWorkspaceCompleteResult, error) {
	return c.SummaryLogsListByWorkspaceCompleteMatchingPredicate(ctx, id, SummaryLogsOperationPredicate{})
}

// SummaryLogsListByWorkspaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SummaryRulesClient) SummaryLogsListByWorkspaceCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate SummaryLogsOperationPredicate) (result SummaryLogsListByWorkspaceCompleteResult, err error) {
	items := make([]SummaryLogs, 0)

	resp, err := c.SummaryLogsListByWorkspace(ctx, id)
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

	result = SummaryLogsListByWorkspaceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
