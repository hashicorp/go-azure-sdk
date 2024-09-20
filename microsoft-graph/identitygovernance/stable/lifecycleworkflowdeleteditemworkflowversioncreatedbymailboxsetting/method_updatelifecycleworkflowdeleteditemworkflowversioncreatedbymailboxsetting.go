package lifecycleworkflowdeleteditemworkflowversioncreatedbymailboxsetting

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateLifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateLifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSettingOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateLifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSettingOperationOptions() UpdateLifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSettingOperationOptions {
	return UpdateLifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSettingOperationOptions{}
}

func (o UpdateLifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateLifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateLifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateLifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSetting - Update property mailboxSettings value.
func (c LifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSettingClient) UpdateLifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSetting(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionId, input stable.MailboxSettings, options UpdateLifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSettingOperationOptions) (result UpdateLifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/createdBy/mailboxSettings", id.ID()),
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
