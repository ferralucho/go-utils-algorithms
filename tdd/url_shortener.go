package main

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

// Step 2: Write implementation code to get the tests to pass
type URLShortener struct {
	store map[string]string
	i     int
}

/*
// Step 4: Write the implementaton to get the tests to pass
func (u *URLShortener) Shorten(in string) (string, error) {
	if !strings.Contains(in, ".com") {
		return "", errors.New("invalid")
	}

	if !strings.Contains(in, "https://") {
		in = fmt.Sprintf("https://%s", in)
	}

	u.i++
	val := strconv.Itoa(u.i)
	u.store[val] = in

	return val, nil
}
*/

// Step 5: Refactor the code and make sure that the tests still pass.
//         Use code coverage tools to ensure that no new behaviour was added.
func (u *URLShortener) Shorten(in string) (string, error) {
	_, err := url.ParseRequestURI(in)
	if err != nil {
		return "", errors.New("invalid")
	}

	if !strings.Contains(in, ".") {
		return "", errors.New("invalid")
	}

	u.i++
	val := strconv.Itoa(u.i)
	u.store[val] = in

	return val, nil
}
