# ООП GoTask

**Описание задачи**

У вас уже есть функция для маскирования ссылок и теперь нужно применить полученные знания в этом модуле и добавить дополнительный поставщик данных (Producer), который будет читать файл и предоставлять его контент и слой который будет выводить результат (Presenter). Их реализацию вынести за интерфейс.

Мы выделим отдельную структуру сервиса у которой будет два интерфейсных значения:

```go
type producer interface {}

type presenter interface {}

type Service struct{
	prod producer
	pres presenter
}
```

Бизнес логика самого сервиса (его методы) должны быть одинаковы для любых реализаций процессора и продюсера.

**Правила выполнения**

- Не изменять логику работы и сигнатуру готовой функцию маскирования

- Создать новую сущность сервиса (структуру), который умеет поставлять нашей функции данные (чтение из файла) и предоставлять результаты работы (записать в файл)

	- Поставщик данных должен реализовать интерфейс`produce()([]string, error)`
	- Обработчик результата должен реализовать интерфейс `present([]string) error`
 
- Ваше приложение будет получать от пользователя путь к файлу  (получить из аргументов запуска приложения) и передать в продюсер через конструктор

- Путь к файлу, в который будем записывать результат мы так же получаем через конструктор (если пользователь не передал его через аргументы запуска то используем значение по-умолчанию) Если файл уже существует - перезаписать его.

- Сервис должен иметь конструктор, который принимает реализацию `presenter` и `producer`

- Сервис должен иметь главный метод запуска `Run`

- Реализация сервиса и его компонентов должна быть в отдельном пакете