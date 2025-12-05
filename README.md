# Expense Tracker CLI
Простой консольный трекер расходов на Golang с использованием библиотеки Cobra

## ТЗ
https://roadmap.sh/projects/expense-tracker

## Установка
### Требования
* Go 1.16 или выше

### Сборка из исходников
```
# Клонирование репозитория
git clone https://github.com/takhv/CLIexpenseTracker.git
cd CLIexpenseTracker

# Сборка проекта
go build

# Для глобальной установки
go install
```

## Использование
```
# Добавление задачи
expenseTracker add --desc "your-expense-description" --amount *amount*

# Удаление задачи
expenseTracker delete --id *id*

# Информация о командах
expenseTracker help

# Список всех задач
expenseTracker list

# Пометить задачу как выполненную 
expenseTracker summary (--month *month-number*)

# Изменить задачу
expenseTracker update --id *id* (--desc "your-new-description") (--amount *your-new-amount*)
```
