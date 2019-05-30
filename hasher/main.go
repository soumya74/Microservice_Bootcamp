// main.go Hasher
//https://stackoverflow.com/questions/10701874/generating-the-sha-hash-of-a-string-using-golang
//https://gist.github.com/andreagrandi/97263aaf7f9344d3ffe6

package main

import (
	"fmt"
	"log"
	"net/http"
	//"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"

	"github.com/gorilla/mux"
)

type test_struct struct {
	Test string
}

//! This function will generate the hash value
//!  and will return the hash as string
func Hasher(token_value string) string {
	h2 := sha256.New()
	h2.Write([]byte(token_value))
	fmt.Println("hashed value ==", h2.Sum(nil))
	sha1_hash := hex.EncodeToString(h2.Sum(nil))

	fmt.Println(token_value, sha1_hash)
	return sha1_hash
}

//! This function will receive the request and will publish the
//! hash value to the redis channel named "hashChannel"
func ReturnHashedNumber(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)
	if r.Method != "POST" {
		fmt.Println("hasher called with Method other than POST")
		return
	}

	fmt.Println("hasher called with Method POST")
	data, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(data))
	dat := make(map[string]string)
	err := json.Unmarshal(data, &dat)
	if err != nil {
		fmt.Println("error_46 : ", err)
		panic(err)
	}
	fmt.Println("Unmarshal done", dat["token"])

	fmt.Println("Hasher Input : ", dat["token"])
	sha1_hash := Hasher(dat["token"])
	m := make(map[string]string)
	m["hash"] = sha1_hash
	js, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error Marshaling")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("JSON Output ", js)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//! will only accept POST request
func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/hasher", ReturnHashedNumber).Methods("POST")
	log.Fatal(http.ListenAndServe(":10001", myRouter))
}

func main() {
	fmt.Println("Hasher Service Running...")
	HandleRequests()
}
