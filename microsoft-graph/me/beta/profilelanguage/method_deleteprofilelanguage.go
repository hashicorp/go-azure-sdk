package profilelanguage

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

type DeleteProfileLanguageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteProfileLanguageOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteProfileLanguageOperationOptions() DeleteProfileLanguageOperationOptions {
	return DeleteProfileLanguageOperationOptions{}
}

func (o DeleteProfileLanguageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteProfileLanguageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteProfileLanguageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteProfileLanguage - Delete languageProficiency. Delete a languageProficiency object from a user's profile.
func (c ProfileLanguageClient) DeleteProfileLanguage(ctx context.Context, id beta.MeProfileLanguageId, options DeleteProfileLanguageOperationOptions) (result DeleteProfileLanguageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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
