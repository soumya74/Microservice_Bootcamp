//https://www.thepolyglotdeveloper.com/2017/07/consume-restful-api-endpoints-golang-application/
// https://redis.io/commands/hmget

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

type test_struct struct {
	Test string
}

func oneWay(w http.ResponseWriter, r *http.Request) {
	var f func()
	var t *time.Timer
	fmt.Println("test worker")
	conn := redis.NewClient(&redis.Options{
		Addr:         "172.25.16.126:6379",
		Password:     "", // no password set
		DB:           0,  // use default DB
		MaxRetries:   3,
		IdleTimeout:  5 * time.Minute,
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 5 * time.Minute,
	})

	pong, err := conn.Ping().Result()
	fmt.Println("pong : ", pong)
	if err != nil {
		fmt.Println("error : ", err)
	}
	defer conn.Close()

	f = func() {
		fmt.Println("Calling STG...")
		response, err := http.Get("http://172.25.16.126:10000/stg/tokens/1")

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))

			fmt.Println("Calling Hasher... Data : ", data)
			url := "http://172.25.16.126:10001/hasher"
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			defer resp.Body.Close()
			if err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", err)
			} else {
				data, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("DATA: ", string(data))

				//! Unmarshaling the data obtained from hasher into a map,
				dat := make(map[string]string)
				err = json.Unmarshal(data, &dat)
				if err != nil {
					fmt.Println("error_46 : ", err)
					panic(err)
				}
				fmt.Println("Unmarshal done", dat["hash"])
				val := strings.HasPrefix(dat["hash"], "0")

				//! publishing 1 if the hash is "Lucky Hash", else 0
				fmt.Println("Lucky Hash : ", val)
				fmt.Println("Hash : ", dat["hash"])
				pubsub := conn.Subscribe("mychannel1")
				defer pubsub.Close()
				resp := conn.Publish("hashChannel", val).Err()
				if resp != nil {
					fmt.Println(resp)
					panic(err)
				}
			}
		}

		t = time.AfterFunc(time.Duration(1)*time.Second, f)
	}
	t = time.AfterFunc(time.Duration(1)*time.Second, f)
	defer t.Stop()
	time.Sleep(time.Minute)
}

func main() {
	http.HandleFunc("/", oneWay)
	http.ListenAndServe(":8080", nil)
}
