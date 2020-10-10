package class2

import (
	"bufio"

	"os"
)

func HashFromFile(path, outPath string) error {

	// in := flag.Bool("in", false, "Sumuje podane argumenty")
	// out := flag.Bool("out", false, "Sumuje podane argumenty")
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return err
	}

	outFile, err := os.OpenFile(outPath, os.O_WRONLY|os.O_CREATE, 0666)
	defer outFile.Close()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		r, err := Hash(scanner.Text())
		println(r)
		if err != nil {
			return err
		}
		_, err = outFile.WriteString(r + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
