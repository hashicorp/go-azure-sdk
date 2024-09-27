package deponboardingsetting

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

type CreateDepOnboardingSettingUploadDepTokenOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateDepOnboardingSettingUploadDepTokenOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDepOnboardingSettingUploadDepTokenOperationOptions() CreateDepOnboardingSettingUploadDepTokenOperationOptions {
	return CreateDepOnboardingSettingUploadDepTokenOperationOptions{}
}

func (o CreateDepOnboardingSettingUploadDepTokenOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDepOnboardingSettingUploadDepTokenOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDepOnboardingSettingUploadDepTokenOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDepOnboardingSettingUploadDepToken - Invoke action uploadDepToken. Uploads a new Device Enrollment Program
// token
func (c DepOnboardingSettingClient) CreateDepOnboardingSettingUploadDepToken(ctx context.Context, id beta.DeviceManagementDepOnboardingSettingId, input CreateDepOnboardingSettingUploadDepTokenRequest, options CreateDepOnboardingSettingUploadDepTokenOperationOptions) (result CreateDepOnboardingSettingUploadDepTokenOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/uploadDepToken", id.ID()),
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
