# avito-backend-trainee-assignment
 Dynamic user segmentation service

 Вопросы:
 -Что возвращать при частично верном запросе "Добавить пользователя в сегмент"? 
 Примеры: 
    один из сегментов в который пытаются добавить пользователя не существует; 
    один из сегментов из которых хотят исключить пользователя изначально не содержал данного пользователя.
 Решение: Полностью отклонить запрос, указав в ответе где именно была ошибка.

