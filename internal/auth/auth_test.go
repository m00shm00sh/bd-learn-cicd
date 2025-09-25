package auth
import (
	"testing"
	"net/http"
)

func TestGetAPIKey(t *testing.T) {
	cases := []struct{
		descr string
		inputFactory func () http.Header
		exp string
		expErr error
	}{
		{
			"no auth",
			func () http.Header {
				h := make(http.Header)
				return h
			},
			"",
			ErrNoAuthHeaderIncluded,
		},{
			"short header",
			func () http.Header {
				h := make(http.Header)
				h.Add("Authorization", "Fred")
				return h
			},
			"",
			ErrMalformedAuthHeader,
		},{
			"invalid header",
			func () http.Header {
				h := make(http.Header)
				h.Add("Authorization", "Foo bar")
				return h
			},
			"",
			ErrMalformedAuthHeader,
		},{
			"valid header",
			func () http.Header {
				h := make(http.Header)
				h.Add("Authorization", "ApiKey 123")
				return h
			},
			"123",
			nil,
		},
	}
	for i, tc := range cases {
		got, err := GetAPIKey(tc.inputFactory())
		if err != tc.expErr || got != tc.exp {
			t.Errorf("[%d] %s: expected (%v,%v) but got (%v,%v)\n", i, tc.descr, tc.exp, tc.expErr, got, err)
		}
	}
}
		

		
