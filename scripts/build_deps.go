package main

import (
       "fmt"
       "os"
       "strings"
       "os/exec"
       "path/filepath"
)

func main() {
     pwd, _ := os.Getwd()

     out, err := exec.Command("find", ".", "-name", "*.ins").Output()
     if err != nil {
     	panic(err)
     }

     files := strings.Split(string(out[:]), "\n")

     for _, file := range files {
     	 if file != "" {
     	    tmp := pwd + "/" + filepath.Dir(file)

	    err := os.Chdir(tmp)
	    if err != nil {
	      panic(err)
	    }

	    ins_files, err := filepath.Glob("*.ins")
	    if err != nil {
	       panic(err)
	    }

	    for _, ins_file := range ins_files {
	    	file_name := strings.ReplaceAll(ins_file, ".ins", "")

	    	if _, err = os.Stat(file_name + ".sty"); !os.IsNotExist(err) {
		   os.Remove(file_name + ".sty")
		}

	    	out, err := exec.Command("pdflatex", "-no-shell-escape", "-interaction nonstopmode", ins_file).Output()
		if err != nil {
		   fmt.Println(string(out[:]))
		   continue
		} else {
		   fmt.Println("Successfully builded", file_name + ".sty")
		}

		if _, err := os.Stat("/usr/share/texmf/tex/latex/" + file_name); os.IsNotExist(err) {
		   err := os.MkdirAll("/usr/share/texmf/tex/latex/" + file_name, os.ModePerm)
		   if err != nil {
		      panic(err)
		   }
	   	}

		_ , error := exec.Command("cp", file_name + ".sty", "/usr/share/texmf/tex/latex/" + file_name + "/").Output()
		if error != nil {
		   panic(error)
     		} else {
		  _ , error := exec.Command("texhash").Output()
		  if error != nil {
		     panic(error)
		  } else {
		     fmt.Println("Updating tex hash ...")
		  }
		}
		
	    }
	    
	 }
     }
}