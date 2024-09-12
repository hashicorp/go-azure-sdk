package userpfxcertificate

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserPfxCertificateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserPFXCertificate
}

type GetUserPfxCertificateOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetUserPfxCertificateOperationOptions() GetUserPfxCertificateOperationOptions {
	return GetUserPfxCertificateOperationOptions{}
}

func (o GetUserPfxCertificateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserPfxCertificateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetUserPfxCertificateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserPfxCertificate - Get userPfxCertificates from deviceManagement. Collection of PFX certificates associated with
// a user.
func (c UserPfxCertificateClient) GetUserPfxCertificate(ctx context.Context, id beta.DeviceManagementUserPfxCertificateId, options GetUserPfxCertificateOperationOptions) (result GetUserPfxCertificateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model beta.UserPFXCertificate
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
