package mobiledevicemanagementpolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMobileDeviceManagementPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.MobilityManagementPolicy
}

type CreateMobileDeviceManagementPolicyOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateMobileDeviceManagementPolicyOperationOptions() CreateMobileDeviceManagementPolicyOperationOptions {
	return CreateMobileDeviceManagementPolicyOperationOptions{}
}

func (o CreateMobileDeviceManagementPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateMobileDeviceManagementPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateMobileDeviceManagementPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateMobileDeviceManagementPolicy - Create new navigation property to mobileDeviceManagementPolicies for policies
func (c MobileDeviceManagementPolicyClient) CreateMobileDeviceManagementPolicy(ctx context.Context, input beta.MobilityManagementPolicy, options CreateMobileDeviceManagementPolicyOperationOptions) (result CreateMobileDeviceManagementPolicyOperationResponse, err error) {
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
		Path:          "/policies/mobileDeviceManagementPolicies",
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

	var model beta.MobilityManagementPolicy
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
