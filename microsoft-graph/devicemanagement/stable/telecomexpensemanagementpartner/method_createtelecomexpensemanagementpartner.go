package telecomexpensemanagementpartner

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateTelecomExpenseManagementPartnerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.TelecomExpenseManagementPartner
}

type CreateTelecomExpenseManagementPartnerOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateTelecomExpenseManagementPartnerOperationOptions() CreateTelecomExpenseManagementPartnerOperationOptions {
	return CreateTelecomExpenseManagementPartnerOperationOptions{}
}

func (o CreateTelecomExpenseManagementPartnerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTelecomExpenseManagementPartnerOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTelecomExpenseManagementPartnerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTelecomExpenseManagementPartner - Create telecomExpenseManagementPartner. Create a new
// telecomExpenseManagementPartner object.
func (c TelecomExpenseManagementPartnerClient) CreateTelecomExpenseManagementPartner(ctx context.Context, input stable.TelecomExpenseManagementPartner, options CreateTelecomExpenseManagementPartnerOperationOptions) (result CreateTelecomExpenseManagementPartnerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/telecomExpenseManagementPartners",
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

	var model stable.TelecomExpenseManagementPartner
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
