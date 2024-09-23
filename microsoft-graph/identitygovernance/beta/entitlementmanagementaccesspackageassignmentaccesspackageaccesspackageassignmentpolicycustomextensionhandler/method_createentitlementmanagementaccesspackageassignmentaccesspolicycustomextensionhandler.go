package entitlementmanagementaccesspackageassignmentaccesspackageaccesspackageassignmentpolicycustomextensionhandler

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

type CreateEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CustomExtensionHandler
}

type CreateEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlerOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlerOperationOptions() CreateEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlerOperationOptions {
	return CreateEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlerOperationOptions{}
}

func (o CreateEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlerOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandler - Create new navigation property
// to customExtensionHandlers for identityGovernance
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageAssignmentPolicyCustomExtensionHandlerClient) CreateEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandler(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageAssignmentPolicyId, input beta.CustomExtensionHandler, options CreateEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlerOperationOptions) (result CreateEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/customExtensionHandlers", id.ID()),
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

	var model beta.CustomExtensionHandler
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
