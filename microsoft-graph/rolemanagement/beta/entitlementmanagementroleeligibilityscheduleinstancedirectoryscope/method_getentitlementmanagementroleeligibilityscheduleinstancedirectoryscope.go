package entitlementmanagementroleeligibilityscheduleinstancedirectoryscope

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.DirectoryObject
}

type GetEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeOperationOptions() GetEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeOperationOptions {
	return GetEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeOperationOptions{}
}

func (o GetEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScope - Get directoryScope from roleManagement. The
// directory object that is the scope of the assignment or role eligibility. Read-only.
func (c EntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeClient) GetEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScope(ctx context.Context, id beta.RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId, options GetEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeOperationOptions) (result GetEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/directoryScope", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalDirectoryObjectImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
