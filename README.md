# ✏️ Д/З 1 - Основы git

## Задание

1. Создать приватный репозиторий на Github: `Your repositories -> New`
2. Добавить ментора: `Repository settings -> Collaborators -> Add people`
3. Склонировать в свою папку ветку с первым д/з: `git clone https://github.com/UserNameShouldBeHere/StudIt-Homework.git --branch homework-1`
4. Создать новую ветку, в которой будет выполняться д/з: `git checkout -b <название>`
5. Дописать функцию Sum из файла [cmd/main.go](./cmd/main.go) так, чтобы возвращалась сумма входных значений
6. Добавить свой репозиторий в список удаленных репозиториев: `git remote add myrepo <ссылка на твой репозиторий>`
7. Закоммитить все изменения: `git add .` и `git commit -m "<комментарий>"`
8. Отправить изменения в свой репозиторий: `git push -u myrepo <название созданной ветки>`
