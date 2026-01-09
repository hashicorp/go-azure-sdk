package securityinformationprotectionsensitivitylabel

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateSecurityInformationProtectionSensitivityLabelOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSecurityInformationProtectionSensitivityLabelOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateSecurityInformationProtectionSensitivityLabelOperationOptions() UpdateSecurityInformationProtectionSensitivityLabelOperationOptions {
	return UpdateSecurityInformationProtectionSensitivityLabelOperationOptions{}
}

func (o UpdateSecurityInformationProtectionSensitivityLabelOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSecurityInformationProtectionSensitivityLabelOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSecurityInformationProtectionSensitivityLabelOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSecurityInformationProtectionSensitivityLabel - Update the navigation property sensitivityLabels in me
func (c SecurityInformationProtectionSensitivityLabelClient) UpdateSecurityInformationProtectionSensitivityLabel(ctx context.Context, id beta.MeSecurityInformationProtectionSensitivityLabelId, input beta.SecuritySensitivityLabel, options UpdateSecurityInformationProtectionSensitivityLabelOperationOptions) (result UpdateSecurityInformationProtectionSensitivityLabelOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
