package directoryadministrativeunitmember

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDirectoryAdministrativeUnitByIdMemberCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// GetDirectoryAdministrativeUnitByIdMemberCount ...
func (c DirectoryAdministrativeUnitMemberClient) GetDirectoryAdministrativeUnitByIdMemberCount(ctx context.Context, id DirectoryAdministrativeUnitId) (result GetDirectoryAdministrativeUnitByIdMemberCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/members/$count", id.ID()),
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
