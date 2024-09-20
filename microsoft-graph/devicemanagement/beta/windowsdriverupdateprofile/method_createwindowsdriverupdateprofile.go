package windowsdriverupdateprofile

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateWindowsDriverUpdateProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.WindowsDriverUpdateProfile
}

type CreateWindowsDriverUpdateProfileOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateWindowsDriverUpdateProfileOperationOptions() CreateWindowsDriverUpdateProfileOperationOptions {
	return CreateWindowsDriverUpdateProfileOperationOptions{}
}

func (o CreateWindowsDriverUpdateProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateWindowsDriverUpdateProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateWindowsDriverUpdateProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateWindowsDriverUpdateProfile - Create new navigation property to windowsDriverUpdateProfiles for deviceManagement
func (c WindowsDriverUpdateProfileClient) CreateWindowsDriverUpdateProfile(ctx context.Context, input beta.WindowsDriverUpdateProfile, options CreateWindowsDriverUpdateProfileOperationOptions) (result CreateWindowsDriverUpdateProfileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/windowsDriverUpdateProfiles",
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

	var model beta.WindowsDriverUpdateProfile
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
