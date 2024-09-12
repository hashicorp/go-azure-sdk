package deviceregistrationpolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeviceRegistrationPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceRegistrationPolicy
}

type GetDeviceRegistrationPolicyOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetDeviceRegistrationPolicyOperationOptions() GetDeviceRegistrationPolicyOperationOptions {
	return GetDeviceRegistrationPolicyOperationOptions{}
}

func (o GetDeviceRegistrationPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeviceRegistrationPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDeviceRegistrationPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDeviceRegistrationPolicy - Get deviceRegistrationPolicy. Read the properties and relationships of a
// deviceRegistrationPolicy object. Represents deviceRegistrationPolicy quota restrictions, additional authentication,
// and authorization policies to register device identities to your organization.
func (c DeviceRegistrationPolicyClient) GetDeviceRegistrationPolicy(ctx context.Context, options GetDeviceRegistrationPolicyOperationOptions) (result GetDeviceRegistrationPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/policies/deviceRegistrationPolicy",
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

	var model beta.DeviceRegistrationPolicy
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
