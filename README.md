# JobTesting
Учебное RestAPI для получения и отправки данных в формате Json. Данные хранятся в файле users.json <br>
/users   <br>
"id"  int<br>
"name" string<br>
POST "id" , "name"  - добавляет пользователя в файл users.json<br>
GET "id" - возвращает "name"  данного id из файла<br>
PUT "id", "name" -  изменяет имя пользователя, если он есть в файле<br>
DELETE "id" - удаляет пользователя из файла<br>
GET "id": 999   - возвращает всё содержимое файла<br>
