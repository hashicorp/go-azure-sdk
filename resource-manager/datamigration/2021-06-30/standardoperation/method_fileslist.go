package standardoperation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ProjectFile
}

type FilesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ProjectFile
}

// FilesList ...
func (c StandardOperationClient) FilesList(ctx context.Context, id ProjectId) (result FilesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/files", id.ID()),
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
		Values *[]ProjectFile `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// FilesListComplete retrieves all the results into a single object
func (c StandardOperationClient) FilesListComplete(ctx context.Context, id ProjectId) (FilesListCompleteResult, error) {
	return c.FilesListCompleteMatchingPredicate(ctx, id, ProjectFileOperationPredicate{})
}

// FilesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c StandardOperationClient) FilesListCompleteMatchingPredicate(ctx context.Context, id ProjectId, predicate ProjectFileOperationPredicate) (result FilesListCompleteResult, err error) {
	items := make([]ProjectFile, 0)

	resp, err := c.FilesList(ctx, id)
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

	result = FilesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
