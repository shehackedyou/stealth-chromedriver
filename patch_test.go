package chromedriver

import (
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"testing"
	"time"

	// TODO: So stupid, get rid of this shit, we just have to do a fucking
	// comparison
	"github.com/stretchr/testify/require"
)

func TestPatcherLatest(t *testing.T) {
	// CI is slow
	RequestTimeout = 30 * time.Second

	p, err := New("", 0)
	require.NoError(t, err, "create patcher")

	path, err := p.Patch()
	require.NoError(t, err, "patch")
	t.Log(path)

	file, err := os.Open(path)
	require.NoError(t, err, "open driver")

	driver, err := io.ReadAll(file)
	require.NoError(t, err, "read driver")

	re := regexp.MustCompile("cdc_.{22}")
	require.Equal(t, false, re.Match(driver))

	_, err = exec.Command(path, "--version").Output()
	require.NoError(t, err, "execute")
}

func TestPatcherVersionPin(t *testing.T) {
	// CI is slow
	RequestTimeout = 30 * time.Second

	p, err := New("", 105)
	require.NoError(t, err, "create patcher")

	path, err := p.Patch()
	require.NoError(t, err, "patch")
	t.Log(path)

	file, err := os.Open(path)
	require.NoError(t, err, "open driver")

	driver, err := io.ReadAll(file)
	require.NoError(t, err, "read driver")

	re := regexp.MustCompile("cdc_.{22}")
	require.Equal(t, false, re.Match(driver))

	output, err := exec.Command(path, "--version").Output()
	require.NoError(t, err, "check version")
	require.Equal(t, true, strings.Contains(string(output), "105"))
}
