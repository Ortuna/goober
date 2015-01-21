package main

import (
    "io/ioutil"
    "net/http"
    "code.google.com/p/gofpdf"
)

type Point struct {
  x float64
  y float64
}

func posToXY(pos int) Point {
  charsPerLine := int((72*36)/8)
  if pos < charsPerLine {
    return Point{float64(pos), 0.0}
  }

  x := float64((pos - (charsPerLine * int(pos/charsPerLine))) * 8)
  y := float64(pos / charsPerLine) * 8

  return Point{x, y}
}

func handler(w http.ResponseWriter, r *http.Request) {
  var width float64 = 72*36
  var height float64 = 72*24
  buf, _ := ioutil.ReadFile("data.txt")
  data := string(buf)

  w.Header().Set("content-type", "application/pdf")
  pdf := gofpdf.New("L", "pt", "A4", "../font")
  pdf.AddPageFormat("P", gofpdf.SizeType{width, height})
  pdf.SetFont("Arial", "B", 8)

  for pos := 0; pos < len(data); pos += 1 {
    point := posToXY(pos)
    pdf.Text(point.x, point.y, "X")
  }

  pdf.Output(w)
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
