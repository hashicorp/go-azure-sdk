package datasecurityandgovernancesensitivitylabel

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDataSecurityAndGovernanceSensitivityLabelComputeRightsAndInheritanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ComputeRightsAndInheritanceResult
}

type CreateDataSecurityAndGovernanceSensitivityLabelComputeRightsAndInheritanceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDataSecurityAndGovernanceSensitivityLabelComputeRightsAndInheritanceOperationOptions() CreateDataSecurityAndGovernanceSensitivityLabelComputeRightsAndInheritanceOperationOptions {
	return CreateDataSecurityAndGovernanceSensitivityLabelComputeRightsAndInheritanceOperationOptions{}
}

func (o CreateDataSecurityAndGovernanceSensitivityLabelComputeRightsAndInheritanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDataSecurityAndGovernanceSensitivityLabelComputeRightsAndInheritanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDataSecurityAndGovernanceSensitivityLabelComputeRightsAndInheritanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDataSecurityAndGovernanceSensitivityLabelComputeRightsAndInheritance - Invoke action
// computeRightsAndInheritance. Computes the rights and inheritance for sensitivity labels based on the input content
// and labels.
func (c DataSecurityAndGovernanceSensitivityLabelClient) CreateDataSecurityAndGovernanceSensitivityLabelComputeRightsAndInheritance(ctx context.Context, input CreateDataSecurityAndGovernanceSensitivityLabelComputeRightsAndInheritanceRequest, options CreateDataSecurityAndGovernanceSensitivityLabelComputeRightsAndInheritanceOperationOptions) (result CreateDataSecurityAndGovernanceSensitivityLabelComputeRightsAndInheritanceOperationResponse, err error) {
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
		Path:          "/me/dataSecurityAndGovernance/sensitivityLabels/computeRightsAndInheritance",
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

	var model beta.ComputeRightsAndInheritanceResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
