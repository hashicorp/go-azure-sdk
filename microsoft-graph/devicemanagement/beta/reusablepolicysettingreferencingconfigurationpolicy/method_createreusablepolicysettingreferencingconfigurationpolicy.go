package reusablepolicysettingreferencingconfigurationpolicy

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

type CreateReusablePolicySettingReferencingConfigurationPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceManagementConfigurationPolicy
}

type CreateReusablePolicySettingReferencingConfigurationPolicyOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateReusablePolicySettingReferencingConfigurationPolicyOperationOptions() CreateReusablePolicySettingReferencingConfigurationPolicyOperationOptions {
	return CreateReusablePolicySettingReferencingConfigurationPolicyOperationOptions{}
}

func (o CreateReusablePolicySettingReferencingConfigurationPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateReusablePolicySettingReferencingConfigurationPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateReusablePolicySettingReferencingConfigurationPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateReusablePolicySettingReferencingConfigurationPolicy - Create new navigation property to
// referencingConfigurationPolicies for deviceManagement
func (c ReusablePolicySettingReferencingConfigurationPolicyClient) CreateReusablePolicySettingReferencingConfigurationPolicy(ctx context.Context, id beta.DeviceManagementReusablePolicySettingId, input beta.DeviceManagementConfigurationPolicy, options CreateReusablePolicySettingReferencingConfigurationPolicyOperationOptions) (result CreateReusablePolicySettingReferencingConfigurationPolicyOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/referencingConfigurationPolicies", id.ID()),
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

	var model beta.DeviceManagementConfigurationPolicy
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
