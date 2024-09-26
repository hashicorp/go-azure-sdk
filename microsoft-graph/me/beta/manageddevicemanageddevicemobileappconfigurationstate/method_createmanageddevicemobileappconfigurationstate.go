package manageddevicemanageddevicemobileappconfigurationstate

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

type CreateManagedDeviceMobileAppConfigurationStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ManagedDeviceMobileAppConfigurationState
}

type CreateManagedDeviceMobileAppConfigurationStateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateManagedDeviceMobileAppConfigurationStateOperationOptions() CreateManagedDeviceMobileAppConfigurationStateOperationOptions {
	return CreateManagedDeviceMobileAppConfigurationStateOperationOptions{}
}

func (o CreateManagedDeviceMobileAppConfigurationStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateManagedDeviceMobileAppConfigurationStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateManagedDeviceMobileAppConfigurationStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateManagedDeviceMobileAppConfigurationState - Create new navigation property to
// managedDeviceMobileAppConfigurationStates for me
func (c ManagedDeviceManagedDeviceMobileAppConfigurationStateClient) CreateManagedDeviceMobileAppConfigurationState(ctx context.Context, id beta.MeManagedDeviceId, input beta.ManagedDeviceMobileAppConfigurationState, options CreateManagedDeviceMobileAppConfigurationStateOperationOptions) (result CreateManagedDeviceMobileAppConfigurationStateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/managedDeviceMobileAppConfigurationStates", id.ID()),
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

	var model beta.ManagedDeviceMobileAppConfigurationState
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
