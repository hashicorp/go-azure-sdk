package kqlscripts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KqlScriptCreateOrUpdateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *KqlScriptResource
}

// KqlScriptCreateOrUpdate ...
func (c KqlScriptsClient) KqlScriptCreateOrUpdate(ctx context.Context, id KqlScriptId, input KqlScriptResource) (result KqlScriptCreateOrUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,
		Path:       id.Path(),
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

	result.Poller, err = dataplane.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// KqlScriptCreateOrUpdateThenPoll performs KqlScriptCreateOrUpdate then polls until it's completed
func (c KqlScriptsClient) KqlScriptCreateOrUpdateThenPoll(ctx context.Context, id KqlScriptId, input KqlScriptResource) error {
	result, err := c.KqlScriptCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing KqlScriptCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after KqlScriptCreateOrUpdate: %+v", err)
	}

	return nil
}
