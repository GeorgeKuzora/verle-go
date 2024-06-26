package weeek

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"verle_go/pkg/tasks"

	"github.com/joho/godotenv"
)

const (
	envVarFileName = ".env"
	tokenVarName   = "WEEEK_API_TOKEN"
)

var client = &http.Client{
	Timeout: 30 * time.Second,
}

var token = getWeeekToken()

func getWeeekToken() string {
	m, err := godotenv.Read(envVarFileName)
	if err != nil {
		log.Fatalf("can't read a file with environment variables %s", envVarFileName)
	}
	token, ok := m[tokenVarName]
	if !ok {
		log.Fatalf("can't find variable %s in a file %s", tokenVarName, envVarFileName)
	}
	token = "Bearer " + token
	return token
}

var projects = map[tasks.ProjectType]int{
	tasks.IMF120:   2,
	tasks.Trobart:  14,
	tasks.Drip:     4,
	tasks.Capsule:  5,
	tasks.Assembly: 6,
}

type TaskFetcher struct {
	Project tasks.ProjectType
}

func (tf TaskFetcher) Fetch(dates []tasks.Date) ([]tasks.Tasks, error) {
	urlPrefix := "https://api.weeek.net/public/v1/tm/tasks?day="
	projPrefix := "&projectId="
	allPrefix := "&all="
	allValue := "0"

	projNum, ok := projects[tf.Project]
	if !ok {
		log.Printf("can't fetch data for project %v, unknown project", tf.Project)
		return nil, fmt.Errorf("can't fetch data for project %v, unknown project", tf.Project)
	}
	p := fmt.Sprint(projNum)

	t := make([]tasks.Tasks, len(dates))

	for i, date := range dates {
		d, err := date.String()
		if err != nil {
			log.Printf("can't fetch data for a date, date is not provided")
			continue
		}
		link := urlPrefix + d + projPrefix + p + allPrefix + allValue
		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, link, nil)
		if err != nil {
			log.Printf("can't create a GET request to link: %s", link)
			continue
		}
		req.Header.Add("Authorization", token)
		res, err := client.Do(req)
		if err != nil {
			log.Printf("can't get a response from a link: %s", link)
			continue
		}
		if res.StatusCode != http.StatusOK {
			log.Printf("unexpected status: got %v in response from link: %s", res.Status, link)
			continue
		}
		defer res.Body.Close()
		resData, err := readResponse(res.Body)
		if err != nil {
			log.Printf("can't read response body from link %s", link)
			continue
		}
		dateTasks, err := UnmarshalDateTasks(resData)
		if err != nil {
			log.Printf("can't unmarshal response data from link %s", link)
			continue
		}
		t[i] = dateTasks
	}
	return t, nil
}

func (tf TaskFetcher) FetchById(id int) (tasks.Task, error) {
	urlPrefix := "https://api.weeek.net/public/v1/tm/tasks/"
	link := urlPrefix + fmt.Sprint(id)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, link, nil)

	var task tasks.Task

	if err != nil {
		log.Printf("can't create a GET request to link: %s", link)
		return task, fmt.Errorf("can't create a GET request to link: %s", link)
	}
	req.Header.Add("Authorization", token)
	res, err := client.Do(req)
	if err != nil {
		log.Printf("can't get a response from a link: %s", link)
		return task, fmt.Errorf("can't get a response from a link: %s", link)
	}
	if res.StatusCode != http.StatusOK {
		log.Printf("unexpected status: got %v in response from link: %s", res.Status, link)
		return task, fmt.Errorf("unexpected status: got %v in response from link: %s", res.Status, link)
	}
	defer res.Body.Close()
	resData, err := readResponse(res.Body)
	if err != nil {
		log.Printf("can't read response body from link %s", link)
		return task, fmt.Errorf("can't read response body from link %s", link)
	}
	task, err = UnmarshalTask(resData)
	if err != nil {
		log.Printf("can't unmarshal response data from link %s", link)
		return task, fmt.Errorf("can't unmarshal response data from link %s", link)
	}
	return task, nil
}

func readResponse(r io.Reader) (string, error) {
	buf := make([]byte, 2048)
	bytes := []byte{}
	for {
		n, err := r.Read(buf)
		bytes = append(bytes, buf[:n]...)
		if err == io.EOF {
			var out string = string(bytes)
			return out, nil
		}
		if err != nil {
			log.Println("can't read response from weeek", err)
			return "", err
		}
	}
}

func UnmarshalDateTasks(data string) (tasks.Tasks, error) {
	var tasks tasks.Tasks
	err := json.Unmarshal([]byte(data), &tasks)
	if err != nil {
		log.Printf("Can't Unmarshal weeek Tasks data from a string %s", data)
		return tasks, err
	}
	return tasks, nil
}

func UnmarshalTask(data string) (tasks.Task, error) {
	type Response struct {
		Success bool       `json:"success"`
		Task    tasks.Task `json:"task"`
	}
	var response Response
	var task tasks.Task
	err := json.Unmarshal([]byte(data), &response)
	if err != nil {
		log.Printf("Can't Unmarshal weeek Task data from a string %s", data)
		return task, err
	}
	if !response.Success {
		log.Printf("unsuccessful Fetch for a string %s", data)
		return task, fmt.Errorf("unsuccessful Fetch for a string %s", data)
	}
	task = response.Task
	return task, nil
}
