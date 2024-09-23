package inferenceclassificationoverride

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateInferenceClassificationOverrideOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.InferenceClassificationOverride
}

type CreateInferenceClassificationOverrideOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateInferenceClassificationOverrideOperationOptions() CreateInferenceClassificationOverrideOperationOptions {
	return CreateInferenceClassificationOverrideOperationOptions{}
}

func (o CreateInferenceClassificationOverrideOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateInferenceClassificationOverrideOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateInferenceClassificationOverrideOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateInferenceClassificationOverride - Create inferenceClassificationOverride. Create a focused Inbox override for a
// sender identified by an SMTP address. Future messages from that SMTP address will be consistently classified as
// specified in the override.
func (c InferenceClassificationOverrideClient) CreateInferenceClassificationOverride(ctx context.Context, input beta.InferenceClassificationOverride, options CreateInferenceClassificationOverrideOperationOptions) (result CreateInferenceClassificationOverrideOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/me/inferenceClassification/overrides",
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

	var model beta.InferenceClassificationOverride
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
