package utils

import "testing"

func TestGetByCodeResolvesLegacyGalicianCode(t *testing.T) {
	for _, code := range []string{"gl", "gy"} {
		if got := GetByCode(code).Language; got != "Galician" {
			t.Errorf("GetByCode(%q) = %s, want Galician", code, got)
		}
	}
}

func TestNormalizeCode(t *testing.T) {
	cases := map[string]string{"gy": "gl", "gl": "gl", "id": "id", "": ""}
	for code, want := range cases {
		if got := NormalizeCode(code); got != want {
			t.Errorf("NormalizeCode(%q) = %q, want %q", code, got, want)
		}
	}
}
