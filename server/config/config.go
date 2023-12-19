package config
import (
 "fmt"
 "os"
 "github.com/joho/godotenv"
)

//Get env value from key
func Config(key string) string {
    err := godotenv.Load(".env") //Load env file

    if err != nil {
        fmt.Print("Error loading .env file")
    }
    
    return os.Getenv(key)
}