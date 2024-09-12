package directoryroleeligibilityscheduleinstanceroledefinition

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

type GetDirectoryRoleEligibilityScheduleInstanceRoleDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UnifiedRoleDefinition
}

type GetDirectoryRoleEligibilityScheduleInstanceRoleDefinitionOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetDirectoryRoleEligibilityScheduleInstanceRoleDefinitionOperationOptions() GetDirectoryRoleEligibilityScheduleInstanceRoleDefinitionOperationOptions {
	return GetDirectoryRoleEligibilityScheduleInstanceRoleDefinitionOperationOptions{}
}

func (o GetDirectoryRoleEligibilityScheduleInstanceRoleDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDirectoryRoleEligibilityScheduleInstanceRoleDefinitionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDirectoryRoleEligibilityScheduleInstanceRoleDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDirectoryRoleEligibilityScheduleInstanceRoleDefinition - Get roleDefinition from roleManagement. Detailed
// information for the roleDefinition object that is referenced through the roleDefinitionId property.
func (c DirectoryRoleEligibilityScheduleInstanceRoleDefinitionClient) GetDirectoryRoleEligibilityScheduleInstanceRoleDefinition(ctx context.Context, id stable.RoleManagementDirectoryRoleEligibilityScheduleInstanceId, options GetDirectoryRoleEligibilityScheduleInstanceRoleDefinitionOperationOptions) (result GetDirectoryRoleEligibilityScheduleInstanceRoleDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/roleDefinition", id.ID()),
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

	var model stable.UnifiedRoleDefinition
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
