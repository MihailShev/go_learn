#Домашнее задание №6
##Функция логирования Otus
Задача: написать функцию логирования LogOtusEvent, на вход которой приходят события типа HwAccepted (домашняя работа принята) и HwSubmitted (студент отправил дз) для этого - спроектировать и реализовать интерфейс OtusEvent. Для события HwAccepted мы хотим логирровать дату, айди и грейд, для HwSubmitter - дату, id и комментарий, например:

2019-01-01 submitted 3456 "please take a look at my homework"
2019-01-01 accepted 3456 4

```
type HwAccepted struct {
Id int
Grade int
}

type HwSubmitted struct {
Id int
Code string
Comment string
}

interface OtusEvent {
}

function LogOtusEvent(e OtusEvent, w io.Writer) {

}
```