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

type CreateReusablePolicySettingReferencingConfigurationPolicyReorderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateReusablePolicySettingReferencingConfigurationPolicyReorderOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateReusablePolicySettingReferencingConfigurationPolicyReorderOperationOptions() CreateReusablePolicySettingReferencingConfigurationPolicyReorderOperationOptions {
	return CreateReusablePolicySettingReferencingConfigurationPolicyReorderOperationOptions{}
}

func (o CreateReusablePolicySettingReferencingConfigurationPolicyReorderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateReusablePolicySettingReferencingConfigurationPolicyReorderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateReusablePolicySettingReferencingConfigurationPolicyReorderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateReusablePolicySettingReferencingConfigurationPolicyReorder - Invoke action reorder
func (c ReusablePolicySettingReferencingConfigurationPolicyClient) CreateReusablePolicySettingReferencingConfigurationPolicyReorder(ctx context.Context, id beta.DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId, input CreateReusablePolicySettingReferencingConfigurationPolicyReorderRequest, options CreateReusablePolicySettingReferencingConfigurationPolicyReorderOperationOptions) (result CreateReusablePolicySettingReferencingConfigurationPolicyReorderOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/reorder", id.ID()),
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

	return
}
