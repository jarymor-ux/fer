Какие юзкейсы логично положить в internal/domain/usecase:

RegisterUser — регистрация пользователя по номеру (у вас уже есть нормализация/валидация телефона в types).
CreatePrivateChat — создать личный чат между 2 пользователями (или вернуть существующий).
CreateGroupChat — создать групповой чат с набором участников.
SendMessage — отправить сообщение в чат с проверкой членства пользователя.
GetChatHistory — получить историю сообщений чата (с пагинацией/лимитом).
ListUserChats — список чатов пользователя (с последним сообщением, временем обновления).
AddMemberToGroup / RemoveMemberFromGroup — управление участниками группового чата.
ConnectClient / DisconnectClient — юзкейс подключения WebSocket-клиента и привязки к UserID/ClientID.
MarkMessagesRead — отметка сообщений прочитанными (понадобится отдельная модель read-state).
DeleteMessage или EditMessage — если планируете это поведение.
Минимальный MVP (с чего начать):

RegisterUser
CreatePrivateChat
SendMessage
GetChatHistory
ListUserChats
ConnectClient (для онлайн-доставки через WS)
Go best practices для слоя usecase:

Держите usecase как application service: оркестрация бизнес-правил, без HTTP/WS деталей.
Зависимости только через интерфейсы репозиториев/транзакций/часов/ID generator.
Один usecase — один входной DTO и один выходной DTO.
Валидация инвариантов в usecase + в доменных типах (как уже сделано с телефоном).
Ошибки домена возвращайте типизированно (например, ErrUserNotInChat, ErrChatNotFound), а маппинг в HTTP-коды делайте в transport-слое.
Если хотите, могу предложить конкретный каркас файлов для internal/domain/usecase (интерфейсы + service + DTO) под ваш текущий код без изменений в проекте.
