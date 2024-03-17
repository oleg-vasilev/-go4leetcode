package lc71

import (
	"path/filepath"
	"strings"
)

// 71. Simplify Path

// Given a string path, which is an absolute path (starting with a slash '/')
// to a file or directory in a Unix-style file system, convert it to the
// simplified canonical path.
//
// In a Unix-style file system, a period '.' refers to the current directory,
// a double period '..' refers to the directory up a level,
// and any multiple consecutive slashes (i.e. '//') are treated
// as a single slash '/'. For this problem, any other format
// of periods such as '...' are treated as file/directory names.
//
// The canonical path should have the following format:
//
// The path starts with a single slash '/'.
// Any two directories are separated by a single slash '/'.
// The path does not end with a trailing '/'.
// The path only contains the directories on the path from the root directory
// to the target file or directory (i.e., no period '.' or double period '..')
// Return the simplified canonical path.

// standard library, you know xD
func simplifyPathStd(path string) string {
	return filepath.Clean(path)
}

func simplifyPath(path string) string {
	var stack = make([]string, 0)
	var parts = strings.Split(path, "/")

	for _, folder := range parts {
		// ignore single dots and empty strings
		if folder == "." || folder == "" {
			continue
		}
		if folder == ".." {
			// double dot removes previous path part
			if len(stack) > 0 { // (if we have at least one)
				stack = stack[:len(stack)-1]
			}
		} else {
			// everything else just appends on top
			stack = append(stack, folder)
		}
	}
	return "/" + strings.Join(stack, "/")
}
