package lifecycleworkflowworkflowuserprocessingresultsubjectmailboxsetting

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

type UpdateLifecycleWorkflowUserProcessingResultSubjectMailboxSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateLifecycleWorkflowUserProcessingResultSubjectMailboxSettingOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateLifecycleWorkflowUserProcessingResultSubjectMailboxSettingOperationOptions() UpdateLifecycleWorkflowUserProcessingResultSubjectMailboxSettingOperationOptions {
	return UpdateLifecycleWorkflowUserProcessingResultSubjectMailboxSettingOperationOptions{}
}

func (o UpdateLifecycleWorkflowUserProcessingResultSubjectMailboxSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateLifecycleWorkflowUserProcessingResultSubjectMailboxSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateLifecycleWorkflowUserProcessingResultSubjectMailboxSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateLifecycleWorkflowUserProcessingResultSubjectMailboxSetting - Update property mailboxSettings value.
func (c LifecycleWorkflowWorkflowUserProcessingResultSubjectMailboxSettingClient) UpdateLifecycleWorkflowUserProcessingResultSubjectMailboxSetting(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultId, input beta.MailboxSettings, options UpdateLifecycleWorkflowUserProcessingResultSubjectMailboxSettingOperationOptions) (result UpdateLifecycleWorkflowUserProcessingResultSubjectMailboxSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/subject/mailboxSettings", id.ID()),
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
