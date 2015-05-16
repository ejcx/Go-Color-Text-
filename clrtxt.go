package clrtxt

import (
        "fmt"
)
const RED    = "\033[31m"
const GREEN  = "\033[32m"
const YELLOW = "\033[33m"
const WHITE  = "\033[37m"
var defaultcolor = WHITE
func colorizer (s, c string) string {
       return fmt.Sprintf("%s%s%s", c, s, defaultcolor)
}
func Setdefaultgreen(){
        defaultcolor = GREEN
}
func Setdefaultyellow(){
        defaultcolor = YELLOW
}
func Setdefaultred(){
        defaultcolor = RED
}
func Setdefaultwhite(){
        defaultcolor = WHITE
}
func Green(t string) string {
        return colorizer(t, GREEN)
}
func Red(t string) string {
        return colorizer(t, RED)
}
func Yellow(t string) string {
        return colorizer(t, YELLOW)
}
func White(t string) string {
        return colorizer(t, WHITE)
}
