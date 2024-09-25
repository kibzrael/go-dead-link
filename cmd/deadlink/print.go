package deadlink

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func PrintLinks(links *[]Link) {
	const colorRed = "\033[0;31m"
	const colorNone = "\033[0m"

	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 4, ' ', 0)
	defer writer.Flush()
	fmt.Fprint(writer, "\nBROKEN LINKS:\t\n\n")
	fmt.Fprintf(writer, "%s   %s\t%s %s\n", colorNone, "Page", colorRed, "Link")
	for _, link := range *links {
		fmt.Fprintf(writer, "%s   %s\t%s %s\n", colorNone, link.page, colorRed, link.url)
	}
	fmt.Fprintf(writer, "\n%s", colorNone)
}
