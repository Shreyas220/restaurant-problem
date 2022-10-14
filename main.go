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
	ans, err := getFood_ids("log.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The top three item ordered are  \n", "1st foodmenu_id:", ans[0], "  2nd foodmenu_id:", ans[1], "  3rd foodmenu_id:", ans[2])
}

func getFood_ids(adr string) ([]string, error) {
	//Open the log file
	jsonFile, err := os.Open(adr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	logs := Logs{}
	json.Unmarshal([]byte(byteValue), &logs)

	//stores key as foodmenu_id and eater_id as value
	food_id := make(map[string][]string)
	//counts the frequency of food_id ordered
	food_freq := make(map[string]int)

	//Goes through logs to make sure there are no repeated entry
	for _, log := range logs.Log {
		if food_id[log.FoodmenuID] != nil {
			//if similar entry found returns err
			for _, eaterid := range food_id[log.FoodmenuID] {
				if eaterid == log.EaterID {
					return []string{}, errors.New("similar entries found")
				}
			}
			food_id[log.FoodmenuID] = append(food_id[log.FoodmenuID], log.EaterID)
			food_freq[log.FoodmenuID] = food_freq[log.FoodmenuID] + 1
		} else if food_id[log.FoodmenuID] == nil {
			food_id[log.FoodmenuID] = append(food_id[log.FoodmenuID], log.EaterID)
			food_freq[log.FoodmenuID] = 1
		}
	}

	//sorting the food_frequency to get top three
	keys := make([]string, 0, len(food_freq))
	for key := range food_freq {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return food_freq[keys[i]] > food_freq[keys[j]]
	})

	ans := []string{keys[0], keys[1], keys[2]}
	return ans, nil

}
