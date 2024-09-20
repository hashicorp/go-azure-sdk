package settingregionalandlanguagesetting

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

type UpdateSettingRegionalAndLanguageSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSettingRegionalAndLanguageSettingOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateSettingRegionalAndLanguageSettingOperationOptions() UpdateSettingRegionalAndLanguageSettingOperationOptions {
	return UpdateSettingRegionalAndLanguageSettingOperationOptions{}
}

func (o UpdateSettingRegionalAndLanguageSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSettingRegionalAndLanguageSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSettingRegionalAndLanguageSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSettingRegionalAndLanguageSetting - Update the navigation property regionalAndLanguageSettings in users
func (c SettingRegionalAndLanguageSettingClient) UpdateSettingRegionalAndLanguageSetting(ctx context.Context, id beta.UserId, input beta.RegionalAndLanguageSettings, options UpdateSettingRegionalAndLanguageSettingOperationOptions) (result UpdateSettingRegionalAndLanguageSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/settings/regionalAndLanguageSettings", id.ID()),
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
