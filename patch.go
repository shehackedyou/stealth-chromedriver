package chromedriver

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"
)

var binaryName = "chromedriver"

var (
	re      = regexp.MustCompile("cdc_.{22}")
	letters = []byte("abcdefghijklmnopqrstuvwxyz")
)

type Binary struct {
	path     string
	platform string
	DataDir  string
}

func Patch(bPath string) (b Binary, err error) {
	b.path = bPath
	log.Printf("PATCHING BINARY")
	//driver = patchDriver(driver)

	//if _, err := os.Stat(bPath); err == nil {
	//	if err := os.Remove(bPath); err != nil {
	//		return b, fmt.Errorf("failed to remove old driver '%s': %w", bPath, err)
	//	} else {
	//		log.Printf("removed existing binary...")
	//	}
	//} else {
	//	log.Printf("no binary to remove, continuing...")
	//}

	// TODO: In our case here, its NOT writeFile, we need to READ the file then
	// maybe write it after we tweak the []byte

	// TODO: We do this to move it to bin folder
	if err = b.dataDir(); err != nil {
		return b, err
	}

	log.Printf("PATCHING BINARY doing the find and replace for the cdc...")
	if err = b.patch(); err != nil {
		log.Printf("patching FAILED!")
		return b, err
	} else {
		log.Printf("patching complete!")
	}

	return b, nil
}

func (b *Binary) dataDir() error {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Printf("failed to locate home folder: %v", err)
		return err
	}

	b.platform = runtime.GOOS
	switch b.platform {
	case "linux":
		b.DataDir = path.Join(home, ".local/share/undetected_chromedriver")
	case "darwin":
		b.DataDir = path.Join(home, "appdata/roaming/undetected_chromedriver")
	case "windows":
		b.DataDir = path.Join(home, "Library/Application Support/undetected_chromedriver")
	default:
		log.Printf("os is not supported")
		return fmt.Errorf("OS not supported: %s", runtime.GOOS)
	}

	if _, err := os.Stat(b.DataDir); os.IsNotExist(err) {
		if err := os.MkdirAll(b.DataDir, 0750); err != nil {
			log.Printf("failed to create data dir: %v", err)
			return fmt.Errorf("failed to create data dir '%s': %w", b.DataDir, err)
		} else {
			log.Printf("created data dir: %v", b.DataDir)
		}
	} else {
		log.Printf("os.Stat(b.DataDir) failed so must exist already")

		log.Printf("data dir already exists, now should be checking for the binary so we can REMOVE it!")
		// os.Stat(b.DataDir + binaryName); os.IsNotExist(err) {\

		//}
	}

	return nil
}

func (b *Binary) patch() error {

	if _, err := os.Stat(b.path); os.IsNotExist(err) {
		log.Printf("binary we just downloaded is MISSING!")
		return fmt.Errorf("binary is missing!!!")
	} else {
		log.Printf("binary EXISTS lets open it!")
	}

	log.Printf("attempting to load chromedriver binary at: %v", b.path)
	binary, err := os.ReadFile(b.path)
	if err != nil {
		log.Printf("binary failed to load from b.path: %v", err)
		return err
	}

	if !re.Match(binary) {
		log.Printf("failed to find cdc in the chromedriver binary!")
		return nil
	}

	newBinary := re.ReplaceAll(binary, randomCDC())
	if bytes.Equal(newBinary, binary) {
		log.Printf("failed to make changes to the binary! something went terribly wrong!")
		panic(fmt.Errorf("failed to patch the binary!"))
	} else {
		log.Printf("success! binary has been patched!")
		log.Printf("Assigning binary path; must come after a move & therefore after a patch!")
		b.path = path.Join(b.DataDir, binaryName)

		if err = os.WriteFile(b.path, newBinary, 0755); err != nil {
			log.Printf("failed to write new binary!!! OH NOES: %v", b.path)
			return err
		}
	}

	log.Printf("new PATCHED binary path b.path len(newBinary)(%v): %v", len(newBinary), b.path)
	log.Printf("SUCCESS!")
	return nil
}

// TODO It should NOT be fixed size, thats too easy stupid
// AND THERE ARE TWO SPOTS that need updating, dont look for just 1 like an
// IDIOT!
func randomCDC() []byte {

	cdc := make([]byte, 26)

	if _, err := rand.Read(cdc); err != nil {

		// Shouldn't happen, but just in case.
		return []byte("xvx_plxklvnobnowmrmiIMvqlb")
	}

	for i, val := range cdc {

		cdc[i] = letters[int(val)%len(letters)]
	}

	cdc[2] = cdc[0]
	cdc[3] = '_'
	cdc[20] = strings.ToUpper(string(cdc[20]))[0]
	cdc[21] = strings.ToUpper(string(cdc[21]))[0]

	log.Printf("new cdc value: %v", string(cdc))

	return cdc
}
