package user

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

type RevokeSignInSessionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *RevokeSignInSessionResult
}

// RevokeSignInSession - Invoke action revokeSignInSessions. Invalidates all the refresh tokens issued to applications
// for a user (as well as session cookies in a user's browser), by resetting the signInSessionsValidFromDateTime user
// property to the current date-time. Typically, this operation is performed (by the user or an administrator) if the
// user has a lost or stolen device. This operation prevents access to the organization's data through applications on
// the device by requiring the user to sign in again to all applications that they have previously consented to,
// independent of device.
func (c UserClient) RevokeSignInSession(ctx context.Context, id stable.UserId) (result RevokeSignInSessionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/revokeSignInSessions", id.ID()),
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

	var model RevokeSignInSessionResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
