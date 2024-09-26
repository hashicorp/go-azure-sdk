package lifecycleworkflowworkflowtemplatetasktaskprocessingresultsubjectmailboxsetting

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

type UpdateLifecycleWorkflowTemplateTaskProcessingResultSubjectMailboxSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateLifecycleWorkflowTemplateTaskProcessingResultSubjectMailboxSettingOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateLifecycleWorkflowTemplateTaskProcessingResultSubjectMailboxSettingOperationOptions() UpdateLifecycleWorkflowTemplateTaskProcessingResultSubjectMailboxSettingOperationOptions {
	return UpdateLifecycleWorkflowTemplateTaskProcessingResultSubjectMailboxSettingOperationOptions{}
}

func (o UpdateLifecycleWorkflowTemplateTaskProcessingResultSubjectMailboxSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateLifecycleWorkflowTemplateTaskProcessingResultSubjectMailboxSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateLifecycleWorkflowTemplateTaskProcessingResultSubjectMailboxSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateLifecycleWorkflowTemplateTaskProcessingResultSubjectMailboxSetting - Update property mailboxSettings value.
func (c LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubjectMailboxSettingClient) UpdateLifecycleWorkflowTemplateTaskProcessingResultSubjectMailboxSetting(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultId, input beta.MailboxSettings, options UpdateLifecycleWorkflowTemplateTaskProcessingResultSubjectMailboxSettingOperationOptions) (result UpdateLifecycleWorkflowTemplateTaskProcessingResultSubjectMailboxSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/subject/mailboxSettings", id.ID()),
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
