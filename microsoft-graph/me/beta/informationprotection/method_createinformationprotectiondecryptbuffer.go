package informationprotection

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateInformationProtectionDecryptBufferOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.BufferDecryptionResult
}

type CreateInformationProtectionDecryptBufferOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateInformationProtectionDecryptBufferOperationOptions() CreateInformationProtectionDecryptBufferOperationOptions {
	return CreateInformationProtectionDecryptBufferOperationOptions{}
}

func (o CreateInformationProtectionDecryptBufferOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateInformationProtectionDecryptBufferOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateInformationProtectionDecryptBufferOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateInformationProtectionDecryptBuffer - Invoke action decryptBuffer
func (c InformationProtectionClient) CreateInformationProtectionDecryptBuffer(ctx context.Context, input CreateInformationProtectionDecryptBufferRequest, options CreateInformationProtectionDecryptBufferOperationOptions) (result CreateInformationProtectionDecryptBufferOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/me/informationProtection/decryptBuffer",
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

	var model beta.BufferDecryptionResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
