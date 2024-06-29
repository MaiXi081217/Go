package main
import ("fmt"
		"os"
		//"log"
		"io"
	)

func read_file(file_path string)(string,error){
	file, err := os.Open(file_path)
	if err != nil{
		return "",err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil{
		return "",err
	}

	return  string(content), nil
}


func main() {  
	var file_path string
	for{
		fmt.Print("path: ")
		_, err := fmt.Scanln(&file_path) // 使用 Scanln 读取，并解析到 file_path 变量中  
		if err != nil {  
			fmt.Println("Error reading file path:", err)  
			continue  
			
		}  
		if file_path == "quit"{
			fmt.Println("Bye")
			return
		}
		content, err := read_file(file_path) 
		if err != nil {  
			fmt.Println("Error reading file:", err)  
			continue 
		}  
		fmt.Println(content) 
	}  

}