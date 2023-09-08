package useronlinemeetingtranscript

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserByIdOnlineMeetingByIdTranscriptsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.CallTranscriptCollectionResponse
}

type ListUserByIdOnlineMeetingByIdTranscriptsCompleteResult struct {
	Items []models.CallTranscriptCollectionResponse
}

// ListUserByIdOnlineMeetingByIdTranscripts ...
func (c UserOnlineMeetingTranscriptClient) ListUserByIdOnlineMeetingByIdTranscripts(ctx context.Context, id UserOnlineMeetingId) (result ListUserByIdOnlineMeetingByIdTranscriptsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/transcripts", id.ID()),
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
		Values *[]models.CallTranscriptCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdOnlineMeetingByIdTranscriptsComplete retrieves all the results into a single object
func (c UserOnlineMeetingTranscriptClient) ListUserByIdOnlineMeetingByIdTranscriptsComplete(ctx context.Context, id UserOnlineMeetingId) (ListUserByIdOnlineMeetingByIdTranscriptsCompleteResult, error) {
	return c.ListUserByIdOnlineMeetingByIdTranscriptsCompleteMatchingPredicate(ctx, id, models.CallTranscriptCollectionResponseOperationPredicate{})
}

// ListUserByIdOnlineMeetingByIdTranscriptsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserOnlineMeetingTranscriptClient) ListUserByIdOnlineMeetingByIdTranscriptsCompleteMatchingPredicate(ctx context.Context, id UserOnlineMeetingId, predicate models.CallTranscriptCollectionResponseOperationPredicate) (result ListUserByIdOnlineMeetingByIdTranscriptsCompleteResult, err error) {
	items := make([]models.CallTranscriptCollectionResponse, 0)

	resp, err := c.ListUserByIdOnlineMeetingByIdTranscripts(ctx, id)
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

	result = ListUserByIdOnlineMeetingByIdTranscriptsCompleteResult{
		Items: items,
	}
	return
}
