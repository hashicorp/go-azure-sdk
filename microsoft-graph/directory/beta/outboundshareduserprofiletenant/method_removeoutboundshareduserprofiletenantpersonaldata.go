package outboundshareduserprofiletenant

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoveOutboundSharedUserProfileTenantPersonalDataOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveOutboundSharedUserProfileTenantPersonalDataOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRemoveOutboundSharedUserProfileTenantPersonalDataOperationOptions() RemoveOutboundSharedUserProfileTenantPersonalDataOperationOptions {
	return RemoveOutboundSharedUserProfileTenantPersonalDataOperationOptions{}
}

func (o RemoveOutboundSharedUserProfileTenantPersonalDataOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RemoveOutboundSharedUserProfileTenantPersonalDataOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveOutboundSharedUserProfileTenantPersonalDataOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveOutboundSharedUserProfileTenantPersonalData - Invoke action removePersonalData. Create a request to remove the
// personal data for an outboundSharedUserProfile.
func (c OutboundSharedUserProfileTenantClient) RemoveOutboundSharedUserProfileTenantPersonalData(ctx context.Context, id beta.DirectoryOutboundSharedUserProfileIdTenantId, options RemoveOutboundSharedUserProfileTenantPersonalDataOperationOptions) (result RemoveOutboundSharedUserProfileTenantPersonalDataOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/removePersonalData", id.ID()),
		RetryFunc:     options.RetryFunc,
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
