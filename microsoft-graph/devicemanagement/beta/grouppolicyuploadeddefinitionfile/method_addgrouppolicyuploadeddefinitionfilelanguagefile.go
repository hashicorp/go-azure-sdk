package grouppolicyuploadeddefinitionfile

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

type AddGroupPolicyUploadedDefinitionFileLanguageFileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// AddGroupPolicyUploadedDefinitionFileLanguageFile - Invoke action addLanguageFiles
func (c GroupPolicyUploadedDefinitionFileClient) AddGroupPolicyUploadedDefinitionFileLanguageFile(ctx context.Context, id beta.DeviceManagementGroupPolicyUploadedDefinitionFileId, input AddGroupPolicyUploadedDefinitionFileLanguageFileRequest) (result AddGroupPolicyUploadedDefinitionFileLanguageFileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/addLanguageFiles", id.ID()),
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
