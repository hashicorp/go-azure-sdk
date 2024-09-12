package settingshiftpreference

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

type GetSettingShiftPreferenceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ShiftPreferences
}

type GetSettingShiftPreferenceOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetSettingShiftPreferenceOperationOptions() GetSettingShiftPreferenceOperationOptions {
	return GetSettingShiftPreferenceOperationOptions{}
}

func (o GetSettingShiftPreferenceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSettingShiftPreferenceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetSettingShiftPreferenceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetSettingShiftPreference - Get shiftPreferences. Retrieve the properties and relationships of a shiftPreferences
// object by ID.
func (c SettingShiftPreferenceClient) GetSettingShiftPreference(ctx context.Context, id stable.UserId, options GetSettingShiftPreferenceOperationOptions) (result GetSettingShiftPreferenceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/settings/shiftPreferences", id.ID()),
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

	var model stable.ShiftPreferences
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
