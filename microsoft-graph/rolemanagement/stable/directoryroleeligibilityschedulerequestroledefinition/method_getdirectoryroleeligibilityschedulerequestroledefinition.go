package directoryroleeligibilityschedulerequestroledefinition

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

type GetDirectoryRoleEligibilityScheduleRequestRoleDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UnifiedRoleDefinition
}

type GetDirectoryRoleEligibilityScheduleRequestRoleDefinitionOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetDirectoryRoleEligibilityScheduleRequestRoleDefinitionOperationOptions() GetDirectoryRoleEligibilityScheduleRequestRoleDefinitionOperationOptions {
	return GetDirectoryRoleEligibilityScheduleRequestRoleDefinitionOperationOptions{}
}

func (o GetDirectoryRoleEligibilityScheduleRequestRoleDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDirectoryRoleEligibilityScheduleRequestRoleDefinitionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDirectoryRoleEligibilityScheduleRequestRoleDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDirectoryRoleEligibilityScheduleRequestRoleDefinition - Get roleDefinition from roleManagement. Detailed
// information for the unifiedRoleDefinition object that is referenced through the roleDefinitionId property. Supports
// $expand.
func (c DirectoryRoleEligibilityScheduleRequestRoleDefinitionClient) GetDirectoryRoleEligibilityScheduleRequestRoleDefinition(ctx context.Context, id stable.RoleManagementDirectoryRoleEligibilityScheduleRequestId, options GetDirectoryRoleEligibilityScheduleRequestRoleDefinitionOperationOptions) (result GetDirectoryRoleEligibilityScheduleRequestRoleDefinitionOperationResponse, err error) {
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
