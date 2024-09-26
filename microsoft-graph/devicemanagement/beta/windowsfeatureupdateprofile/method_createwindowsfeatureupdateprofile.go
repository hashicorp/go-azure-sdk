package windowsfeatureupdateprofile

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateWindowsFeatureUpdateProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.WindowsFeatureUpdateProfile
}

type CreateWindowsFeatureUpdateProfileOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateWindowsFeatureUpdateProfileOperationOptions() CreateWindowsFeatureUpdateProfileOperationOptions {
	return CreateWindowsFeatureUpdateProfileOperationOptions{}
}

func (o CreateWindowsFeatureUpdateProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateWindowsFeatureUpdateProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateWindowsFeatureUpdateProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateWindowsFeatureUpdateProfile - Create new navigation property to windowsFeatureUpdateProfiles for
// deviceManagement
func (c WindowsFeatureUpdateProfileClient) CreateWindowsFeatureUpdateProfile(ctx context.Context, input beta.WindowsFeatureUpdateProfile, options CreateWindowsFeatureUpdateProfileOperationOptions) (result CreateWindowsFeatureUpdateProfileOperationResponse, err error) {
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
		Path:          "/deviceManagement/windowsFeatureUpdateProfiles",
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

	var model beta.WindowsFeatureUpdateProfile
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
