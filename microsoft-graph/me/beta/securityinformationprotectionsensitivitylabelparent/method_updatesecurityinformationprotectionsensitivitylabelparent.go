package securityinformationprotectionsensitivitylabelparent

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

type UpdateSecurityInformationProtectionSensitivityLabelParentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSecurityInformationProtectionSensitivityLabelParentOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateSecurityInformationProtectionSensitivityLabelParentOperationOptions() UpdateSecurityInformationProtectionSensitivityLabelParentOperationOptions {
	return UpdateSecurityInformationProtectionSensitivityLabelParentOperationOptions{}
}

func (o UpdateSecurityInformationProtectionSensitivityLabelParentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSecurityInformationProtectionSensitivityLabelParentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSecurityInformationProtectionSensitivityLabelParentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSecurityInformationProtectionSensitivityLabelParent - Update the navigation property parent in me
func (c SecurityInformationProtectionSensitivityLabelParentClient) UpdateSecurityInformationProtectionSensitivityLabelParent(ctx context.Context, id beta.MeSecurityInformationProtectionSensitivityLabelId, input beta.SecuritySensitivityLabel, options UpdateSecurityInformationProtectionSensitivityLabelParentOperationOptions) (result UpdateSecurityInformationProtectionSensitivityLabelParentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/parent", id.ID()),
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
