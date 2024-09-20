package securityinformationprotectionsensitivitylabel

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabelOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.SecurityContentLabel
}

type CreateSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabelOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabelOperationOptions() CreateSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabelOperationOptions {
	return CreateSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabelOperationOptions{}
}

func (o CreateSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabelOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabelOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabelOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabel - Invoke action extractContentLabel.
// Use the metadata that exists on an already-labeled piece of information to resolve the metadata to a specific
// sensitivity label. The contentInfo input is resolved to informationProtectionContentLabel.
func (c SecurityInformationProtectionSensitivityLabelClient) CreateSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabel(ctx context.Context, input CreateSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabelRequest, options CreateSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabelOperationOptions) (result CreateSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabelOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/me/security/informationProtection/sensitivityLabels/security.extractContentLabel",
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

	var model beta.SecurityContentLabel
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
