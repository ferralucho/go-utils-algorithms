package json

import (
	"encoding/json"
	"fmt"
)

const resp = `
{
    "id": "MLB1076588183",
    "statistics": {
        "1": {
            "id": "MLB1076588183",
            "statistics": {
                "standard": {
                    "weight": {
                        "standardDeviation": 112.0,
                        "mean": 682.0
                    },
                    "height": {
                        "standardDeviation": 20.0,
                        "mean": 28.17
                    },
                    "width": {
                        "standardDeviation": 7.02,
                        "mean": 68.0
                    },
                    "length": {
                        "standardDeviation": 8.16,
                        "mean": 80.5
                    },
                    "n": 30
                },
                "anomalous": {
                    "weight": {
                        "standardDeviation": 119.0,
                        "mean": 241.0
                    },
                    "height": {
                        "standardDeviation": 5.22,
                        "mean": 6.85
                    },
                    "width": {
                        "standardDeviation": 7.37,
                        "mean": 23.5
                    },
                    "length": {
                        "standardDeviation": 9.49,
                        "mean": 26.3
                    },
                    "n": 3
                }
            },
            "statsLastUpdate": "01-08-2018T20:09:47.000Z",
            "categoryId": null,
            "inItem": false
        },
        "2": {
            "id": "MLB1076588183",
            "statistics": {
                "standard": {
                    "weight": {
                        "standardDeviation": 112.0,
                        "mean": 1102.0
                    },
                    "height": {
                        "standardDeviation": 20.0,
                        "mean": 12.17
                    },
                    "width": {
                        "standardDeviation": 7.02,
                        "mean": 15.2
                    },
                    "length": {
                        "standardDeviation": 8.16,
                        "mean": 30.5
                    },
                    "n": 30
                },
                "anomalous": {
                    "weight": {
                        "standardDeviation": 119.0,
                        "mean": 241.0
                    },
                    "height": {
                        "standardDeviation": 5.22,
                        "mean": 6.85
                    },
                    "width": {
                        "standardDeviation": 7.37,
                        "mean": 23.5
                    },
                    "length": {
                        "standardDeviation": 9.49,
                        "mean": 26.3
                    },
                    "n": 3
                }
            },
            "statsLastUpdate": "02-08-2018T13:16:51.000Z",
            "categoryId": null,
            "inItem": false
        }
    },
    "nextKey": null
}
`

type ItemResponse struct {
	ID         string               `json:"id"`
	Statistics map[string]Statistic `json:"statistics"`
	NextKey    interface{}          `json:"nextKey"`
}

type Statistic struct {
	Key        string `json:"key"`
	ID         string `json:"id"`
	Statistics struct {
		Standard struct {
			Weight struct {
				StandardDeviation float64 `json:"standardDeviation"`
				Mean              float64 `json:"mean"`
			} `json:"weight"`
			Height struct {
				StandardDeviation float64 `json:"standardDeviation"`
				Mean              float64 `json:"mean"`
			} `json:"height"`
			Width struct {
				StandardDeviation float64 `json:"standardDeviation"`
				Mean              float64 `json:"mean"`
			} `json:"width"`
			Length struct {
				StandardDeviation float64 `json:"standardDeviation"`
				Mean              float64 `json:"mean"`
			} `json:"length"`
			N int `json:"n"`
		} `json:"standard"`
		Anomalous struct {
			Weight struct {
				StandardDeviation float64 `json:"standardDeviation"`
				Mean              float64 `json:"mean"`
			} `json:"weight"`
			Height struct {
				StandardDeviation float64 `json:"standardDeviation"`
				Mean              float64 `json:"mean"`
			} `json:"height"`
			Width struct {
				StandardDeviation float64 `json:"standardDeviation"`
				Mean              float64 `json:"mean"`
			} `json:"width"`
			Length struct {
				StandardDeviation float64 `json:"standardDeviation"`
				Mean              float64 `json:"mean"`
			} `json:"length"`
			N int `json:"n"`
		} `json:"anomalous"`
	} `json:"statistics"`
	StatsLastUpdate string      `json:"statsLastUpdate"`
	CategoryID      interface{} `json:"categoryId"`
	InItem          bool        `json:"inItem"`
}

func main() {
	var ir ItemResponse
	if err := json.Unmarshal([]byte(resp), &ir); err != nil {
		panic(err)
	}

	stats := make([]Statistic, 0)
	for k, s := range ir.Statistics {
		s.Key = k

		stats = append(stats, s)
	}

	fmt.Println(stats)
}
