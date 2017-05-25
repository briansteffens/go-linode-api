package main

//import "log"
import "fmt"
import "strings"
import "io/ioutil"

func main() {
	token, err := ioutil.ReadFile("linode-token")

	if err != nil {
		panic(err)
	}

	linode := LinodeClient{strings.TrimSpace(string(token))}

	filter := And(
		Comparison{"label", Eq, "bws"},
	)

	fmt.Println(filter.Json())

	var result LinodeResult
	err = linode.Request("linode/instances", filter, &result)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)

/*

	x := Filter {LogicalAnd, []FilterNode{
		Comparison{"Label", Contains, "some"},
		Filter {LogicalOr, []FilterNode{
			Comparison{"label", Neq, "dorp"},
			Comparison{"TotalTransfer", Gt, "123123"},
		}},
	}}

	y := And(
		Comparison{"Label", Contains, "some"},
		Comparison{"Label", Eq, "yup"},
		Or(
			Comparison{"label", Neq, "dorp"},
			Comparison{"TotalTransfer", Gt, "123123"},
		),
	)



	fmt.Println(x)
	fmt.Println(y)

	fmt.Println(y.Json());



*/


	//var result LinodeResult
	//err = linode.Request("linode/instances", &result)

	//var result LinodeResult
	//err := request("linode/instances", &result)

	//var result RegionsResult
	//err = linode.Request("regions", &result)

	//var result DistributionResult
	//err = linode.Request("linode/distributions", &result)

	//var result KernelResult
	//err = linode.Request("linode/kernels", &result)

	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}

	//fmt.Println(result)
	//fmt.Println(result.TotalPages)
}
