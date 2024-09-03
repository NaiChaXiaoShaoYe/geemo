package routers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ActionGetSeasonList(c *gin.Context) {
	mid := c.Query("mid")

	seasonInfoArray := make([]map[string]interface{}, 0)

	page := 1
	for {
		data := getSeasonList(mid, page)
		page++
		if len(data.Data.ItemList.SeasonsList) == 0 {
			break
		}
		for _, value := range data.Data.ItemList.SeasonsList {
			seasonInfo := make(map[string]interface{})
			seasonInfo["id"] = value.Meta.SeasonId
			seasonInfo["name"] = value.Meta.Name
			seasonInfo["count"] = value.Meta.Total

			seasonInfoArray = append(seasonInfoArray, seasonInfo)
		}
	}

	c.JSON(200, gin.H{
		"data": gin.H{
			"list": seasonInfoArray,
		},
	})
}

func getSeasonList(mid string, page int) SeasonListResponseData {

	params := url.Values{}
	params.Set("mid", mid)
	params.Set("page_num", strconv.Itoa(page))
	params.Set("page_size", "1")
	u, _ := url.Parse("https://api.bilibili.com/x/polymer/web-space/seasons_series_list")
	u.RawQuery = params.Encode()
	fmt.Println(u.String())
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	for key, value := range FakeHeaders {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	body, _ := io.ReadAll(res.Body)
	var data SeasonListResponseData
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}
	return data
}
