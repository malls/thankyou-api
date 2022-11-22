package printful

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// {
//     "sync_product": {
//         "name": "string"
//     },
//     "sync_variants": [
//         {
//             "variant_id": 1,
//             "files": [
//                 {
//                     "url": "string"
//                 }
//             ]
//         }
//     ]
// }

func GetCategories() json.RawMessage {
	authorization := "Bearer " + os.Getenv("PRINTFULTOKEN")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.printful.com/categories", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", authorization)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var objs json.RawMessage
	if err := json.Unmarshal(data, &objs); err != nil {
		panic(err)
	}
	return objs
}

func GetProduct(id string) json.RawMessage {
	authorization := "Bearer " + os.Getenv("PRINTFULTOKEN")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.printful.com/products/"+id, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", authorization)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var objs json.RawMessage
	if err := json.Unmarshal(data, &objs); err != nil {
		panic(err)
	}
	return objs
}

func GetProducts(id string) json.RawMessage {
	authorization := "Bearer " + os.Getenv("PRINTFULTOKEN")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.printful.com/products?category_id="+id, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", authorization)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var objs json.RawMessage
	if err := json.Unmarshal(data, &objs); err != nil {
		panic(err)
	}
	return objs
}

func CreateShirt() {
	authorization := "Bearer " + os.Getenv("PRINTFULTOKEN")

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.printful.com/store/products", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", authorization)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
