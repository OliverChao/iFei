package forever

import (
	"fmt"
	"github.com/b3log/lute"
	"github.com/sirupsen/logrus"
)

var (
	isLoad bool = false
)
var Markdown *lute.Lute

// maybe lute engine can be a kind of mini service in docker
//thus you can use requests to access it
func LoadMarkdownEngine() {
	Markdown = lute.New()
	isLoad = true
	logrus.Info("[lute] load successfully")
}

func GetMarkdownEngine() (*lute.Lute, error) {
	if isLoad {
		return Markdown, nil
	} else {
		logrus.Fatal("[Lute] does not Load...please call LoadMarkdownEngine function first...")
		return nil, fmt.Errorf("engine does not load error")
	}
}
