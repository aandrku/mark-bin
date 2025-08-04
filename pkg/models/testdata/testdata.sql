-- sample users

INSERT INTO users (username, email, hashed_password, created) VALUES
('alice',   'alice@example.com',   '$2a$12$C6UzMDM.H6dfI/f/IKcEe.EqDkfQ6VVphX1A0rCwbyXU0JvCk4g2e', '2024-09-15 10:30:00');

INSERT INTO users (username, email, hashed_password, created) VALUES
('bob',     'bob@example.com',     '$2a$10$7EqJtq98hPqEX7fNZaFWoO.J8z6NDF8HiK8l2r3L1GxFz0x0miY3C', '2024-10-01 15:45:00');

INSERT INTO users (username, email, hashed_password, created) VALUES
('charlie', 'charlie@example.com', '$2a$10$MT8V/Y/FmKMn2nG6zOKDeO0DQoVO.Z5uXMxtrS50YgG3k7HhM86eW', '2025-01-20 09:00:00');

INSERT INTO users (username, email, hashed_password, created) VALUES
('diana',   'diana@example.com',   '$2a$10$wHh3lDwXMiEvVLZbSGeiSewP0gytsZjwOqG14Dp7PYhCFLXpF1pBi', '2025-02-10 11:15:00');

INSERT INTO users (username, email, hashed_password, created) VALUES
('eve',     'eve@example.com',     '$2a$12$eImiTXuWVxfM37uY4JANjQ==qwertyuiopasdfghjklzxcvbnm12', '2025-03-05 17:20:00');

-- sample snippets

INSERT INTO snippets (title, content, created, updated, user_id) VALUES
('Morning Light', 'Golden sun rises\nShadows stretch across the field\nDay begins anew', '2024-01-01 08:00:00', '2024-01-01 08:00:00', 1);

INSERT INTO snippets (title, content, created, updated, user_id) VALUES
('Snowy Silence', 'Soft flakes gently fall\nThe world muffled into peace\nWinter whispers near', '2024-01-15 14:23:00', '2024-01-15 14:23:00', 2);

INSERT INTO snippets (title, content, created, updated, user_id) VALUES
('Evening Rain', 'Raindrops kiss the roof\nDistant thunder rolls like drums\nCandles flicker low', '2024-02-10 19:45:00', '2024-02-10 19:45:00', 1);

INSERT INTO snippets (title, content, created, updated, user_id) VALUES
('Spring Bloom', 'Cherry trees explode\nPetals float like pink snowflakes\nAir tastes soft and new', '2024-03-12 11:00:00', '2024-03-12 11:00:00', 3);

INSERT INTO snippets (title, content, created, updated, user_id) VALUES
('Night Thoughts', 'Moonlight through my blinds\nIdeas drift like smoke through dreams\nSleep forgets my name', '2024-04-05 23:59:00', '2024-04-05 23:59:00', 2);

INSERT INTO snippets (title, content, created, updated, user_id) VALUES
('Quiet Library', 'A turning page hums\nIn silence knowledge blossoms\nTime folds into ink', '2024-05-20 09:12:00', '2024-05-20 09:12:00', 1);

INSERT INTO snippets (title, content, created, updated, user_id) VALUES
('Summer Wind', 'Warm breeze through the trees\nLeaves whisper forgotten songs\nStillness comes alive', '2024-06-02 17:30:00', '2024-06-02 17:30:00', 3);

INSERT INTO snippets (title, content, created, updated, user_id) VALUES
('Broken Clock', 'Hands spin endlessly\nTicking time that won’t return\nMoments out of sync', '2024-07-14 02:00:00', '2024-07-14 02:00:00', 2);

INSERT INTO snippets (title, content, created, updated, user_id) VALUES
('Autumn Walk', 'Crisp leaves underfoot\nGold and red in fading light\nThe earth takes a breath', '2024-08-29 15:15:00', '2024-08-29 15:15:00', 1);

INSERT INTO snippets (title, content, created, updated, user_id) VALUES
('Final Snowfall', 'White blankets the street\nFootprints disappear at dusk\nAll begins again', '2024-12-01 07:45:00', '2024-12-01 07:45:00', 3);

