package main

import (
  linuxproc "github.com/c9s/goprocinfo/linux"
  etcd "github.com/coreos/go-etcd/etcd"
  "os/exec"
  "strconv"
  "log"
)

type Machine struct{
  cpus string
  memory string
}

func (machine *Machine) AvailableCpus() string {
  nproc, _ := exec.Command("nproc").Output()
  cpus := string(nproc)
  machine.cpus = cpus
  return machine.cpus
}

func (machine *Machine) AvailableMemory() string {
  memInfo, err :=  linuxproc.ReadMemInfo("/proc/meminfo")
  if err != nil {
    log.Fatal("Error on Read MemInfo")
  }
  machine.memory = strconv.FormatUint(memInfo["MemFree"], 10)
  return machine.memory
}

func MachineInfo() *Machine {
  machine := new(Machine)
  machine.AvailableCpus()
  machine.AvailableMemory()
  return machine
}



func main(){

  machine := MachineInfo()

  log.Println("Machine Info> ", machine)
  log.Printf("CPU Core: %s", machine.cpus)
  log.Println("Meminfo: ", machine.memory)

  etcd := etcd.NewClient(nil)
  etcd.SetDir("/my_org/my_app", 2000)
  etcd.Set("/my_org/my_app/config", "{'ip_address': 127.0.0.1, 'port': 8181}", 10000)

}
