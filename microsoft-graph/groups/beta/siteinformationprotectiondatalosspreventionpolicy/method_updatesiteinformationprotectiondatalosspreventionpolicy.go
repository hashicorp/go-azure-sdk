package siteinformationprotectiondatalosspreventionpolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateSiteInformationProtectionDataLossPreventionPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSiteInformationProtectionDataLossPreventionPolicyOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateSiteInformationProtectionDataLossPreventionPolicyOperationOptions() UpdateSiteInformationProtectionDataLossPreventionPolicyOperationOptions {
	return UpdateSiteInformationProtectionDataLossPreventionPolicyOperationOptions{}
}

func (o UpdateSiteInformationProtectionDataLossPreventionPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSiteInformationProtectionDataLossPreventionPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSiteInformationProtectionDataLossPreventionPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSiteInformationProtectionDataLossPreventionPolicy - Update the navigation property dataLossPreventionPolicies
// in groups
func (c SiteInformationProtectionDataLossPreventionPolicyClient) UpdateSiteInformationProtectionDataLossPreventionPolicy(ctx context.Context, id beta.GroupIdSiteIdInformationProtectionDataLossPreventionPolicyId, input beta.DataLossPreventionPolicy, options UpdateSiteInformationProtectionDataLossPreventionPolicyOperationOptions) (result UpdateSiteInformationProtectionDataLossPreventionPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
