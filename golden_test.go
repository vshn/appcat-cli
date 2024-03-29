package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/vshn/appcat-cli/internal/applications"
)

const (
	GOLDEN_PATH = "tests/golden/"
)

var ext = regexp.MustCompile(`\.json$`)

func init() {
	logrus.SetOutput(io.Discard)
}

func TestGolden(t *testing.T) {
	testCases := findTests()
	for testsName, testInstances := range testCases {
		t.Run(testsName, func(t *testing.T) {
			testCase(t, testsName, testInstances)
		})
	}
}

func findTests() map[string]map[string]string {
	c, err := os.ReadDir(GOLDEN_PATH)
	check(err, "finding instances")

	testsMap := make(map[string]map[string]string)
	for _, e := range c {
		if ext.MatchString(e.Name()) && !e.IsDir() {
			testMap := make(map[string]string)
			jsonBuf, err := os.ReadFile(filepath.Join(GOLDEN_PATH, e.Name()))
			check(err, "Reading Json Test File")

			// Unmarshal the JSON data into the map
			err = json.Unmarshal(jsonBuf, &testMap)
			check(err, "Error unmarshalling tests")
			testsMap[ext.ReplaceAllString(e.Name(), "")] = testMap
		}
	}
	return testsMap
}

// Splits and aggregates args to correctly read in json input and space separated values
func splitStringArgs(args string) []string {
	prelimArgs := strings.Fields(args)
	var result []string
	var current string
	isSpecial := false
	for _, arg := range prelimArgs {
		if strings.HasPrefix(arg, "{") || strings.HasPrefix(arg, "\"") {
			isSpecial = true
		}

		if isSpecial {
			if strings.HasSuffix(arg, "}") || strings.HasSuffix(arg, "\"") {
				if strings.HasSuffix(arg, "\"") {
					current = strings.TrimSpace(current)
					current = current[1:]
					current += " " + arg[:len(arg)-1]
					result = append(result, current)
					current = ""
					isSpecial = false
				} else {
					if current == "" {
						result = append(result, arg)
					} else {
						current = strings.TrimSpace(current)
						result = append(result, current+" "+arg)
						current = ""
					}

					isSpecial = false
				}
			} else {
				current += " " + arg
			}

		} else {
			current = strings.TrimSpace(current)
			result = append(result, arg)
		}
	}
	return result
}

func testCase(t *testing.T, instanceName string, instanceFile map[string]string) {
	apps := applications.MakeAppMap()
	old, err := os.Getwd()
	check(err, "determining current working directory")
	dirName := ext.ReplaceAllString(instanceName, "")
	root := filepath.Join(GOLDEN_PATH, dirName)
	check(os.Chdir(root), "changing working directory")

	for testName, testParams := range instanceFile {

		t.Run(testName, func(t *testing.T) {
			args := []string{"appcat-cli"}
			seperatedParams := splitStringArgs(testParams)
			args = append(args, seperatedParams...)
			var logs bytes.Buffer
			logrus.SetOutput(&logs)

			outFile, err := os.Create(testName + ".yaml")
			check(err, "Could not create outFile for test")
			if c := Main(apps, args, outFile); c != 0 {
				t.Errorf("\nappcat-cli exited with code %v while compiling with args '%v' \r\n %s", c, args, logs.String())
			}
		})
	}
	//compare golden test directories with with the directories in the Git index.
	cmd := exec.Command("git", "diff", "--exit-code", "--minimal", "--", ".")
	if _, err := cmd.CombinedOutput(); err != nil {
		t.Errorf("error from git diff: %v", err)
	}
	check(os.Chdir(old), "resetting working directory")
}

func check(err error, context string) {
	if err != nil {
		_, cf, cl, _ := runtime.Caller(0)
		cf = filepath.Base(cf)
		log.Fatalf("%s:%d: Error %s: %v", cf, cl, context, err)
	}
}
