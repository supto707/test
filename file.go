package main

import ("os/exec")

func main(){
	
	file := exec.Command("open" , "CV.pdf")

	file.Start()
	
	
}
