package settingregionalandlanguagesetting

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteSettingRegionalAndLanguageSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteSettingRegionalAndLanguageSettingOperationOptions struct {
	IfMatch  *string
	Metadata *odata.Metadata
}

func DefaultDeleteSettingRegionalAndLanguageSettingOperationOptions() DeleteSettingRegionalAndLanguageSettingOperationOptions {
	return DeleteSettingRegionalAndLanguageSettingOperationOptions{}
}

func (o DeleteSettingRegionalAndLanguageSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteSettingRegionalAndLanguageSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteSettingRegionalAndLanguageSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteSettingRegionalAndLanguageSetting - Delete navigation property regionalAndLanguageSettings for me
func (c SettingRegionalAndLanguageSettingClient) DeleteSettingRegionalAndLanguageSetting(ctx context.Context, options DeleteSettingRegionalAndLanguageSettingOperationOptions) (result DeleteSettingRegionalAndLanguageSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
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

	return
}
