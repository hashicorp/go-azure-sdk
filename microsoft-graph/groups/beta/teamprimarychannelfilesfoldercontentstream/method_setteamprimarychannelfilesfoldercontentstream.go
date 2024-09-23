package teamprimarychannelfilesfoldercontentstream

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

type SetTeamPrimaryChannelFilesFolderContentStreamOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetTeamPrimaryChannelFilesFolderContentStreamOperationOptions struct {
	ContentType string
	Metadata    *odata.Metadata
	RetryFunc   client.RequestRetryFunc
}

func DefaultSetTeamPrimaryChannelFilesFolderContentStreamOperationOptions() SetTeamPrimaryChannelFilesFolderContentStreamOperationOptions {
	return SetTeamPrimaryChannelFilesFolderContentStreamOperationOptions{}
}

func (o SetTeamPrimaryChannelFilesFolderContentStreamOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetTeamPrimaryChannelFilesFolderContentStreamOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetTeamPrimaryChannelFilesFolderContentStreamOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetTeamPrimaryChannelFilesFolderContentStream - Update contentStream for the navigation property filesFolder in
// groups. The content stream, if the item represents a file.
func (c TeamPrimaryChannelFilesFolderContentStreamClient) SetTeamPrimaryChannelFilesFolderContentStream(ctx context.Context, id beta.GroupId, input []byte, options SetTeamPrimaryChannelFilesFolderContentStreamOperationOptions) (result SetTeamPrimaryChannelFilesFolderContentStreamOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: options.ContentType,
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/team/primaryChannel/filesFolder/contentStream", id.ID()),
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

	return
}
