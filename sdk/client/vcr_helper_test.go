package client

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestIsVCRRecordedResponse(t *testing.T) {
	testCases := []struct {
		name     string
		response *http.Response
		expected bool
	}{
		{
			name:     "nil response",
			response: nil,
			expected: false,
		},
		{
			name: "nil headers",
			response: &http.Response{
				Header: nil,
			},
			expected: false,
		},
		{
			name: "header false",
			response: &http.Response{
				Header: http.Header{
					"X-Azure-SDK-VCR-Test": []string{"true"},
				},
			},
			expected: false,
		},
		{
			name: "header true",
			response: &http.Response{
				Header: http.Header{
					http.CanonicalHeaderKey(VCRReplayHeader): []string{"true"},
				},
			},
			expected: true,
		},
		{
			name: "header true mixed case",
			response: &http.Response{
				Header: http.Header{
					http.CanonicalHeaderKey(VCRReplayHeader): []string{"TrUe"},
				},
			},
			expected: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := IsVCRRecordedResponse(testCase.response)
			if actual != testCase.expected {
				t.Fatalf("expected %t, got %t", testCase.expected, actual)
			}
		})
	}
}

func TestIsVCRReplayMissError(t *testing.T) {
	testCases := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "nil error",
			err:      nil,
			expected: false,
		},
		{
			name:     "plain error",
			err:      errors.New(VCRInteractionNotFoundErrMsg),
			expected: false,
		},
		{
			name: "url error replay miss",
			err: &url.Error{
				Op:  http.MethodGet,
				URL: "https://example.test/resource",
				Err: errors.New(VCRInteractionNotFoundErrMsg),
			},
			expected: true,
		},
		{
			name: "wrapped url error replay miss",
			err: fmt.Errorf("wrapped: %w", &url.Error{
				Op:  http.MethodGet,
				URL: "https://example.test/resource",
				Err: errors.New("Requested Interaction Not Found"),
			}),
			expected: true,
		},
		{
			name: "url error different message",
			err: &url.Error{
				Op:  http.MethodGet,
				URL: "https://example.test/resource",
				Err: errors.New("dial tcp timeout"),
			},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := IsVCRReplayMissError(testCase.err)
			if actual != testCase.expected {
				t.Fatalf("expected %t, got %t (err=%v)", testCase.expected, actual, testCase.err)
			}
		})
	}
}

func TestIsVCRReplayMissErrorDeprecated(t *testing.T) {
	testCases := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "nil error",
			err:      nil,
			expected: false,
		},
		{
			name:     "wrapped replay miss message",
			err:      fmt.Errorf("polling failed: %+v", errors.New(VCRInteractionNotFoundErrMsg)),
			expected: true,
		},
		{
			name: "url error replay miss",
			err: &url.Error{
				Op:  http.MethodGet,
				URL: "https://example.test/resource",
				Err: errors.New(VCRInteractionNotFoundErrMsg),
			},
			expected: true,
		},
		{
			name:     "different error",
			err:      errors.New("unsupported protocol scheme"),
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := IsVCRReplayMissErrorDeprecated(testCase.err)
			if actual != testCase.expected {
				t.Fatalf("expected %t, got %t (err=%v)", testCase.expected, actual, testCase.err)
			}
		})
	}
}

func TestIsVCRReplaying(t *testing.T) {
	testCases := []struct {
		name     string
		client   *Client
		expected bool
	}{

		{
			name:     "default transport mode",
			client:   NewClient("https://localhost", "Example", "2020-01-01"),
			expected: false,
		},
		{
			name: "record mode",
			client: &Client{
				TransportMode: TransportModeVCRRecord,
			},
			expected: false,
		},
		{
			name: "replay mode",
			client: &Client{
				TransportMode: TransportModeVCRReplay,
			},
			expected: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := IsVCRReplaying(testCase.client)
			if actual != testCase.expected {
				t.Fatalf("expected %t, got %t", testCase.expected, actual)
			}
		})
	}
}
