package oexec

import (
	"strings"
	"sync"
	"testing"
	"time"
)

func eq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestExecuteWithOutArguments(t *testing.T) {
	cmd := "echo"
	tmpOutput, returnedError := execute(cmd)
	returnedOutput := strings.TrimSpace(string(tmpOutput))
	if returnedOutput != "" || returnedError != nil {
		t.Fatalf("Incorrect output generated")
	}
}

func TestExecuteWithArguments(t *testing.T) {
	outputStr := "cmd executed"
	cmd := "echo"
	tmpOutput, returnedError := execute(cmd, outputStr)
	returnedOutput := strings.TrimSpace(string(tmpOutput))
	if returnedOutput != outputStr || returnedError != nil {
		t.Fatalf("Incorrect output generated")
	}
}

func TestExecuteWithInvalidCmd(t *testing.T) {
	cmd := "dfgffd"
	returnedOutput, returnedError := execute(cmd)
	if returnedOutput != nil || returnedError == nil {
		t.Fatalf("Incorrect output generated")
	}
}

func TestRemoveEmptyStr(t *testing.T) {
	strs := [][]string{
		{"cd", "", " ", " ", "/folder-name"},
		{"cd", "/folder-name"},
		{"ls", "", " ", "-l", " ", "/folder-name"},
		{"mv", "", " ", "/old-folder-name", " ", " ", " ", "/new-folder-name/path1/path2"},
	}
	expectedStrs := [][]string{
		{"cd", "/folder-name"},
		{"cd", "/folder-name"},
		{"ls", "-l", "/folder-name"},
		{"mv", "/old-folder-name", "/new-folder-name/path1/path2"},
	}

	for i := range strs {
		cleanedStrs := removeEmptyStr(strs[i])
		if !eq(cleanedStrs, expectedStrs[i]) {
			t.Fatalf("Incorrect result returned: expected %v got %v", expectedStrs[i], cleanedStrs)
		}
	}
}

func TestProcessCmdStr(t *testing.T) {
	cmds := []string{
		"ls -l",
		"ls    -l",
		"mv /old-folder-name /new-folder-name/path1/path2",
		"mv /old-folder-name     /new-folder-name/path1/path2",
	}
	cmdNames := []string{
		"ls",
		"ls",
		"mv",
		"mv",
	}
	cmdArgs := [][]string{
		{"-l"},
		{"-l"},
		{"/old-folder-name", "/new-folder-name/path1/path2"},
		{"/old-folder-name", "/new-folder-name/path1/path2"},
	}

	for i := range cmds {
		cmdName, args := processCmdStr(cmds[i])
		if cmdName != cmdNames[i] {
			t.Fatalf("Incorrect command name generated")
		}
		for j := range cmdArgs[i] {
			if args[j] != cmdArgs[i][j] {
				t.Fatalf("Incorrect command argument generated")
			}
		}
	}
}

func TestRun(t *testing.T) {
	cmd := "echo 'test'"
	out := run(cmd)
	if strings.TrimSpace(string(out.Stdout)) != "'test'" || out.Stderr != nil {
		t.Fatalf("Incorrect output generated")
	}
}
func TestRunInvalidCmd(t *testing.T) {
	cmd := "slkmfdlkds"
	out := run(cmd)
	if out.Stdout != nil || out.Stderr == nil {
		t.Fatalf("Incorrect output generated")
	}
}

func TestRunParallel(t *testing.T) {
	var wg sync.WaitGroup
	outputs := make([]*Output, 4)
	wg.Add(1)
	go runParallel(&wg, "echo test", 3, outputs)
	wg.Wait()
	if strings.TrimSpace(string(outputs[3].Stdout)) != "test" || outputs[3].Stderr != nil {
		t.Fatalf("Incorrect output generated")
	}
}

func TestSeries(t *testing.T) {
	startTime := time.Now()
	outs := Series("sleep 2s", "sleep 3s", "sdfds", "sleep 2s")
	elapsedTime := time.Since(startTime)
	if elapsedTime < 7 || outs[2].Stderr == nil {
		t.Fatalf("Incorrect output generated")
	}
}

func TestParallel(t *testing.T) {
	startTime := time.Now()
	outs := Parallel("sleep 2s", "sleep 3s", "sdfds", "sleep 2s")
	elapsedTime := time.Since(startTime)
	if elapsedTime.Seconds() > 3.1 || outs[2].Stderr == nil {
		t.Fatalf("Incorrect output generated")
	}
}
