package meonlinemeetingtranscript

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

type ListMeOnlineMeetingByIdTranscriptsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.CallTranscriptCollectionResponse
}

type ListMeOnlineMeetingByIdTranscriptsCompleteResult struct {
	Items []models.CallTranscriptCollectionResponse
}

// ListMeOnlineMeetingByIdTranscripts ...
func (c MeOnlineMeetingTranscriptClient) ListMeOnlineMeetingByIdTranscripts(ctx context.Context, id MeOnlineMeetingId) (result ListMeOnlineMeetingByIdTranscriptsOperationResponse, err error) {
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

// ListMeOnlineMeetingByIdTranscriptsComplete retrieves all the results into a single object
func (c MeOnlineMeetingTranscriptClient) ListMeOnlineMeetingByIdTranscriptsComplete(ctx context.Context, id MeOnlineMeetingId) (ListMeOnlineMeetingByIdTranscriptsCompleteResult, error) {
	return c.ListMeOnlineMeetingByIdTranscriptsCompleteMatchingPredicate(ctx, id, models.CallTranscriptCollectionResponseOperationPredicate{})
}

// ListMeOnlineMeetingByIdTranscriptsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOnlineMeetingTranscriptClient) ListMeOnlineMeetingByIdTranscriptsCompleteMatchingPredicate(ctx context.Context, id MeOnlineMeetingId, predicate models.CallTranscriptCollectionResponseOperationPredicate) (result ListMeOnlineMeetingByIdTranscriptsCompleteResult, err error) {
	items := make([]models.CallTranscriptCollectionResponse, 0)

	resp, err := c.ListMeOnlineMeetingByIdTranscripts(ctx, id)
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

	result = ListMeOnlineMeetingByIdTranscriptsCompleteResult{
		Items: items,
	}
	return
}
