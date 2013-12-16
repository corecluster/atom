package main

import (
  linuxproc "github.com/c9s/goprocinfo/linux"
  "os/exec"
  "log"
)

type MachineInfo struct{
  cpus int
  memory int
}

func main(){

  cpus, _ := exec.Command("nproc").Output()
  memInfo, err :=  linuxproc.ReadMemInfo("/proc/meminfo")
  if err != nil {
    log.Println("Error on Read MemInfo")
  }

  log.Printf("CPU Core: %s", cpus)
  log.Println("Meminfo: ", memInfo["MemFree"])//.Free)


}
