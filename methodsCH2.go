package main

import (
    "fmt"
)

type VirtualMachine struct {
    ip int
    hostname string
    ram int
}

func (vm VirtualMachine) DisplayDetails() {
    fmt.Printf("ip: %v\n", vm.ip)
    fmt.Printf("hostname: %v\n", vm.hostname)
    fmt.Printf("ram: %v\n", vm.ram)
}

func (vm *VirtualMachine) AddRam() {
    fmt.Println("Adding 8gb of ram")
    vm.ram += 8
}

func main() {
    myVm := VirtualMachine{
        ip: 123455678,
        hostname: "myHost",
        ram: 8,
    }

    myVm.DisplayDetails()

    myVm.AddRam()

    myVm.DisplayDetails()
}
