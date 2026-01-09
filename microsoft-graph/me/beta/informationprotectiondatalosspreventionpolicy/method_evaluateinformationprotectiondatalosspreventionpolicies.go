package informationprotectiondatalosspreventionpolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EvaluateInformationProtectionDataLossPreventionPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DlpEvaluatePoliciesJobResponse
}

type EvaluateInformationProtectionDataLossPreventionPoliciesOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultEvaluateInformationProtectionDataLossPreventionPoliciesOperationOptions() EvaluateInformationProtectionDataLossPreventionPoliciesOperationOptions {
	return EvaluateInformationProtectionDataLossPreventionPoliciesOperationOptions{}
}

func (o EvaluateInformationProtectionDataLossPreventionPoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o EvaluateInformationProtectionDataLossPreventionPoliciesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o EvaluateInformationProtectionDataLossPreventionPoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// EvaluateInformationProtectionDataLossPreventionPolicies - Invoke action evaluate
func (c InformationProtectionDataLossPreventionPolicyClient) EvaluateInformationProtectionDataLossPreventionPolicies(ctx context.Context, input EvaluateInformationProtectionDataLossPreventionPoliciesRequest, options EvaluateInformationProtectionDataLossPreventionPoliciesOperationOptions) (result EvaluateInformationProtectionDataLossPreventionPoliciesOperationResponse, err error) {
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
		Path:          "/me/informationProtection/dataLossPreventionPolicies/evaluate",
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

	var model beta.DlpEvaluatePoliciesJobResponse
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
