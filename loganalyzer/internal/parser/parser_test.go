package parser

import "testing"

func TestCleanAndExtractFields(t *testing.T) {
	// Table of test cases
	tests := []struct {
		name       string
		logline    string
		shouldPass bool
		wantMethod string
		wantCode   int
	}{
		{
			name:       "valid GET request",
			logline:    `192.168.1.105 - - [01/Jan/2026:08:02:48 +0000] "GET /index.html HTTP/1.1" 200 1289`,
			shouldPass: true,
			wantMethod: "GET",
			wantCode:   200,
		},
		{
			name:       "invalid IP address",
			logline:    `999.999.999.999 - - [01/Jan/2026:08:02:48 +0000] "GET /index.html HTTP/1.1" 200 1289`,
			shouldPass: false,
		},
		{
			name:       "invalid status code",
			logline:    `192.168.1.105 - - [01/Jan/2026:08:02:48 +0000] "GET /index.html HTTP/1.1" ABC 1289`,
			shouldPass: false,
		},
		{
			name:       "empty line",
			logline:    "",
			shouldPass: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cleanAndExtractFields(tt.logline)

			if !tt.shouldPass {
				if err == nil {
					t.Fatalf("expected error for invalid line, got nil")
				}
				if got != nil {
					t.Fatalf("expected nil for invalid line, got %+v", got)
				}
				return
			}

			// If it should pass, verify extracted fields
			if got == nil {
				t.Fatalf("expected valid LogFields, got nil")
			}
			if got.Method != tt.wantMethod {
				t.Errorf("Method: got %s, want %s", got.Method, tt.wantMethod)
			}
			if got.StatusCode != tt.wantCode {
				t.Errorf("StatusCode: got %d, want %d", got.StatusCode, tt.wantCode)
			}
		})
	}
}
