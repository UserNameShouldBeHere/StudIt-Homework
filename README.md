# ✏️ Д/З 2 - Основы Go

## Задания:

1. Описать интерфейс для работы с некоторым API пользователей в файле cmd/main.go.
   >  Методы интерфейса:
   * получение имени пользователя `GetUserName(id uint16) string`
   * удаление `RemoveUser(id uint16) bool`
   * редактирование имени пользователя `ChangeUserName(id uint16, newName string) bool`

   Написать простую реализацию интерфейса в файле [api.go](./internal/api/api.go). Реализация должна возвращать заранее заготовленные или рандомные значения нужных типов. Проверить работу функций в [main.go](./cmd/task_1/main.go)

2. Необходимо разработать два структуры:
   * `User`, в которой содержится
      * `Name` - имя пользователя
      * `Login` - логин
      * `Password` - пароль
      * `LastSeenAt` - дата последнего захода
   * `Admin`, представляющая того же пользователя (включает в себя `User`), но с дополнительным полем `Access`, являющимя uint8
   
   После этого реализовать функцию `renameUser`, которая принимает любого пользователя и новое имя, а затем изменяет имя обычного пользователя или админа, если у последнего значение поля `Access` ниже 3

3. Необходимо написать реализации к трем функциям:
   * `getFiles` - получает все файлы из переданной директории, включая файлы в поддиректориях
   * `filterFiles` - фильтрует полученные строки следующим образом:
      * если переданная строка начинается с `*.`, то фильтрует по переданному расширению, например, `*.html`
      * в любом другом случае фильтрует по частичному совпадению с названием файла
   * `saveToFile` - сохраняет переданные строки в указанный файл

## Требования:

1. Задания не связаны и каждое задание выполняется в отдельной папке проекта