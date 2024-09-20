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

type CreateHardwarePasswordInfoOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.HardwarePasswordInfo
}

type CreateHardwarePasswordInfoOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateHardwarePasswordInfoOperationOptions() CreateHardwarePasswordInfoOperationOptions {
	return CreateHardwarePasswordInfoOperationOptions{}
}

func (o CreateHardwarePasswordInfoOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateHardwarePasswordInfoOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateHardwarePasswordInfoOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateHardwarePasswordInfo - Create new navigation property to hardwarePasswordInfo for deviceManagement
func (c HardwarePasswordInfoClient) CreateHardwarePasswordInfo(ctx context.Context, input beta.HardwarePasswordInfo, options CreateHardwarePasswordInfoOperationOptions) (result CreateHardwarePasswordInfoOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/hardwarePasswordInfo",
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

	var model beta.HardwarePasswordInfo
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
