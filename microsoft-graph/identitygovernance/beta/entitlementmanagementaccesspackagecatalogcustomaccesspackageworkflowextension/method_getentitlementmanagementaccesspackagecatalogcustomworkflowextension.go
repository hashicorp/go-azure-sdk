package entitlementmanagementaccesspackagecatalogcustomaccesspackageworkflowextension

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CustomAccessPackageWorkflowExtension
}

type GetEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionOperationOptions() GetEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionOperationOptions {
	return GetEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionOperationOptions{}
}

func (o GetEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementAccessPackageCatalogCustomWorkflowExtension - Get customAccessPackageWorkflowExtension. Read
// the properties and relationships of a customAccessPackageWorkflowExtension object for an accessPackageCatalog object.
func (c EntitlementManagementAccessPackageCatalogCustomAccessPackageWorkflowExtensionClient) GetEntitlementManagementAccessPackageCatalogCustomWorkflowExtension(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId, options GetEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionOperationOptions) (result GetEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model beta.CustomAccessPackageWorkflowExtension
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
