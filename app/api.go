package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	notionDatabaseURL := loadEnv("NOTION_DATABASE_URL")
	url := "https://api.notion.com/v1/databases/" + notionDatabaseURL + "/query"

	payload := strings.NewReader("{\"page_size\":3}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+loadEnv("NOTION_API"))
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

}

func loadEnv(keyName string) string {
	err := godotenv.Load(".env")
	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
	// .envの SAMPLE_MESSAGEを取得して、messageに代入します。
	message := os.Getenv(keyName)

	return message
}
