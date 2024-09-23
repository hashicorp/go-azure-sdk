package enterpriseapproleeligibilityschedulerequest

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

type CreateEnterpriseAppRoleEligibilityScheduleRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UnifiedRoleEligibilityScheduleRequest
}

type CreateEnterpriseAppRoleEligibilityScheduleRequestOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEnterpriseAppRoleEligibilityScheduleRequestOperationOptions() CreateEnterpriseAppRoleEligibilityScheduleRequestOperationOptions {
	return CreateEnterpriseAppRoleEligibilityScheduleRequestOperationOptions{}
}

func (o CreateEnterpriseAppRoleEligibilityScheduleRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEnterpriseAppRoleEligibilityScheduleRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEnterpriseAppRoleEligibilityScheduleRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEnterpriseAppRoleEligibilityScheduleRequest - Create new navigation property to roleEligibilityScheduleRequests
// for roleManagement
func (c EnterpriseAppRoleEligibilityScheduleRequestClient) CreateEnterpriseAppRoleEligibilityScheduleRequest(ctx context.Context, id beta.RoleManagementEnterpriseAppId, input beta.UnifiedRoleEligibilityScheduleRequest, options CreateEnterpriseAppRoleEligibilityScheduleRequestOperationOptions) (result CreateEnterpriseAppRoleEligibilityScheduleRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/roleEligibilityScheduleRequests", id.ID()),
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

	var model beta.UnifiedRoleEligibilityScheduleRequest
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
