package inferenceclassification

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

type UpdateInferenceClassificationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateInferenceClassificationOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateInferenceClassificationOperationOptions() UpdateInferenceClassificationOperationOptions {
	return UpdateInferenceClassificationOperationOptions{}
}

func (o UpdateInferenceClassificationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateInferenceClassificationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateInferenceClassificationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateInferenceClassification - Update the navigation property inferenceClassification in users
func (c InferenceClassificationClient) UpdateInferenceClassification(ctx context.Context, id beta.UserId, input beta.InferenceClassification, options UpdateInferenceClassificationOperationOptions) (result UpdateInferenceClassificationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/inferenceClassification", id.ID()),
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
