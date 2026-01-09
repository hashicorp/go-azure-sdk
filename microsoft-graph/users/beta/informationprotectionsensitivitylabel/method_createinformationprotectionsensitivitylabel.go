package informationprotectionsensitivitylabel

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateInformationProtectionSensitivityLabelOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.SensitivityLabel
}

type CreateInformationProtectionSensitivityLabelOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateInformationProtectionSensitivityLabelOperationOptions() CreateInformationProtectionSensitivityLabelOperationOptions {
	return CreateInformationProtectionSensitivityLabelOperationOptions{}
}

func (o CreateInformationProtectionSensitivityLabelOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateInformationProtectionSensitivityLabelOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateInformationProtectionSensitivityLabelOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateInformationProtectionSensitivityLabel - Create new navigation property to sensitivityLabels for users
func (c InformationProtectionSensitivityLabelClient) CreateInformationProtectionSensitivityLabel(ctx context.Context, id beta.UserId, input beta.SensitivityLabel, options CreateInformationProtectionSensitivityLabelOperationOptions) (result CreateInformationProtectionSensitivityLabelOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/informationProtection/sensitivityLabels", id.ID()),
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

	var model beta.SensitivityLabel
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
