package file

/*
	INFORMATION ABOUT CODE
	-------------------------

	All of the important method that i used in this code you'll find it in the documentation of OS.
	https://golang.org/pkg/os/
*/

import (
	"os"
)


func ReadAndGetFile(typelist string) string {

	if compareTypeList(typelist){
		file, errFile := os.Open("./file/lists/list_" + typelist + ".txt")
		defer file.Close()
		checkError(errFile)

		//Reading a file

		fileInfo, errFileInfo := file.Stat()
		checkError(errFileInfo)

		//fmt.Printf("The file is %d bytes long \n", fileInfo.Size())

		dataByte := make([]byte, fileInfo.Size())
		_, err := file.Read(dataByte)
		checkError(err)
		
		return string(dataByte)

	}else{
		return ""
	}
}

func compareTypeList (typelist string) bool{
	switch typelist{
	case "english", "spanish":
		return true
	default:
		return false
	}
}

func checkError (err error){
	if err != nil {
		panic(err)
	}
}

