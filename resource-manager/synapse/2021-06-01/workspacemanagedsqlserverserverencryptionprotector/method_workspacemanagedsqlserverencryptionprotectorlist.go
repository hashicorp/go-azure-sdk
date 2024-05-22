package workspacemanagedsqlserverserverencryptionprotector

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagedSqlServerEncryptionProtectorListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]EncryptionProtector
}

type WorkspaceManagedSqlServerEncryptionProtectorListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []EncryptionProtector
}

// WorkspaceManagedSqlServerEncryptionProtectorList ...
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) WorkspaceManagedSqlServerEncryptionProtectorList(ctx context.Context, id WorkspaceId) (result WorkspaceManagedSqlServerEncryptionProtectorListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/encryptionProtector", id.ID()),
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
		Values *[]EncryptionProtector `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WorkspaceManagedSqlServerEncryptionProtectorListComplete retrieves all the results into a single object
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) WorkspaceManagedSqlServerEncryptionProtectorListComplete(ctx context.Context, id WorkspaceId) (WorkspaceManagedSqlServerEncryptionProtectorListCompleteResult, error) {
	return c.WorkspaceManagedSqlServerEncryptionProtectorListCompleteMatchingPredicate(ctx, id, EncryptionProtectorOperationPredicate{})
}

// WorkspaceManagedSqlServerEncryptionProtectorListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) WorkspaceManagedSqlServerEncryptionProtectorListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate EncryptionProtectorOperationPredicate) (result WorkspaceManagedSqlServerEncryptionProtectorListCompleteResult, err error) {
	items := make([]EncryptionProtector, 0)

	resp, err := c.WorkspaceManagedSqlServerEncryptionProtectorList(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = WorkspaceManagedSqlServerEncryptionProtectorListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
