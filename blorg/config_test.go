package blorg

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"testing"
)

func TestBlorg(t *testing.T) {
	config, err := ReadConfig("testdata/blorg.org")
	if err != nil {
		t.Errorf("Could not read config: %s", err)
		return
	}
	committedHashBs, err := ioutil.ReadFile("testdata/public.md5")
	if err != nil {
		t.Errorf("Could not read hash bytes: %s", err)
		return
	}
	if err := config.Render(); err != nil {
		t.Errorf("Could not render: %s", err)
		return
	}
	renderedHashBs, err := exec.Command("bash", "-c", fmt.Sprintf("find %s -type f | sort -u | xargs md5sum", config.PublicDir)).Output()
	if err != nil {
		t.Errorf("Could not hash PublicDir: %s", err)
		return
	}
	rendered, committed := strings.TrimSpace(string(renderedHashBs)), strings.TrimSpace(string(committedHashBs))
	if rendered != committed {
		t.Errorf("PublicDir hashes do not match: '%s' -> '%s'", committed, rendered)
		return
	}
}
