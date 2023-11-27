package longtermretentionbackup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LtrBackupOperationsListByServerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]LtrServerBackupOperation
}

type LtrBackupOperationsListByServerCompleteResult struct {
	Items []LtrServerBackupOperation
}

// LtrBackupOperationsListByServer ...
func (c LongTermRetentionBackupClient) LtrBackupOperationsListByServer(ctx context.Context, id FlexibleServerId) (result LtrBackupOperationsListByServerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/ltrBackupOperations", id.ID()),
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
		Values *[]LtrServerBackupOperation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// LtrBackupOperationsListByServerComplete retrieves all the results into a single object
func (c LongTermRetentionBackupClient) LtrBackupOperationsListByServerComplete(ctx context.Context, id FlexibleServerId) (LtrBackupOperationsListByServerCompleteResult, error) {
	return c.LtrBackupOperationsListByServerCompleteMatchingPredicate(ctx, id, LtrServerBackupOperationOperationPredicate{})
}

// LtrBackupOperationsListByServerCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LongTermRetentionBackupClient) LtrBackupOperationsListByServerCompleteMatchingPredicate(ctx context.Context, id FlexibleServerId, predicate LtrServerBackupOperationOperationPredicate) (result LtrBackupOperationsListByServerCompleteResult, err error) {
	items := make([]LtrServerBackupOperation, 0)

	resp, err := c.LtrBackupOperationsListByServer(ctx, id)
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

	result = LtrBackupOperationsListByServerCompleteResult{
		Items: items,
	}
	return
}
