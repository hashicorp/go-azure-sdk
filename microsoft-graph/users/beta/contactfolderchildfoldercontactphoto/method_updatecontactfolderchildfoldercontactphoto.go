package contactfolderchildfoldercontactphoto

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

type UpdateContactFolderChildFolderContactPhotoOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// UpdateContactFolderChildFolderContactPhoto - Update the navigation property photo in users
func (c ContactFolderChildFolderContactPhotoClient) UpdateContactFolderChildFolderContactPhoto(ctx context.Context, id beta.UserIdContactFolderIdChildFolderIdContactId, input beta.ProfilePhoto) (result UpdateContactFolderChildFolderContactPhotoOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPatch,
		Path:       fmt.Sprintf("%s/photo", id.ID()),
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

	return
}
