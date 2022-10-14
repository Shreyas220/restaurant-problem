package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

type Logs struct {
	Log []struct {
		EaterID    string `json:"eater_id,"`
		FoodmenuID string `json:"foodmenu_id"`
	} `json:"Logs"`
}

func main() {
	//read the log
	ans, err := restaurant("log.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ans)

}

func restaurant(adr string) ([]string, error) {
	//open the log file
	jsonFile, err := os.Open(adr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	logs := Logs{}
	json.Unmarshal([]byte(byteValue), &logs)

	food_id := make(map[string][]string)
	freq := make(map[string]int)

	//goes through logs
	for _, log := range logs.Log {
		if food_id[log.FoodmenuID] != nil {
			//if similar entry found returns err
			for _, eaterid := range food_id[log.FoodmenuID] {
				if eaterid == log.EaterID {
					return []string{}, errors.New("similar entries found")
				}
			}
			food_id[log.FoodmenuID] = append(food_id[log.FoodmenuID], log.EaterID)
			freq[log.FoodmenuID] = freq[log.FoodmenuID] + 1
		} else if food_id[log.FoodmenuID] == nil {
			food_id[log.FoodmenuID] = append(food_id[log.FoodmenuID], log.EaterID)
			freq[log.FoodmenuID] = 1
		}
	}

	keys := make([]string, 0, len(freq))
	for key := range freq {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return freq[keys[i]] > freq[keys[j]]
	})

	ans := []string{keys[0], keys[1], keys[2]}
	return ans, nil

}
