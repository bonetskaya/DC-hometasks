# Домашнее задание #1

Скелет приложения принимает запросы на изменение и получение данных о товарах в интернет-магазине.

localhost:8000/items GET --- получить список товаров
localhost:8000/items PUT --- добавить товар (json вида {"category": "str", "category":"str"}, без пропуска полей!)
localhost:8000/items/{id} GET --- получить товар по id
localhost:8000/items/{id} PUT --- обновить товар по id (json вида {"category": "str", "category":"str"})
localhost:8000/items/{id} DELETE --- удалить товар по id

Пока приложение не использует бд, соответственно есть один процесс main.
