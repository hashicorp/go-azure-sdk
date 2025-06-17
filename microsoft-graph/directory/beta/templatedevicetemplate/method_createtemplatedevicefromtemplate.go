package templatedevicetemplate

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

type CreateTemplateDeviceFromTemplateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.Device
}

type CreateTemplateDeviceFromTemplateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateTemplateDeviceFromTemplateOperationOptions() CreateTemplateDeviceFromTemplateOperationOptions {
	return CreateTemplateDeviceFromTemplateOperationOptions{}
}

func (o CreateTemplateDeviceFromTemplateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTemplateDeviceFromTemplateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTemplateDeviceFromTemplateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTemplateDeviceFromTemplate - Invoke action createDeviceFromTemplate. Create a new device from a deviceTemplate.
func (c TemplateDeviceTemplateClient) CreateTemplateDeviceFromTemplate(ctx context.Context, id beta.DirectoryTemplateDeviceTemplateId, input CreateTemplateDeviceFromTemplateRequest, options CreateTemplateDeviceFromTemplateOperationOptions) (result CreateTemplateDeviceFromTemplateOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/createDeviceFromTemplate", id.ID()),
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

	var model beta.Device
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
