package windowsupdatecatalogitem

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

type DeleteWindowsUpdateCatalogItemOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteWindowsUpdateCatalogItemOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteWindowsUpdateCatalogItemOperationOptions() DeleteWindowsUpdateCatalogItemOperationOptions {
	return DeleteWindowsUpdateCatalogItemOperationOptions{}
}

func (o DeleteWindowsUpdateCatalogItemOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteWindowsUpdateCatalogItemOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteWindowsUpdateCatalogItemOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteWindowsUpdateCatalogItem - Delete navigation property windowsUpdateCatalogItems for deviceManagement
func (c WindowsUpdateCatalogItemClient) DeleteWindowsUpdateCatalogItem(ctx context.Context, id beta.DeviceManagementWindowsUpdateCatalogItemId, options DeleteWindowsUpdateCatalogItemOperationOptions) (result DeleteWindowsUpdateCatalogItemOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
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

	return
}
