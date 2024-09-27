package virtualendpointgalleryimage

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateVirtualEndpointGalleryImageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CloudPCGalleryImage
}

type CreateVirtualEndpointGalleryImageOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateVirtualEndpointGalleryImageOperationOptions() CreateVirtualEndpointGalleryImageOperationOptions {
	return CreateVirtualEndpointGalleryImageOperationOptions{}
}

func (o CreateVirtualEndpointGalleryImageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateVirtualEndpointGalleryImageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateVirtualEndpointGalleryImageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateVirtualEndpointGalleryImage - Create new navigation property to galleryImages for deviceManagement
func (c VirtualEndpointGalleryImageClient) CreateVirtualEndpointGalleryImage(ctx context.Context, input beta.CloudPCGalleryImage, options CreateVirtualEndpointGalleryImageOperationOptions) (result CreateVirtualEndpointGalleryImageOperationResponse, err error) {
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
		Path:          "/deviceManagement/virtualEndpoint/galleryImages",
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

	var model beta.CloudPCGalleryImage
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
