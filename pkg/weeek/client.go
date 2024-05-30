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
	if ok == false {
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
	project tasks.ProjectType
}

func (tf TaskFetcher) Fetch(dates []tasks.Date) ([]tasks.Tasks, error) {
	urlPrefix := "https://api.weeek.net/public/v1/tm/tasks?day="
	projPrefix := "&projectId="
	allPrefix := "&all="
	allValue := "0"

	projNum, ok := projects[tf.project]
	if ok == false {
		log.Printf("can't fetch data for project %v, unknown project", tf.project)
		return nil, fmt.Errorf("can't fetch data for project %v, unknown project", tf.project)
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
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			log.Printf("unexpected status: got %v in response from link: %s", res.Status, link)
			continue
		}
		resData := readResponse(res.Body)
		dateTasks := UnmarshalDateTasks(resData)
		t
	}

}

func readResponse(r io.Reader) string {
	buf := make([]byte, 2048)
	bytes := []byte{}
	for {
		n, err := r.Read(buf)
		bytes = append(bytes, buf[:n]...)
		if err == io.EOF {
			var out string = string(bytes)
			return out
		}
		if err != nil {
			panic(err)
		}
	}
}

func UnmarshalDateTasks(data string) DateTasks {
	var tasks DateTasks
	err := json.Unmarshal([]byte(data), &tasks)
	if err != nil {
		log.Fatalln("Can't Unmarshal weeek data", err)
	}
	return tasks
}
