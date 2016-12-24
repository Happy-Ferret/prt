package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// Checkout checks out a repo
func Checkout(branch, loc string) error {
	cmd := exec.Command("git", "checkout", branch)
	cmd.Dir = loc

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Could not git checkout repo!")
	}

	return nil
}

// Clean cleans a repo
func Clean(loc string) error {
	cmd := exec.Command("git", "clean", "-f")
	cmd.Dir = loc

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Could not git clean repo!")
	}

	return nil
}

// Clone clones a repo
func Clone(url, branch, loc string) error {
	cmd := exec.Command("git", "clone", "--depth", "1", "-b", branch, url)
	cmd.Dir = loc

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Could not git clone repo!")
	}

	return nil
}

// Diff checks a repo for differences
func Diff(branch, loc string) []string {
	cmd := exec.Command("diff", "--name-status", branch)
	cmd.Dir = loc
	b := new(bytes.Buffer)
	cmd.Stdout = b

	cmd.Run()
	// TODO: For some reason this cmd always exits with a 2 exit code
	//if err != nil {
	//	return fmt.Errorf("Could not git clone repo!")
	//}

	diff := b.String()
	if len(diff) < 1 {
		return []string{}
	}

	// Make output pretty
	diff = strings.Replace(diff, "A", "Adding", 1)
	diff = strings.Replace(diff, "C", "Copying", 1)
	diff = strings.Replace(diff, "D", "Deleting", 1)
	diff = strings.Replace(diff, "M", "Modifying", 1)
	diff = strings.Replace(diff, "R", "Renaming", 1)

	return strings.Split(diff, "\n")
}

// Fetch fetches a repo
func Fetch(loc string) error {
	cmd := exec.Command("git", "fetch", "--depth", "1")
	cmd.Dir = loc

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Could not git fetch repo!")
	}

	return nil
}

// Reset resets a repo
func Reset(branch, loc string) error {
	cmd := exec.Command("git", "reset", "--hard", "origin/"+branch)
	cmd.Dir = loc

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Could not git reset repo!")
	}

	return nil
}