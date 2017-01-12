package main

import "log"
import "fmt"
import "strings"
import "io/ioutil"

func main() {
	token, err := ioutil.ReadFile("linode-token")

	if err != nil {
		panic(err)
	}

	linode := LinodeClient{strings.TrimSpace(string(token))}

	var result LinodeResult
	err = linode.Request("linode/instances", &result)

	//var result LinodeResult
	//err := request("linode/instances", &result)

	var datacenters DatacentersResult
	err2 := linode.Request("datacenters", &datacenters)

	//var result DistributionResult
	//err := request("distributions", &result)

	//var result KernelResult
	//err := request("kernels", &result)

	if err != nil {
		log.Fatal(err)
		return
	}

	if err2 != nil {
		log.Fatal(err2)
		return
	}

	fmt.Println(result)
	fmt.Println(result.TotalPages)
	fmt.Println(" ")
	fmt.Println(datacenters)
	fmt.Println(datacenters.TotalPages)
}
