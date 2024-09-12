package authenticationtemporaryaccesspassmethod

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

type CreateAuthenticationTemporaryAccessPassMethodOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.TemporaryAccessPassAuthenticationMethod
}

// CreateAuthenticationTemporaryAccessPassMethod - Create temporaryAccessPassMethod. Create a new
// temporaryAccessPassAuthenticationMethod object on a user. A user can only have one Temporary Access Pass that's
// usable within its specified lifetime. If the user requires a new Temporary Access Pass while the current Temporary
// Access Pass is valid, the admin can create a new Temporary Access Pass for the user, the previous Temporary Access
// Pass will be deleted, and a new Temporary Access Pass will be created.
func (c AuthenticationTemporaryAccessPassMethodClient) CreateAuthenticationTemporaryAccessPassMethod(ctx context.Context, id beta.UserId, input beta.TemporaryAccessPassAuthenticationMethod) (result CreateAuthenticationTemporaryAccessPassMethodOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/authentication/temporaryAccessPassMethods", id.ID()),
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

	var model beta.TemporaryAccessPassAuthenticationMethod
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
