package main

func main() {

	// //Первый способ, завершение горутины закрытием канала
	// done := make(chan struct{})
	// fmt.Println("Открывается канал done")
	// fmt.Println("Запускается горутина")
	// go func(done chan struct{}) {
	// 	for {
	// 		select {
	// 		case <-done:
	// 			fmt.Println("Горутина закрыта")
	// 			return
	// 		default:
	// 			fmt.Println("В работе")
	// 			time.Sleep(200 * time.Millisecond)
	// 		}
	// 	}
	// }(done)
	// time.Sleep(time.Second * 3)
	// close(done)
	// time.Sleep(10 * time.Millisecond)

	// //Второй способ завершения с помощью context
	// cont, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	// defer cancel()

	// go func(cont context.Context) {
	// 	for {
	// 		select {
	// 		case <-cont.Done():
	// 			fmt.Println("Горутина завершена", cont.Err())
	// 			return
	// 		default:
	// 			fmt.Println("В работе")
	// 			time.Sleep(time.Millisecond * 200)
	// 		}
	// 	}
	// }(cont)
	// time.Sleep(2 * time.Second)

	// // Третий способ завершения - просто через return
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("Горутина завершена")
	// 	return
	// }()
	// time.Sleep(3 * time.Second)
	// fmt.Println("Программа завершена")

	// //Четвёртый способ panic()
	// go func() {
	// 	panic("Goroutine panic")
	// }()
	// time.Sleep(time.Second * 1)
}
