package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

// Summary recorde Summary file string
var Summary string
var err error

func postmanMkdir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		check(err)
	}
}
func readMe(item *Item) {
	for index := 0; index <= item.level; index++ {
		Summary += "  "
	}
	Summary += "* [" + item.Name + "](" + item.path + "/README.md)\n"
}
func readJSON(Collection *Item) {
	var Items = root(Collection)
	postmanMkdir(Collection.path)
	Summary += "\n## " + Collection.Name + "\n\n"
	Summary += "* [" + Collection.Name + "](" + Collection.path + "/README.md)\n"
	writeReadME(Collection)
	for _, Item := range Items {
		itemNode(Item, Collection.path)
	}
	writeSummary()
}
func itemNode(Node *Item, path string) {
	var Items = child(Node, path)
	// Node.printName()
	switch ntype := nodeTpye(Node); ntype {
	case "floder":
		postmanMkdir(Node.path)
		readMe(Node)
		for _, Item := range Items {
			itemNode(Item, Node.path)
		}
	case "apis":
		var path = Node.path + ".md"
		writeFile(Node, path)
	case "api":
		var path = Node.path + "/README.md"
		writeFile(Node, path)
	case "emptyfloder":
		var path = Node.path + ".md"
		writeFile(Node, path)
	default:
		fmt.Printf("%s.????????\n", ntype)
	}

}

func writeFile(Parent *Item, path string) {
	var file string
	file = "# "
	file += Parent.Name
	file += "\n"
	if _, err := os.Stat(path); os.IsExist(err) {
		fmt.Println("IsExist...", path)
	}
	if !strings.Contains(path, "README") {

		for index := 0; index <= Parent.level; index++ {
			Summary += "  "
		}
		Summary += "* [" + Parent.Name + "](" + path + ")\n"
	}

	if childHasRequest(Parent) {
		f, err := os.Create(path)
		check(err)
		defer f.Close()
		f.Sync()
		Parent.writeTitle(f)
		for _, Item := range Parent.Items {
			if hasRequest(Item) {
				Item.writing(f)
			}
		}
	}
}

func writeSummary() {
	var insert string
	if _, err := os.Stat("Summary.md"); err == nil {
		r, err := ioutil.ReadFile("Summary.md")
		check(err)
		insert = string(r)
		insert += Summary
	} else if os.IsNotExist(err) {
		insert = Summary
	}
	f, err := os.Create("Summary.md")
	check(err)
	defer f.Close()
	f.Sync()
	f.WriteString(insert)
}

func translateText(input string) string {
	ctx := context.Background()
	client, err := translate.NewClient(ctx)
	check(err)
	// Use the client.
	translations, err := client.Translate(ctx,
		[]string{input}, language.English,
		&translate.Options{
			Source: language.TraditionalChinese,
			Format: translate.Text,
		})
	// Close the client when finished.
	err = client.Close()
	check(err)
	return translations[0].Text
}

func writeReadME(root *Item) {
	f, err := os.Create(root.path + "/README.md")
	check(err)
	defer f.Close()
	f.Sync()
	root.writeReadME(f)
	if childHasRequest(root) {
		for _, Item := range root.Items {
			if hasRequest(Item) {
				Item.writing(f)
			}
		}
	}
}
