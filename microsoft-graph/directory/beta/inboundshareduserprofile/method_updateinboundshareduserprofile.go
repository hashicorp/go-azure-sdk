package inboundshareduserprofile

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateInboundSharedUserProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateInboundSharedUserProfileOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateInboundSharedUserProfileOperationOptions() UpdateInboundSharedUserProfileOperationOptions {
	return UpdateInboundSharedUserProfileOperationOptions{}
}

func (o UpdateInboundSharedUserProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateInboundSharedUserProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateInboundSharedUserProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateInboundSharedUserProfile - Update the navigation property inboundSharedUserProfiles in directory
func (c InboundSharedUserProfileClient) UpdateInboundSharedUserProfile(ctx context.Context, id beta.DirectoryInboundSharedUserProfileId, input beta.InboundSharedUserProfile, options UpdateInboundSharedUserProfileOperationOptions) (result UpdateInboundSharedUserProfileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
