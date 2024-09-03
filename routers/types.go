package routers

type SeasonMeta struct {
	Category    int    `json:"category"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	Mid         int    `json:"mid"`
	Name        string `json:"name"`
	SeasonId    int    `json:"season_id"`
	Total       int    `json:"total"`
}

type Archive struct {
	Bvid     string `json:"bvid"`
	Pic      string `json:"pic"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
}

type Page struct {
	PageNum  int `json:"page_num"`
	PageSize int `json:"page_size"`
	Total    int `json:"total"`
}

type VedioListResponseData struct {
	Data struct {
		Archives []Archive  `json:"archives"`
		Page     Page       `json:"page"`
		Meta     SeasonMeta `json:"meta"`
	} `json:"data"`
}
type Season struct {
	Archives   []Archive  `json:"archives"`
	Meta       SeasonMeta `json:"meta"`
	RecentAids []int      `json:"recent_aids"`
}

type SeasonListResponseData struct {
	Data struct {
		ItemList struct {
			Page        Page     `json:"page"`
			SeasonsList []Season `json:"seasons_list"`
		} `json:"items_lists"`
	} `json:"data"`
}
