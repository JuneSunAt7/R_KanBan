package main
import (
    "net/http"
    "github.com/pterm/pterm"
	"fmt"
    "os/exec"
    "runtime"
)
func openBrowser(url string) error {
    var cmd *exec.Cmd

    switch runtime.GOOS {
    case "linux":
        cmd = exec.Command("xdg-open", url)
    case "windows":
        cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
    case "darwin":
        cmd = exec.Command("open", url)
    default:
        return fmt.Errorf("unsupported platform")
    }

    return cmd.Start()
}
func main(){
	fs := http.FileServer(http.Dir("web"))
    http.Handle("/", fs)

    
	url := "http://localhost:500" // Замените на нужный URL
    if err := openBrowser(url); err != nil {
        fmt.Printf("Ошибка при открытии браузера: %s\n", err)
    }
	http.ListenAndServe(":500", nil)
    pterm.Success.Println("Запущен клиент на http://localhost:500/")

}
