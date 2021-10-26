package dragon

import (
	"log"
	"net/http"
	"os"
)

type (
	static struct {
		fileSystem http.FileSystem
	}
)

var s static

// ServeStaticFile will reading and serving index1.html file in specific folder as requested in argument
// 	example :
// 		r := dragon.NewRouter()
//		r.ServeHTTPFile("./static")
func ServeStaticFile(dir string) {
	if isDir(dir) {
		s = static{fileSystem: http.Dir(dir)}
	}
}

// setupStaticFile
func (s *static) setupStaticFile(rw http.ResponseWriter, req *http.Request) {
	fs := http.FileServer(s.fileSystem)
	fs.ServeHTTP(rw, req)
}

// isDir validating directory, returning true or false
func isDir(dir string) bool {
	stat, err := os.Stat(dir)
	if err != nil {
		log.Fatal(err)
		return false
	}

	if stat.IsDir() {
		return true
	}

	return false
}
