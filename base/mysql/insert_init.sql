INSERT INTO user_table (user_name, follow_count, follower_count)
values ('Judy', 4, 3);
INSERT INTO user_table (user_name, follow_count, follower_count)
values ('John', 1, 4);
INSERT INTO user_table (user_name, follow_count, follower_count)
values ('Michel', 2, 2);
INSERT INTO user_table (user_name, follow_count, follower_count)
values ('Bonjour', 1, 3);
INSERT INTO user_table (user_name, follow_count, follower_count)
values ('Ben', 3, 1);
INSERT INTO video_table(user_id, play_url, title)
values(1, 'www.test.com/Judy/video', 'Judy first video');
INSERT INTO video_table(user_id, play_url, title)
values(
        3,
        'www.test.com/Michel/video',
        'Michel first video'
    );
INSERT INTO video_table(user_id, play_url, title)
values(5, 'www.test.com/Ben/video', 'Ben first video');
INSERT INTO comment_table(video_id, user_id, content)
values(1, 2, '这简直是太酷了！');
INSERT INTO comment_table(video_id, user_id, content)
values(2, 4, '这完全符合我对MySQL的想象！');
INSERT INTO comment_table(video_id, user_id, content)
values(3, 1, '感觉不如原神');
INSERT INTO follower_table(user_id, fans_id)
values(2, 1);
INSERT INTO follower_table(user_id, fans_id)
values(3, 1);
INSERT INTO follower_table(user_id, fans_id)
values(4, 1);
INSERT INTO follower_table(user_id, fans_id)
values(5, 1);
INSERT INTO follower_table(user_id, fans_id)
values(1, 2);
INSERT INTO follower_table(user_id, fans_id)
values(1, 3);
INSERT INTO follower_table(user_id, fans_id)
values(1, 5);
INSERT INTO follower_table(user_id, fans_id)
values(2, 3);
INSERT INTO follower_table(user_id, fans_id)
values(2, 4);
INSERT INTO follower_table(user_id, fans_id)
values(2, 5);
INSERT INTO follower_table(user_id, fans_id)
values(3, 5);
INSERT INTO follower_table(user_id, fans_id)
values(4, 2);
INSERT INTO follower_table(user_id, fans_id)
values(4, 3);
INSERT INTO likes_table(user_id, video_id)
values(1, 1);
INSERT INTO likes_table(user_id, video_id)
values(2, 1);
INSERT INTO likes_table(user_id, video_id)
values(2, 3);
INSERT INTO likes_table(user_id, video_id)
values(4, 2);