package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func GetTime() string {
	return time.Now().Format("00:00:00")
}

func GetEnv(key string) (val string) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	val = os.Getenv(key)
	return
}

func PrintBanner(version string) {
	banner := ` 
 _   _ ______ _       _____ _______    _______ _____ 
| \ | |  ____| |     / ____|__   __|/\|__   __/ ____|
|  \| | |__  | |    | (___    | |  /  \  | | | (___  
| .   |  __| | |     \___ \   | | / /\ \ | |  \___ \ 
| |\  | |    | |____ ____) |  | |/ ____ \| |  ____) |
|_| \_|_|    |______|_____/   |_/_/    \_\_| |_____/  
                                Version: ` + version
	fmt.Println(banner)
}
