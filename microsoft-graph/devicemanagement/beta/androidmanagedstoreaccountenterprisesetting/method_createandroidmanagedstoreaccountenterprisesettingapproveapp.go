package androidmanagedstoreaccountenterprisesetting

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAndroidManagedStoreAccountEnterpriseSettingApproveAppOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateAndroidManagedStoreAccountEnterpriseSettingApproveAppOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAndroidManagedStoreAccountEnterpriseSettingApproveAppOperationOptions() CreateAndroidManagedStoreAccountEnterpriseSettingApproveAppOperationOptions {
	return CreateAndroidManagedStoreAccountEnterpriseSettingApproveAppOperationOptions{}
}

func (o CreateAndroidManagedStoreAccountEnterpriseSettingApproveAppOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAndroidManagedStoreAccountEnterpriseSettingApproveAppOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAndroidManagedStoreAccountEnterpriseSettingApproveAppOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAndroidManagedStoreAccountEnterpriseSettingApproveApp - Invoke action approveApps
func (c AndroidManagedStoreAccountEnterpriseSettingClient) CreateAndroidManagedStoreAccountEnterpriseSettingApproveApp(ctx context.Context, input CreateAndroidManagedStoreAccountEnterpriseSettingApproveAppRequest, options CreateAndroidManagedStoreAccountEnterpriseSettingApproveAppOperationOptions) (result CreateAndroidManagedStoreAccountEnterpriseSettingApproveAppOperationResponse, err error) {
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
		Path:          "/deviceManagement/androidManagedStoreAccountEnterpriseSettings/approveApps",
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
