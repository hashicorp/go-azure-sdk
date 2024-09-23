package compliancemanagementpartner

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateComplianceManagementPartnerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ComplianceManagementPartner
}

type CreateComplianceManagementPartnerOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateComplianceManagementPartnerOperationOptions() CreateComplianceManagementPartnerOperationOptions {
	return CreateComplianceManagementPartnerOperationOptions{}
}

func (o CreateComplianceManagementPartnerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateComplianceManagementPartnerOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateComplianceManagementPartnerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateComplianceManagementPartner - Create new navigation property to complianceManagementPartners for
// deviceManagement
func (c ComplianceManagementPartnerClient) CreateComplianceManagementPartner(ctx context.Context, input beta.ComplianceManagementPartner, options CreateComplianceManagementPartnerOperationOptions) (result CreateComplianceManagementPartnerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/complianceManagementPartners",
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

	var model beta.ComplianceManagementPartner
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
