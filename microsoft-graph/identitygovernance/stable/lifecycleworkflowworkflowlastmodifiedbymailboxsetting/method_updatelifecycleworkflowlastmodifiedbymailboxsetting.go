package lifecycleworkflowworkflowlastmodifiedbymailboxsetting

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

type UpdateLifecycleWorkflowLastModifiedByMailboxSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateLifecycleWorkflowLastModifiedByMailboxSettingOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateLifecycleWorkflowLastModifiedByMailboxSettingOperationOptions() UpdateLifecycleWorkflowLastModifiedByMailboxSettingOperationOptions {
	return UpdateLifecycleWorkflowLastModifiedByMailboxSettingOperationOptions{}
}

func (o UpdateLifecycleWorkflowLastModifiedByMailboxSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateLifecycleWorkflowLastModifiedByMailboxSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateLifecycleWorkflowLastModifiedByMailboxSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateLifecycleWorkflowLastModifiedByMailboxSetting - Update property mailboxSettings value.
func (c LifecycleWorkflowWorkflowLastModifiedByMailboxSettingClient) UpdateLifecycleWorkflowLastModifiedByMailboxSetting(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowWorkflowId, input stable.MailboxSettings, options UpdateLifecycleWorkflowLastModifiedByMailboxSettingOperationOptions) (result UpdateLifecycleWorkflowLastModifiedByMailboxSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/lastModifiedBy/mailboxSettings", id.ID()),
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
