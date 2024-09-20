package lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresultsubjectmailboxsetting

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

type UpdateLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskSubjectMailboxSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskSubjectMailboxSettingOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskSubjectMailboxSettingOperationOptions() UpdateLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskSubjectMailboxSettingOperationOptions {
	return UpdateLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskSubjectMailboxSettingOperationOptions{}
}

func (o UpdateLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskSubjectMailboxSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskSubjectMailboxSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskSubjectMailboxSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskSubjectMailboxSetting - Update property
// mailboxSettings value.
func (c LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient) UpdateLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskSubjectMailboxSetting(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultId, input beta.MailboxSettings, options UpdateLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskSubjectMailboxSettingOperationOptions) (result UpdateLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskSubjectMailboxSettingOperationResponse, err error) {
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
