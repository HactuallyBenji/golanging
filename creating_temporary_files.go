package main

import (
   "fmt"
   "io/ioutil"
   "os"
)

func main() {
   // create a new temporary file
   f, err := ioutil.TempFile("U:/github/golanging/test_directory", "file")
   if err != nil {
      panic(err)
   }

   // the OS will clean up this temporary file by itself at some point but we can do it anyway for safety
   defer os.Remove(f.Name())

   // we can see the pattern added to the end of the temp filename
   fmt.Println("File name:", f.Name())

   // add data to the file
   f.WriteString("Hello!\n")
   f.WriteString("This file will be deleted once the program is done executing.\n")
   f.WriteString("The advantage of using temp files is that they don't pollute the file system.\n")
   f.WriteString("Don't use temp files to persist data because they will be deleted by the OS at some point.\n")

   // read the new file
   tempContents, _ := ioutil.ReadFile(f.Name())
   fmt.Print(string(tempContents))

   // createa new temporary directory
   d, err := ioutil.TempDir("U:/github/golanging/test_directory", "tempdir")
   if err != nil {
      panic(err)
   }

   // the OS will delete temp directory at some point but we can do it anyway.
   defer os.RemoveAll(d)

   // print the temporary directory name
   fmt.Println("Directory name:", d)
}
