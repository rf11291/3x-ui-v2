package web

import (
	"strings"
	"testing"
)

func TestIndexDoesNotRenderOrCheckPanelUpdates(t *testing.T) {
	page, err := htmlFS.ReadFile("html/index.html")
	if err != nil {
		t.Fatalf("read index template: %v", err)
	}

	for _, forbidden := range []string{
		"getPanelUpdateInfo",
		"panelUpdateModal",
		"panel-update-modal",
	} {
		if strings.Contains(string(page), forbidden) {
			t.Errorf("index template must not include %q", forbidden)
		}
	}
}
