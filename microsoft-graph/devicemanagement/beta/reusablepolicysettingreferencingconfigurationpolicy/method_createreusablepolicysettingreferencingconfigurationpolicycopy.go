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

type CreateReusablePolicySettingReferencingConfigurationPolicyCopyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceManagementConfigurationPolicy
}

type CreateReusablePolicySettingReferencingConfigurationPolicyCopyOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateReusablePolicySettingReferencingConfigurationPolicyCopyOperationOptions() CreateReusablePolicySettingReferencingConfigurationPolicyCopyOperationOptions {
	return CreateReusablePolicySettingReferencingConfigurationPolicyCopyOperationOptions{}
}

func (o CreateReusablePolicySettingReferencingConfigurationPolicyCopyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateReusablePolicySettingReferencingConfigurationPolicyCopyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateReusablePolicySettingReferencingConfigurationPolicyCopyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateReusablePolicySettingReferencingConfigurationPolicyCopy - Invoke action createCopy
func (c ReusablePolicySettingReferencingConfigurationPolicyClient) CreateReusablePolicySettingReferencingConfigurationPolicyCopy(ctx context.Context, id beta.DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId, input CreateReusablePolicySettingReferencingConfigurationPolicyCopyRequest, options CreateReusablePolicySettingReferencingConfigurationPolicyCopyOperationOptions) (result CreateReusablePolicySettingReferencingConfigurationPolicyCopyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/createCopy", id.ID()),
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
