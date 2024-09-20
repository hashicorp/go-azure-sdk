package drivebundlecontent

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetDriveBundleContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetDriveBundleContentOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSetDriveBundleContentOperationOptions() SetDriveBundleContentOperationOptions {
	return SetDriveBundleContentOperationOptions{}
}

func (o SetDriveBundleContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetDriveBundleContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetDriveBundleContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetDriveBundleContent - Update content for the navigation property bundles in groups. The content stream, if the item
// represents a file.
func (c DriveBundleContentClient) SetDriveBundleContent(ctx context.Context, id stable.GroupIdDriveIdBundleId, input []byte, options SetDriveBundleContentOperationOptions) (result SetDriveBundleContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/content", id.ID()),
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
