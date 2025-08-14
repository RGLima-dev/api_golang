USE devbook;

INSERT INTO users (id, username, nickname, email, passwd)
VALUES (25, 'user1', 'usernick1', 'user1@gmail.com', '$2a$10$FhOoWA1jbkNM7ZQb37tHNe4gSWQqZeetepG6hTfvkGiw8y5QYoF3u');

INSERT INTO users (id, username, nickname, email, passwd)
VALUES (26, 'user2', 'usernick2', 'user2@gmail.com', '$2a$10$FhOoWA1jbkNM7ZQb37tHNe4gSWQqZeetepG6hTfvkGiw8y5QYoF3u');

INSERT INTO users (id, username, nickname, email, passwd)
VALUES (27, 'user3', 'usernick3', 'user3@gmail.com', '$2a$10$FhOoWA1jbkNM7ZQb37tHNe4gSWQqZeetepG6hTfvkGiw8y5QYoF3u');

INSERT INTO users (id, username, nickname, email, passwd)
VALUES (28, 'user4', 'usernick4', 'user4@gmail.com', '$2a$10$FhOoWA1jbkNM7ZQb37tHNe4gSWQqZeetepG6hTfvkGiw8y5QYoF3u');

INSERT INTO followers (user_id, follower_id)
VALUES (25,26);

INSERT INTO followers (user_id, follower_id)
VALUES (25,27);

INSERT INTO followers (user_id, follower_id)
VALUES (26,27);
