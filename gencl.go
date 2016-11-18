package main


import (
	"fmt"
	"strings"
	"regexp"
	"io/ioutil"
	"path/filepath"
)

func rep(s string) string {
	return strings.Title(s[1:len(s)])
}

func main() {
	matches, _ := filepath.Glob("./tmp/src/github.com/jhaynie/go-github-protobuf/github/*")
	messageRe := regexp.MustCompile("message (\\w+)")

	fmt.Println("package github")
	fmt.Println("import (")
	fmt.Println(`	"fmt"`)
	fmt.Println(`	"github.com/golang/protobuf/descriptor"`)
	fmt.Println(`	d "google.golang.org/genproto/protobuf"`)
	fmt.Println(")")

	fmt.Println("func GetGithubDescriptors() (map[string]*d.DescriptorProto, map[string]string) {")
	fmt.Println("	var results = make(map[string]*d.DescriptorProto)")
	fmt.Println("	var filenames = make(map[string]string)")

	for _, match := range matches {
		if strings.HasSuffix(match, ".proto") {
			buf, _ := ioutil.ReadFile(match)
			str := string(buf)
			for _, m := range messageRe.FindAllString(str, -1) {
				cn := strings.Replace(m, "message ", "", -1)
				vn := strings.ToLower(cn)
				fmt.Printf("	_, %s := descriptor.ForMessage(&%s{})\n", vn, cn)
				fmt.Printf(`	v%s := fmt.Sprintf(".github.%%`, vn)
				fmt.Printf(`s", *%s.Name)`, vn)
				fmt.Println()
				fmt.Printf(`	results[v%s] = %s`, vn, vn)
				fmt.Println("\n")
				fmt.Printf(`	filenames[v%s] = "%s"`, vn, filepath.Base(match))
				fmt.Println("\n")
			}
		}
	}

	fmt.Println("	return results, filenames")
	fmt.Println("}")
}
