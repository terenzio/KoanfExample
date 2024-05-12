package main

//func main() {
//	// Load .env file from the same directory
//	//if err := godotenv.Load(); err != nil {
//	//	log.Fatalf("Error loading .env file: %v", err)
//	//}
//	//
//	//// Retrieve the DB_HOST environment variable
//	//dbHost := os.Getenv("DB_HOST")
//	//dbUser := os.Getenv("DB_USER")
//	//dbPass := os.Getenv("DB_PASS")
//	//
//	//fmt.Println("DB Host:", dbHost)
//	//fmt.Println("DB User:", dbUser)
//	//fmt.Println("DB Password:", dbPass)
//
//	// Create a new koanf instance with a default delimiter "_" to separate nested keys in environment variables.
//	k := koanf.New(".")
//
//	// Load environment variables using the env provider.
//	// The env provider takes an optional config to specify a delimiter used for key separation.
//	// For example, if your environment variables are of the form "DB_HOST", use "_" as the delimiter.
//	if err := k.Load(env.Provider("", ".", func(s string) string {
//		return s
//	}), env.Parser()); err != nil {
//		log.Fatalf("error loading config: %v", err)
//	}
//
//	// Print the loaded DB configuration
//	fmt.Println("DB Host:", k.String("DB_HOST"))
//	fmt.Println("DB User:", k.String("DB_USER"))
//	fmt.Println("DB Password:", k.String("DB_PASS"))
//
//}
