package query

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type ApiError struct {
	Error []string `json:"error"`
}

func host() string {

	api_host := "api.odrone.eu"
	if env := os.ExpandEnv("$ODRONE_API_HOST"); len(env) != 0 {
		api_host = env
	}

	return api_host
}

func Get(path string, result interface{}) (err error) {

	resp, err := http.Get("http://" + host() + path)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {

		var apiError ApiError
		if err = json.Unmarshal(body, &apiError); err != nil {
			return
		}

		errorText := ""
		for _, e := range apiError.Error {
			errorText += e + "\n"
		}

		return errors.New(strings.TrimRight(errorText, "\n"))
	}

	if result != nil {
		if err = json.Unmarshal(body, &result); err != nil {
			return
		}
	}

	return nil
}

func Delete(path string, result interface{}) (err error) {

	req, err := http.NewRequest("DELETE", "http://"+host()+path, nil)
	if err != nil {
		return
	}

	// Submit the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {

		var apiError ApiError
		if err = json.Unmarshal(respBody, &apiError); err != nil {
			return
		}

		errorText := ""
		for _, e := range apiError.Error {
			errorText += e + "\n"
		}

		return errors.New(strings.TrimRight(errorText, "\n"))
	}

	if result != nil {
		if err = json.Unmarshal(respBody, &result); err != nil {
			return
		}
	}

	return nil
}

func Post(path string, form map[string]string, file io.Reader, result interface{}) (err error) {

	var body io.Reader
	contentType := ""

	if file != nil {

		var b bytes.Buffer
		w := multipart.NewWriter(&b)

		fw, err := w.CreateFormFile("file", "file")
		if err != nil {
			return err
		}
		if _, err = io.Copy(fw, file); err != nil {
			return err
		}

		// Add the other fields
		for key, val := range form {

			if fw, err = w.CreateFormField(key); err != nil {
				return err
			}
			if _, err = fw.Write([]byte(val)); err != nil {
				return err
			}
		}

		w.Close()

		body = &b

		contentType = w.FormDataContentType()
	} else {

		values := url.Values{}

		for key, val := range form {

			values.Add(key, val)
		}

		body = strings.NewReader(values.Encode())
	}

	req, err := http.NewRequest("POST", "http://"+host()+path, body)
	if err != nil {
		return
	}

	if len(contentType) != 0 {

		req.Header.Set("Content-Type", contentType)
	}

	// Submit the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {

		var apiError ApiError
		if err = json.Unmarshal(respBody, &apiError); err != nil {
			return
		}

		errorText := ""
		for _, e := range apiError.Error {
			errorText += e + "\n"
		}

		return errors.New(strings.TrimRight(errorText, "\n"))
	}

	if result != nil {
		if err = json.Unmarshal(respBody, &result); err != nil {
			return
		}
	}

	return nil
}

func Put(path string, form map[string]string, file io.Reader, result interface{}) (err error) {

	var body io.Reader
	contentType := ""

	if file != nil {

		var b bytes.Buffer
		w := multipart.NewWriter(&b)

		fw, err := w.CreateFormFile("file", "file")
		if err != nil {
			return err
		}
		if _, err = io.Copy(fw, file); err != nil {
			return err
		}

		// Add the other fields
		for key, val := range form {

			if fw, err = w.CreateFormField(key); err != nil {
				return err
			}
			if _, err = fw.Write([]byte(val)); err != nil {
				return err
			}
		}

		w.Close()

		body = &b

		contentType = w.FormDataContentType()
	} else {

		values := url.Values{}

		for key, val := range form {

			values.Add(key, val)
		}

		body = strings.NewReader(values.Encode())
	}

	req, err := http.NewRequest("PUT", "http://"+host()+path, body)
	if err != nil {
		return
	}

	if len(contentType) != 0 {

		req.Header.Set("Content-Type", contentType)
	}

	// Submit the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {

		var apiError ApiError
		if err = json.Unmarshal(respBody, &apiError); err != nil {
			return
		}

		errorText := ""
		for _, e := range apiError.Error {
			errorText += e + "\n"
		}

		return errors.New(strings.TrimRight(errorText, "\n"))
	}

	if result != nil {
		if err = json.Unmarshal(respBody, &result); err != nil {
			return
		}
	}

	return nil
}
