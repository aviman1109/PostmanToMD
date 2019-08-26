package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`(\()([a-zA-Z0-9-_ ]*)(\))`)

func root(root *Item) []*Item {
	format := translateText(root.Info.Name)
	// f := strings.ReplaceAll(format, "(", "_")
	f := re.ReplaceAllString(format, "_${2}_")
	root.path = strings.ReplaceAll(f, " ", "-")
	root.Name = root.Info.Name
	root.parent = root
	root.level = 1
	if len(root.Items) > 0 {
		return root.Items
	}
	return nil
}
func child(Node *Item, path string) []*Item {
	// Node.level = Node.parent.level + 1
	Node.path += strings.ReplaceAll(path, " ", "-")
	if hasChild(Node) {
		Node.path += "/"
		var format = translateText(Node.Name)
		f := re.ReplaceAllString(format, "_${2}_")
		Node.path += strings.ReplaceAll(f, " ", "-")
		for _, Item := range Node.Items {
			Item.parent = Node
			Item.level = Node.level + 1
		}
		return Node.Items
	}
	return nil
}

func getParent(Item *Item) *Item {
	return Item.parent
}
func (Item *Item) hasChild() bool {
	if Item.Items != nil {
		return true
	}
	return false
}
func (Item *Item) hasItems() bool {
	if Item.Items != nil {
		return true
	}
	return false
}
func hasChild(Item *Item) bool {
	if Item.Items != nil {
		return true
	}
	return false
}
func hasItems(Item *Item) bool {
	if Item.Items != nil {
		return true
	}
	return false
}
func hasRequest(Item *Item) bool {
	if Item.Request != nil {
		return true
	}
	return false
}
func childHasRequest(Item *Item) bool {
	if hasChild(Item) {
		for _, i := range Item.Items {
			if hasRequest(i) {
				return true
			}
		}
	}
	return false
}
func childHasItem(Item *Item) bool {
	if hasChild(Item) {
		for _, i := range Item.Items {
			if hasItems(i) {
				return true
			}
		}
	}
	return false
}

func (Item *Item) printName() {
	if Item.Name != "" {
		fmt.Println(Item.Name)
	} else if Item.Info != nil {
		fmt.Println(Item.Info.Name)
	} else {
		fmt.Println("there are some err")
	}
}
func (Item *Item) printPath() {
	fmt.Println(Item.path)
}

func nodeTpye(Item *Item) string {
	if hasChild(Item) && !hasRequest(Item) && childHasItem(Item) && childHasRequest(Item) {
		return "floder"
	} else if hasChild(Item) && !hasRequest(Item) && !childHasItem(Item) && childHasRequest(Item) {
		return "apis"
	} else if !hasChild(Item) && hasRequest(Item) && !childHasItem(Item) && !childHasRequest(Item) {
		return "api"
	} else if hasChild(Item) && !hasRequest(Item) && !childHasItem(Item) && !childHasRequest(Item) {
		return "emptyfloder"
	}
	fmt.Printf("wrong %t %t %t %t", hasChild(Item), hasRequest(Item), childHasItem(Item), childHasRequest(Item))
	return "wrong"

}

func (Item *Item) writing(fp *os.File) {
	w := bufio.NewWriter(fp)
	fmt.Fprint(w, "{% api \""+Item.Name+"\", ")
	fmt.Fprint(w, "method=\""+Item.Request.Method+"\",")
	fmt.Fprint(w, "url=\""+Item.Request.URL.Raw+"\" %}\n")
	fmt.Fprintln(w, Item.Description)
	fmt.Fprintln(w, "## Request")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "```json")
	fmt.Fprintln(w, Item.Request.Body.Raw)
	fmt.Fprintln(w, "```")
	fmt.Fprintln(w, "")
	if Item.Response != nil {
		for _, Response := range Item.Response {
			fmt.Fprintln(w, "## Response")
			fmt.Fprintln(w, "")
			fmt.Fprintf(w, "status: %s\n", Response.Status)
			fmt.Fprintf(w, "code: %d\n", Response.Code)
			fmt.Fprintln(w, "```json")
			fmt.Fprintln(w, Response.Body)
			fmt.Fprintln(w, "```")
			fmt.Fprintln(w, "")
		}
	}

	fmt.Fprintln(w, "{% endapi %}\n")
	w.Flush() // Don't forget to flush!
}

func (Item *Item) writeTitle(fp *os.File) {
	w := bufio.NewWriter(fp)
	fmt.Fprintf(w, "# %s\n\n", Item.Name)
	w.Flush() // Don't forget to flush!
}

func (Item *Item) writeReadME(fp *os.File) {
	w := bufio.NewWriter(fp)
	fmt.Fprintf(w, "# %s\n\n", Item.Info.Name)
	fmt.Fprintln(w, Item.Info.Description)
	fmt.Fprintln(w, "")
	w.Flush() // Don't forget to flush!
}
