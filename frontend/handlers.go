package frontend

import (
  "os"
  "fmt"
  "io"
	"net/http"

	"github.com/spf13/viper"

	"github.com/pdxjohnny/telem/log"
)

func Example(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<html>
<head>
    <title>Upload file</title>
</head>
<body>
<form enctype="multipart/form-data" action="/upload" method="post">
      <input type="file" name="file" />
      <input type="submit" value="upload" />
</form>
</body>
</html>`))
}

func Upload(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(32 << 20)

	file, handler, err := req.FormFile("file")
	log.PrintError("frontend upload", err)
	defer file.Close()

	fmt.Fprintf(w, "%v", handler.Header)

  os.MkdirAll(viper.GetString("upload"), 0700)
	f, err := os.OpenFile(
		viper.GetString("upload")+handler.Filename,
		os.O_WRONLY|os.O_CREATE,
		0600,
	)
	log.PrintError("frontend upload", err)
	defer f.Close()

	io.Copy(f, file)
}
