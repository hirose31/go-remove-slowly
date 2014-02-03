package main

import(
	"fmt"
	"os"
	"flag"
)

const TruncateSize int64 = 1024 * 1024 * 128

func main() {
	var (
		force bool
	)
	flag.BoolVar(&force, "force", false, "never prompt")
	flag.BoolVar(&force, "f",     false, "never prompt")
	flag.Parse()

	for _, fn := range flag.Args() {

		fi, err := os.Stat(fn)
		if err != nil {
			fmt.Println(fn, err);
			continue
		}
		fmt.Printf("File: %s\n", fi.Name())
		fmt.Printf("Size: %d\n", fi.Size());

		if !force {
			// prompt
			fmt.Print("Remove this file? (y/[n]) ")
			yn := "n"
			fmt.Scan(&yn);
			if (yn != "y") {
				fmt.Println("Skip");
				continue
			}
		}

		n := ( fi.Size() - 1 ) / TruncateSize - 1
		fmt.Println(n)
		for n > 0 {
			fmt.Println("left: chunk", n);
			if err := os.Truncate(fn, n * TruncateSize); err != nil {
				fmt.Println(fn, err);
				continue
			}
			n--
		}
		if err := os.Remove(fn); err != nil {
			fmt.Println(fn, err);
			continue
		}
	}
}
