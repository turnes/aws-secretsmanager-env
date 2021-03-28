package tools

import (
	"os"
	"regexp"
)

var validAwsSecretName = `[/,+,=,@]`

func renameFile(file string) (renamedFile string) {
	var re = regexp.MustCompile(validAwsSecretName)
	renamedFile = re.ReplaceAllString(file, "-")
	return
}

func Save(secretName, content string) (Rerr error) {
	secretName = renameFile(secretName)
	var f *os.File
	f, err := os.OpenFile(secretName, os.O_WRONLY, os.ModePerm)
	if os.IsNotExist(err) {
		f, err = os.Create(secretName)
		if err != nil {
			Rerr = err
			return
		}
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		Rerr = err
	}
	f.Chmod(0640)
	f.Sync()
	return
}

func Load() {

}
