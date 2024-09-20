package settingstoragequota

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

type UpdateSettingStorageQuotaOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSettingStorageQuotaOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateSettingStorageQuotaOperationOptions() UpdateSettingStorageQuotaOperationOptions {
	return UpdateSettingStorageQuotaOperationOptions{}
}

func (o UpdateSettingStorageQuotaOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSettingStorageQuotaOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSettingStorageQuotaOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSettingStorageQuota - Update the navigation property quota in users
func (c SettingStorageQuotaClient) UpdateSettingStorageQuota(ctx context.Context, id stable.UserId, input stable.UnifiedStorageQuota, options UpdateSettingStorageQuotaOperationOptions) (result UpdateSettingStorageQuotaOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/settings/storage/quota", id.ID()),
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
