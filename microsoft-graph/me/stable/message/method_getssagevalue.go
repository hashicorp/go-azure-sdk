package message

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

type GetssageValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetssageValueOperationOptions struct {
	IncludeHiddenMessages *string
}

func DefaultGetssageValueOperationOptions() GetssageValueOperationOptions {
	return GetssageValueOperationOptions{}
}

func (o GetssageValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetssageValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o GetssageValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.IncludeHiddenMessages != nil {
		out.Append("includeHiddenMessages", fmt.Sprintf("%v", *o.IncludeHiddenMessages))
	}
	return &out
}

// GetssageValue - Get open extension. Get an open extension (openTypeExtension object) identified by name or fully
// qualified name. The table in the Permissions section lists the resources that support open extensions. The following
// table lists the three scenarios where you can get an open extension from a supported resource instance.
func (c MessageClient) GetssageValue(ctx context.Context, id stable.MeMessageId, options GetssageValueOperationOptions) (result GetssageValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/$value", id.ID()),
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
