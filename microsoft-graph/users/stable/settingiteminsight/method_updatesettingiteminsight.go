package settingiteminsight

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

type UpdateSettingItemInsightOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSettingItemInsightOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateSettingItemInsightOperationOptions() UpdateSettingItemInsightOperationOptions {
	return UpdateSettingItemInsightOperationOptions{}
}

func (o UpdateSettingItemInsightOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSettingItemInsightOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSettingItemInsightOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSettingItemInsight - Update the navigation property itemInsights in users
func (c SettingItemInsightClient) UpdateSettingItemInsight(ctx context.Context, id stable.UserId, input stable.UserInsightsSettings, options UpdateSettingItemInsightOperationOptions) (result UpdateSettingItemInsightOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/settings/itemInsights", id.ID()),
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
