package informationprotectionsensitivitylabel

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

type EvaluateInformationProtectionSensitivityLabelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.EvaluateLabelJobResponse
}

type EvaluateInformationProtectionSensitivityLabelsOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultEvaluateInformationProtectionSensitivityLabelsOperationOptions() EvaluateInformationProtectionSensitivityLabelsOperationOptions {
	return EvaluateInformationProtectionSensitivityLabelsOperationOptions{}
}

func (o EvaluateInformationProtectionSensitivityLabelsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o EvaluateInformationProtectionSensitivityLabelsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o EvaluateInformationProtectionSensitivityLabelsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// EvaluateInformationProtectionSensitivityLabels - Invoke action evaluate
func (c InformationProtectionSensitivityLabelClient) EvaluateInformationProtectionSensitivityLabels(ctx context.Context, id beta.UserId, input EvaluateInformationProtectionSensitivityLabelsRequest, options EvaluateInformationProtectionSensitivityLabelsOperationOptions) (result EvaluateInformationProtectionSensitivityLabelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/informationProtection/sensitivityLabels/evaluate", id.ID()),
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

	var model beta.EvaluateLabelJobResponse
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
