/*
    Первая программа 
    на языке Go
*/
package main
import (
  "fmt"
  "time"
  "os"
  "math"
) 
 
func main() {

	var num int = 0
	
	for {

		fmt.Print("\nМеню:\n")
		fmt.Print("1 - Узнать кол-во прошедших дней\n")
		fmt.Print("0 - Выход\n\n")
		fmt.Fscan(os.Stdin, &num)

		switch {
		case num == 1:
			getDate()
		case num == 0:
			return
		}
	}  
}

func getDate() {
	var (
		day int 
		month int
		year int
	)
	  
	fmt.Print("Введите дату отсчета:дд.мм.гггг\n")
	fmt.Print("дд: ")
	fmt.Fscan(os.Stdin, &day)
	fmt.Print("мм: ")
	fmt.Fscan(os.Stdin, &month)
	fmt.Print("гггг: ")
	fmt.Fscan(os.Stdin, &year)  

	t1 := Date(year, month, day)
	today := time.Now() 
	days := today.Sub(t1).Hours() / 24
	days = math.Floor(days)
  fmt.Println("Прошло месяцев: ", math.Floor((days/30)*10)/10)  
	fmt.Println("Прошло дней: ", days)  
	fmt.Println("День сейчас: ", days + 1) 
  
}

func Date(year, month, day int) time.Time {
    return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}