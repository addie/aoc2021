/*

Copyright (c) 2021 - Present. Blend Labs, Inc. All rights reserved
Blend Confidential - Restricted

*/

package data

import (
	"aoc2021/secret"
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	URL = "https://adventofcode.com/%d/day/%d"
	Filename = "data/day_%d_%d_data"
)

type Configuration struct {
	SessionCookie string `json:"session-cookie"`
	Output        string `json:"output"`
	Payload       string `json:"data"`
	Year          int    `json:"year"`
	Day           int    `json:"day"`
	Force         bool   `json:"-"`
	Wait          bool   `json:"-"`
}

func Post(day, year, level, result int) error {
	payload := fmt.Sprintf(`{"level": %d, "answer": %d}`, level, result)
	config := &Configuration{
		SessionCookie: secret.SessionID,
		Year:          year,
		Day:           day,
		Payload:       payload,
	}
	_, err := query(http.MethodPost, config)
	if err != nil {
		return err
	}
	return nil
}

func Get(day, year int, filename string) ([]int, error) {
	config := &Configuration{
		SessionCookie: secret.SessionID,
		Output:        filename,
		Year:          year,
		Day:           day,
	}
	resp, err := query(http.MethodGet, config)
	if err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			return nil, err
		}
		log.Println(err)
	}
	defer resp.Body.Close()

	flags := os.O_WRONLY | os.O_CREATE
	if config.Force {
		flags |= os.O_TRUNC
	} else {
		flags |= os.O_EXCL
	}

	file, err := os.OpenFile(config.Output, flags, 0666)
	if os.IsExist(err) {
		log.Printf("file '%s' already exists; use '-force' to overwrite\n", config.Output)
		file, _ := os.OpenFile(config.Output, os.O_RDONLY, 0666)
		defer file.Close()
		return readFile(file)
	} else if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return nil, err
	}

	return readFile(file)
}

func readFile(file *os.File) ([]int, error) {
	var res []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		digit, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		res = append(res, digit)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

func query(method string, config *Configuration) (*http.Response, error) {
	client := new(http.Client)

	path := "/input"
	var body io.Reader
	if method == http.MethodPost {
		path = "/answer"
		body = strings.NewReader(config.Payload)
	}

	req, err := http.NewRequest(method, fmt.Sprintf(URL+path, config.Year, config.Day), body)
	if err != nil {
		return nil, err
	}

	cookie := new(http.Cookie)
	cookie.Name, cookie.Value = "session", config.SessionCookie
	req.AddCookie(cookie)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	return resp, nil
}
