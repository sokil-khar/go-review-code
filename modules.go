package modules

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

/*
	req: url address
	object should be fill after function execution
*/
func GetRequest(req string, dto interface{}) {
	resp, err := http.Get(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &dto)
}

func Contains(slice []string, item string) bool {
    set := make(map[string]struct{}, len(slice))
    for _, s := range slice {
        set[s] = struct{}{}
    }

    _, ok := set[item] 
    return ok
}
