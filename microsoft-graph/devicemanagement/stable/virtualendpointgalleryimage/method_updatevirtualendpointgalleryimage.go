package virtualendpointgalleryimage

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateVirtualEndpointGalleryImageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateVirtualEndpointGalleryImageOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateVirtualEndpointGalleryImageOperationOptions() UpdateVirtualEndpointGalleryImageOperationOptions {
	return UpdateVirtualEndpointGalleryImageOperationOptions{}
}

func (o UpdateVirtualEndpointGalleryImageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateVirtualEndpointGalleryImageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateVirtualEndpointGalleryImageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateVirtualEndpointGalleryImage - Update the navigation property galleryImages in deviceManagement
func (c VirtualEndpointGalleryImageClient) UpdateVirtualEndpointGalleryImage(ctx context.Context, id stable.DeviceManagementVirtualEndpointGalleryImageId, input stable.CloudPCGalleryImage, options UpdateVirtualEndpointGalleryImageOperationOptions) (result UpdateVirtualEndpointGalleryImageOperationResponse, err error) {
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
