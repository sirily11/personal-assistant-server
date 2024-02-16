package environments

import "os"

var (
	// AdminApiKey is the API key/root password for the admin.
	AdminApiKey = os.Getenv("ADMIN_API_KEY")
	// GptEndpoint is the endpoint for the GPT-3/4 API.
	GptEndpoint = os.Getenv("GPT_ENDPOINT")
	// GptKey is the key for the GPT-3/4 API.
	GptKey = os.Getenv("GPT_KEY")
	// Version is the version of the application.
	Version = os.Getenv("VERSION")
	// DatabaseUrl is the URL of the database.
	DatabaseUrl  = os.Getenv("DATABASE_URL")
	DatabaseName = os.Getenv("DATABASE_NAME")
)
