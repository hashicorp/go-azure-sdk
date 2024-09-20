package securityinformationprotectionsensitivitylabel

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSecurityInformationProtectionSensitivityLabelsCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetSecurityInformationProtectionSensitivityLabelsCountOperationOptions struct {
	Filter   *string
	Metadata *odata.Metadata
	Search   *string
}

func DefaultGetSecurityInformationProtectionSensitivityLabelsCountOperationOptions() GetSecurityInformationProtectionSensitivityLabelsCountOperationOptions {
	return GetSecurityInformationProtectionSensitivityLabelsCountOperationOptions{}
}

func (o GetSecurityInformationProtectionSensitivityLabelsCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSecurityInformationProtectionSensitivityLabelsCountOperationOptions) ToOData() *odata.Query {
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

func (o GetSecurityInformationProtectionSensitivityLabelsCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetSecurityInformationProtectionSensitivityLabelsCount - Get the number of the resource
func (c SecurityInformationProtectionSensitivityLabelClient) GetSecurityInformationProtectionSensitivityLabelsCount(ctx context.Context, options GetSecurityInformationProtectionSensitivityLabelsCountOperationOptions) (result GetSecurityInformationProtectionSensitivityLabelsCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/me/security/informationProtection/sensitivityLabels/$count",
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
