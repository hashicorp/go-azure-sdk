package workflowenvelopes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListInstanceWorkflowsSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WorkflowEnvelope
}

type WebAppsListInstanceWorkflowsSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WorkflowEnvelope
}

type WebAppsListInstanceWorkflowsSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListInstanceWorkflowsSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListInstanceWorkflowsSlot ...
func (c WorkflowEnvelopesClient) WebAppsListInstanceWorkflowsSlot(ctx context.Context, id SlotId) (result WebAppsListInstanceWorkflowsSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListInstanceWorkflowsSlotCustomPager{},
		Path:       fmt.Sprintf("%s/workflows", id.ID()),
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
		Values *[]WorkflowEnvelope `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListInstanceWorkflowsSlotComplete retrieves all the results into a single object
func (c WorkflowEnvelopesClient) WebAppsListInstanceWorkflowsSlotComplete(ctx context.Context, id SlotId) (WebAppsListInstanceWorkflowsSlotCompleteResult, error) {
	return c.WebAppsListInstanceWorkflowsSlotCompleteMatchingPredicate(ctx, id, WorkflowEnvelopeOperationPredicate{})
}

// WebAppsListInstanceWorkflowsSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WorkflowEnvelopesClient) WebAppsListInstanceWorkflowsSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate WorkflowEnvelopeOperationPredicate) (result WebAppsListInstanceWorkflowsSlotCompleteResult, err error) {
	items := make([]WorkflowEnvelope, 0)

	resp, err := c.WebAppsListInstanceWorkflowsSlot(ctx, id)
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

	result = WebAppsListInstanceWorkflowsSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
