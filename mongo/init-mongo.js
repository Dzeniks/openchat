db.createUser({
    user: 'root',
    pwd: 'toor',
    roles: [
        {
            role: 'readWrite',
            db: 'openchat',
        },
    ],
});

db = new Mongo().getDB("openchat");

db.createCollection('users', { capped: false });
db.createCollection('chats', { capped: false });
