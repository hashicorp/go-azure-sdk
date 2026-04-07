package managednetwork

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutboundRulesPostOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]OutboundRuleBasicResource
}

type OutboundRulesPostCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []OutboundRuleBasicResource
}

type OutboundRulesPostCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *OutboundRulesPostCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// OutboundRulesPost ...
func (c ManagedNetworkClient) OutboundRulesPost(ctx context.Context, id ManagedNetworkId, input ManagedNetworkSettingsBasicResource) (result OutboundRulesPostOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &OutboundRulesPostCustomPager{},
		Path:       fmt.Sprintf("%s/batchOutboundRules", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// OutboundRulesPostThenPoll performs OutboundRulesPost then polls until it's completed
func (c ManagedNetworkClient) OutboundRulesPostThenPoll(ctx context.Context, id ManagedNetworkId, input ManagedNetworkSettingsBasicResource) error {
	result, err := c.OutboundRulesPost(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing OutboundRulesPost: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after OutboundRulesPost: %+v", err)
	}

	return nil
}
