package settingcontactmergesuggestion

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

type UpdateSettingContactMergeSuggestionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSettingContactMergeSuggestionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateSettingContactMergeSuggestionOperationOptions() UpdateSettingContactMergeSuggestionOperationOptions {
	return UpdateSettingContactMergeSuggestionOperationOptions{}
}

func (o UpdateSettingContactMergeSuggestionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSettingContactMergeSuggestionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSettingContactMergeSuggestionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSettingContactMergeSuggestion - Update the navigation property contactMergeSuggestions in users
func (c SettingContactMergeSuggestionClient) UpdateSettingContactMergeSuggestion(ctx context.Context, id beta.UserId, input beta.ContactMergeSuggestions, options UpdateSettingContactMergeSuggestionOperationOptions) (result UpdateSettingContactMergeSuggestionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/settings/contactMergeSuggestions", id.ID()),
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
