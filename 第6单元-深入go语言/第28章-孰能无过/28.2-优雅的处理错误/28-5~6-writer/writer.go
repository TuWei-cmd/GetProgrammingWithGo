package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	err := proverbs("proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return // 如果前面有错误，就跳过写入操作
	}
	_, sw.err = fmt.Fprintln(sw.w, s) // 尝试写入文本并储存可能出现的错误
}

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	sw := safeWriter{w: f}

	sw.writeln("Don’t communicate by sharing memory, share memory by communicating.")
	sw.writeln("Concurrency is not parallism.")
	sw.writeln("Channels orchestrate; mutexes serialize.")
	sw.writeln("The bigger the interface, the weaker the abstraction.")
	sw.writeln("Make the zero value useful. (e.g. bytes.Buffer or sync.Mutex)")
	sw.writeln("interface{} says nothing.")
	sw.writeln("Gofmt’s style is no one’s favorite, yet gofmt is everyone’s favorite.")
	sw.writeln("A little copying is better than a little dependency.")
	sw.writeln("Syscall must always be guarded with build tags.")
	sw.writeln("Cgo must always be guarded with build tags.")
	sw.writeln("Cgo is not Go.")
	sw.writeln("With the unsafe package, there are no guarantees.")
	sw.writeln("Clear is better than clever.")
	sw.writeln("Reflection is never clear.")
	sw.writeln("Errors are values.")
	sw.writeln("Don’t just check errors, handle them gracefully.")
	sw.writeln("Design the architecture, name the components, document the details.")
	sw.writeln("Documentation is for users.")
	sw.writeln("Don’t panic.")

	return sw.err
}
