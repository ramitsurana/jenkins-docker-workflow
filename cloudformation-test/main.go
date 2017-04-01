package main

import (
 "fmt"
 "os/exec"
 "log"
 "os"
 "bufio"
)

func main(){

  arg0 := "jq"
  arg1 := "'.Parameters'" 
  arg2 := "file" 
  arg3 := "'.Resources'"
  arg4 := "'.Mappings'"
  arg5 := "'.Metadata'"  

reader := bufio.NewReader(os.Stdin)
fmt.Print("\n Enter path of your file: ")
file, _ := reader.ReadString('\n')
				
		fmt.Println("Checking Parameters")
        cmd1 := exec.Command(arg0, arg1, file)
        err1 := cmd1.Start()
		if err1 != nil {
            log.Fatal(err1)
			log.Printf("Parameters failed")
			os.Exit(1)			
        }
		err1 = cmd1.Wait()	
	
  fmt.Println("Parameters OK") 

        fmt.Println("Checking Resources")
        cmd2 := exec.Command(arg0, arg3, file)
        err2 := cmd2.Start()
		if err2 != nil {
            log.Fatal(err2)
			log.Printf("Resources failed")
			os.Exit(1)			
        }
		err1 = cmd2.Wait()	
	
  fmt.Println("Resources OK")

  fmt.Println("Checking Mappings")
        cmd3 := exec.Command(arg0, arg4, file)
        err3 := cmd3.Start()
		if err3 != nil {
            log.Fatal(err3)
			log.Printf("Mappings failed")
			os.Exit(1)			
        }
		err3 = cmd3.Wait()	
	
  fmt.Println("Mappings OK")

fmt.Println("Checking Metadata")
        cmd4 := exec.Command(arg0, arg5, file)
        err4 := cmd4.Start()
		if err4 != nil {
            log.Fatal(err4)
			log.Printf("Metadata failed")
			os.Exit(1)			
        }
		err4 = cmd4.Wait()	

fmt.Println("Metadata OK")

fmt.Printf("Basic Checks completed.\nOPTIONS [.Metadata][.Resources] \nList of services: \n * .EcsCluster \n * .EcsElb \n * .ECSAutoScalingGroup \n * .ContainerInstances")	
reader1 := bufio.NewReader(os.Stdin)
fmt.Print("\n Enter your service: ")
service, _ := reader1.ReadString('\n')
fmt.Println(service)  

        cmdservice := exec.Command(arg0, service, arg2)               
        //fmt.Println(cmdservice.Start())        
        errservice := cmdservice.Start()                       
		if errservice != nil {
            log.Fatal(errservice)
			log.Printf("Service failed.Please check the list.")
			os.Exit(1)			
        }
		errservice = cmdservice.Wait()	


}
