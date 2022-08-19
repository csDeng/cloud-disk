package test

import (
	"encoding/base64"
	"testing"
)

func TestBase64(t *testing.T) {
	s := "jTZFP8cjKpeGKkvDS0X0B+vLwbDnsQFfaImHuQYxq+NN/I3Q9kq7By9A3MhIJPolFvY+hCgOGLH/zVdQRInrHOBT5Hiq7zIyLZjG6U2QF0fHjog3jsANy8YanJyXNDitn0iJ6NIusD7z2+r8NaJOZ31ric3Iqm4q1rtp0rXi0oaBvp19vDLQDGFrStrZ8NtYLQ8u29rYApqADF92gcnUXjarp7KMbV30TT/2JQO3ut3I3Zd6AQRMpzSBYveUD7MM65q/IN0twUxl1+yGV+T2T5SMqulVUWfZlj0AzK6R5Iu7y12tx3gQxDRwX5xZgwk/JkgvfTUChi5rHmWRKCW560n9CL0S7tjBMuWvt6M0iaEgN42r4w=="
	res, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
	t.Logf("%s \r\n %x", res, res)
}
