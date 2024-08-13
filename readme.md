## How to Run
You can grab a copy of a Go compiler from their official website: https://go.dev

Or use a package manager like "snap" https://snapcraft.io/go

To confirm successfult instalation of Go you can run command "go version".

To run the wheel of fortune you should navigate to the root directory of the project and run "go run ."

Then in the console you should see two lines being printed informing you that the API server is running on port 8081 and the front-end is running on 8080.

All of the environment variables have default values, to check them in a case you wish to modify them navigate to "config/env.go"
