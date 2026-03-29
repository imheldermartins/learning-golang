curl -X POST "http://localhost:4001/db/execute?pretty" \
-H "Content-Type: application/json" \
-d @- <<EOF
[
  "CREATE TABLE IF NOT EXISTS Users (
    id INTEGER PRIMARY KEY,
    name TEXT,
    email TEXT NOT NULL UNIQUE,
    password TEXT,
    created_at TEXT DEFAULT (datetime('now','localtime')),
    updated_at TEXT DEFAULT (datetime('now','localtime'))
  )"
]
EOF

curl -X POST "http://localhost:4001/db/execute?pretty" \
-H "Content-Type: application/json" \
-d @- <<EOF
[
"CREATE TABLE IF NOT EXISTS publishers (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    website TEXT,
    country TEXT,
    created_at TEXT DEFAULT (datetime('now','localtime')),
    updated_at TEXT DEFAULT (datetime('now','localtime'))
  )",
  "CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    publisher_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    content TEXT,
    status TEXT DEFAULT 'draft',
    created_at TEXT DEFAULT (datetime('now','localtime')),
    updated_at TEXT DEFAULT (datetime('now','localtime')),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (publisher_id) REFERENCES publishers(id)
  )"
]
EOF

curl -X POST "http://localhost:4001/db/execute?pretty" \
  -H "Content-Type: application/json" \
  -d @- <<EOF
[
  "INSERT INTO users (name, email, password) VALUES ('Helder', 'helder@email.com', '123456')",
  "INSERT INTO users (name, email, password) VALUES ('Heloisa', 'helo@email.com', 'abcdef')",
  "INSERT INTO publishers (name, website, country) VALUES ('Tech News BR', 'https://technews.example', 'Brasil')",
  "INSERT INTO publishers (name, website, country) VALUES ('Open Dev Media', 'https://opendev.example', 'Portugal')",
  "INSERT INTO posts (user_id, publisher_id, title, content, status) VALUES (1, 1, 'Primeiro post', 'Conteúdo do primeiro post', 'published')",
  "INSERT INTO posts (user_id, publisher_id, title, content, status) VALUES (2, 2, 'Segundo post', 'Conteúdo do segundo post', 'draft')"
]
EOF

curl -X POST "http://localhost:4001/db/query?pretty" \
  -H "Content-Type: application/json" \
  -d @- <<EOF
[
  "SELECT * FROM users"
]
EOF

curl -X POST "http://localhost:4001/db/query?pretty" \
  -H "Content-Type: application/json" \
  -d @- <<EOF
[
  "SELECT * FROM publishers"
]
EOF

curl -X POST "http://localhost:4001/db/query?pretty" \
  -H "Content-Type: application/json" \
  -d @- <<EOF
[
  "SELECT
      p.id,
      p.title,
      p.status,
      u.name AS author_name,
      pub.name AS publisher_name,
      p.created_at
   FROM posts p
   INNER JOIN users u ON u.id = p.user_id
   INNER JOIN publishers pub ON pub.id = p.publisher_id
   ORDER BY p.id DESC"
]
EOF