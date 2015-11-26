package main

import(
	
	"fmt"
	"net/http"
	"encoding/json"
	"encoding/hex"
	"crypto/sha256"
)

type KeyPair struct	{

	Key int `json:"key"`
	Value string `json:"value"`
}

func main() {
	fmt.Println("Client Running")
	pair := KeyPair{}
	i := 0
	j := 0
	ch := 'a'
	var temp uint8
	temp = 0
	for i < 10	{
		pair = KeyPair{i+1,string(ch)}
		pairJson,_ := json.Marshal(pair)
		hash := sha256.New()
    	hash.Write([]byte(pairJson))
    	hexValue := hex.EncodeToString(hash.Sum(nil))
    	for j<64	{
    		temp+=hexValue[j]
    		j++
    	}
    	temp = temp%3
    	url := fmt.Sprintf("http://localhost:300%d/keys/%d/%s",temp,pair.Key,pair.Value)
    	fmt.Println(url)
    	client := http.Client{}
    	req,err := http.NewRequest("PUT",url,nil)
    	if err != nil	{
			fmt.Print(err)
		}
		res,err := client.Do(req)
		defer res.Body.Close()
    	if err != nil	{
			fmt.Print(err)
		}
		ch++
		i++
		j=0
		temp=0
	}

	// Sample Get Request from Server on Port 3002
	res,_ := http.Get("http://localhost:3002/keys/1")
    jsonDecoder := json.NewDecoder(res.Body)
    jsonDecoder.Decode(&pair)
    fmt.Printf("Output from Server on [Port:3002] for 1st KeyPair::%+v\n",pair)

}