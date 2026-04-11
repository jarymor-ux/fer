# ARCH

## Доменные сущности

### User

**Fields**
ID = uuid
WS_CONNECTION - *ws.Conn
messages - chan []byte
chats - []Chat
group_chats - []GroupChat

### Credentials

**Fields**
user_uuid - *User.ID
phone - Russian Phone Number(string)
password - User Password(string)

### Message

**Fields**

Sender - *User.ID
Reciver - //TODO:
Time - unix timestamp

### Chat
FirstMessageSender - *User.ID
Member1 - *User.ID
Member2 - *User.ID
History - HashSet<Message>


### GroupChat

//TODO:
