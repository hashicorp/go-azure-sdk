package settingregionalandlanguagesetting

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSettingRegionalAndLanguageSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.RegionalAndLanguageSettings
}

type GetSettingRegionalAndLanguageSettingOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetSettingRegionalAndLanguageSettingOperationOptions() GetSettingRegionalAndLanguageSettingOperationOptions {
	return GetSettingRegionalAndLanguageSettingOperationOptions{}
}

func (o GetSettingRegionalAndLanguageSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSettingRegionalAndLanguageSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetSettingRegionalAndLanguageSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetSettingRegionalAndLanguageSetting - Get regionalAndLanguageSettings. Retrieve the properties of a user's
// regionalAndLanguageSettings.
func (c SettingRegionalAndLanguageSettingClient) GetSettingRegionalAndLanguageSetting(ctx context.Context, options GetSettingRegionalAndLanguageSettingOperationOptions) (result GetSettingRegionalAndLanguageSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/me/settings/regionalAndLanguageSettings",
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

	var model beta.RegionalAndLanguageSettings
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
