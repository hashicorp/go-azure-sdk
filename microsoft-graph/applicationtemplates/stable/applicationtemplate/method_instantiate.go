package applicationtemplate

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

type InstantiateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ApplicationServicePrincipal
}

// Instantiate - Invoke action instantiate. Add an instance of an application from the Microsoft Entra application
// gallery into your directory. The application template with ID 8adf8e6e-67b2-4cf2-a259-e3dc5476c621 can be used to add
// a non-gallery app that you can configure different single-sign on (SSO) modes like SAML SSO and password-based SSO.
func (c ApplicationTemplateClient) Instantiate(ctx context.Context, id stable.ApplicationTemplateId, input InstantiateRequest) (result InstantiateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/instantiate", id.ID()),
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

	var model stable.ApplicationServicePrincipal
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
