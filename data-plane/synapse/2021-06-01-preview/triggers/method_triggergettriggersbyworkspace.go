package triggers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggerGetTriggersByWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]TriggerResource
}

type TriggerGetTriggersByWorkspaceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []TriggerResource
}

type TriggerGetTriggersByWorkspaceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *TriggerGetTriggersByWorkspaceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// TriggerGetTriggersByWorkspace ...
func (c TriggersClient) TriggerGetTriggersByWorkspace(ctx context.Context) (result TriggerGetTriggersByWorkspaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &TriggerGetTriggersByWorkspaceCustomPager{},
		Path:       "/triggers",
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
		Values *[]TriggerResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// TriggerGetTriggersByWorkspaceComplete retrieves all the results into a single object
func (c TriggersClient) TriggerGetTriggersByWorkspaceComplete(ctx context.Context) (TriggerGetTriggersByWorkspaceCompleteResult, error) {
	return c.TriggerGetTriggersByWorkspaceCompleteMatchingPredicate(ctx, TriggerResourceOperationPredicate{})
}

// TriggerGetTriggersByWorkspaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TriggersClient) TriggerGetTriggersByWorkspaceCompleteMatchingPredicate(ctx context.Context, predicate TriggerResourceOperationPredicate) (result TriggerGetTriggersByWorkspaceCompleteResult, err error) {
	items := make([]TriggerResource, 0)

	resp, err := c.TriggerGetTriggersByWorkspace(ctx)
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

	result = TriggerGetTriggersByWorkspaceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
