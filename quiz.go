package main

import (
    "fmt"
    "log"
    "time"

    "github.com/jroimartin/gocui"
)

var wynik int = 0
var questnum int = 1
 
func main() {
  
    g, err := gocui.NewGui(gocui.OutputNormal)
    if err != nil {
        log.Panicln(err)
    }
    defer g.Close()
 
    g.Cursor = true
    g.Mouse = true
 
    g.SetManagerFunc(layout)
 
    if err := keybindings(g); err != nil {
        log.Panicln(err)
    }
 
    if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
        log.Panicln(err)
    }

}
 
func layout(g *gocui.Gui) error {

    if v, err := g.SetView("wynik", 50, 0, 66, 2); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelBgColor = gocui.ColorGreen
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "WYNIK =",wynik)
    }
    if v, err := g.SetView("questnum", 30, 0, 48, 2); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "PYTANIE -", questnum ,"z 6")
    }

    if v, err := g.SetView("quest1", 2, 3, 66, 5); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "Ktore z tych miast jest stolica Polski?")
    }
    if v, err := g.SetView("but1", 2, 7, 22, 9); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "A - POZNAN")
    }
    if v, err := g.SetView("but2", 24, 7, 44, 9); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "B - WROCLAW")
    }
    if v, err := g.SetView("but3", 46, 7, 66, 9); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "C - WARSZAWA")
    }
    return nil
}
 
 
func keybindings(g *gocui.Gui) error {
    if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
        return err
    }
    for _, n := range []string{"but3"} {
        if err := g.SetKeybinding(n, gocui.MouseLeft, gocui.ModNone, YshowMsg); err != nil {
            return err
        }
    }
    for _, n := range []string{"but1", "but2"} {
        if err := g.SetKeybinding(n, gocui.MouseLeft, gocui.ModNone, NshowMsg); err != nil {
            return err
        }
    }
    for _, n := range []string{"abut1"} {
        if err := g.SetKeybinding(n, gocui.MouseLeft, gocui.ModNone, aYshowMsg); err != nil {
            return err
        }
    }
    for _, n := range []string{"abut2", "abut3"} {
        if err := g.SetKeybinding(n, gocui.MouseLeft, gocui.ModNone, aNshowMsg); err != nil {
            return err
        }
    }
    for _, n := range []string{"bbut1", "bbut2",} {
        if err := g.SetKeybinding(n, gocui.MouseLeft, gocui.ModNone, bYshowMsg); err != nil {
            return err
        }
    }
    for _, n := range []string{"bbut3"} {
        if err := g.SetKeybinding(n, gocui.MouseLeft, gocui.ModNone, bNshowMsg); err != nil {
            return err
        }
    }
    for _, n := range []string{"cbut2",} {
        if err := g.SetKeybinding(n, gocui.MouseLeft, gocui.ModNone, cYshowMsg); err != nil {
            return err
        }
    }
    for _, n := range []string{"cbut1", "cbut3"} {
        if err := g.SetKeybinding(n, gocui.MouseLeft, gocui.ModNone, cNshowMsg); err != nil {
            return err
        }
    }
    for _, n := range []string{"dbut1",} {
        if err := g.SetKeybinding(n, gocui.MouseLeft, gocui.ModNone, dYshowMsg); err != nil {
            return err
        }
    }
    for _, n := range []string{"dbut2", "dbut3"} {
        if err := g.SetKeybinding(n, gocui.MouseLeft, gocui.ModNone, dNshowMsg); err != nil {
            return err
        }
    }
    for _, n := range []string{"ebut3",} {
        if err := g.SetKeybinding(n, gocui.MouseLeft, gocui.ModNone, eYshowMsg); err != nil {
            return err
        }
    }
    for _, n := range []string{"ebut1", "ebut2"} {
        if err := g.SetKeybinding(n, gocui.MouseLeft, gocui.ModNone, eNshowMsg); err != nil {
            return err
        }
    }
    return nil
}
 
func quit(g *gocui.Gui, v *gocui.View) error {
    return gocui.ErrQuit
}
 
func YshowMsg(g *gocui.Gui, v *gocui.View) error {
/*
    maxX, maxY := g.Size()
    if v, err := g.SetView("cos", maxX/2-10, maxY/2, maxX/2+10, maxY/2+2); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "Dobrze!")
    }
    return nil
*/

    ch := make(chan bool, 1)
    timeout := make(chan bool, 1)

    go func() {
        fmt.Println("Dobrze")
        time.Sleep(2 * time.Second)
        timeout <- true
    }()

    select {
    case <-ch:
        fmt.Println("Read from ch")
    case <-timeout:
    }
 
    wynik = wynik + 1
    questnum = questnum + 1
    g, err := gocui.NewGui(gocui.OutputNormal)
    if err != nil {
        log.Panicln(err)
    }
    defer g.Close()
 
    g.Cursor = true
    g.Mouse = true
 
    g.SetManagerFunc(alayout)
 
    if err := keybindings(g); err != nil {
        log.Panicln(err)
    }
 
    if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
        log.Panicln(err)
    }
    return nil
}
 
func NshowMsg(g *gocui.Gui, v *gocui.View) error {
    ch := make(chan bool, 1)
    timeout := make(chan bool, 1)

    go func() {
        fmt.Println("Zle")
        time.Sleep(2 * time.Second)
        timeout <- true
    }()

    select {
    case <-ch:
        fmt.Println("Read from ch")
    case <-timeout:
    }

 
    wynik = wynik
    questnum = questnum + 1
    g, err := gocui.NewGui(gocui.OutputNormal)
    if err != nil {
        log.Panicln(err)
    }
    defer g.Close()
 
    g.Cursor = true
    g.Mouse = true
 
    g.SetManagerFunc(alayout)
 
    if err := keybindings(g); err != nil {
        log.Panicln(err)
    }
 
    if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
        log.Panicln(err)
    }
    return nil
}
 
func alayout(g *gocui.Gui) error {
    if v, err := g.SetView("wynik", 50, 0, 66, 2); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelBgColor = gocui.ColorGreen
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "WYNIK =",wynik)
    }
    if v, err := g.SetView("questnum", 30, 0, 48, 2); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "PYTANIE -", questnum ,"z 6")
    }
    if v, err := g.SetView("aquest1", 2, 3, 66, 5); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "W ktorym kraju znajduje sie Sekwana?")
    }
    if v, err := g.SetView("abut1", 2, 7, 22, 9); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "A - FRANCJA")
    }
    if v, err := g.SetView("abut2", 24, 7, 44, 9); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "B - UKRAINA")
    }
    if v, err := g.SetView("abut3", 46, 7, 66, 9); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "C - MALTA")
    }
    return nil
}
 
func aYshowMsg(g *gocui.Gui, v *gocui.View) error {
    ch := make(chan bool, 1)
    timeout := make(chan bool, 1)

    go func() {
        fmt.Println("Dobrze")
        time.Sleep(2 * time.Second)
        timeout <- true
    }()

    select {
    case <-ch:
        fmt.Println("Read from ch")
    case <-timeout:
    }

    wynik = wynik + 1
    questnum = questnum + 1

    g, err := gocui.NewGui(gocui.OutputNormal)
    if err != nil {
        log.Panicln(err)
    }
    defer g.Close()
 
    g.Cursor = true
    g.Mouse = true
 
    g.SetManagerFunc(blayout)
 
    if err := keybindings(g); err != nil {
        log.Panicln(err)
    }
 
    if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
        log.Panicln(err)
    }
    return nil
}
 
func aNshowMsg(g *gocui.Gui, v *gocui.View) error {
    ch := make(chan bool, 1)
    timeout := make(chan bool, 1)

    go func() {
        fmt.Println("Zle")
        time.Sleep(2 * time.Second)
        timeout <- true
    }()

    select {
    case <-ch:
        fmt.Println("Read from ch")
    case <-timeout:
    }

 
    wynik = wynik
    questnum = questnum + 1

    g, err := gocui.NewGui(gocui.OutputNormal)
    if err != nil {
        log.Panicln(err)
    }
    defer g.Close()
 
    g.Cursor = true
    g.Mouse = true
 
    g.SetManagerFunc(blayout)
 
    if err := keybindings(g); err != nil {
        log.Panicln(err)
    }
 
    if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
        log.Panicln(err)
    }
    return nil
}
 
func blayout(g *gocui.Gui) error {
    if v, err := g.SetView("wynik", 50, 0, 66, 2); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelBgColor = gocui.ColorGreen
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "WYNIK =",wynik)
    }
    if v, err := g.SetView("questnum", 30, 0, 48, 2); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "PYTANIE -", questnum ,"z 6")
    }
    if v, err := g.SetView("bquest1", 2, 3, 66, 5); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "Czy Wielka Brytania jest w UE?")
    }
    if v, err := g.SetView("bbut1", 2, 7, 22, 9); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "A - TAK")
    }
    if v, err := g.SetView("bbut2", 24, 7, 44, 9); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "B - CHYBA TAK")
    }
    if v, err := g.SetView("bbut3", 46, 7, 66, 9); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "C - NIE")
    }
    return nil
}
 
func bYshowMsg(g *gocui.Gui, v *gocui.View) error {
    ch := make(chan bool, 1)
    timeout := make(chan bool, 1)

    go func() {
        fmt.Println("Dobrze")
        time.Sleep(2 * time.Second)
        timeout <- true
    }()

    select {
    case <-ch:
        fmt.Println("Read from ch")
    case <-timeout:
    }

    wynik = wynik + 1
    questnum = questnum + 1
    g, err := gocui.NewGui(gocui.OutputNormal)
    if err != nil {
        log.Panicln(err)
    }
    defer g.Close()
 
    g.Cursor = true
    g.Mouse = true
 
    g.SetManagerFunc(clayout)
 
    if err := keybindings(g); err != nil {
        log.Panicln(err)
    }
 
    if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
        log.Panicln(err)
    }
    return nil
}
 
func bNshowMsg(g *gocui.Gui, v *gocui.View) error {
    ch := make(chan bool, 1)
    timeout := make(chan bool, 1)

    go func() {
        fmt.Println("Zle")
        time.Sleep(2 * time.Second)
        timeout <- true
    }()

    select {
    case <-ch:
        fmt.Println("Read from ch")
    case <-timeout:
    }

 
    wynik = wynik
    questnum = questnum + 1
    g, err := gocui.NewGui(gocui.OutputNormal)
    if err != nil {
        log.Panicln(err)
    }
    defer g.Close()
 
    g.Cursor = true
    g.Mouse = true
 
    g.SetManagerFunc(clayout)
 
    if err := keybindings(g); err != nil {
        log.Panicln(err)
    }
 
    if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
        log.Panicln(err)
    }
    return nil
}
 
func clayout(g *gocui.Gui) error {
    if v, err := g.SetView("wynik", 50, 0, 66, 2); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelBgColor = gocui.ColorGreen
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "WYNIK =",wynik)
    }
    if v, err := g.SetView("questnum", 30, 0, 48, 2); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "PYTANIE -", questnum ,"z 6")
    }
    if v, err := g.SetView("cquest1", 2, 3, 66, 5); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "Nad jakim morzem lezy Polska?")
    }
    if v, err := g.SetView("cbut1", 2, 7, 22, 9); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "A - CZARNYM")
    }
    if v, err := g.SetView("cbut2", 24, 7, 44, 9); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "B - BALTYCKIM")
    }
    if v, err := g.SetView("cbut3", 46, 7, 66, 9); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "C - NORWESKIM")
    }
    return nil
}
 
func cYshowMsg(g *gocui.Gui, v *gocui.View) error {
    ch := make(chan bool, 1)
    timeout := make(chan bool, 1)

    go func() {
        fmt.Println("Dobrze")
        time.Sleep(2 * time.Second)
        timeout <- true
    }()

    select {
    case <-ch:
        fmt.Println("Read from ch")
    case <-timeout:
    }

    wynik = wynik + 1
    questnum = questnum + 1
    g, err := gocui.NewGui(gocui.OutputNormal)
    if err != nil {
        log.Panicln(err)
    }
    defer g.Close()
 
    g.Cursor = true
    g.Mouse = true
 
    g.SetManagerFunc(dlayout)
 
    if err := keybindings(g); err != nil {
        log.Panicln(err)
    }
 
    if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
        log.Panicln(err)
    }
    return nil
}
 
func cNshowMsg(g *gocui.Gui, v *gocui.View) error {
    ch := make(chan bool, 1)
    timeout := make(chan bool, 1)

    go func() {
        fmt.Println("Zle")
        time.Sleep(2 * time.Second)
        timeout <- true
    }()

    select {
    case <-ch:
        fmt.Println("Read from ch")
    case <-timeout:
    }

 
    wynik = wynik
    questnum = questnum + 1
    g, err := gocui.NewGui(gocui.OutputNormal)
    if err != nil {
        log.Panicln(err)
    }
    defer g.Close()
 
    g.Cursor = true
    g.Mouse = true
 
    g.SetManagerFunc(dlayout)
 
    if err := keybindings(g); err != nil {
        log.Panicln(err)
    }
 
    if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
        log.Panicln(err)
    }
    return nil
}
 
func dlayout(g *gocui.Gui) error {
    if v, err := g.SetView("wynik", 50, 0, 66, 2); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelBgColor = gocui.ColorGreen
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "WYNIK =",wynik)
    }
    if v, err := g.SetView("questnum", 30, 0, 48, 2); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "PYTANIE -", questnum ,"z 6")
    }
    if v, err := g.SetView("dquest1", 2, 3, 66, 5); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "Ktory kraj jest wiekszy powierzchniowo")
    }
    if v, err := g.SetView("dbut1", 2, 7, 22, 9); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "A - NORWEGIA")
    }
    if v, err := g.SetView("dbut2", 24, 7, 44, 9); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "B - POLSKA")
    }
    if v, err := g.SetView("dbut3", 46, 7, 66, 9); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "C - WLOCHY")
    }
    return nil
}
 
func dYshowMsg(g *gocui.Gui, v *gocui.View) error {
    ch := make(chan bool, 1)
    timeout := make(chan bool, 1)

    go func() {
        fmt.Println("Dobrze")
        time.Sleep(2 * time.Second)
        timeout <- true
    }()

    select {
    case <-ch:
        fmt.Println("Read from ch")
    case <-timeout:
    }

    wynik = wynik + 1
    questnum = questnum + 1
    g, err := gocui.NewGui(gocui.OutputNormal)
    if err != nil {
        log.Panicln(err)
    }
    defer g.Close()
 
    g.Cursor = true
    g.Mouse = true
 
    g.SetManagerFunc(elayout)
 
    if err := keybindings(g); err != nil {
        log.Panicln(err)
    }
 
    if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
        log.Panicln(err)
    }
    return nil
}
 
func dNshowMsg(g *gocui.Gui, v *gocui.View) error {
    ch := make(chan bool, 1)
    timeout := make(chan bool, 1)

    go func() {
        fmt.Println("Zle")
        time.Sleep(2 * time.Second)
        timeout <- true
    }()

    select {
    case <-ch:
        fmt.Println("Read from ch")
    case <-timeout:
    }

 
    wynik = wynik
    questnum = questnum + 1
    g, err := gocui.NewGui(gocui.OutputNormal)
    if err != nil {
        log.Panicln(err)
    }
    defer g.Close()
 
    g.Cursor = true
    g.Mouse = true
 
    g.SetManagerFunc(elayout)
 
    if err := keybindings(g); err != nil {
        log.Panicln(err)
    }
 
    if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
        log.Panicln(err)
    }
    return nil
}
 
func elayout(g *gocui.Gui) error {
    if v, err := g.SetView("wynik", 50, 0, 66, 2); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelBgColor = gocui.ColorGreen
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "WYNIK =",wynik)
    }
    if v, err := g.SetView("questnum", 30, 0, 48, 2); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "PYTANIE -", questnum ,"z 6")
    }
    if v, err := g.SetView("equest1", 2, 3, 66, 5); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "Baku to stolica")
    }
    if v, err := g.SetView("ebut1", 2, 7, 22, 9); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "A - GRUZJI")
    }
    if v, err := g.SetView("ebut2", 24, 7, 44, 9); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "B - ARMENII")
    }
    if v, err := g.SetView("ebut3", 46, 7, 66, 9); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "C - AZERBEJDZANU")
    }
    return nil
}
 
func eYshowMsg(g *gocui.Gui, v *gocui.View) error {
    ch := make(chan bool, 1)
    timeout := make(chan bool, 1)

    go func() {
        fmt.Println("Dobrze")
        time.Sleep(2 * time.Second)
        timeout <- true
    }()

    select {
    case <-ch:
        fmt.Println("Read from ch")
    case <-timeout:
    }

    wynik = wynik + 1
    questnum = questnum + 1
 
    g, err := gocui.NewGui(gocui.OutputNormal)
    if err != nil {
        log.Panicln(err)
    }
    defer g.Close()
 
    g.Cursor = true
    g.Mouse = true
 
    g.SetManagerFunc(endlayout)
 
    if err := keybindings(g); err != nil {
        log.Panicln(err)
    }
 
    if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
        log.Panicln(err)
    }
    return nil
}
 
func eNshowMsg(g *gocui.Gui, v *gocui.View) error {
    ch := make(chan bool, 1)
    timeout := make(chan bool, 1)

    go func() {
        fmt.Println("Zle")
        time.Sleep(2 * time.Second)
        timeout <- true
    }()

    select {
    case <-ch:
        fmt.Println("Read from ch")
    case <-timeout:
    }

 
    wynik = wynik
    questnum = questnum + 1
 
    g, err := gocui.NewGui(gocui.OutputNormal)
    if err != nil {
        log.Panicln(err)
    }
    defer g.Close()
 
    g.Cursor = true
    g.Mouse = true
 
    g.SetManagerFunc(endlayout)
 
    if err := keybindings(g); err != nil {
        log.Panicln(err)
    }
 
    if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
        log.Panicln(err)
    }
    return nil
}
 
func endlayout(g *gocui.Gui) error {
    if v, err := g.SetView("equest1", 2, 3, 66, 5); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelBgColor = gocui.ColorRed
        v.SelFgColor = gocui.ColorWhite
        fmt.Fprintln(v, "TWOJ WYNIK TO ", wynik ,"/6")
    }
    return nil
}
