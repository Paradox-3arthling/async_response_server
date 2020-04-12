package hook

// 1. func to for making hook server
//    make sure it's https for safety
func CreateServer(port, path string) (string, error){
	new_str := port + path
	// generate url to be used
	// make server
	return new_str, nil
}
// 2. func handler for setting up hook