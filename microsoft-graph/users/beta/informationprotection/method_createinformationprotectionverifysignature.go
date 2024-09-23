package informationprotection

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

type CreateInformationProtectionVerifySignatureOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.VerificationResult
}

type CreateInformationProtectionVerifySignatureOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateInformationProtectionVerifySignatureOperationOptions() CreateInformationProtectionVerifySignatureOperationOptions {
	return CreateInformationProtectionVerifySignatureOperationOptions{}
}

func (o CreateInformationProtectionVerifySignatureOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateInformationProtectionVerifySignatureOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateInformationProtectionVerifySignatureOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateInformationProtectionVerifySignature - Invoke action verifySignature
func (c InformationProtectionClient) CreateInformationProtectionVerifySignature(ctx context.Context, id beta.UserId, input CreateInformationProtectionVerifySignatureRequest, options CreateInformationProtectionVerifySignatureOperationOptions) (result CreateInformationProtectionVerifySignatureOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/informationProtection/verifySignature", id.ID()),
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

	var model beta.VerificationResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
