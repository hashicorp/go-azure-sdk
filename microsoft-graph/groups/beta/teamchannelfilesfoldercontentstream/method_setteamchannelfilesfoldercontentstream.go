package teamchannelfilesfoldercontentstream

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

type SetTeamChannelFilesFolderContentStreamOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetTeamChannelFilesFolderContentStreamOperationOptions struct {
	ContentType string
	Metadata    *odata.Metadata
	RetryFunc   client.RequestRetryFunc
}

func DefaultSetTeamChannelFilesFolderContentStreamOperationOptions() SetTeamChannelFilesFolderContentStreamOperationOptions {
	return SetTeamChannelFilesFolderContentStreamOperationOptions{}
}

func (o SetTeamChannelFilesFolderContentStreamOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetTeamChannelFilesFolderContentStreamOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetTeamChannelFilesFolderContentStreamOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetTeamChannelFilesFolderContentStream - Update contentStream for the navigation property filesFolder in groups. The
// content stream, if the item represents a file.
func (c TeamChannelFilesFolderContentStreamClient) SetTeamChannelFilesFolderContentStream(ctx context.Context, id beta.GroupIdTeamChannelId, input []byte, options SetTeamChannelFilesFolderContentStreamOperationOptions) (result SetTeamChannelFilesFolderContentStreamOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: options.ContentType,
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/filesFolder/contentStream", id.ID()),
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
