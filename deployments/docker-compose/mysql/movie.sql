CREATE DATABASE IF NOT EXISTS movie;
USE movie;
CREATE TABLE movie.metadata
(
    id          bigint UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    title       varchar(30)  NOT NULL COMMENT '''名字''',
    description varchar(255) NOT NULL COMMENT '''描述''',
    director    varchar(50)  NOT NULL COMMENT '''导演'''
);
INSERT INTO movie.metadata (id, title, description, director)
VALUES (1, 'The Shawshank Redemption',
        'Two imprisoned men bond over a number of years, finding solace and eventual redemption through acts of common decency.',
        'Frank Darabont'),
       (2, 'The Godfather',
        'An organized crime dynasty''s aging patriarch transfers control of his clandestine empire to his reluctant son.',
        'Francis Ford Coppola'),
       (3, 'The Dark Knight',
        'When the menace known as the Joker wreaks havoc and chaos on the people of Gotham, Batman must accept one of the greatest psychological and physical tests of his ability to fight injustice.',
        'Christopher Nolan'),
       (4, 'Schindler''s List',
        'In German-occupied Poland during World War II, industrialist Oskar Schindler gradually becomes concerned for his Jewish workforce after witnessing their persecution by the Nazis.',
        'Steven Spielberg'),
       (5, 'Forrest Gump',
        'The presidencies of Kennedy and Johnson, the events of Vietnam, Watergate, and other historical events unfold through the perspective of an Alabama man with an IQ of 75.',
        'Robert Zemeckis');

CREATE TABLE movie.ratings
(
    record_id   bigint UNSIGNED NOT NULL,
    record_type bigint UNSIGNED NOT NULL COMMENT '''评论类型1电影 2电视剧''',
    user_id     bigint UNSIGNED NOT NULL COMMENT '''用户ID''',
    value       int UNSIGNED    NOT NULL COMMENT '''评分''',
    PRIMARY KEY (record_id, record_type)
);
INSERT INTO movie.ratings (record_type, record_id, user_id, value)
VALUES (1, 1, 1234, 9),
       (1, 2, 5678, 7),
       (1, 3, 9876, 10),
       (1, 4, 1111, 8),
       (1, 5, 1111, 8);
