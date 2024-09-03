package routers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetVideoList(c *gin.Context) {
	mid := c.Query("mid")
	season_id := c.Query("season_id")

	params := url.Values{}
	params.Set("mid", mid)
	params.Set("season_id", season_id)
	params.Set("page_num", "1")
	params.Set("page_size", "100")
	params.Set("wts", strconv.FormatInt(time.Now().Unix(), 10))

	u, _ := url.Parse("https://api.bilibili.com/x/polymer/web-space/seasons_archives_list")
	u.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
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
	var data VedioListResponseData
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	videoInfoArray := make([]map[string]string, 0)

	for _, value := range data.Data.Archives {
		videoInfo := make(map[string]string)
		videoInfo["bvid"] = value.Bvid
		videoInfo["pic"] = value.Pic
		videoInfo["title"] = value.Title
		videoInfo["duration"] = formatDuration(value.Duration)

		videoInfoArray = append(videoInfoArray, videoInfo)
	}

	c.JSON(200, gin.H{
		"data": gin.H{
			"list": videoInfoArray,
		},
	})
}
