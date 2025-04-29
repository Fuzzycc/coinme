package ledger

import (
	"encoding/csv"
	"io"
	"sync"
)

/*
Coin Rules
Name must be 50 characters = 50 bytes,so
Version;ID;Name;Value
Version 1 byte
ID 4 bytes
Name 50 bytes
Value 10 bytes (uint32 is 4 billion)
semicolons 3 bytes
line is 1+4+50+10+3
Total is 68 bytes

What I need:
TextAppender, LineReader
*/


type BaseReader struct {
	Comma   rune
	Comment rune
	cR      *csv.Reader
}

func NewBaseReader(r io.Reader) *BaseReader {
	rr := csv.NewReader(r)
	br := &BaseReader{
		Comma:   ';',
		Comment: '0',
		cR:      rr,
	}
	br.cR.Comma = br.Comma
	br.cR.Comment = br.Comment
	return br
}

// func (r *Reader) FieldPos(field int) (line, column int)
// func (r *Reader) InputOffset() int64

func (r *BaseReader) Read() (record []string, err error) {
	return r.cR.Read()
}

func (r *BaseReader) ReadAll() (records [][]string, err error) {
	for {
		record, err := r.Read()
		if err == io.EOF {
			return records, nil
		}
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	// taken from r.cR.ReadAll()
	// why do this though?
}

type BaseWriter struct {
	// comma rune
	cW *csv.Writer
	m  sync.Mutex
}

func (bw *BaseWriter) SetComma(comma rune) {
	// bw.comma = comma
	bw.cW.Comma = comma
}

func NewBaseWriter(w io.Writer) *BaseWriter {
	ww := csv.NewWriter(w)
	bw := &BaseWriter{
		// comma: ';',
		cW: ww,
	}
	bw.cW.Comma = ';'
	return bw
}

func (bw *BaseWriter) SetCRLF(useCRLF bool) {
	bw.cW.UseCRLF = useCRLF
}

// func (bw *BaseWriter) Error() error

func (bw *BaseWriter) Flush() {
	bw.m.Lock()
	defer bw.m.Unlock()
	bw.cW.Flush()
}

func (bw *BaseWriter) Write(record []string) error {
	bw.m.Lock()
	defer bw.m.Unlock()
	record[2] = "\"" + record[2] + "\""
	return bw.cW.Write(record)
}

func (bw *BaseWriter) WriteAll(records [][]string) error {
	bw.m.Lock()
	defer bw.m.Unlock()
	// for i := range len(records) {
	// 	records[i][2] = `'` + records[i][2] + `'`
	// 	// for some reason, using " results in 3 quotes on each side. I have no idea why
	// 	// turns out quotes, Comma, and space at the start are quoted. So a quote is quoted and becomes three.
	// }
	return bw.cW.WriteAll(records)
}


func CleanBase(r io.Reader, w io.Writer) {
	// read the cache
	// read the file
	// loop the file, matching it against the cache
}