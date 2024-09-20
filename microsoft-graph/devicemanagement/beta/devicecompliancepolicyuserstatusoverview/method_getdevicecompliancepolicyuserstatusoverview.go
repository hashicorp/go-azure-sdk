package devicecompliancepolicyuserstatusoverview

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

type GetDeviceCompliancePolicyUserStatusOverviewOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceComplianceUserOverview
}

type GetDeviceCompliancePolicyUserStatusOverviewOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetDeviceCompliancePolicyUserStatusOverviewOperationOptions() GetDeviceCompliancePolicyUserStatusOverviewOperationOptions {
	return GetDeviceCompliancePolicyUserStatusOverviewOperationOptions{}
}

func (o GetDeviceCompliancePolicyUserStatusOverviewOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeviceCompliancePolicyUserStatusOverviewOperationOptions) ToOData() *odata.Query {
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

func (o GetDeviceCompliancePolicyUserStatusOverviewOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDeviceCompliancePolicyUserStatusOverview - Get userStatusOverview from deviceManagement. Device compliance users
// status overview
func (c DeviceCompliancePolicyUserStatusOverviewClient) GetDeviceCompliancePolicyUserStatusOverview(ctx context.Context, id beta.DeviceManagementDeviceCompliancePolicyId, options GetDeviceCompliancePolicyUserStatusOverviewOperationOptions) (result GetDeviceCompliancePolicyUserStatusOverviewOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/userStatusOverview", id.ID()),
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

	var model beta.DeviceComplianceUserOverview
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
