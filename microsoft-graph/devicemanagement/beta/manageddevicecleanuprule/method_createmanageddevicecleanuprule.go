package manageddevicecleanuprule

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateManagedDeviceCleanupRuleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ManagedDeviceCleanupRule
}

type CreateManagedDeviceCleanupRuleOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateManagedDeviceCleanupRuleOperationOptions() CreateManagedDeviceCleanupRuleOperationOptions {
	return CreateManagedDeviceCleanupRuleOperationOptions{}
}

func (o CreateManagedDeviceCleanupRuleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateManagedDeviceCleanupRuleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateManagedDeviceCleanupRuleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateManagedDeviceCleanupRule - Create new navigation property to managedDeviceCleanupRules for deviceManagement
func (c ManagedDeviceCleanupRuleClient) CreateManagedDeviceCleanupRule(ctx context.Context, input beta.ManagedDeviceCleanupRule, options CreateManagedDeviceCleanupRuleOperationOptions) (result CreateManagedDeviceCleanupRuleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/managedDeviceCleanupRules",
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

	var model beta.ManagedDeviceCleanupRule
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
