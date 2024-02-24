package chromedriver

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
)

var driverPath = filepath.Join(os.TempDir(), "chromedriver")

var chromeBrowser [][]string = [][]string{
	{"chromium-browser", "--version"},
	{"chromium", "--version"},
	{"google-chrome", "--version"},
}

func parseChromeVersion(line []byte) string {
	re := regexp.MustCompile(`.* (\d+)\.(\d+)\.(\d+).*`)
	ss := re.FindSubmatch(line)

	fmt.Printf("%q\n", ss)
	if len(ss) != 4 {
		return ""
	}
	return string(ss[1])
}

func chromeVersion() (version string) {
	var line []byte
	var err error
	for _, chrome := range chromeBrowser {
		cmd := exec.Command(chrome[0], chrome[1:]...)
		line, err = cmd.CombinedOutput()
		if err == nil {
			break
		}
	}
	if err != nil {
		return "120"
	}
	return parseChromeVersion(line)
}

func latestRelease() (version string) {
	var url = "https://googlechromelabs.github.io/chrome-for-testing/LATEST_RELEASE_"
	chromeVersion := chromeVersion()

	res, err := http.Get(url + chromeVersion)
	if err != nil {
		return ""
	}
	defer res.Body.Close()

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return ""
	}
	return string(buf)
}

func targetArch() (target string, err error) {
	archTable := map[string]map[string]string{
		"linux": {
			"amd64": "linux64",
		},
		"darwin": {
			"arm64": "mac-arm64",
			"amd64": "mac-x64",
		},
		"windows": {
			"amd64": "win64",
			"386":   "win32",
		},
	}

	archs, osSupported := archTable[runtime.GOOS]
	if !osSupported {
		return "", fmt.Errorf("not supported: %s", runtime.GOOS)
	}

	target, archSupported := archs[runtime.GOARCH]
	if !archSupported {
		return "", fmt.Errorf("not supported on %s: %s", runtime.GOOS, runtime.GOARCH)
	}

	return target, nil
}

func Download() (string, error) {
	version := latestRelease()
	target, err := targetArch()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://storage.googleapis.com/chrome-for-testing-public/%s/%s/chromedriver-%s.zip", version, target, target)
	log.Printf("download from: %s", url)

	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("failed to http.Get(url): %v\n", err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("failed to io.ReadAll(res.Body): %v\n", err)
		return "", err
	}

	r, err := zip.NewReader(bytes.NewReader(body), res.ContentLength)
	if err != nil {
		log.Printf("failed to zip.NewReader(...): %v\n", err)
		return "", err
	}

	var binaryPath string
	var licensePath string
	for _, file := range r.File {
		savePath := filepath.Join(os.TempDir(), filepath.Base(file.Name))
		dst, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return "", err
		}
		defer dst.Close()
		src, err := file.Open()
		if err != nil {
			log.Printf("failed to file.Open(): %v", err)
			return "", err
		}
		defer src.Close()

		io.Copy(dst, src)

		log.Printf("saved: %s", savePath)
		switch savePath {
		case "/tmp/chromedriver":
			binaryPath = savePath
		case "/tmp/LICENSE.chromedriver":
			licensePath = savePath
		}

	}
	if err := os.Remove(licensePath); err == nil {
		log.Printf("deleted: %v\n", licensePath)
	} else {
		log.Printf("failed to delete: %v\n", licensePath)
	}
	return binaryPath, nil
}
