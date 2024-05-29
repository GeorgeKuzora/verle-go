package weeek

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"verle_go/pkg/config"
	"verle_go/pkg/tasks"

	"github.com/joho/godotenv"
)

type WeeekTaskFetcher struct {
	project tasks.ProjectType
}

func GetWeekDayTasks(day string, workplace config.Workplace) string {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	link := "https://api.weeek.net/public/v1/tm/tasks?day=" + day + "&projectId=" + fmt.Sprint(workplace.WeeekProject.ProjectNumber) + "&all=0"
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, link, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", getWeeekToken())
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status: got %v", res.Status))
	}
	fmt.Println(res.Header.Get("Content-Type"))
	return readResponse(res.Body)
}

func getWeeekToken() string {
	m, err := godotenv.Read(".env")
	if err != nil {
		panic(err)
	}
	token := m["WEEEK_API_TOKEN"]
	token = "Bearer " + token
	return token
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
