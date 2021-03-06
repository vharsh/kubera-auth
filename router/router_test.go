package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	// TODO: Understand why this hack doesn't work

	/*for _, key := range []string{"JWT_SECRET", "ADMIN_USERNAME", "ADMIN_PASSWORD", "CONFIGMAP_NAME", "DB_SERVER", "PORTAL_URL"} {
		_ = os.Setenv(key, "dummy")
	}
	*/
}

func TestHealthCheck(t *testing.T) {
	// prevents gin from polluting the logs
	gin.SetMode(gin.TestMode)
	// create a dummy context
	fakeReq := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(fakeReq)
	HealthCheck(c)
	if c.Writer.Status() != http.StatusOK {
		t.Errorf("Expected: %v, Got: %v", http.StatusOK, c.Writer.Status())
	}
}

/*
func TestCallbackRequest(t *testing.T) {
	type args struct {
		state string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "GitHub login",
			args: args{
				types.GithubState,
			},
		},
	}
	for _, tt := range tests {
		os.Setenv("PORTAL_URL", "https://blog.google/")
		// TODO: Fake the provider later
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
*/
