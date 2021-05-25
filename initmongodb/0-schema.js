db.createUser({
    user: "db_user",
    pwd: "pass",
    roles: [
        {
            role: "readWrite",
            db: "db",
        }
    ]
});

db.createCollection("calculators");