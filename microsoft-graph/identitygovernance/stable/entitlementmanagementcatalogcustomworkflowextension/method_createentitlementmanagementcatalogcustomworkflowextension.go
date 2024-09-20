package entitlementmanagementcatalogcustomworkflowextension

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateEntitlementManagementCatalogCustomWorkflowExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.CustomCalloutExtension
}

type CreateEntitlementManagementCatalogCustomWorkflowExtensionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateEntitlementManagementCatalogCustomWorkflowExtensionOperationOptions() CreateEntitlementManagementCatalogCustomWorkflowExtensionOperationOptions {
	return CreateEntitlementManagementCatalogCustomWorkflowExtensionOperationOptions{}
}

func (o CreateEntitlementManagementCatalogCustomWorkflowExtensionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementCatalogCustomWorkflowExtensionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementCatalogCustomWorkflowExtensionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementCatalogCustomWorkflowExtension - Create accessPackageCustomWorkflowExtension. Create a new
// accessPackageAssignmentRequestWorkflowExtension or accessPackageAssignmentWorkflowExtension object and add it to an
// existing accessPackageCatalog object. You must explicitly provide an @odata.type property that indicates whether the
// object is an accessPackageAssignmentRequestWorkflowExtension or an accessPackageAssignmentWorkflowExtension.
func (c EntitlementManagementCatalogCustomWorkflowExtensionClient) CreateEntitlementManagementCatalogCustomWorkflowExtension(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementCatalogId, input stable.CustomCalloutExtension, options CreateEntitlementManagementCatalogCustomWorkflowExtensionOperationOptions) (result CreateEntitlementManagementCatalogCustomWorkflowExtensionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/customWorkflowExtensions", id.ID()),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalCustomCalloutExtensionImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
