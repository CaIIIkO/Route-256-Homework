
## Домашнее задание №1 «Утилита для управления ПВЗ»
Необходимо реализовать консольную утилиту для менеджера ПВЗ.

Тулза запускается каждый раз и принимает команды на вход.

Тулза должна обладать командой help, благодаря которой можно получить список доступных команд с кратким описанием. 

Тулза должна обрабатывать кейс запуска без файла.

Список команд для реализации:

 1. **Принять заказ от курьера.** На вход принимается ID заказа, ID получателя и срок хранения. Товар нельзя принять дважды. Если срок хранения в прошлом, приложение должно выдать ошибку. Список принятых товаров необходимо сохранять в файл. Формат файла остается на выбор автора.
 2. **Вернуть заказ курьеру.** На вход принимается ID заказа. Метод должен удалять поставку из вашего файла. Можно вернуть только те поставки, у которых вышел срок хранения и которые не были выданы клиенту.
 3. **Выдать заказ клиенту.** На вход принимается слайс ID заказов. Можно выдавать только те заказы, которые были приняты от курьера (есть в файле) и чей срок хранения меньше текущей даты. Все ID заказов должны принадлежать только одному клиенту. Если заказы не найдены или не все заказы принадлежат одному клиенту - выдавать ошибку.
 5. **Получить список заказов.** На вход принимается ID пользователя как обязательный параметр и опциональные параметры, которые позволяют получать только последние N (передаваемый параметр) товаров или параметр, который позволяет получить только заказы клиента находящиеся в нашем ПВЗ.
 6. **Принять возврат от клиента.** На вход принимается  ID пользователя и ID заказа. Заказ может быть возвращен в течении 2х дней с момента выдачи. Так же необходимо проверить, что заказ выдавался с нашего ПВЗ.
 7. **Получить список возвратов.** Метод должен выдавать список пагинированно.


### Подсказки
- `os.OpenFile`, `os.Create`,`os.ReadALL`,`os.Args`