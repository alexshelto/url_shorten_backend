CREATE TABLE url(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  original_url TEXT NOT NULL,
  shortened_url TEXT NOT NULL,
  visit_count INTEGER
);

INSERT INTO url (
    id, original_url, shortened_url, visit_count
) VALUES (
    1, "https://actix.rs/docs/databases/", "http://localhost:8080/foobar", 12
)
