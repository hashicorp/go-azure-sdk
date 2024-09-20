package synchronizationtemplate

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

type CreateSynchronizationTemplateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.SynchronizationTemplate
}

type CreateSynchronizationTemplateOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateSynchronizationTemplateOperationOptions() CreateSynchronizationTemplateOperationOptions {
	return CreateSynchronizationTemplateOperationOptions{}
}

func (o CreateSynchronizationTemplateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateSynchronizationTemplateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateSynchronizationTemplateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateSynchronizationTemplate - Create new navigation property to templates for applications
func (c SynchronizationTemplateClient) CreateSynchronizationTemplate(ctx context.Context, id stable.ApplicationId, input stable.SynchronizationTemplate, options CreateSynchronizationTemplateOperationOptions) (result CreateSynchronizationTemplateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/synchronization/templates", id.ID()),
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

	var model stable.SynchronizationTemplate
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
