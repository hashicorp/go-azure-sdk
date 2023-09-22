package backups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByVaultOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Backup
}

type ListByVaultCompleteResult struct {
	Items []Backup
}

type ListByVaultOperationOptions struct {
	Filter *string
}

func DefaultListByVaultOperationOptions() ListByVaultOperationOptions {
	return ListByVaultOperationOptions{}
}

func (o ListByVaultOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByVaultOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByVaultOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// ListByVault ...
func (c BackupsClient) ListByVault(ctx context.Context, id BackupVaultId, options ListByVaultOperationOptions) (result ListByVaultOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/backups", id.ID()),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]Backup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByVaultComplete retrieves all the results into a single object
func (c BackupsClient) ListByVaultComplete(ctx context.Context, id BackupVaultId, options ListByVaultOperationOptions) (ListByVaultCompleteResult, error) {
	return c.ListByVaultCompleteMatchingPredicate(ctx, id, options, BackupOperationPredicate{})
}

// ListByVaultCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BackupsClient) ListByVaultCompleteMatchingPredicate(ctx context.Context, id BackupVaultId, options ListByVaultOperationOptions, predicate BackupOperationPredicate) (result ListByVaultCompleteResult, err error) {
	items := make([]Backup, 0)

	resp, err := c.ListByVault(ctx, id, options)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListByVaultCompleteResult{
		Items: items,
	}
	return
}
