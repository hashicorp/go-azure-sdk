package entitlementmanagementaccesspackageassignmentpolicycustomextensionstagesetting

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

type GetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCountOperationOptions struct {
	Filter    *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Search    *string
}

func DefaultGetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCountOperationOptions() GetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCountOperationOptions {
	return GetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCountOperationOptions{}
}

func (o GetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCount - Get the number of the
// resource
func (c EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingClient) GetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCount(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId, options GetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCountOperationOptions) (result GetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/customExtensionStageSettings/$count", id.ID()),
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
