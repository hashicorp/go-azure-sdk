package lifecycleworkflowdeleteditemworkflowcreatedbymailboxsetting

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateLifecycleWorkflowDeletedItemWorkflowCreatedByMailboxSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateLifecycleWorkflowDeletedItemWorkflowCreatedByMailboxSettingOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateLifecycleWorkflowDeletedItemWorkflowCreatedByMailboxSettingOperationOptions() UpdateLifecycleWorkflowDeletedItemWorkflowCreatedByMailboxSettingOperationOptions {
	return UpdateLifecycleWorkflowDeletedItemWorkflowCreatedByMailboxSettingOperationOptions{}
}

func (o UpdateLifecycleWorkflowDeletedItemWorkflowCreatedByMailboxSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateLifecycleWorkflowDeletedItemWorkflowCreatedByMailboxSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateLifecycleWorkflowDeletedItemWorkflowCreatedByMailboxSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateLifecycleWorkflowDeletedItemWorkflowCreatedByMailboxSetting - Update property mailboxSettings value.
func (c LifecycleWorkflowDeletedItemWorkflowCreatedByMailboxSettingClient) UpdateLifecycleWorkflowDeletedItemWorkflowCreatedByMailboxSetting(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId, input beta.MailboxSettings, options UpdateLifecycleWorkflowDeletedItemWorkflowCreatedByMailboxSettingOperationOptions) (result UpdateLifecycleWorkflowDeletedItemWorkflowCreatedByMailboxSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/createdBy/mailboxSettings", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	return
}
