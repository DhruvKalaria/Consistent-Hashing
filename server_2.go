package main

import(
	
	"log"
	"fmt"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
)

type KeyPair struct	{

	Key int `json:"key"`
	Value string `json:"value"`
}

var KV = make(map[int]KeyPair)

func main() {
	
	fmt.Println("Starting server on Port : 3001")
	router := httprouter.New()
	router.Handle("PUT","/keys/:keypair_id/:keypair_value", InsertKeyPair)
	router.Handle("GET","/keys/:keypair_id",GetKeyPair)
	router.Handle("GET","/keys", GetAllKeyPair)
	log.Fatal(http.ListenAndServe(":3001", router))
}

func InsertKeyPair(w http.ResponseWriter, r *http.Request, ps httprouter.Params)	{

	var key int
	var value string

	key,_ = strconv.Atoi(ps.ByName("keypair_id"))
	value = ps.ByName("keypair_value")
	
	keyPair := KeyPair{key,value}
	KV[key] = keyPair
	//fmt.Println(KV[key])
}

func GetKeyPair(w http.ResponseWriter, r *http.Request, ps httprouter.Params)	{

	key,_ := strconv.Atoi(ps.ByName("keypair_id"))
	elem, ok := KV[key]
	if ok {

		w.WriteHeader(http.StatusCreated)
    	json.NewEncoder(w).Encode(elem)
		fmt.Println(elem)
	}
}

func GetAllKeyPair(w http.ResponseWriter, r* http.Request, _ httprouter.Params)	{
	
	store := make([]KeyPair, len(KV))
	index := 0
	for _, value := range KV {
    	store[index] = value
    	index++
	}
	json.NewEncoder(w).Encode(store)
}