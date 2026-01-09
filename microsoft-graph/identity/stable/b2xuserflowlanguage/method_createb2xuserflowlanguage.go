package b2xuserflowlanguage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateB2xUserFlowLanguageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UserFlowLanguageConfiguration
}

type CreateB2xUserFlowLanguageOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateB2xUserFlowLanguageOperationOptions() CreateB2xUserFlowLanguageOperationOptions {
	return CreateB2xUserFlowLanguageOperationOptions{}
}

func (o CreateB2xUserFlowLanguageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateB2xUserFlowLanguageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateB2xUserFlowLanguageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateB2xUserFlowLanguage - Create new navigation property to languages for identity
func (c B2xUserFlowLanguageClient) CreateB2xUserFlowLanguage(ctx context.Context, id stable.IdentityB2xUserFlowId, input stable.UserFlowLanguageConfiguration, options CreateB2xUserFlowLanguageOperationOptions) (result CreateB2xUserFlowLanguageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/languages", id.ID()),
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

	var model stable.UserFlowLanguageConfiguration
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
