package employeeexperienceassignedrolememberusermailboxsetting

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

type UpdateEmployeeExperienceAssignedRoleMemberUserMailboxSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEmployeeExperienceAssignedRoleMemberUserMailboxSettingOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEmployeeExperienceAssignedRoleMemberUserMailboxSettingOperationOptions() UpdateEmployeeExperienceAssignedRoleMemberUserMailboxSettingOperationOptions {
	return UpdateEmployeeExperienceAssignedRoleMemberUserMailboxSettingOperationOptions{}
}

func (o UpdateEmployeeExperienceAssignedRoleMemberUserMailboxSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEmployeeExperienceAssignedRoleMemberUserMailboxSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEmployeeExperienceAssignedRoleMemberUserMailboxSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEmployeeExperienceAssignedRoleMemberUserMailboxSetting - Update property mailboxSettings value.
func (c EmployeeExperienceAssignedRoleMemberUserMailboxSettingClient) UpdateEmployeeExperienceAssignedRoleMemberUserMailboxSetting(ctx context.Context, id beta.UserIdEmployeeExperienceAssignedRoleIdMemberId, input beta.MailboxSettings, options UpdateEmployeeExperienceAssignedRoleMemberUserMailboxSettingOperationOptions) (result UpdateEmployeeExperienceAssignedRoleMemberUserMailboxSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/user/mailboxSettings", id.ID()),
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
