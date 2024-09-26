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

type ClearReusablePolicySettingReferencingConfigurationPolicyEnrollmentTimeDeviceMembershipTargetOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *ClearReusablePolicySettingReferencingConfigurationPolicyEnrollmentTimeDeviceMembershipTargetResult
}

type ClearReusablePolicySettingReferencingConfigurationPolicyEnrollmentTimeDeviceMembershipTargetOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultClearReusablePolicySettingReferencingConfigurationPolicyEnrollmentTimeDeviceMembershipTargetOperationOptions() ClearReusablePolicySettingReferencingConfigurationPolicyEnrollmentTimeDeviceMembershipTargetOperationOptions {
	return ClearReusablePolicySettingReferencingConfigurationPolicyEnrollmentTimeDeviceMembershipTargetOperationOptions{}
}

func (o ClearReusablePolicySettingReferencingConfigurationPolicyEnrollmentTimeDeviceMembershipTargetOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ClearReusablePolicySettingReferencingConfigurationPolicyEnrollmentTimeDeviceMembershipTargetOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ClearReusablePolicySettingReferencingConfigurationPolicyEnrollmentTimeDeviceMembershipTargetOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ClearReusablePolicySettingReferencingConfigurationPolicyEnrollmentTimeDeviceMembershipTarget - Invoke action
// clearEnrollmentTimeDeviceMembershipTarget
func (c ReusablePolicySettingReferencingConfigurationPolicyClient) ClearReusablePolicySettingReferencingConfigurationPolicyEnrollmentTimeDeviceMembershipTarget(ctx context.Context, id beta.DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId, options ClearReusablePolicySettingReferencingConfigurationPolicyEnrollmentTimeDeviceMembershipTargetOperationOptions) (result ClearReusablePolicySettingReferencingConfigurationPolicyEnrollmentTimeDeviceMembershipTargetOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/clearEnrollmentTimeDeviceMembershipTarget", id.ID()),
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

	var model ClearReusablePolicySettingReferencingConfigurationPolicyEnrollmentTimeDeviceMembershipTargetResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
