package directoryadministrativeunitscopedrolemember

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteDirectoryAdministrativeUnitByIdScopedRoleMemberByIdOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// DeleteDirectoryAdministrativeUnitByIdScopedRoleMemberById ...
func (c DirectoryAdministrativeUnitScopedRoleMemberClient) DeleteDirectoryAdministrativeUnitByIdScopedRoleMemberById(ctx context.Context, id DirectoryAdministrativeUnitScopedRoleMemberId) (result DeleteDirectoryAdministrativeUnitByIdScopedRoleMemberByIdOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodDelete,
		Path:       id.ID(),
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
