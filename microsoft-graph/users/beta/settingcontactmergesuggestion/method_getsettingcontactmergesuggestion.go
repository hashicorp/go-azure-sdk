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

type GetSettingContactMergeSuggestionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ContactMergeSuggestions
}

type GetSettingContactMergeSuggestionOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetSettingContactMergeSuggestionOperationOptions() GetSettingContactMergeSuggestionOperationOptions {
	return GetSettingContactMergeSuggestionOperationOptions{}
}

func (o GetSettingContactMergeSuggestionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSettingContactMergeSuggestionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetSettingContactMergeSuggestionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetSettingContactMergeSuggestion - Get contactMergeSuggestions from users. The user's settings for the visibility of
// merge suggestion for the duplicate contacts in the user's contact list.
func (c SettingContactMergeSuggestionClient) GetSettingContactMergeSuggestion(ctx context.Context, id beta.UserId, options GetSettingContactMergeSuggestionOperationOptions) (result GetSettingContactMergeSuggestionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/settings/contactMergeSuggestions", id.ID()),
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

	var model beta.ContactMergeSuggestions
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
