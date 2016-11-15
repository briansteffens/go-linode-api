package main

import "log"
import "fmt"
import "strings"
import "io/ioutil"

func main() {
    token, err := ioutil.ReadFile("linode-token")

    if (err != nil) {
        panic(err)
    }

    linode := LinodeClient{strings.TrimSpace(string(token))}

    var result LinodeResult
    err = linode.Request("linode/instances", &result)

    //var result LinodeResult
    //err := request("linode/instances", &result)

    //var result DatacentersResult
    //err := linode.Request("datacenters", &result)

    //var result DistributionResult
    //err := request("distributions", &result)

    //var result KernelResult
    //err := request("kernels", &result)

    if (err != nil) {
        log.Fatal(err)
        return
    }

    fmt.Println(result)
    fmt.Println(result.TotalPages)
}
