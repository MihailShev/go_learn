Реализовать двусвязанный список. 
Что такое двусвязный список: https://en.wikipedia.org/wiki/Doubly_linked_list

Ожидаемые типы (псевдокод):

```
List      // тип контейнер
  Len()   // длинна списка
  First() // первый Item
  Last()  // последний Item
  PushFront(v interface{}) // добавить значение в начало
  PushBack(v interface{})  // добавить значение в конец

Item   // элемент списка
  Value() interface{}  // возвращает значение
  Nex() *Item          // следующий Item
  Prev() *Item         // предыдущий
  Remove()             // удалить Item из списка
```