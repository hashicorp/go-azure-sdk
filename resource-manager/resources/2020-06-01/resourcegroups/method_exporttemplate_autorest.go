package resourcegroups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExportTemplateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ExportTemplate ...
func (c ResourceGroupsClient) ExportTemplate(ctx context.Context, id commonids.ResourceGroupId, input ExportTemplateRequest) (result ExportTemplateOperationResponse, err error) {
	req, err := c.preparerForExportTemplate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcegroups.ResourceGroupsClient", "ExportTemplate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForExportTemplate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcegroups.ResourceGroupsClient", "ExportTemplate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ExportTemplateThenPoll performs ExportTemplate then polls until it's completed
func (c ResourceGroupsClient) ExportTemplateThenPoll(ctx context.Context, id commonids.ResourceGroupId, input ExportTemplateRequest) error {
	result, err := c.ExportTemplate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ExportTemplate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ExportTemplate: %+v", err)
	}

	return nil
}

// preparerForExportTemplate prepares the ExportTemplate request.
func (c ResourceGroupsClient) preparerForExportTemplate(ctx context.Context, id commonids.ResourceGroupId, input ExportTemplateRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/exportTemplate", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForExportTemplate sends the ExportTemplate request. The method will close the
// http.Response Body if it receives an error.
func (c ResourceGroupsClient) senderForExportTemplate(ctx context.Context, req *http.Request) (future ExportTemplateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}
