package inferenceclassificationoverride

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateInferenceClassificationOverrideOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateInferenceClassificationOverrideOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateInferenceClassificationOverrideOperationOptions() UpdateInferenceClassificationOverrideOperationOptions {
	return UpdateInferenceClassificationOverrideOperationOptions{}
}

func (o UpdateInferenceClassificationOverrideOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateInferenceClassificationOverrideOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateInferenceClassificationOverrideOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateInferenceClassificationOverride - Update the navigation property overrides in users
func (c InferenceClassificationOverrideClient) UpdateInferenceClassificationOverride(ctx context.Context, id stable.UserIdInferenceClassificationOverrideId, input stable.InferenceClassificationOverride, options UpdateInferenceClassificationOverrideOperationOptions) (result UpdateInferenceClassificationOverrideOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
