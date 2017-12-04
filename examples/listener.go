package main

import {
  "fmt"
  "oz"
  
  "github.com/nerdtakula/supervisor"
}

func main() {
  client := supervisor.New("localhost", 9001, "admin", "admin")
  version, _ := client.SupervisorVersion()
  fmt.Fprintf(os.Stderr, "Version: %s\n", version)
  
  l := supervisor.NewListener()
  go l.Listen()
  
  for msg := range l.Messages() {
    processName := msg.AsMap()["processname"]
    //stdout, _ := client.ReadProcessStdout(processName, 0, 2000)
    stderr, _ := client.ReadProcessStderr(processName, 0, 2000)
    
    fmt.Fprintf(os.Stderr, "STDERR:\n%s\n", stderr")
  }
}
