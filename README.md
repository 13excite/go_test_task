## Постановка

Реализовать REST апи приложение с двумя методами:

* Запись события определенного типа. Считаем, что события фиксированные и могут условно называться: a, b, c. Приложение должно записывать в любую удобную БД факт и время наступления события.

* Получение кол-ва событий каждого типа. За указанный интервал времени должен возвращаться json вида:
  ```json
   {"a": 100, "b": 200, "c": 0}
  ```

 
## Условия работы приложения:

Приложение должно обеспечить throttling регистрируемых событий.
В один момент времени приложение должно принимать не более N запросов на регистрацию события. Если этот порог превышен, то 	выдавать http ответ 509
Запуск приложения в докере
 
## Тестирование и бенчмарки:

Обеспечить тестирование приложения таким образом, чтобы была полностью “замокана” функциональность работы с базой (те для проведения тестов не требовалась физически поднятая БД)

throttling также должен быть проверен тестами.

Продемонстрировать бенчмарк и результаты проверки производительности метода, регистрирующего события.
