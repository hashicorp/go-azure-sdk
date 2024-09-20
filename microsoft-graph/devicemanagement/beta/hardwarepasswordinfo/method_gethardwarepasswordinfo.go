package hardwarepasswordinfo

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetHardwarePasswordInfoOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.HardwarePasswordInfo
}

type GetHardwarePasswordInfoOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetHardwarePasswordInfoOperationOptions() GetHardwarePasswordInfoOperationOptions {
	return GetHardwarePasswordInfoOperationOptions{}
}

func (o GetHardwarePasswordInfoOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetHardwarePasswordInfoOperationOptions) ToOData() *odata.Query {
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

func (o GetHardwarePasswordInfoOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetHardwarePasswordInfo - Get hardwarePasswordInfo from deviceManagement. Intune will provide customer the ability to
// configure hardware/bios settings on the enrolled windows 10 Azure Active Directory joined devices. Starting from
// June, 2024 (Intune Release 2406), this type will no longer be supported and will be marked as deprecated
func (c HardwarePasswordInfoClient) GetHardwarePasswordInfo(ctx context.Context, id beta.DeviceManagementHardwarePasswordInfoId, options GetHardwarePasswordInfoOperationOptions) (result GetHardwarePasswordInfoOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model beta.HardwarePasswordInfo
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
