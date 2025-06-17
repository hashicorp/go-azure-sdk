package templatedevicetemplateowner

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTemplateDeviceTemplateOwnerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.DirectoryObject
}

type GetTemplateDeviceTemplateOwnerOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetTemplateDeviceTemplateOwnerOperationOptions() GetTemplateDeviceTemplateOwnerOperationOptions {
	return GetTemplateDeviceTemplateOwnerOperationOptions{}
}

func (o GetTemplateDeviceTemplateOwnerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetTemplateDeviceTemplateOwnerOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetTemplateDeviceTemplateOwnerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetTemplateDeviceTemplateOwner - Get owners from directory. Collection of directory objects that can manage the
// device template and the related deviceInstances. Owners can be represented as service principals, users, or
// applications. An owner has full privileges over the device template and doesn't require other administrator roles to
// create, update, or delete devices from this template, as well as to add or remove template owners. There can be a
// maximum of 100 owners on a device template. Supports $expand.
func (c TemplateDeviceTemplateOwnerClient) GetTemplateDeviceTemplateOwner(ctx context.Context, id beta.DirectoryTemplateDeviceTemplateIdOwnerId, options GetTemplateDeviceTemplateOwnerOperationOptions) (result GetTemplateDeviceTemplateOwnerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalDirectoryObjectImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
