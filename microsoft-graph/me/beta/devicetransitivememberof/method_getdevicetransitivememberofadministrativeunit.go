package devicetransitivememberof

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

type GetDeviceTransitiveMemberOfAdministrativeUnitOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AdministrativeUnit
}

type GetDeviceTransitiveMemberOfAdministrativeUnitOperationOptions struct {
	ConsistencyLevel *odata.ConsistencyLevel
	Expand           *odata.Expand
	Metadata         *odata.Metadata
	RetryFunc        client.RequestRetryFunc
	Select           *[]string
}

func DefaultGetDeviceTransitiveMemberOfAdministrativeUnitOperationOptions() GetDeviceTransitiveMemberOfAdministrativeUnitOperationOptions {
	return GetDeviceTransitiveMemberOfAdministrativeUnitOperationOptions{}
}

func (o GetDeviceTransitiveMemberOfAdministrativeUnitOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeviceTransitiveMemberOfAdministrativeUnitOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.ConsistencyLevel != nil {
		out.ConsistencyLevel = *o.ConsistencyLevel
	}
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

func (o GetDeviceTransitiveMemberOfAdministrativeUnitOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDeviceTransitiveMemberOfAdministrativeUnit - Get the item of type microsoft.graph.directoryObject as
// microsoft.graph.administrativeUnit
func (c DeviceTransitiveMemberOfClient) GetDeviceTransitiveMemberOfAdministrativeUnit(ctx context.Context, id beta.MeDeviceIdTransitiveMemberOfId, options GetDeviceTransitiveMemberOfAdministrativeUnitOperationOptions) (result GetDeviceTransitiveMemberOfAdministrativeUnitOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/administrativeUnit", id.ID()),
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

	var model beta.AdministrativeUnit
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
