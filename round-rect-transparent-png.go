
  package main

  import (
    "os"
    "strconv"
    "image"
    "image/color"
    "github.com/llgcode/draw2d/draw2dimg"
    "github.com/llgcode/draw2d/draw2dkit"
  )

  func main() {

    args := os.Args[1:] ; if len(args) < 6 { panic(1) }

    filename := args[0]

    height, _ := strconv.ParseFloat(args[1], 64)
    width, _ := strconv.ParseFloat(args[2], 64)
    radius, _ := strconv.ParseFloat(args[3], 64)
    opacity, _ := strconv.Atoi(args[4])
    padding, _ := strconv.Atoi(args[5])

    r := float64(radius)
    o := uint8(opacity)
    p := float64(padding)

    img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
    gc := draw2dimg.NewGraphicContext(img)

    fill := color.RGBA{0x0, 0x0, 0x0, o}

    gc.SetFillColor(fill)
    gc.SetStrokeColor(image.Transparent)
    gc.SetLineWidth(1)

    draw2dkit.RoundedRectangle(gc, 0, 0, width, height, r, r)
    gc.FillStroke()

    gc.SetStrokeColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
    gc.SetFillColor(image.Transparent)

    draw2dkit.RoundedRectangle(gc, p, p, width-p, height-p, 2, 2)
    gc.FillStroke()

    draw2dimg.SaveToPngFile(filename, img)

  }