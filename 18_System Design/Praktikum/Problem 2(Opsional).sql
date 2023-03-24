-- Terdapat sebuah query pada SQL yaitu `SELECT * FROM users;`

-- Dengan tujuan yang sama, tuliskan dalam bentuk perintah:
-- in SQL
SELECT * FROM users;

-- 1. Redis
 GET * FROM users;
KEYS user:*;

-- 2. Neo4j
MATCH (n) RETURN n

-- 3. Cassandra
SELECT * FROM users;