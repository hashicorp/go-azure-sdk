package virtualendpointsupportedregion

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateVirtualEndpointSupportedRegionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CloudPCSupportedRegion
}

type CreateVirtualEndpointSupportedRegionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateVirtualEndpointSupportedRegionOperationOptions() CreateVirtualEndpointSupportedRegionOperationOptions {
	return CreateVirtualEndpointSupportedRegionOperationOptions{}
}

func (o CreateVirtualEndpointSupportedRegionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateVirtualEndpointSupportedRegionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateVirtualEndpointSupportedRegionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateVirtualEndpointSupportedRegion - Create new navigation property to supportedRegions for deviceManagement
func (c VirtualEndpointSupportedRegionClient) CreateVirtualEndpointSupportedRegion(ctx context.Context, input beta.CloudPCSupportedRegion, options CreateVirtualEndpointSupportedRegionOperationOptions) (result CreateVirtualEndpointSupportedRegionOperationResponse, err error) {
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
		Path:          "/deviceManagement/virtualEndpoint/supportedRegions",
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

	var model beta.CloudPCSupportedRegion
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
