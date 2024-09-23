package settingstoragequotaservice

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

type CreateSettingStorageQuotaServiceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ServiceStorageQuotaBreakdown
}

type CreateSettingStorageQuotaServiceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateSettingStorageQuotaServiceOperationOptions() CreateSettingStorageQuotaServiceOperationOptions {
	return CreateSettingStorageQuotaServiceOperationOptions{}
}

func (o CreateSettingStorageQuotaServiceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateSettingStorageQuotaServiceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateSettingStorageQuotaServiceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateSettingStorageQuotaService - Create new navigation property to services for users
func (c SettingStorageQuotaServiceClient) CreateSettingStorageQuotaService(ctx context.Context, id stable.UserId, input stable.ServiceStorageQuotaBreakdown, options CreateSettingStorageQuotaServiceOperationOptions) (result CreateSettingStorageQuotaServiceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/settings/storage/quota/services", id.ID()),
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

	var model stable.ServiceStorageQuotaBreakdown
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
