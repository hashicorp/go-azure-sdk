package appserviceenvironments

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResumeOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
	Model        *[]Site

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ResumeOperationResponse, error)
}

type ResumeCompleteResult struct {
	Items []Site
}

func (r ResumeOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ResumeOperationResponse) LoadMore(ctx context.Context) (resp ResumeOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// Resume ...
func (c AppServiceEnvironmentsClient) Resume(ctx context.Context, id HostingEnvironmentId) (result ResumeOperationResponse, err error) {
	req, err := c.preparerForResume(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "Resume", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForResume(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "Resume", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ResumeThenPoll performs Resume then polls until it's completed
func (c AppServiceEnvironmentsClient) ResumeThenPoll(ctx context.Context, id HostingEnvironmentId) error {
	result, err := c.Resume(ctx, id)
	if err != nil {
		return fmt.Errorf("performing Resume: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Resume: %+v", err)
	}

	return nil
}

// preparerForResume prepares the Resume request.
func (c AppServiceEnvironmentsClient) preparerForResume(ctx context.Context, id HostingEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/resume", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForResumeWithNextLink prepares the Resume request with the given nextLink token.
func (c AppServiceEnvironmentsClient) preparerForResumeWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForResume sends the Resume request. The method will close the
// http.Response Body if it receives an error.
func (c AppServiceEnvironmentsClient) senderForResume(ctx context.Context, req *http.Request) (future ResumeOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
